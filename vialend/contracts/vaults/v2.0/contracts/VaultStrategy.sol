// SPDX-License-Identifier: MIT

pragma solidity >=0.8.8 <0.9.0;

import "./@openzeppelin/contracts/math/Math.sol";
import "./@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "./@uniswap/v3-core/contracts/libraries/FullMath.sol";
import "./@uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol";
import "./@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";
import "./@uniswap/v3-periphery/contracts/libraries/LiquidityAmounts.sol";
import "./@uniswap/v3-periphery/contracts/libraries/PositionKey.sol";
import "./@uniswap/v3-core/contracts/libraries/TickMath.sol";
import "./@uniswap/v3-core/contracts/interfaces/callback/IUniswapV3MintCallback.sol";
import "./@uniswap/v3-core/contracts/interfaces/callback/IUniswapV3SwapCallback.sol";
import  { ICErc20, ICEth,IWETH9 }  from "./interfaces/IViaProtocols.sol";

import "./UniCompFees.sol";
//import "./libraries/UniCompHelper.sol"; 

import "./ViaVault.sol";

/// @author  ViaLend
/// @title   strategy Uni + Compound
/// @notice  A Smart Contract that helps liquidity providers managing their funds on Uniswap V3 .

contract VaultStrategy is ReentrancyGuard , UniCompFees  {
	
	using SafeERC20 for IERC20;


	address constant public UNIV3_FACTORY = 0x1F98431c8aD98523631AE4a59f267346ea31F984;
	address public creator;
	address public immutable factory;
	address public immutable 	admin;
	address payable public immutable  vault;
	address payable public immutable 	_WETH;
    address public immutable token0;         // underlying token0
    address public immutable token1;         // underlying token1
	address public protocol;			// where fee cuts to protocol 

	ICEth public immutable _CETH;
    IUniswapV3Pool public immutable pool;        // get by uni factory, token0, token1, feetier
    uint128 public immutable quoamount;  		// for calc price, based on token0 decimal, ie: 1e18 for eth, 1e8 for wbtc
   
// 	uint8 public decimal0;
//    uint8 public decimal1;
    uint8 public uniPortion;       // uniswap portion ratio
    uint8 public compPortion;       // compound portion ratio
    uint8 public protocolFee;		// 0 - 20% of profit
    uint32 public twapDuration;        // oracle twap durantion
//    uint32 public threshold;	    // initial range

    int24 public tickSpacing;
    int24 public cLow;
    int24 public cHigh;
    int24 public maxTwapDeviation;     // for twap     

    mapping (address => address) public _CTOKEN;
    
	uint256 public curComp0;	// current amount of token0 in Compound pool
	uint256 public curComp1;	// current amount of token1 in Compound pool
	    
    uint256 public vaultCap;	   	// 0: no cap
    uint256 public individualCap;	//  0 : no cap      

    bool private isEmergency = false;  // only canbe changed within emergency

    mapping (uint => address[] ) public decenter;		// who helped to triggering buttons e.g. decenter[1].push( address )  1=rebalance
    
 	
    enum ARRY {
		    PROTOCOL,
			CREATOR,
			WETH,
			CETH,
			CTOKEN0,
			CTOKEN1,
			TOKEN0,
			TOKEN1,
			VAULT,
			ADMIN
	}

    constructor( 
    	address[10] memory _contracts,
		uint8  _uniPortion,
		uint8  _compPortion, 
		uint8  _protocolFee,
		uint24 _feetier,
		uint128  _quoamount ) 
	{
		factory = msg.sender;

		protocol =  _contracts[uint(ARRY.PROTOCOL)];
		creator = _contracts[uint(ARRY.CREATOR)];
		admin =  _contracts[uint(ARRY.ADMIN)];
        vault =   payable(_contracts[uint(ARRY.VAULT)]);

        pool = IUniswapV3Pool(IUniswapV3Factory(UNIV3_FACTORY).getPool( _contracts[uint(ARRY.TOKEN0)],_contracts[uint(ARRY.TOKEN1)], _feetier)); 
        // token0 & token1 could be changed order by the pool 
        token0 = pool.token0();  
        token1 = pool.token1();
        
        if ( token0 == _contracts[uint(ARRY.TOKEN1)] ) { 	
        	// tokens order changed , the ctokens order must change accordingly.
			_CTOKEN[token0] = _contracts[uint(ARRY.CTOKEN1)];	
			_CTOKEN[token1] =_contracts[uint(ARRY.CTOKEN0)]; 
        } else {
			_CTOKEN[token0] = _contracts[uint(ARRY.CTOKEN0)];
			_CTOKEN[token1] =_contracts[uint(ARRY.CTOKEN1)];
        }

		_WETH = payable (_contracts[uint(ARRY.WETH)]);	
        _CETH = ICEth(_contracts[uint(ARRY.CETH)]);	
		
		tickSpacing = pool.tickSpacing();
        twapDuration = 2;
        maxTwapDeviation = 20000;
		quoamount = _quoamount;
        require( _uniPortion + _compPortion <= 100, "portion uni+com>100");
        uniPortion =  _uniPortion;
        compPortion =  _compPortion;
        protocolFee = _protocolFee;
    }
    
    event MintUniV3Liquidity(int24 indexed newLow, int24 indexed newHigh, uint128 indexed liquidity);
    event BurnUniV3Liquidity(int24 indexed cLow, int24 indexed cHigh, uint128 indexed liquidity);
	event VaultWithdraw( address indexed to, uint256 amount0,  uint256 amount1);
    event Rebalance(address indexed,uint256 u0,  uint256 u1, uint256 c0,  uint256 c1);
    event MyLog(string, uint256);
    
    
    /// check status == 1
    modifier onlyActive {
        require (IVaultFactory(factory).checkActive( address(this) ), 'not active');
        _;
    }
    
    
    modifier onlyAdmin {
        require (msg.sender == admin,"not admin");  //only admin 
        _;
    }

    modifier onlyVault {
    	//require (IVaultFactory(factory).pairVault(address(this)) == msg.sender, "vault" );
        require (msg.sender == vault, 'not vault');
        _;
    }    
    
    
	receive() external payable {}
	fallback() external payable {}


    	
 /// @dev Callback for Uniswap V3 pool
    function uniswapV3MintCallback(
        uint256 amount0,
        uint256 amount1,
        bytes calldata data
    ) external  {
        require(msg.sender == address(pool),"PS");

        IERC20(token0).safeTransfer(msg.sender, amount0);
        IERC20(token1).safeTransfer(msg.sender, amount1);

 		assert(data.length == 0);
    }

    /// @dev Callback for Uniswap V3 pool
    function uniswapV3SwapCallback(
        int256 amount0Delta,
        int256 amount1Delta,
        bytes calldata data
    ) external  {
        require(msg.sender == address(pool),"PS2");
        
        IERC20(token0).safeTransfer(msg.sender, uint256(amount0Delta));
        IERC20(token1).safeTransfer(msg.sender, uint256(amount1Delta));

        assert(data.length == 0);
    }
	
	
    /// external & public functions:

	///@notice low level calls to pull all positions backt to vault
	///@param redeemType is true amount is ctoken, false amount is underlying 
	function emergency(		 
		int24 tickLower,
        int24 tickUpper,
        uint128 liquidity,
        uint256 amount0,
        uint256 amount1,
        bool redeemType
	) external onlyAdmin {
		
        isEmergency = true;
        
        pool.burn(tickLower, tickUpper, liquidity);
        pool.collect(address(this), tickLower, tickUpper, type(uint128).max, type(uint128).max);
        
        // from outside CEth.balanceOf(address(this)
        if (amount0>0) redeemCErc20(token0, amount0, redeemType);
        if (amount1>0) redeemCErc20(token1, amount1, redeemType);
        
        //transfer to vault or hold asssets here for user to withdraw
        IERC20(token0).safeTransfer(vault, IERC20(token0).balanceOf(address(this)) );
        IERC20(token1).safeTransfer(vault, IERC20(token1).balanceOf(address(this)) );

	}
	
	///@notice send funds back to vault
	function callFunds() external onlyVault {
		alloc();
		uint256 a0 =IERC20(token0).balanceOf(address(this));
		uint256 a1 =IERC20(token1).balanceOf(address(this));
        if(a0>0) IERC20(token0).safeTransfer(vault, a0);
        if(a1>0) IERC20(token1).safeTransfer(vault, a1);
	}

	
	function rebalance(
		int24 newLow,
        int24 newHigh
		) external onlyActive returns(uint256 u0,uint256 u1,uint256 c0,uint256 c1)	
	{
		alloc();
		ViaVault(vault).moveFunds();
        ( u0, u1, c0, c1) =  _rebalance(newLow, newHigh); 
        emit Rebalance(address(this), u0,u1,c0,c1);
	}

	// make sure all position has been removed before doing rebalance. 
	// when newLow and newHigh is 0, calc new range with current cLow and cHigh
	function _rebalance(int24 newLow, int24 newHigh ) 
		internal 
		returns (uint256 u0, uint256 u1, uint256 c0,uint256 c1)
	{
		(	,int24 tick, , , , , ) 	= pool.slot0();

    	if (newLow==0 && newHigh==0) {
			if (cHigh == 0 && cLow ==0) {
                return(0,0,0,0);  // cannot do rebalance if cLow and cHigh is 0
            }
			int24 hRange = ( cHigh - cLow ) / 2;
			newHigh = (tick + hRange) / tickSpacing * tickSpacing;
			newLow  = (tick - hRange) / tickSpacing * tickSpacing;
		}

  		_validRange(newLow, newHigh, tick);  // passed 1200 , 2100, 18382

		uint256 currentBalance0 = getBalance(token0);
		uint256 currentBalance1 = getBalance(token1);

        //#calculate the amounts to supply into uniswap
        (u0, u1) = calcUniPortionAmounts(currentBalance0, currentBalance1);

        if ( uniPortion > 0 ) {
			mintUniV3( newLow, newHigh, u0, u1 );
            (cLow, cHigh) = (newLow, newHigh);    
		}

		if (compPortion> 0 ) {
			currentBalance0 = getBalance(token0);
			currentBalance1 = getBalance(token1);
			c0 = currentBalance0 * compPortion / 100 ;
			c1 = currentBalance1 * compPortion / 100 ;

	        mintCompound(token0,c0);
	        mintCompound(token1,c1);
		}	

		//if (aavePortion> 0 ) {}
		
		return(u0,u1,c0,c1);
	}
	
	function getBalance(address _token) internal view returns(uint256){
		return (IERC20(_token).balanceOf(address(this)));
	}
	
    function getPrice() public view returns(uint256) {
    	return(getQuoteAtTick(getTwap(), uint128(quoamount), token0, token1) );
    }
    
    function getTwap() public view returns (int24 tick) {
        
        require(twapDuration != 0, 'TWAP');   
        
        uint32[] memory secondsAgo = new uint32[](2);
        secondsAgo[0] = twapDuration;
        secondsAgo[1] = 0;

        (int56[] memory tickCumulatives, ) = pool.observe(secondsAgo);
        int56 tickDelta = tickCumulatives[1] - tickCumulatives[0];
        tick = int24(tickDelta / int56(int32( twapDuration )));
        
        if (tickDelta < 0 && (tickDelta % int56(int32(twapDuration)) != 0)) tick--;
    }

   function getQuoteAtTick(
        int24 tick,
        uint128 baseAmount,
        address baseToken,
        address quoteToken
    ) public pure returns (uint256 quote) {
    	
        uint160 sqrtRatioX96 = TickMath.getSqrtRatioAtTick(tick);

        // Calculate quoamount with better precision if it doesn't overflow when multiplied by itself
        if (sqrtRatioX96 <= type(uint128).max) {
            uint256 ratioX192 = uint256(sqrtRatioX96) * sqrtRatioX96;
            quote = baseToken < quoteToken
                ? FullMath.mulDiv(ratioX192, baseAmount, 1 << 192)
                : FullMath.mulDiv(1 << 192, baseAmount, ratioX192);
        } else {
            uint256 ratioX128 = FullMath.mulDiv(sqrtRatioX96, sqrtRatioX96, 1 << 64);
            quote = baseToken < quoteToken
                ? FullMath.mulDiv(ratioX128, baseAmount, 1 << 128)
                : FullMath.mulDiv(1 << 128, baseAmount, ratioX128);
        }
    }
    
    // calculate withdraw amount requested by vault user
    function vaultWithdraw(address to, uint256 amount0, uint256 amount1) 
    	public 
		onlyVault
    	returns (bool)
	{
		// remove position (alloc fees)
		// take fund from vault
		// distract withdrawal amounts
		// call rebalance. 	
		// send funds to vault
		//return success
		
		bool succ = alloc();
		require(succ, "alloc");
		
      	if (amount0 > 0) IERC20(token0).safeTransfer(vault, amount0);
	    if (amount1 > 0) IERC20(token1).safeTransfer(vault, amount1);	
		
    	_rebalance(0,0);
		
		emit VaultWithdraw( to, amount0,  amount1);

		return true;
		
    }

    function setMaxTwapDeviation(int24 _maxTwapDeviation) external onlyAdmin {
         require(_maxTwapDeviation > 0, "maxTwapDeviation");
        maxTwapDeviation = _maxTwapDeviation;
    }

    function setTwapDuration(uint32 _twapDuration) external onlyAdmin {
         require(_twapDuration > 0, "twapDuration");
        twapDuration = _twapDuration;
    }
    
    function setUniPortionRatio(uint8 ratio) external onlyAdmin {
    	require (ratio <= 100,"ratio");
		uniPortion = ratio;
    }
    
    function setCompPortionRatio(uint8 ratio) external onlyAdmin {
    	require (ratio <= 100,"ratio");
		compPortion = ratio;
    }

	function setprotocol(address _protocol) external onlyAdmin {
        protocol = _protocol;
    }
   

 	function mintUniV3(
        int24 newLow,
        int24 newHigh,
        uint256 amount0,
        uint256 amount1
    ) internal  { 
	  
        (uint160 sqrtRatioX96, , , , , , ) = pool.slot0();
        
        uint128 liquidity =
            LiquidityAmounts.getLiquidityForAmounts(
                sqrtRatioX96,
                TickMath.getSqrtRatioAtTick(newLow),
                TickMath.getSqrtRatioAtTick(newHigh),
                amount0,
                amount1
            );
            
        pool.mint(address(this), newLow, newHigh, liquidity, "");
        emit MintUniV3Liquidity(newLow, newHigh, liquidity);
	}

	

	
	function getUniAmounts(int24 tickLower, int24 tickUpper)
        public
        view
        returns (uint256 amount0, uint256 amount1)
    {
        (uint128 liquidity, , , uint128 tokensOwed0, uint128 tokensOwed1) =
            _position(tickLower, tickUpper);
        (amount0, amount1) = _amountsForLiquidity(tickLower, tickUpper, liquidity);
        amount0 = amount0 + uint256(tokensOwed0);
        amount1 = amount1 + uint256(tokensOwed1);
    }
    
    
    function redeemCEth(uint256 amount, bool redeemType) internal {
		uint r;
        if (redeemType == true) {
            // amount=cETH的数量
            r = _CETH.redeem(amount);
			if (r==0) _wrap(address(this).balance);
        } else {
            // amount=要赎回的ETH的数量
            r = _CETH.redeemUnderlying(amount);
			if (r==0) _wrap(amount);
        }

		
        if (r != 0) {	//something wrong. Ctoken may be stuck in Comp
			if (!isEmergency) {
				revert("Ceth not 0");  
			} else {
				// emergency 
			}
        }	
	}

	function redeemCErc20(address underlying, uint256 amount, bool redeemType) internal {
	
		if (underlying == _WETH) {
			redeemCEth(amount, redeemType);
			return;
		}
		
        uint256 r;
        if (redeemType == true) {
            // amount=归还cERC20的数量
            r = ICErc20(_CTOKEN[underlying]).redeem(amount);
        } else {
            // amount=要赎回的ERC20的数量
            r = ICErc20(_CTOKEN[underlying]).redeemUnderlying(amount);
        }
        
        if (r != 0) {	//something wrong. Ctoken may be stuck in Comp
			if (!isEmergency) {
				revert('Ct');  
			} else {
				// emergency 
			}
        }	
	}	

    /// remove all positions from protocols, if any failed, an emergency maybe required.
    function removePositions() internal  {
	    removeUniswap();
		removeLending();
	}
	
	function removeUniswap() internal {
       
    	if (cLow == 0 && cHigh ==0)  return;
		
       (uint128 liquidity, , , , ) = _position(cLow, cHigh);   // get liquidity by current ticklow, tickhigh
       
 		(uint256 burned0, uint256 burned1, uint256 collect0, uint256 collect1) = _burnAndCollectUnis(cLow, cHigh, liquidity);
    
    	if (liquidity > 0) {
	        uFees0 = collect0 - burned0;
	       	uFees1 = collect1 - burned1;
       }
	   	(liquidity, , , , ) = _position(cLow, cHigh);	// should be 0 otherwise, there is problem
	}

        	
	/// remove lending position; gether fees
	function removeLending() internal {
        (uint256 c0, uint256 c1 ) = getCAmounts();
        if(c0>0) redeemCErc20(token0, c0, true);
        if(c1>0) redeemCErc20(token1, c1, true);
	}
	
	function _position(int24 tickLower, int24 tickUpper)
        internal
        view
        returns ( uint128, uint256, uint256,  uint128, uint128)
    {
        bytes32 positionKey = PositionKey.compute(address(this), tickLower, tickUpper);
        return pool.positions(positionKey);
    }
    

 function _burnAndCollectUnis(
        int24 tickLow,
        int24 tickHigh,
        uint128 liquidity 
        )
        internal
        returns ( uint256 burned0, uint256 burned1, uint256 collect0,uint256 collect1)
    {
		
        if ( liquidity > 0) {
        	( burned0, burned1) =  pool.burn(tickLow, tickHigh, liquidity) ;
        } 
        
         // Collect all owed tokens including earned fees
        (collect0,  collect1) =
            pool.collect(
                address(this),
                tickLow,
                tickHigh,
                type(uint128).max,
                type(uint128).max
            );
    }   

	///@notice calculate best portion to put into uniswap
	function calcUniPortionAmounts(uint256 total0, uint256 total1) internal view returns(uint256 amount0, uint256 amount1){

		uint256 p = getPrice();
		uint256 total0In1 = total0 * p;   // to have same quantity of token0 as token1 upon current price
		
		if (total0In1 > total1)  {    
			// amount of token0 > amount of token 1, then use token0 to calculate the larger part of supply  
			amount0 = total0 * uniPortion / 100;
			amount1 = Math.min( total1 * uniPortion /100, amount0 * p );
		} else {
			// amount of token1 > amount of token 0, then use token1 to calculate the larger part of supply  
			amount1 = total1 * uniPortion / 100;	
			amount0 = Math.min( total0 * uniPortion/100, amount1/p );
		}
		
	}


	function mintCompound(address _underlying, uint256 amount) internal {
        if (_underlying == _WETH ) {
   	        _unwrap(amount);  
			_CETH.mint{gas:250000,value:amount}();
        } else {

			IERC20(_underlying).approve(_CTOKEN[_underlying], amount);
			uint mintResult = ICErc20(_CTOKEN[_underlying]).mint(amount);
			if (mintResult != 0) {
				emit MyLog("mintResult is not 0, there was an error", mintResult);
			}
        }
	}

	function _wrap(uint256 amount) internal {
	    if (amount > 0) {
	        IWETH9(_WETH).deposit{value:amount}();
	    }   
	}

	function _unwrap(uint256 amount) internal {
	    if (amount > 0) {
	        IWETH9(_WETH).withdraw(amount);
	    }
	}   
  

    function _amountsForLiquidity(
        int24 tickLow,
        int24 tickHigh,
        uint128 liquidity
    ) internal view returns (uint256, uint256) {
        (uint160 sqrtRatioX96, , , , , , ) = pool.slot0();
        return
            LiquidityAmounts.getAmountsForLiquidity(
                sqrtRatioX96,
                TickMath.getSqrtRatioAtTick(tickLow),
                TickMath.getSqrtRatioAtTick(tickHigh),
                liquidity
            );
    }
    
    function collectFees() internal {
    	
    }

    function alloc() internal returns ( bool ){
    	
		removePositions();		// get all assets back to vault
		//collectFees();
		return(true);  // fees are allocated
    }

	function getCAmounts() internal view returns (uint256 amount0, uint256 amount1) {
		amount0 = ICErc20(_CTOKEN[token0]).balanceOf( address(this) );
		amount1 = ICErc20(_CTOKEN[token1]).balanceOf( address(this) );
	}

	function getCtokenUnderlyingAmounts() internal returns (uint256 amount0, uint256 amount1) {
		amount0 = ICErc20(_CTOKEN[token0]).balanceOfUnderlying( address(this) );
		amount1 = ICErc20(_CTOKEN[token1]).balanceOfUnderlying( address(this) );
	}
	
	function getTotalAmounts() public returns(uint256 , uint256) {
		(uint256 a0, uint256 a1) = getUniAmounts(cLow, cHigh);
		(uint256 b0, uint256 b1) = getCtokenUnderlyingAmounts();
		(uint256 c0, uint256 c1) = ( IERC20(token0).balanceOf(address(this)),IERC20(token1).balanceOf(address(this)) ); 
		return (a0+b0+c0, a1+b1+c1);
	}

  	function _validRange(int24 _lower, int24 _upper, int24 _tick) internal view {
        require(_lower < _upper, "V1");
        require(_lower < _tick , "V2");
        require(_upper > _tick , "V3");
        require(_lower >= TickMath.MIN_TICK, "V4");
        require(_upper <= TickMath.MAX_TICK, "V5");
        require(_lower % tickSpacing == 0, "V6");
        require(_upper % tickSpacing == 0, "V7");
    }
    
    
}