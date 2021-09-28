// SPDX-License-Identifier: MIT

/*
personal deposit cap

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



import "../interfaces/IFeeMakerEvents.sol";


/// @author  ViaLend
/// @title   FeeMaker
/// @notice  A Smart Contract that helps liquidity providers managing their funds on Uniswap V3 .

contract FeeMaker is 
    ERC20,
    IFeeMakerEvents, 
    IUniswapV3MintCallback,
    IUniswapV3SwapCallback,
    ReentrancyGuard
{
    using SafeERC20 for IERC20;
    using SafeMath for uint256;
	

    IUniswapV3Pool public immutable pool;

	IERC20 public immutable token0;
    IERC20 public immutable token1;
    IERC20 public immutable ttoken;
    int24 public immutable tickSpacing;
    
    uint256 public protocolFee;
    uint256 public maxTotalSupply;
    address public strategist;
    address public governance;
    

    int24 public cLow;
    int24 public cHigh;
    uint256 public accruedProtocolFees0;
    uint256 public accruedProtocolFees1;

    uint256 public AccumulateFees0;
    uint256 public AccumulateFees1;
    
  	uint256 public lastRebalance;
    int24 public lastTick;

    uint32 public twapDuration;
	int24 public maxTwapDeviation;    
    
    
    /// @dev 
    /// @param _pool Uniswap V3 pool address
    ///	@param _ttoken SWSM team token
    /// @param _protocolFee Protocol fee expressed as multiple of 1e-6
    /// @param _maxTotalSupply Cap on total supply
    
    constructor(
        address _pool,
        address _ttoken,
        uint256 _protocolFee,
        uint256 _maxTotalSupply,
        int24 _maxTwapDeviation,
        uint32 _twapDuration
        
    ) ERC20("ViaLend Token","VLT") {

		pool = IUniswapV3Pool(_pool);	
		emit GeneralA("pool", _pool);
		
        token0 = IERC20(IUniswapV3Pool(_pool).token0());
        emit GeneralS("token0", "token0");
        
        token1 = IERC20(IUniswapV3Pool(_pool).token1());
        emit GeneralS("token1", "token1");
        
        ttoken = IERC20(_ttoken);
        emit GeneralS("ttoken", "ttoken");

        protocolFee = _protocolFee;
        
        
        governance = msg.sender;

		tickSpacing = IUniswapV3Pool(_pool).tickSpacing();

        require(_protocolFee < 1e6, "protocolFee");

        maxTotalSupply = _maxTotalSupply;

        maxTwapDeviation = _maxTwapDeviation;

        twapDuration = _twapDuration;

		require(_maxTwapDeviation > 0, "maxTwapDeviation");
		
        require(_twapDuration > 0, "twapDuration");
    }

    
    /// - Deposit tokens 
    /// @notice tokens get deposited in this smart contract will be held until next rebalance. 
    /// @param amountToken0 amount of token0 to deposit
    /// @param amountToken0 amount of token1 to deposit
    /// @param staker  The staker's address of the tokens 
    /// @return shares Number of shares minted
    /// @return amount0 Amount of token0 deposited
    /// @return amount1 Amount of token1 deposited
     
    function deposit(
        uint256 amountToken0,
        uint256 amountToken1,
        address staker
    )
        external
        
        nonReentrant
        
        returns (
            uint256 shares,
            uint256 amount0,
            uint256 amount1
        )
    {
        require(amountToken0 > 0 || amountToken1 > 0, _hint2("amountToken0,1,",amountToken0,amountToken1,0,",") );
       
        require(staker != address(0) && staker != address(this), "staker");

  		// Poke positions so to get uniswap v3 fees up to date. 
        _poke(cLow, cHigh);
		

        (shares, amount0, amount1) = _calcShares(amountToken0, amountToken1); 
        //_calcSharesAndAmounts(amountToken0, amountToken1); 
		
        require(shares > 0, _hint2("shares ", shares, 0, 0,"" ) ); 
        
        //todo
        require(amountMin(amount0,amount1), "amountMIn");

        // transfer tokens from sender
        if (amount0 > 0) token0.safeTransferFrom(msg.sender, address(this), amount0);
        if (amount1 > 0) token1.safeTransferFrom(msg.sender, address(this), amount1);

        _mint(staker, shares);

		//#debug  test send staker some ttoken tokenGiveAwayRate.div(100).mul(shares)
		// if ( ttoken.balanceOf(address(this) ) > 1000000000000000000 )
		// {
		// 	//uint tokenGiveAwayRate = 10; 
	 //        ttoken.safeTransfer(staker, 1000000000000000000);
  //       }

        emit Deposit(msg.sender, staker, shares, amount0, amount1);

        require(totalSupply() <= maxTotalSupply, "CAP");
    }

    // poke to update fees from uniswap. 
    function _poke(int24 tickLower, int24 tickUpper) internal {
        (uint128 liquidity, , , , ) = _position(tickLower, tickUpper);
        if (liquidity > 0) {
            pool.burn(tickLower, tickUpper, 0);
        }
    }


 
    
    
    /// - staker Withdraw 
    /// @param staker address of Recipient of tokens
    /// @param shares number of Shares owned by sender to be burned
    /// @return amount0 Amount of token0 sent to staker 
    /// @return amount1 Amount of token1 sent to staker
    /// @param amount0Min Revert if resulting `amount0` is smaller than this
    /// @param amount1Min Revert if resulting `amount1` is smaller than this
    
    function withdraw(
        uint256 shares,
        uint256 amount0Min,
        uint256 amount1Min,
        address staker
    ) external  nonReentrant returns (uint256 amount0, uint256 amount1) {
        
        require(shares > 0, "shares <= 0");

        require(staker != address(0) && staker != address(this), "staker");

        uint256 totalSupply = totalSupply();

        // Burn shares
        _burn(msg.sender, shares);

        require(totalSupply > 0, "t1");
        
        // Calculate token amounts proportional to unused balances
        uint256 unusedAmount0 = getBalance0().mul(shares).div(totalSupply);
        uint256 unusedAmount1 = getBalance1().mul(shares).div(totalSupply);

        // Withdraw proportion of liquidity directly from Uniswap pool
        (uint256 poolamount0, uint256 poolamount1) =
            _burnLiquidityShare(cLow, cHigh, shares, totalSupply);

        // Sum up total amounts owed to recipient
        amount0 = unusedAmount0.add(poolamount0);
        amount1 = unusedAmount1.add(poolamount1);


        require(amount0 >= amount0Min, _hint2(" amount0<amount0Min: ",amount0, 0,0,"") ) ;
        require(amount1 >= amount1Min, _hint2(" amount1<amount1Min: ",amount1, 0,0,"") );
        
        // Push tokens to recipient
        if (amount0 > 0) token0.safeTransfer(staker, amount0);
        if (amount1 > 0) token1.safeTransfer(staker, amount1);

        emit Withdraw(msg.sender, staker, shares, amount0, amount1,token0.name(),token1.name());
    }

    function _calcShares(uint256 amountToken0, uint256 amountToken1)
        internal
        view
        returns (
            uint256 shares,
            uint256 amount0,
            uint256 amount1
        )
    {
			
    	// total Amount of the shares in vault
        uint256 totalSupply = totalSupply();
        
        //get total amount from current balance + position liquidity for each token
        (uint256 total0, uint256 total1) = getTVL();

       	assert( totalSupply == 0   || total0 >0 ||  total1>0 );
		
        ( uint160 sqrtPriceX96, 	int24 tick, , , , , ) 	= pool.slot0();
        
        //( amount0,  amount1) = _swap(sqrtPriceX96, amountToken0, amountToken1);
        (amount0, amount1 ) = (amountToken0, amountToken1);

        shares =  amount0.mul(amount1).div(1e18);
        
    }    

	/// @dev Calculates the largest possible `amount0` and `amount1` such that
    /// they're in the same proportion as total amounts, but not greater than
    /// `amount0Desired` and `amount1Desired` respectively.
    function _calcSharesAndAmounts(uint256 amount0Desired, uint256 amount1Desired)
        internal
        view
        returns (
            uint256 shares,
            uint256 amount0,
            uint256 amount1
        )
    {
        uint256 totalSupply = totalSupply();
        (uint256 total0, uint256 total1) = getTVL();

        // If total supply > 0, vault can't be empty
        assert(totalSupply == 0 || total0 > 0 || total1 > 0);

        if (totalSupply == 0) {
            // For first deposit, just use the amounts desired
            amount0 = amount0Desired;
            amount1 = amount1Desired;
            shares = Math.max(amount0, amount1);
        } else if (total0 == 0) {
            amount1 = amount1Desired;
            shares = amount1.mul(totalSupply).div(total1);
        } else if (total1 == 0) {
            amount0 = amount0Desired;
            shares = amount0.mul(totalSupply).div(total0);
        } else {
            uint256 cross = Math.min(amount0Desired.mul(total1), amount1Desired.mul(total0));
            require(cross > 0, "cross");

            // Round up amounts
            amount0 = cross.sub(1).div(total1).add(1);
            amount1 = cross.sub(1).div(total0).add(1);
            shares = cross.mul(totalSupply).div(total0).div(total1);
        }
    }		 
	/// collect fees and remove liquidity from Uniswap pool.
	/// @param amount0 burned amount of token0
	/// @param amount1 burned amount of token1
    function _burnLiquidityShare(
        int24 tickLower,
        int24 tickUpper,
        uint256 shares,
        uint256 totalSupply
    ) internal returns (uint256 amount0, uint256 amount1) {
        (uint128 totalLiquidity, , , , ) = _position(tickLower, tickUpper);
        
        require(totalSupply > 0, _hint2("t2",totalLiquidity,0,0,"") ) ;
        
        uint256 liquidity = uint256(totalLiquidity).mul(shares).div(totalSupply);

        if (liquidity > 0) {
            (uint256 burned0, uint256 burned1, uint256 fees0, uint256 fees1) =
                _burnAndCollect(tickLower, tickUpper, _toUint128(liquidity));

            // Add share of fees
            amount0 = burned0.add(fees0.mul(shares).div(totalSupply));
            amount1 = burned1.add(fees1.mul(shares).div(totalSupply));
        }
    }
    
    
    /// @dev Casts uint256 to uint128 with overflow check.
    function _toUint128(uint256 x) internal pure returns (uint128) {
        assert(x <= type(uint128).max);
        return uint128(x);
    }
    
    function getTVL() 
    	public 
    	view  
    	returns (uint256 total0, uint256 total1) {
         
        (uint256 Amount0, uint256 Amount1) =  getPositionAmounts(cLow, cHigh);
        
        total0 = getBalance0().add(Amount0);
        total1 = getBalance1().add(Amount1);

    }
    
    

	function removePosition() internal {
        
        // Withdraw all current liquidity from Uniswap pool
      
       (uint128 allLiquidity, , , , ) = _position(cLow, cHigh); 
        
  		_burnAndCollect(cLow, cHigh, allLiquidity);

	}

    function rebalance(
        int24 newLow,
        int24 newHigh,
        int24 tick
        
    ) internal { 

        // get new balance
        uint256 balance0 = getBalance0();
        uint256 balance1 = getBalance1();

        
		// Place position on Uniswap
        uint128 liquidity = _liquidityForAmounts( newLow, newHigh, balance0, balance1);
        
        require(liquidity > 0 ,append("liquidity: ",uint2str(liquidity),"","","")) ;

   
        pool.mint(address(this), newLow, newHigh, liquidity, "");

        uint256 newBalance0 = getBalance0();
        uint256 newBalance1 = getBalance1();

        emit RebalanceLog(liquidity, balance0,balance1, newBalance0, newBalance1);

        (cLow, cHigh) = (newLow, newHigh);
        
    }
    
   
	///@notice Todo lending uni mix
	function strategy1(
		int24 newLow,
        int24 newHigh,
        int256 swapAmount,
        bool zeroForOne
        
		) external nonReentrant onlyGovernance  {
			
			
	}
	
	
	
	///@notice simple univ3 with swap
	function strategy0(
		int24 newLow,
        int24 newHigh,
        int256 swapAmount,
        bool zeroForOne
        
		) external nonReentrant onlyGovernance  {
		
        (	uint160 sqrtPriceX96, 	int24 tick, , , , , ) 	= pool.slot0();

  		_validRange(newLow, newHigh, tick);  // passed 1200 , 2100, 18382
        
        // Check price is not too close to min/max allowed by Uniswap. Price
        // shouldn't be this extreme unless something was wrong with the pool.

        int24 range = newHigh - newLow ;
            
        require(tick > TickMath.MIN_TICK + range  + tickSpacing, "tick too low");
        require(tick < TickMath.MAX_TICK - range  - tickSpacing, "tick too high");

        // Check price has not moved a lot recently. This mitigates price
        // manipulation during rebalance and also prevents placing orders
        // when it's too volatile.
        int24 twap = getTwap();
        int24 deviation = tick > twap ? tick - twap : twap - tick;
        require(deviation <= maxTwapDeviation, _hint2("deviation,maxTwapDeviation", uint256(deviation),uint256(maxTwapDeviation), 0,""));

        
        removePosition();
        
        uint256 balance0 = getBalance0();
        uint256 balance1 = getBalance1();
        
		// to be optimized outside
		uint160 sqrtPriceLimitX96 = zeroForOne ? TickMath.MIN_SQRT_RATIO + 1 : TickMath.MAX_SQRT_RATIO - 1;

		if (swapAmount > 0) {
	        swap( 
	        	 swapAmount, 
				 zeroForOne ,
				sqrtPriceLimitX96 );
		}
        
        

//        int24 tickFloor = _floor(tick);
//        int24 tickCeil = tickFloor + tickSpacing;

		//add position
		rebalance(
	        newLow,
	        newHigh,
			tick
			);

        lastRebalance = block.timestamp;
        lastTick = tick;
        		
	}
	
	/// to be optimized , or moved offchain
	function swap( 
		int256 swapAmount, 
		bool zeroForOne ,
		uint160 sqrtPriceLimitX96 
		) public returns (int256 , int256) {


    	return pool.swap(
               address(this),
               zeroForOne,
               swapAmount,
               sqrtPriceLimitX96,
               abi.encode(msg.sender)
        );
	}
	

  
	    /// @dev Fetches time-weighted average price in ticks from Uniswap pool.
    function getTwap() public view returns (int24) {
        uint32 _twapDuration = twapDuration;
        uint32[] memory secondsAgo = new uint32[](2);
        secondsAgo[0] = _twapDuration;
        secondsAgo[1] = 0;

        (int56[] memory tickCumulatives, ) = pool.observe(secondsAgo);
        return int24((tickCumulatives[1] - tickCumulatives[0]) / _twapDuration);
    }

    /// @dev Rounds tick down towards negative infinity so that it's a multiple
    /// of `tickSpacing`.
    function _floor(int24 tick) internal view returns (int24) {
        int24 compressed = tick / tickSpacing;
        if (tick < 0 && tick % tickSpacing != 0) compressed--;
        return compressed * tickSpacing;
    }

	
	
 /// @dev Callback for Uniswap V3 pool
    function uniswapV3MintCallback(
        uint256 amount0,
        uint256 amount1,
        bytes calldata data
    ) external override {
        require(msg.sender == address(pool));
        if (amount0 > 0) token0.safeTransfer(msg.sender, amount0);
        if (amount1 > 0) token1.safeTransfer(msg.sender, amount1);
    }

    /// @dev Callback for Uniswap V3 pool
    function uniswapV3SwapCallback(
        int256 amount0Delta,
        int256 amount1Delta,
        bytes calldata data
    ) external override {
        require(msg.sender == address(pool),"sender = pool?");

        //require(false, _hint2("deltaamount", uint256(amount0Delta),uint256(amount1Delta),0,""));
        
        if (amount0Delta > 0) token0.safeTransfer(msg.sender, uint256(amount0Delta));
        if (amount1Delta > 0) token1.safeTransfer(msg.sender, uint256(amount1Delta));
    }
    
	/// @dev Withdraws liquidity from a range and collects all fees in the
    /// process.
    function _burnAndCollect(
        int24 tickLower,
        int24 tickUpper,
        uint128 liquidity
    )
        internal
        returns (
            uint256 burned0,
            uint256 burned1,
            uint256 feesToVault0,
            uint256 feesToVault1
        )
    {
        if (liquidity > 0) {
	        ( burned0, burned1) =  pool.burn(tickLower, tickUpper, liquidity) ;
        }
        

        // Collect all owed tokens including earned fees
        (uint256 collect0, uint256 collect1) =
            pool.collect(
                address(this),
                tickLower,
                tickUpper,
                type(uint128).max,
                type(uint128).max
            );

        feesToVault0 = collect0.sub(burned0);
        feesToVault1 = collect1.sub(burned1);

        AccumulateFees0 = AccumulateFees0 + feesToVault0;
        AccumulateFees1 = AccumulateFees1 + feesToVault1;
        
        uint256 feesToProtocol0;
        uint256 feesToProtocol1;

        // Update accrued protocol fees
        uint256 _protocolFee = protocolFee;
        if (_protocolFee > 0) {
            feesToProtocol0 = feesToVault0.mul(_protocolFee).div(1e6);
            feesToProtocol1 = feesToVault1.mul(_protocolFee).div(1e6);
            feesToVault0 = feesToVault0.sub(feesToProtocol0);
            feesToVault1 = feesToVault1.sub(feesToProtocol1);
            accruedProtocolFees0 = accruedProtocolFees0.add(feesToProtocol0);
            accruedProtocolFees1 = accruedProtocolFees1.add(feesToProtocol1);
        }
        emit CollectFees(feesToVault0, feesToVault1, feesToProtocol0, feesToProtocol1);
    }    
	
    /// @dev Wrapper around `LiquidityAmounts.getLiquidityForAmounts()`.
    function _liquidityForAmounts(
        int24 tickLower,
        int24 tickUpper,
        uint256 amount0,
        uint256 amount1
    ) internal view returns (uint128) {
        (uint160 sqrtRatioX96, , , , , , ) = pool.slot0();
        return
            LiquidityAmounts.getLiquidityForAmounts(
                sqrtRatioX96,
                TickMath.getSqrtRatioAtTick(tickLower),
                TickMath.getSqrtRatioAtTick(tickUpper),
                amount0,
                amount1
            );
    }
    
    
     function _validRange(int24 _lower, int24 _upper, int24 _tick) internal view {
    	
        
        
        require(_lower < _upper, _hint2("lower < _upper ", uint(_lower), uint(_upper), uint(_tick),"")  );
        require(_lower < _tick , _hint2("lower > _tick ", uint(_lower), uint(_upper), uint(_tick),""));
        require(_upper > _tick , _hint2("_upper > _tick ", uint(_lower), uint(_upper), uint(_tick),""));
        
        require(_lower >= TickMath.MIN_TICK, "Lower too low");
        require(_upper <= TickMath.MAX_TICK, "Upper too high");

        
        require(_lower % tickSpacing == 0, "Lower % tickSpacing");
        require(_upper % tickSpacing == 0, "Upper % tickSpacing");
    }

     /// @notice Used to collect accumulated protocol fees.
    function collectProtocol(
        uint256 amount0,
        uint256 amount1,
        address to
    ) external onlyGovernance {
    	
    	require (accruedProtocolFees0 >= amount0 && accruedProtocolFees1 >= amount1,"protocolfees");

		if (amount0 > 0) {
	        accruedProtocolFees0 = accruedProtocolFees0.sub(amount0);
	        token0.safeTransfer(to, amount0);
		}
							
        if (amount1 > 0) {
        	accruedProtocolFees1 = accruedProtocolFees1.sub(amount1);
        	token1.safeTransfer(to, amount1);
        }
    }
    
	/**
     * @notice Amounts of token0 and token1 held in vault's position. Includes
     * owed fees but excludes the proportion of fees that will be paid to the
     * protocol. Doesn't include fees accrued since last poke.
    /// @param tickLower lower line price from last rebalance, 
    /// @param tickUpper upper line price from last rebalance
     */
    function getPositionAmounts(int24 tickLower, int24 tickUpper)
        public
        view
        returns (uint256 amount0, uint256 amount1)
    {
        (uint128 liquidity, , , uint128 tokensOwed0, uint128 tokensOwed1) =
            _position(tickLower, tickUpper);
        (amount0, amount1) = _amountsForLiquidity(tickLower,tickUpper, liquidity);

        // Subtract protocol fees
        uint256 oneMinusFee = uint256(1e6).sub(protocolFee);
        amount0 = amount0.add(uint256(tokensOwed0).mul(oneMinusFee).div(1e6));
        amount1 = amount1.add(uint256(tokensOwed1).mul(oneMinusFee).div(1e6));
    }


	/// @dev Wrapper around `IUniswapV3Pool.positions()`.
    function _position(int24 tickLower, int24 tickUpper)
        internal
        view
        returns (
            uint128,
            uint256,
            uint256,
            uint128,
            uint128
        )
    {

        bytes32 positionKey = PositionKey.compute(address(this), tickLower, tickUpper);
        return pool.positions(positionKey);
        
    }
    
    /// @dev Wrapper around `LiquidityAmounts.getAmountsForLiquidity()`.
    function _amountsForLiquidity(
        int24 tickLower,
        int24 tickUpper,
        uint128 liquidity
    ) internal view returns (uint256, uint256) {
        (uint160 sqrtRatioX96, , , , , , ) = pool.slot0();
        return
            LiquidityAmounts.getAmountsForLiquidity(
                sqrtRatioX96,
                TickMath.getSqrtRatioAtTick(tickLower),
                TickMath.getSqrtRatioAtTick(tickUpper),
                liquidity
            );
    }

	///calculate the minimum amount for token0 and token1 to deposit
	/// todo
	function amountMin(uint256 amount0, uint256 amount1) internal pure returns (bool){
		return true; 
		
	}
    /// @notice return Balance of available token0.
     
    function getBalance0() public view returns (uint256) {
        return token0.balanceOf(address(this)).sub(accruedProtocolFees0);
    }


    /// @notice return Balance of available token1.
    
    function getBalance1() public view returns (uint256) {
        return token1.balanceOf(address(this)).sub(accruedProtocolFees1);
    }   

	/// @notice vault liquidity in uniswap
    function getSSLiquidity(int24 tickLower, int24 tickUpper) external view returns(uint128 liquidity) {
    	( liquidity , , , , ) = _position(tickLower, tickUpper);
    }

	
    /// @notice Removes tokens accidentally sent to this vault.
    function sweep(
        IERC20 token,
        uint256 amount,
        address to
    ) external onlyGovernance {
        require(token != token0 && token != token1, "token");
        token.safeTransfer(to, amount);
    }
    
	///@notice set new maxTotalSupply
	function setMaxTotalSupply(uint256 newMax) external nonReentrant onlyGovernance {
			maxTotalSupply = newMax;
	}
	
    modifier onlyGovernance {
        require(msg.sender == governance, "governance");
        _;
    }
    
    
    /// @notice set the strategist contrat that send rebalance to the vault
    function setStrategist(address _strategist) external onlyGovernance {
        strategist = _strategist;
    }


    function setMaxTwapDeviation(int24 _maxTwapDeviation) external onlyGovernance {
        require(_maxTwapDeviation > 0, "maxTwapDeviation");
        maxTwapDeviation = _maxTwapDeviation;
    }

    function setTwapDuration(uint32 _twapDuration) external onlyGovernance {
        require(_twapDuration > 0, "twapDuration");
        twapDuration = _twapDuration;
    }

    ///#debug 888888888888888888888888888 dummy functions 
    
    
	/// @notice 3 functiosn for demo
	/// @dev 
	///		Hello(), Greet(), plus(x,y)
	
	function Hello() public pure returns (string memory) {
        return "Hello World";
    }
    
    function Greet(string memory str) public pure returns (string memory) {
        return str;
    }
    
    function plus(uint256 x, uint256 y) public pure returns (uint256) {
        return x.add(y);
    }    


	function _hint2(string memory msg1, uint v1, uint v2, uint v3,string memory msg2) internal pure returns (string memory)	{
		
		string memory sv1;
		string memory sv2;
		string memory sv3;
		
		if (v1 > 0 )
			sv1 = uint2str(v1);
		
		if (v2 > 0 )
			sv2 = uint2str(v2);

		if (v3 > 0 )
			sv3 = uint2str(v3);
			
		if (true)  // #debug, turn to false when launch
			return append( msg1 , ",", append(uint2str(v1),",", uint2str(v2),",",uint2str(v3)), ",", msg2) ;
		else
			return msg1;
	}
	
	

	function uint2str(uint i) internal pure returns (string memory ){
	    if (i == 0) return "0";
	    uint j = i;
	    uint length;
	    while (j != 0){
	        length++;
	        j /= 10;
	    }
	    bytes memory bstr = new bytes(length);
	    uint k = length - 1;
	    while (i != 0){
	        bstr[k--] = byte(uint8(48 + i % 10));
	        i /= 10;
	    }
	    return string(bstr);
	}


	function append(string memory a1, string  memory a2, string memory  a3, string memory  a4, string memory  a5) internal pure returns (string memory) {
	
	    return string(abi.encodePacked(a1, a2, a3, a4, a5));
	
	}

	


}

