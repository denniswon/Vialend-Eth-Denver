// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.5.0;

import "./@openzeppelin/contracts/math/Math.sol";
import "./@openzeppelin/contracts/math/SafeMath.sol";
import "./@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "./@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "./@uniswap/v3-core/contracts/interfaces/callback/IUniswapV3MintCallback.sol";
import "./@uniswap/v3-core/contracts/interfaces/callback/IUniswapV3SwapCallback.sol";
import "./@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";
import "./@uniswap/v3-core/contracts/libraries/TickMath.sol";
import "./@uniswap/v3-core/contracts/libraries/Position.sol";
import "./@uniswap/v3-periphery/contracts/libraries/LiquidityAmounts.sol";
import "./@uniswap/v3-periphery/contracts/libraries/PositionKey.sol";

//import "../interfaces/IFeeMaker.sol";
//import "../interfaces/IFeeMakerEvents.sol";



contract Tester 
{
    using SafeERC20 for IERC20;
    using SafeMath for uint256;

	    uint128 public p1;
	    uint256  public p2;
	    uint256 public  p3;
	    uint128  public p4;
	    uint128  public p5;
	    uint128  public p6;
		uint128  public amount0;
		uint128  public amount1;
        
    
    constructor() {

    }


 
   
    function callPosition
    (	address pool, 
    	address LP,
    	int24 tickLower, 
		int24 tickUpper
	)
        external
        
        returns (
            uint128,
            uint256,
            uint256,
            uint128,
            uint128,
            uint128 ,
            uint128 
        )
    {

        bytes32 positionKey = PositionKey.compute(LP, tickLower, tickUpper);
        
        (amount0, amount1) = IUniswapV3Pool(pool).collect(LP, tickLower, tickUpper,0,0);
        
        (p1,p2,p3,p4,p5) = IUniswapV3Pool(pool).positions(positionKey);
        return(p1,p2,p3,p4,p5 , amount0, amount1);

        
    }

    
    function setPosition
    (	address pool, 
    	address LP,
    	int24 tickLower, 
		int24 tickUpper
	)
        external
        
        returns (
            uint128,
            uint256,
            uint256,
            uint128,
            uint128
        )
    {
		
        bytes32 positionKey = PositionKey.compute(LP, tickLower, tickUpper);
        
        (p1,p2,p3,p4,p5) = IUniswapV3Pool(pool).positions(positionKey) ;
        p6 = 6;
        return(p1,p2,p3,p4,p5);

        
    }

	function  getPosition() external view returns (
            uint128,
            uint256,
            uint256,
            uint128,
            uint128,
            uint128,
            uint128,
            uint128
        ) {
		return (p1,p2,p3,p4,p5,p6,amount0,amount1);
    }

}


