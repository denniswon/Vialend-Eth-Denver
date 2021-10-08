// SPDX-License-Identifier: MIT
/*
personal deposit cap

Contract code size exceeds 24576 bytes 
(a limit introduced in Spurious Dragon). 
This contract may not be deployable on mainnet. 
Consider enabling the optimizer (with a low "runs" value!), 
turning off revert strings, or using libraries.

where the share used:
_burnLiquidityShare(cLow, cHigh, shares, totalSupply);
_burnLendingShare(shares, totalSupply);
unusedbalance (share, totalsupply)
withdraw
deposit

*/
pragma solidity >=0.5.0;

import "@openzeppelin/contracts/math/Math.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@uniswap/v3-core/contracts/interfaces/callback/IUniswapV3MintCallback.sol";
import "@uniswap/v3-core/contracts/interfaces/callback/IUniswapV3SwapCallback.sol";
import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";
import "@uniswap/v3-core/contracts/libraries/TickMath.sol";
import "@uniswap/v3-periphery/contracts/libraries/LiquidityAmounts.sol";
import "@uniswap/v3-periphery/contracts/libraries/PositionKey.sol";

//import "libraries/libVialendCompound.sol";
import "libraries/lib.sol";

import "interfaces/IFeeMakerEvents.sol";

import "./ownable.sol";
import "./viaCompound.sol";
import "./viaUniswap.sol";


/// @author  ViaLend FeeMaker
/// @title   ViaLend FeeMaker
/// @notice  A Smart Contract that helps liquidity providers managing their funds on Uniswap V3 .

contract ViaLendFeeMaker is 
	Ownable,
	ViaCompound,
	ViaUniswap,
    ERC20,
    IUniswapV3MintCallback,
    IUniswapV3SwapCallback
    
{
	
    using SafeERC20 for IERC20;
    using SafeMath for uint256;
	

  
	IERC20 public immutable token0;
    IERC20 public immutable token1;
    
    uint256 prev0;
    uint256 prev1;

	
    /// @dev 
    /// @param _pool Uniswap V3 pool address
    /// @param _protocolFeeRate Protocol fee expressed as multiple of 1e-6
    /// @param _maxTotalSupply Cap on total supply
    
    constructor(
        address _pool,
        address payable _weth,
        address _cToken0,
        address _cToken1,
        address _cEth,
        uint256 _protocolFeeRate,
        uint256 _maxTotalSupply,
        int24 _maxTwapDeviation,
        uint32 _twapDuration,
		uint8 _uniPortionRate
        
    ) ERC20("ViaLend Token0","VLT0") {

        governance = msg.sender;
        team = msg.sender;  
        // temporary team and governance are the same

		pool = IUniswapV3Pool(_pool);	
		
        token0 = IERC20(IUniswapV3Pool(_pool).token0());
        
        token1 = IERC20(IUniswapV3Pool(_pool).token1());

		//require (address(token0) == _weth || address(token1) == _weth)
		
        CToken0 = IcErc20(_cToken0);
        CToken1 = IcErc20(_cToken1);
        CEther = IcEther(_cEth);

        
        require(_weth != address(0), "WETH");
	
	    WETH = IWETH9(_weth);

        protocolFeeRate = _protocolFeeRate;
        

		tickSpacing = IUniswapV3Pool(_pool).tickSpacing();

        require(_protocolFeeRate < 1e6, "PFR");

        maxTotalSupply = _maxTotalSupply;

        maxTwapDeviation = _maxTwapDeviation;

        twapDuration = _twapDuration;
        
        uniPortionRate =  _uniPortionRate ;

		require(_maxTwapDeviation > 0, "MTD");
		
        require(_twapDuration > 0, "TD");
    }


    
 /// - Deposit2  deposit + strategy1
    /// @notice tokens get deposited in this smart contract will be held until next rebalance. 
    /// @param amountToken0 amount of token0 to deposit
    /// @param amountToken0 amount of token1 to deposit
    function deposit(
        uint256 amountToken0,
        uint256 amountToken1
    )
        external
        
        nonReentrant
        
        returns (
            uint256 shares,
            uint256 amount0,
            uint256 amount1
        )
    {
    	

        address to = msg.sender;

		require(amountToken0>0 || amountToken1>0,"amt0");

        require(to != address(0) && to != address(this), "to");

	    Assetholder[to].capital0 += int(amountToken0);
		Assetholder[to].capital1 += int(amountToken1);
//		Assetholder[to].deposittime.push(block.timestamp);

        push(to);

  		// Poke positions so to get uniswap v3 fees up to date. 
        //_poke(cLow, cHigh);
		
		//to be optimized
		//removePositions();
		//reCalcShares();
		

		//now calc the sender's share
        (shares, amount0, amount1) = _calcShares(amountToken0, amountToken1); 
		

        // transfer tokens from sender
        if (amount0 > 0) token0.safeTransferFrom(msg.sender, address(this), amount0);
        if (amount1 > 0) token1.safeTransferFrom(msg.sender, address(this), amount1);

        _mint(to, shares);

        emit Deposit(msg.sender, to, shares, amount0, amount1);

        require(totalSupply() <= maxTotalSupply, "CAP");


    }     
    
    /// recalcShares and collect protocol fees
    function reCalcShares() internal {

		return ;// to do 
		
		uint256 new0 = token0.balanceOf(address(this));
		uint256 new1 = token1.balanceOf(address(this));
		
		if (prev0 <= 0 ) prev0 = new0;
		if (prev1 <= 0 ) prev1 = new1;
		
		(uint256 fees0 , uint256 fees1 )  = ( new0 - prev0 ,  new1 - prev1 ) ;
		
		emit MyLog2("fees",fees0,fees1);

		if (fees0 > 0 || fees1 > 0) {
			
			(uint256 netfees0, uint256 netfees1) = withdrawProtocolFees(fees0, fees1);
			
			uint cnt = accounts.length;
			uint256 _totalSupply = totalSupply();
			
			if (_totalSupply > 0 ) {
				// mint fees for each holder
				for (uint i=0; i< cnt; i++) {
					uint256 _share = balanceOf(accounts[i]);
					
					require(accounts[i]!=address(0) && accounts[i]!=address(this),"A0");
					
					uint256 myfees0 = netfees0.mul(_share).div(_totalSupply);
					uint256 myfees1 = netfees1.mul(_share).div(_totalSupply);
					
					(uint256 feeshares, , ) = _calcShares(myfees0,myfees1); 
					
					_mint(accounts[i], feeshares);
					
					//record the fees in each account's total
					Assetholder[accounts[i]].fees0 += myfees0;
					Assetholder[accounts[i]].fees1 += myfees1;
				}
			} 
		}
		    	
		
    }
    
	function withdrawProtocolFees(uint256 fees0, uint256 fees1) internal returns(uint256 netfees0, uint256 netfees1){
		
		uint256 _protocolFeeRate = protocolFeeRate;
        if (_protocolFeeRate > 0) {
        	
            uint256 feesToProtocol0 = fees0.mul(_protocolFeeRate).div(1e6);
            uint256 feesToProtocol1 = fees1.mul(_protocolFeeRate).div(1e6);
            
            netfees0 = fees0.sub(feesToProtocol0);
            netfees1 = fees1.sub(feesToProtocol1);
            
            accruedProtocolFees0 = accruedProtocolFees0.add(feesToProtocol0);
            accruedProtocolFees1 = accruedProtocolFees1.add(feesToProtocol1);
            
            if (feesToProtocol0 > 0) token0.safeTransfer(team, feesToProtocol0);
        	if (feesToProtocol1 > 0) token1.safeTransfer(team, feesToProtocol1);

        } else {
	        netfees0 = fees0;
            netfees1 = fees1;

        }
        
	}



    
    
    /// - staker Withdraw 
    /// @param percent number of percentage of Shares owned by sender to be burned
    /// @return amount0 Amount of token0 sent to staker 
    /// @return amount1 Amount of token1 sent to staker
    
   
    function withdraw(
        uint256 percent
    ) external  nonReentrant returns (uint256 amount0, uint256 amount1) {
        
        
        // shares <=  cantor_pair_calculate(deposit0,deposit1) 
        // totalShares <= mint(user, shares) 
        //
        //(deposit0,deposit1) <= cantor_pair_reverse(shares);
        // (storedtotal0, storedtotal1) <= cantor_pair_reverse(totalShares);
        // currentTotal0 , currentTotal1 = token0.balanceOf(vault), token1.balanceOf(vault), 
        // withdrawAmount0 = currentTotal0*deposit0/storedtotal0 
        
        //percent = 100;  //for now

		address to = msg.sender;
		
        require(to != address(0) && to != address(this), "to2");
        
        require(percent > 0 && percent <= 100, "pcent");
        
        uint256 shares = balanceOf(to).mul(percent).div(100);

        uint256 totalSupply = totalSupply();


        require(totalSupply > 0, "ts0");

		require(shares <= totalSupply , "sh1");
        

		//(uint256 share0, uint256 share1) = lib._fetchAmounts(shares);


        _burn(to, shares);

       

        // Withdraw proportion of liquidity directly from Uniswap pool
        (uint256 poolamount0, uint256 poolamount1) =
            _burnLiquidityShare(cLow, cHigh, shares, totalSupply);

        (uint256 lendingamount0, uint256 lendingamount1) =
            _burnLendingShare(shares, totalSupply);

        // Calculate token amounts proportional to unused balances
        //amount0 = getBalance0().mul(share0).div(100);
        //amount1 = getBalance1().mul(share1).div(100);

        amount0 = getBalance0().mul(shares).div(totalSupply);
        amount1 = getBalance1().mul(shares).div(totalSupply);

       
        // Push tokens to recipient
        if (amount0 > 0) {
        	token0.safeTransfer(to, amount0);
			Assetholder[to].capital0 -= int(amount0);
		}
        if (amount1 > 0) {
        	token1.safeTransfer(to, amount1);
			Assetholder[to].capital1 -= int(amount1);
		}


        emit Withdraw(msg.sender, to, shares, amount0, amount1, poolamount0,poolamount1,lendingamount0,lendingamount1);
    }
    
    
	function _burnLendingShare (uint256 shares, uint256 totalShares) internal returns(uint256,uint256) {

      	(uint256 cAmount0, uint256 cAmount1) = getCAmounts();

		uint256 myCamount0 = cAmount0.mul(shares).div(totalShares);
		uint256 myCamount1 = cAmount1.mul(shares).div(totalShares);

		removeLending(address(token0), address(token1),myCamount0,myCamount1);

		uint32 decimal0 = 18; // to do
		uint32 decimal1 = 6;  // to do 
		//oneCTokenInUnderlying = exchangeRateCurrent / (1 * 10 ^ (18 + underlyingDecimals - cTokenDecimals))

		//#debug #review
        uint256 myamount0 = myCamount0.mul(CToken0.exchangeRateCurrent()).div(10**decimal0) ;
        uint256 myamount1 = myCamount1.mul(CToken1.exchangeRateCurrent()).div(10** (18 + decimal1 - 8) ) ;

		
		emit MyLog2("_burnLendingShare ", myamount0,myamount1);
		
		return (myamount0,myamount1);
	}
	
    function _calcShares(uint256 amountToken0, uint256 amountToken1)
        internal
        pure
        returns (
            uint256 shares,
            uint256 amount0,
            uint256 amount1
        )
    {

		
        (amount0, amount1 ) = (amountToken0, amountToken1);

		shares = lib._sqrt (amountToken0.mul(amountToken1) );
		//shares = lib.cantor_pair_calculate(amount0,amount1); 

        
    }    

    
 /*// COMMENT OUT BECAUSE CODE SIZE EXCEEDS DEPLOYER LIMIT. MOVE TO FRONTEND

    function getTVL() public view returns (uint256 total0, uint256 total1) {
         
        (uint256 uniliq0, uint256 uniliq1) =  getPositionAmounts(cLow, cHigh);
        (uint256 lending0, uint256 lending1) =  getLendingAmounts();

		// balance remaining + liquidity + lending supply
        total0 = getBalance0().add(uniliq0).add(lending0);
        total1 = getBalance1().add(uniliq1).add(lending1);

    }

 
    function getLendingAmounts() public view returns(uint256 , uint256 ){

    	(uint256 cAmount0, uint256 cAmount1) = getCAmounts();
		
		
        //uint8 decimals0 = EIP20Interface(CToken0.underlying()).decimals();
        
        //uint8 decimals1 = EIP20Interface(CToken1.underlying()).decimals();
            
		//require(token0.decimals() ==       underlyingDecimals = EIP20Interface(cErc20.underlying()).decimals();
       //oneCTokenInUnderlying  = exchangeRateCurrent / (1 * 10 ** (18 + underlyingDecimals - cTokenDecimals))
        uint256 amount0 = cAmount0.mul(CToken0.exchangeRateStored().div(10 ** (18 + 18 - 8))  );
        uint256 amount1 = cAmount1.mul(CToken1.exchangeRateStored().div(10 ** (18 + 6 - 8) ) );

		return(amount0,amount1);
    }
    
*/    
	function removePositions() internal {
		

		(uint256 amount0, uint256 amount1) = getCAmounts();
		
		removeLending( address(token0), address(token1),amount0, amount1);
		
		removeUniswap();
	}


	function removeUniswap() internal {
        
        // Withdraw all current liquidity from Uniswap pool
      
       (uint128 allLiquidity, , , , ) = _position(cLow, cHigh); 
        
  		_burnAndCollectUnis(cLow, cHigh, allLiquidity);

	}
	
	

	function strategy0(
		int24 newLow,
        int24 newHigh,
        int256 swapAmount,
        bool zeroForOne
        
		) external nonReentrant onlyTeam  {
		
        		//moved due to  Contract code size exceeds 24576 bytes (a limit introduced in Spurious Dragon). This contract may not be deployable on mainnet. Consider enabling the optimizer (with a low "runs" value!), turning off revert strings, or using libraries.
	}
	

	function strategy1(
		int24 newLow,
        int24 newHigh
		) external nonReentrant onlyTeam  {
		
		
		require(totalSupply() > 0,"t0");
		
        (	,int24 tick, , , , , ) 	= pool.slot0();

  		_validRange(newLow, newHigh, tick);  // passed 1200 , 2100, 18382
        
        // Check price is not too close to min/max allowed by Uniswap. Price
        // shouldn't be this extreme unless something was wrong with the pool.

        //int24 range = newHigh - newLow ;
            
        // int24 twap = getTwap();
        // int24 deviation = tick > twap ? tick - twap : twap - tick;
        
        // // avoid high slipage
        // require(deviation <= maxTwapDeviation, "deviation");

		prev0 = token0.balanceOf(address(this));
		prev1 = token1.balanceOf(address(this));

        // remove positions from uniswap and lending pool get back to vault
        removePositions();
        
        reCalcShares();
        // rebalance, 90% lending, 10% liquidity


        //add xx% assets to uniswap
        uint256 uniPortion0 =  getBalance0().mul(uniPortionRate).div(100);
        uint256 uniPortion1 = getBalance1().mul(uniPortionRate).div(100);

		//add rest portion to uniswap
		rebalance(
	        newLow,
	        newHigh,
			uniPortion0,
			uniPortion1
			);


		// get remainting assets to lending
        uint256 unUsedbalance0 = getBalance0();
        uint256 unUsedbalance1 = getBalance1();
        
        bool result = lendingSupply(address(token0), address(token1), unUsedbalance0,  unUsedbalance1);
        
         require(result, "LDf");
       


        lastRebalance = block.timestamp;
        lastTick = tick;
        		
	}

	

 /// @dev Callback for Uniswap V3 pool
    function uniswapV3MintCallback(
        uint256 amount0,
        uint256 amount1,
        bytes calldata data
    ) external override {
         require(msg.sender == address(pool),"PS");
        if (amount0 > 0) token0.safeTransfer(msg.sender, amount0);
        if (amount1 > 0) token1.safeTransfer(msg.sender, amount1);
    }

    /// @dev Callback for Uniswap V3 pool
    function uniswapV3SwapCallback(
        int256 amount0Delta,
        int256 amount1Delta,
        bytes calldata data
    ) external override {
         
        require(msg.sender == address(pool),"PS2");
        

        ////#debug require(false, _hint2("deltaamount", uint256(amount0Delta),uint256(amount1Delta),0,""));
        
        if (amount0Delta > 0) token0.safeTransfer(msg.sender, uint256(amount0Delta));
        if (amount1Delta > 0) token1.safeTransfer(msg.sender, uint256(amount1Delta));
    }
    
     /// @notice Used to collect accumulated protocol fees.
    function collectProtocol(
        uint256 amount0,
        uint256 amount1,
        address to
    ) external onlyGovernance {
    	
    	require (accruedProtocolFees0 >= amount0 && accruedProtocolFees1 >= amount1,"CP");

		if (amount0 > 0) {
	        accruedProtocolFees0 = accruedProtocolFees0.sub(amount0);
	        token0.safeTransfer(to, amount0);
		}
							
        if (amount1 > 0) {
        	accruedProtocolFees1 = accruedProtocolFees1.sub(amount1);
        	token1.safeTransfer(to, amount1);
        }
    }
    
	///calculate the minimum amount for token0 and token1 to deposit
	/// todo
	// function amountMin(uint256 amount0, uint256 amount1) internal pure returns (bool){
	// 	return true; 
		
	// }
    /// @notice return Balance of available token0.
     
    function getBalance0() public view returns (uint256) {
        return token0.balanceOf(address(this));
    }


    /// @notice return Balance of available token1.
    
    function getBalance1() public view returns (uint256) {
        return token1.balanceOf(address(this));
    }   

/* CODE SIZE
    /// @notice Removes tokens accidentally sent to this vault.
    function sweep(
        IERC20 token,
        uint256 amount,
        address to
    ) external onlyGovernance {
        require(token != token0 && token != token1, "Sw");
        token.safeTransfer(to, amount);
    }

*/
     
	/**
     * @notice low lever Removes positions for emergency. 
    */
    function emergencyBurn() external onlyGovernance {
			removePositions();

		uint cnt = accounts.length;
		
		uint256 total0 = token0.balanceOf(address(this));
		uint256 total1 = token1.balanceOf(address(this));
        
		uint256 tt = totalSupply();
		
		
		for (uint i=0; i< cnt; i++) {
			
			uint256 share = balanceOf(accounts[i]);
	        
			_burn(accounts[i], share);

			
			uint256 amount0 = total0.mul( share ).div(tt);
			uint256 amount1 = total1.mul( share ).div(tt);
			

	        // if (amount0 > 0) token0.safeTransfer(accounts[i], amount0);
	        // if (amount1 > 0) token1.safeTransfer(accounts[i], amount1);
	        
	        if (amount0 > 0) {
	        	token0.safeTransfer(accounts[i], amount0);
				Assetholder[accounts[i]].capital0 -= int(amount0);
			}
	        if (amount1 > 0) {
	        	token1.safeTransfer(accounts[i], amount1);
				Assetholder[accounts[i]].capital1 -= int(amount1);
			}
        }
        
		//#review
		token0.safeTransfer(governance, token0.balanceOf(address(this))) ;
		token1.safeTransfer(governance, token1.balanceOf(address(this))) ;

		
    }
    
	
	/// fallback function has been split into receive() and fallback(). It is a new change of the compiler.
	fallback() external payable {}
	receive() external payable {}
}