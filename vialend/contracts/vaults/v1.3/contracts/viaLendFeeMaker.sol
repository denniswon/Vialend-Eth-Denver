// SPDX-License-Identifier: MIT
/*
todo: personal deposit cap

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

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@uniswap/v3-core/contracts/interfaces/callback/IUniswapV3MintCallback.sol";
import "@uniswap/v3-core/contracts/interfaces/callback/IUniswapV3SwapCallback.sol";
import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";
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
    
	uint256 private Price;		//uniswap v3 oracle price twap based
	
    
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
        uint128 _quoteAmount,
		uint8 _uniPortion,
		address _team
        
    ) ERC20("ViaLend Token0","VLT0") {

        governance = msg.sender;
        team = _team;  
       
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
        
        quoteAmount = _quoteAmount;
        
        uniPortion =  _uniPortion ;

		require(_maxTwapDeviation > 0, "MTD");
		
        require(_twapDuration > 0, "TD");
    }


    
    /// @notice deposit amount will be adjusted to the current ratio
    function deposit(
        uint256 amountToken0,		// deposit amount of token0
        uint256 amountToken1,		// deposit amount of token1
        bool doRebalance			// whether do rebalance or not
    )
        external
        nonReentrant
    {

        address to = msg.sender;	//  to address = msg.sender

		require(amountToken0>0 || amountToken1>0,"amt0");

        require(to != address(0) && to != address(this) && to !=team, "to");

	
        _push(to);  			// push new address into accounts

  		
       // _poke(cLow, cHigh);		// Poke v3 positions so to get uniswap v3 fees up to date. do not need if position removed

		//uint256 P = getSpotPrice();
		Price = getQuoteAtTick(getTwap(address(pool),twapDuration), uint128(quoteAmount),address(token0), address(token1));		

        (uint256 shares, uint256 amount0, uint256 amount1) = _calcShares(amountToken0, amountToken1);
		

		_alloc();				// dividen fees if any for prev users
		
		updateShares();	

        // transfer tokens from sender
        if (amount0 > 0) token0.safeTransferFrom(msg.sender, address(this), amount0);
        if (amount1 > 0) token1.safeTransferFrom(msg.sender, address(this), amount1);

        _mint(to, shares);

        require(totalSupply() <= maxTotalSupply, "CAP");

	    Assetholder[to].deposit0 += amount0;		// add or increase msg.sender's asset0
		Assetholder[to].deposit1 += amount1;		// add or increase msg.sender's asset1
		Assetholder[to].current0 += amount0;
		Assetholder[to].current1 += amount1;

		if (doRebalance )  rebalance(0,0, totalSupply());		// rebalance

        emit Deposit(msg.sender, to, shares, amount0, amount1);

    }     
    
	
    //update current users' share with current price
    function updateShares() internal {
	
			uint256 tt = totalSupply();

			if (tt <= 0 ) return;		// return totalsupply is 0 

			uint cnt = accounts.length;		// get array size 

			uint256 total0 = token0.balanceOf(address(this));
			uint256 total1 = token1.balanceOf(address(this));
		   

			for (uint i=0; i< cnt; i++) {	
				
				uint256 share = balanceOf(accounts[i]);  //user's share
				
				if (share > 0 ) {
					
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
		
		// add up all earned fees from univ3 and lending pool
		( uint256 fees0, uint256 fees1) = ( uFees0.add(lFees0),  uFees1.add(lFees1) ) ;

		// log current collected fees		
		emit CollectFees(address(this),uFees0,uFees1,lFees0,lFees1);


		//update fees record
		Fees.U3Fees0 =  Fees.U3Fees0.add(uFees0); 
		Fees.U3Fees1 =  Fees.U3Fees1.add(uFees1); 
		Fees.LcFees0 =  Fees.LcFees0.add(lFees0); 
		Fees.LcFees1 =  Fees.LcFees1.add(lFees1); 

		// clear temp variables
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
				_push(team);   // push team to accounts for the first time.
				
            	//token0.safeTransfer(team, feesToProtocol0);
				//token1.safeTransfer(team, feesToProtocol1);
			}

        } 
 
        
	}
    
   // percent is the percentage of range is 0 - 100 
    function withdraw(
        uint8 percent
    ) external  nonReentrant returns (uint256 amount0, uint256 amount1) {
        
		address to = msg.sender;
		
		require ( percent <= 100, "pc");
		
        require(to != address(0) && to != address(this), "toW");
        
		// get total shares
        uint256 totalSupply = totalSupply();

        // calc percent shares to withdraw
        uint256 shares = balanceOf(to).mul(percent).div(100);   	

		require(shares <= totalSupply , "sh1");
   
		_alloc();  //remove positions and dividen fees
		
        amount0 = getBalance0().mul(shares).div(totalSupply);
        amount1 = getBalance1().mul(shares).div(totalSupply);

		// burn user's share
        _burn(to, shares);
        
        // deduct withdrawal fees from the total fees. for frontend view and calc only
        Fees.U3Fees0 = Fees.U3Fees0.sub(Fees.U3Fees0.mul(shares).div(totalSupply) ) ;
        Fees.U3Fees1 = Fees.U3Fees1.sub(Fees.U3Fees1.mul(shares).div(totalSupply) ) ;
        Fees.LcFees0 = Fees.LcFees0.sub(Fees.LcFees0.mul(shares).div(totalSupply) ) ;
        Fees.LcFees1 = Fees.LcFees1.sub(Fees.LcFees1.mul(shares).div(totalSupply) ) ;

        // if amount > 0 , do transfer and update user Assets record. 
        if (amount0 > 0) {
        	token0.safeTransfer(to, amount0);
			
			// update assetsholder record
			Assetholder[to].current0 -= (Assetholder[to].current0 > amount0 ) ? amount0 : Assetholder[to].current0;	
		}
        if (amount1 > 0) {
        	token1.safeTransfer(to, amount1);
			
			// update assetsholder record
			Assetholder[to].current1 -= (Assetholder[to].current1 > amount1 ) ? amount1 : Assetholder[to].current1;
		}
		
		// clean up assets record if user's share is 0
		if ( balanceOf(to) == 0 ) {

			_delme(to);		

			Assetholder[to].deposit0 = 0;
			Assetholder[to].deposit1 = 0;
			Assetholder[to].current0 = 0;
			Assetholder[to].current1 = 0;
		}
		
		rebalance(0,0, totalSupply)	;
		
		//log
        emit Withdraw(msg.sender, shares, amount0, amount1);
    }
    
   
	// remove the address from accounts array
	// Move the last element to the deleted spot.
	// Remove the last element.
	function _delme(address _addr ) internal {


		//require( accId[_addr] > 0, "invalid address");
		if ( !_exist( _addr ) ) return;


		uint index = accId[_addr]-1;

        accounts[index] = accounts[accounts.length - 1];	// move last itme to the delete one.
        
        accId[accounts[index]] = index + 1;		// update accId

        accounts.pop();  // delete last item in account array
        
		delete accId[_addr]; // delete the address mapping 
	
	}



	// calculate user's share
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
		
		uint256 _totalSupply = totalSupply();
		
		require(_totalSupply > 0,"t0");
        
        _alloc();
        
        rebalance(newLow,newHigh,_totalSupply);
        		
	}
	
	// make sure all position has been removed before doing rebalance. 
	// when newLow and newHigh is 0, calc new range with current cLow and cHigh
	function rebalance(
			int24 newLow,
			int24 newHigh,
			uint256 total
			) internal  {
	
      	(	,int24 tick, , , , , ) 	= pool.slot0();
    
    	
		if ( total <= 0 ) return;
    	
    	if (newLow==0 && newHigh==0) {
			
			if (cHigh == 0 && cLow ==0) return;  // cannot do rebalance if cLow and cHigh is 0
			
			int24 hRange = ( cHigh - cLow ) / 2;
			
			newHigh = (tick + hRange) / tickSpacing * tickSpacing;
			newLow  = (tick - hRange) / tickSpacing * tickSpacing;
			
		}

  		lib._validRange(newLow, newHigh, tick, tickSpacing);  // passed 1200 , 2100, 18382
        
            
        // int24 twap = getTwap();
        // int24 deviation = tick > twap ? tick - twap : twap - tick;
        
        // // avoid high slipage
        // require(deviation <= maxTwapDeviation, "deviation");


        //#review todo   calc best ratio to 50:50 in uni v3
        
        if ( uniPortion > 0 ) {
        	
	        uint256 uniPortion0 =  getBalance0().mul(uniPortion).div(100);
	        uint256 uniPortion1 = getBalance1().mul(uniPortion).div(100);
	
			//add rest portion to uniswap
			_mintUniV3(
		        newLow,
		        newHigh,
				uniPortion0,
				uniPortion1
				);
	
		}

        if ( uniPortion < 100 ) {   // when uniPortion==100 means no assets go to lending pool
			// get remainting assets to lending
	        uint256 unUsedbalance0 = getBalance0();
	        uint256 unUsedbalance1 = getBalance1();
	        
	        bool result = _mintCompound(address(token0), address(token1), unUsedbalance0,  unUsedbalance1);
	        
	        require(result, "MCp");
		}


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
		
		// check if there is stuck cTokens		
		(uint256 cAmount0, uint256 cAmount1) = getCAmounts();
		bool checkCtoken = ( cAmount0>0 || cAmount1 >0) ;

		//loop user's account array to withdraw 
		uint cnt = accounts.length;		// get array size 
		
		uint256 total0 = token0.balanceOf(address(this));
		uint256 total1 = token1.balanceOf(address(this));
        
		uint256 tt = totalSupply();
		
		for (uint i=0; i< cnt; i++) {	
			
			uint256 share = balanceOf(accounts[i]);
			
			if ( checkCtoken ) {	// withdraw cToken first if there is any locked
				unmintCtoken(accounts[i],share, tt, cAmount0, cAmount1 );   // withdraw stuck ctokens
			}
			
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


	// send ctoken to user
	function unmintCtoken(address to, uint256 share, uint256 tt, uint256 cAmount0, uint256 cAmount1) internal {
			
		uint256 amount0 = cAmount0.mul( share ).div(tt);  // calc ctoken0 amount belong to user by share
		uint256 amount1 = cAmount1.mul( share ).div(tt);  // calc ctoken1 amount belong to user by share

        if (amount0 > 0) withdrawCTokens(to, address(CToken0), amount0);  // if >0, send ctoken0 to user
        if (amount1 > 0) withdrawCTokens(to, address(CToken1), amount1); // if >0, send ctoken1 to user
		
	}

		
 	function sweep(
        IERC20 token,
        uint256 amount
    ) external onlyGovernance {
        require(token != token0 && token != token1, "token");
        token.safeTransfer(msg.sender, amount);
    }	

	/// fallback function has been split into receive() and fallback(). It is a new change of the compiler.
	fallback() external payable {}
	receive() external payable {}
}