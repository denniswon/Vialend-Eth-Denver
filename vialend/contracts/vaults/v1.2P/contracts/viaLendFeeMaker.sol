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

	uint256 private Price;
	
    
    constructor(
        address _pool,
        address payable _weth,
        address _cToken0,
        address _cToken1,
        address _cEth,
        uint256 _protocolFee,
        uint256 _maxTotalSupply,
        int24 _maxTwapDeviation,
        uint32 _twapDuration,
		uint8 _uniPortion,
		address _team
        
    ) ERC20("ViaLend Token0","VLT0") {

        governance = msg.sender;
        team = _team;  
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

        protocolFee = _protocolFee;

		tickSpacing = IUniswapV3Pool(_pool).tickSpacing();

        require(_protocolFee < 100, "PFR");

        maxTotalSupply = _maxTotalSupply;

        maxTwapDeviation = _maxTwapDeviation;

        twapDuration = _twapDuration;
        
        uniPortion =  _uniPortion ;

		require(_maxTwapDeviation > 0, "MTD");
		
        require(_twapDuration > 0, "TD");
    }


    
    /// @notice deposit amount will be adjusted to the current ratio
    function deposit(
        uint256 amountToken0,		// deposit amount of token0
        uint256 amountToken1,		// deposit amount of token1
        bool doRebalence			// whether do rebalance or not
    )
        external
        nonReentrant
    {

        address to = msg.sender;	//  to address = msg.sender

		require(amountToken0>0 || amountToken1>0,"amt0");

        require(to != address(0) && to != address(this) && to !=team, "to");

	
        _push(to);  			// push new address into accounts

  		
       // _poke(cLow, cHigh);		// Poke v3 positions so to get uniswap v3 fees up to date. do not need if position removed

		//uint256 P = getUniswapPrice();
		Price = getQuoteAtTick(getTwap(address(pool),twapDuration), uint128(1e18),address(token0), address(token1));		

        (uint256 shares, uint256 amount0, uint256 amount1) = _calcShares(amountToken0, amountToken1);
		

		_alloc();				// dividen fees if any for prev users
		
		updateExists();	

        // transfer tokens from sender
        if (amount0 > 0) token0.safeTransferFrom(msg.sender, address(this), amount0);
        if (amount1 > 0) token1.safeTransferFrom(msg.sender, address(this), amount1);

        _mint(to, shares);

        require(totalSupply() <= maxTotalSupply, "CAP");

	    Assetholder[to].deposit0 += amount0;		// add or increase msg.sender's asset0
		Assetholder[to].deposit1 += amount1;		// add or increase msg.sender's asset1
		Assetholder[to].current0 += amount0;
		Assetholder[to].current1 += amount1;

		if (doRebalence )  rebalance(0,0);		// rebalance

        emit Deposit(msg.sender, to, shares, amount0, amount1);

    }     
    
	
    //update Others share with current price
    function updateExists() internal {
	
			uint256 tt = totalSupply();

			if (tt <= 0 ) return;		// return totalsupply is 0 

			uint cnt = accounts.length;		// get array size 

			uint256 total0 = token0.balanceOf(address(this));
			uint256 total1 = token1.balanceOf(address(this));
		   

			for (uint i=0; i< cnt; i++) {	
				
				uint256 share = balanceOf(accounts[i]);
				
				if (share > 0) {
					
					// get new amount based on new balance of each user's share
					uint256 myamount0 = total0.mul(share).div(tt);
					uint256 myamount1 = total1.mul(share).div(tt);
			        
					
					// calc new shares with new balance of tokens for each user
					(uint256 nshare, , )  = _calcShares(myamount0,myamount1); //myamount0.add(myamount1.mul(1e18).div(P) ) ;
					
					// could be smaller, so use != 
					if (nshare != share) {
					

						// burn old share
						_burn(accounts[i], share);
						
						// mint new share
						_mint(accounts[i], nshare) ;
		
						// update user assets with new amounts
						Assetholder[accounts[i]].current0 = myamount0;		
						Assetholder[accounts[i]].current1 = myamount1;
					}
				}
        	}
    	
    }
    
    /// allocate fees for external
    function alloc() external onlyGovernance {
    	_alloc();
    }

    /// allocate fees and adjust user's share
    function _alloc() internal returns ( bool ){
		

		uint256 _totalSupply = totalSupply();
		
		if (_totalSupply <=0) return false;    // if totalsupply is 0,  there is no asset in vault

		removePositions();		// get all assets back to vault

		if (_totalSupply > 0 ) collectFees();

		
		return(true);  // fees are allocated
		    	
    }
    
    
    //calculate net fees and protocol fees 
	function collectFees() internal 	{
		
		uint256 _protocolFee = protocolFee;
		
		// add up all earned fees
		( uint256 fees0, uint256 fees1) = ( uFees0.add(lFees0),  uFees1.add(lFees1) ) ;

		// log current collected fees		
		emit CollectFees(address(this),uFees0,uFees1,lFees0,lFees1);


		//update total fees
		Fees.U3Fees0 =  Fees.U3Fees0.add(uFees0); 
		Fees.U3Fees1 =  Fees.U3Fees1.add(uFees1); 
		Fees.LcFees0 =  Fees.LcFees0.add(lFees0); 
		Fees.LcFees1 =  Fees.LcFees1.add(lFees1); 

		// clear temp fees variables
		( uFees0,uFees1, lFees0, lFees1) = (0,0,0,0);
		
		uint256 feesToProtocol0;
		uint256 feesToProtocol1;
		
        if (_protocolFee > 0 )  {
        	
			// calculate protocol fees
			if (fees0 > 0 ) {
	            feesToProtocol0 = fees0.mul(_protocolFee).div(100);
	            //netfees0 = fees0.sub(feesToProtocol0);
	            Fees.AccruedProtocolFees0 = Fees.AccruedProtocolFees0.add(feesToProtocol0);
			}
			
			if (fees1 > 0 ) {
	            feesToProtocol1 = fees1.mul(_protocolFee).div(100);
	            // generate net fees
	            //netfees1 = fees1.sub(feesToProtocol1);
	            //record the protocol fees
	            Fees.AccruedProtocolFees1 = Fees.AccruedProtocolFees1.add(feesToProtocol1);
			}
            
            //collect protocol fees to team
            if (feesToProtocol0 > 0 || feesToProtocol1 > 0) {

            	(uint256 teamshare, , ) = _calcShares(feesToProtocol0, feesToProtocol1);

				_mint(team, teamshare );  // mint fees to team
				_push(team);   // push team to accounts 
				
            	//token0.safeTransfer(team, feesToProtocol0);
				//token1.safeTransfer(team, feesToProtocol1);
			}

        } 
 
        
	}



    
    
   
    function withdraw(
        uint256 percent
    ) external  nonReentrant returns (uint256 amount0, uint256 amount1) {
        
        

		address to = msg.sender;
		
        require(to != address(0) && to != address(this), "toW");
        
        require(percent > 0 && percent <= 100, "Pcnt");
        
		_alloc();  //remove positions and dividen fees

        // calc percent shares to withdraw
        uint256 shares = balanceOf(to).mul(percent).div(100);   	
        //uint256 shares = balanceOf(to);

		// get total shares
        uint256 totalSupply = totalSupply();

        require(totalSupply > 0, "ts0");

		require(shares <= totalSupply , "sh1");

		
        amount0 = getBalance0().mul(shares).div(totalSupply);
        amount1 = getBalance1().mul(shares).div(totalSupply);

		// burn user's share
        _burn(to, shares);
        
        // update total fees record
        Fees.U3Fees0 = Fees.U3Fees0.sub(Fees.U3Fees0.mul(shares).div(totalSupply) ) ;
        Fees.U3Fees1 = Fees.U3Fees1.sub(Fees.U3Fees1.mul(shares).div(totalSupply) ) ;
        Fees.LcFees0 = Fees.LcFees0.sub(Fees.LcFees0.mul(shares).div(totalSupply) ) ;
        Fees.LcFees1 = Fees.LcFees1.sub(Fees.LcFees1.mul(shares).div(totalSupply) ) ;

        // if amount > 0 , do transfer and update user Assets record. 
        if (amount0 > 0) {
        	token0.safeTransfer(to, amount0);
			Assetholder[to].current0 -= (Assetholder[to].current0 > amount0 ) ? amount0 : Assetholder[to].current0;	
		}
        if (amount1 > 0) {
        	token1.safeTransfer(to, amount1);
			Assetholder[to].current1 -= (Assetholder[to].current1 > amount1 ) ? amount1 : Assetholder[to].current1;
		}
		
		// clean up assets record if user's share is 0
		if ( balanceOf(to) == 0 ) {
			Assetholder[to].deposit0 = 0;
			Assetholder[to].deposit1 = 0;
			Assetholder[to].current0 = 0;
			Assetholder[to].current1 = 0;
		}
			
		//log
        emit Withdraw(msg.sender, shares, amount0, amount1);
    }
    
    // commented because currently using remove all
	// function _burnLendingShare (uint256 shares, uint256 totalShares) internal returns(uint256,uint256) {

 //      	(uint256 cAmount0, uint256 cAmount1) = getCAmounts();
		
		

	// 	uint256 myCamount0 = cAmount0.mul(shares).div(totalShares);
	// 	uint256 myCamount1 = cAmount1.mul(shares).div(totalShares);

	// 	removeLending(address(token0), address(token1),myCamount0,myCamount1);

	// 	uint32 decimal0 = 18; // to do
	// 	uint32 decimal1 = 6;  // to do 
	// 	//oneCTokenInUnderlying = exchangeRateCurrent / (1 * 10 ^ (18 + underlyingDecimals - cTokenDecimals))

	// 	//#debug #review
 //        uint256 myamount0 = myCamount0.mul(CToken0.exchangeRateCurrent()).div(10**decimal0) ;
 //        uint256 myamount1 = myCamount1.mul(CToken1.exchangeRateCurrent()).div(10** (18 + decimal1 - 8) ) ;

		
	// 	emit MyLog2("_burnLendingShare ", myamount0,myamount1);
		
	// 	return (myamount0,myamount1);
	// }


    function _calcShares(uint256 amountIn0, uint256 amountIn1)
        internal
        view
        returns (
            uint256 shares,
            uint256 amount0,
            uint256 amount1
            
        )
    {

			// mint the new deposit 
			amount0 = amountIn0;
			amount1 = amountIn1;

			require (Price > 0, "P");
			
			// based on previous price, only new deposit will need to fetch new price 
			shares = amount0 + (amount1 * 1e18 ) / Price;   
			
			
			
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
		

		(uint256 camount0, uint256 camount1) = getCAmounts();
		
		if (camount0 > 0 || camount1 > 0) {
			removeLending( address(token0), address(token1), camount0,camount1 );
		}
		
		removeUniswap();
	}
	
	

	function strategy1(
		int24 newLow,
        int24 newHigh
		) external nonReentrant onlyGovernance  {
		
		
		require(totalSupply() > 0,"t0");
        
        _alloc();
        
        rebalance(newLow,newHigh);
        		
	}
	
	// make sure all position has been removed before doing rebalance. 
	// when newLow and newHigh is 0, calc new range with current cLow and cHigh
	function rebalance(
			int24 newLow,
			int24 newHigh
			) internal  {
	
      	(	,int24 tick, , , , , ) 	= pool.slot0();
    
    	if (newLow==0 && newHigh==0) {
			
			if (cHigh == 0 && cLow ==0) return;  // cannot do rebalance if cLow and cHigh is 0
			
			int24 hRange = ( cHigh - cLow ) / 2;
			
			newHigh = (tick + hRange) / tickSpacing * tickSpacing;
			newLow  = (tick - hRange) / tickSpacing * tickSpacing;
			
		}

  		lib._validRange(newLow, newHigh, tick, tickSpacing);  // passed 1200 , 2100, 18382
        
        // Check price is not too close to min/max allowed by Uniswap. Price
        // shouldn't be this extreme unless something was wrong with the pool.

        //int24 range = newHigh - newLow ;
            
        // int24 twap = getTwap();
        // int24 deviation = tick > twap ? tick - twap : twap - tick;
        
        // // avoid high slipage
        // require(deviation <= maxTwapDeviation, "deviation");


        //#review todo   calc best ratio to 50:50 in uni v3
        uint256 uniPortion0 =  getBalance0().mul(uniPortion).div(100);
        uint256 uniPortion1 = getBalance1().mul(uniPortion).div(100);

		//add rest portion to uniswap
		_mintUniV3(
	        newLow,
	        newHigh,
			uniPortion0,
			uniPortion1
			);


		// get remainting assets to lending
        uint256 unUsedbalance0 = getBalance0();
        uint256 unUsedbalance1 = getBalance1();
        
        bool result = _mintCompound(address(token0), address(token1), unUsedbalance0,  unUsedbalance1);
        
        require(result, "MCp");


        // lastRebalance = block.timestamp;
        // lastTick = tick;

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
    
 
	
    /// @notice return Balance of available token0.
     
    function getBalance0() internal view returns (uint256) {
        return token0.balanceOf(address(this));
    }


    /// @notice return Balance of available token1.
    
    function getBalance1() internal view returns (uint256) {
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
 

    
	// remove all positions and send fund to each user
	// send any left over to governance
    function emergencyBurn() external onlyGovernance {
			
		removePositions();	// remove all positions from uniswap and lending pool
		
		removeCTokens(); // in case assets stuck in compound


		//should remove the code below after testing

		uint cnt = accounts.length;		// get array size 
		
		uint256 total0 = token0.balanceOf(address(this));
		uint256 total1 = token1.balanceOf(address(this));
        
		uint256 tt = totalSupply();
		
		for (uint i=0; i< cnt; i++) {	
			
			uint256 share = balanceOf(accounts[i]);
			
			_burn(accounts[i], share);

		    Assetholder[accounts[i]].deposit0 = 0;
			Assetholder[accounts[i]].deposit1 = 0;
			Assetholder[accounts[i]].current0 = 0;
			Assetholder[accounts[i]].current1 = 0;

		    
			uint256 amount0 = total0.mul( share ).div(tt);
			uint256 amount1 = total1.mul( share ).div(tt);
			

	        if (amount0 > 0) token0.safeTransfer(accounts[i], amount0);
	        if (amount1 > 0) token1.safeTransfer(accounts[i], amount1);
			
			
        }
        

		//send rest tokens to governance
		if ( token0.balanceOf(address(this)) > 0 ) 
			token0.safeTransfer(governance, token0.balanceOf(address(this))) ;

		if ( token1.balanceOf(address(this)) > 0 ) 
			token1.safeTransfer(governance, token1.balanceOf(address(this))) ;

		

    }

/*    due to size
 	function sweep(
        IERC20 token,
        uint256 amount
    ) external onlyGovernance {
        require(token != token0 && token != token1, "token");
        token.safeTransfer(msg.sender, amount);
    }	
*/

	
    	
	/// fallback function has been split into receive() and fallback(). It is a new change of the compiler.
	fallback() external payable {}
	receive() external payable {}
}