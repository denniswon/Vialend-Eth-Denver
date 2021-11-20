// SPDX-License-Identifier: MIT
pragma solidity >=0.5.0;

import "@openzeppelin/contracts/math/Math.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";
import "@uniswap/v3-core/contracts/libraries/TickMath.sol";
import "@uniswap/v3-periphery/contracts/libraries/LiquidityAmounts.sol";
import "@uniswap/v3-periphery/contracts/libraries/PositionKey.sol";

import "libraries/lib.sol";
import "./ownable.sol";
import "interfaces/IFeeMakerEvents.sol";

contract ViaUniswap is 
	Ownable,
    IFeeMakerEvents
    
{
	
    using SafeMath for uint256;


    // // poke to update fees from uniswap. 
    // function _poke(int24 tickLower, int24 tickUpper) internal {
    //     (uint128 liquidity, , , , ) = _position(tickLower, tickUpper);
    //     if (liquidity > 0) {
    //         pool.burn(tickLower, tickUpper, 0);
    //     }
    // }



 
    function _mintUniV3(
        int24 newLow,
        int24 newHigh,
        uint256 amount0,
        uint256 amount1
        
    ) internal { 

        
		// Place position on Uniswap
        //uint128 liquidity = _liquidityForAmounts( newLow, newHigh, amount0, amount1);
        
         (uint160 sqrtRatioX96, , , , , , ) = pool.slot0();
        
        uint128 liquidity =
            LiquidityAmounts.getLiquidityForAmounts(
                sqrtRatioX96,
                TickMath.getSqrtRatioAtTick(newLow),
                TickMath.getSqrtRatioAtTick(newHigh),
                amount0,
                amount1
            );
            
        
        //#debug require(liquidity > 0 ,append("liquidity: ",uint2str(liquidity),"","","")) ;

   
        pool.mint(address(this), newLow, newHigh, liquidity, "");

      

        // emit RebalanceLog(liquidity, newBalance0, newBalance1);

        (cLow, cHigh) = (newLow, newHigh);
        
    }
   
	
	///remove uniswap position; generate fees
	function removeUniswap() internal {
       
    	if (cLow == 0 && cHigh ==0)  return;
		
       (uint128 liquidity, , , , ) = _position(cLow, cHigh);   // get liquidity by current ticklow, tickhigh
       
 		(uint256 burned0, uint256 burned1, uint256 collect0, uint256 collect1) = _burnAndCollectUnis(cLow, cHigh, liquidity);
    
    	if (liquidity > 0) {
	        uFees0 = collect0.sub(burned0);
	       	uFees1 = collect1.sub(burned1);
       }
    
    	(liquidity, , , , ) = _position(cLow, cHigh);	// should be 0 otherwise, there is problem
		
		emit MyLog("uni liquidity = If this is not 0, there was an error", liquidity);
	}
	

	///@notice burn and collect 
    function _burnAndCollectUnis(
        int24 tickLower,
        int24 tickUpper,
        uint128 liquidity
    )
        internal
        
        returns (
            uint256 burned0,
            uint256 burned1,
   			uint256 collect0,
			uint256 collect1 

        )
    {
    	
		
        if ( liquidity > 0) {
        	( burned0, burned1) =  pool.burn(tickLower, tickUpper, liquidity) ;
        } 
        
         // Collect all owed tokens including earned fees
        (collect0,  collect1) =
            pool.collect(
                address(this),
                tickLower,
                tickUpper,
                type(uint128).max,
                type(uint128).max
            );

        
     
        
    }    
	
    // /// @dev Wrapper around `LiquidityAmounts.getLiquidityForAmounts()`.
    // function _liquidityForAmounts(
    //     int24 tickLower,
    //     int24 tickUpper,
    //     uint256 amount0,
    //     uint256 amount1
    // ) internal view returns (uint128) {
    //     (uint160 sqrtRatioX96, , , , , , ) = pool.slot0();
    //     return
    //         LiquidityAmounts.getLiquidityForAmounts(
    //             sqrtRatioX96,
    //             TickMath.getSqrtRatioAtTick(tickLower),
    //             TickMath.getSqrtRatioAtTick(tickUpper),
    //             amount0,
    //             amount1
    //         );
    // }
    
    
   
     /// @notice Used to collect accumulated protocol fees.
    
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
        uint256 oneMinusFee = uint256(100).sub(protocolFee);
        amount0 = amount0.add(uint256(tokensOwed0).mul(oneMinusFee).div(100));
        amount1 = amount1.add(uint256(tokensOwed1).mul(oneMinusFee).div(100));
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

	

	/// @notice vault liquidity in uniswap
    function getSSLiquidity(int24 tickLower, int24 tickUpper) external view returns(uint128 liquidity) {
    	( liquidity , , , , ) = _position(tickLower, tickUpper);
    }


	function poolAddress() external view returns (address) {
         return address(pool);
    }


    function getTwap(address pool, uint32 period ) public view returns (int24 tick) {
        
        require(period != 0, 'xBP');   

        uint32[] memory secondsAgo = new uint32[](2);
        secondsAgo[0] = period;
        secondsAgo[1] = 0;

        (int56[] memory tickCumulatives, ) = IUniswapV3Pool(pool).observe(secondsAgo);
        int56 tickDelta = tickCumulatives[1] - tickCumulatives[0];
        tick = int24(tickDelta / period);
        if (tickDelta < 0 && (tickDelta % period != 0)) tick--;
    }
    
 

   function getQuoteAtTick(
        int24 tick,
        uint128 baseAmount,
        address baseToken,
        address quoteToken
    ) public pure returns (uint256 quoteAmount) {
    	
        uint160 sqrtRatioX96 = TickMath.getSqrtRatioAtTick(tick);

        // Calculate quoteAmount with better precision if it doesn't overflow when multiplied by itself
        if (sqrtRatioX96 <= type(uint128).max) {
            uint256 ratioX192 = uint256(sqrtRatioX96) * sqrtRatioX96;
            quoteAmount = baseToken < quoteToken
                ? FullMath.mulDiv(ratioX192, baseAmount, 1 << 192)
                : FullMath.mulDiv(1 << 192, baseAmount, ratioX192);
        } else {
            uint256 ratioX128 = FullMath.mulDiv(sqrtRatioX96, sqrtRatioX96, 1 << 64);
            quoteAmount = baseToken < quoteToken
                ? FullMath.mulDiv(ratioX128, baseAmount, 1 << 128)
                : FullMath.mulDiv(1 << 128, baseAmount, ratioX128);
        }
    }
    
  
	// function getPriceBySQRTP(uint160 sqrtPriceX96)
	//         external
	//         pure
	//         returns (uint256 price)
 //    {
	//         return uint(sqrtPriceX96).mul(uint(sqrtPriceX96)).mul(1e18) >> (96 * 2);
 //    }    

}