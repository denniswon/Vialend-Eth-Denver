// SPDX-License-Identifier: MIT
pragma solidity >=0.5.0;

import "../interfaces/IWETH9.sol";
import "../@openzeppelin/contracts/math/SafeMath.sol";
import "../@uniswap/v3-core/contracts/libraries/TickMath.sol";
import "../@uniswap/v3-core/contracts/libraries/FullMath.sol";
import "../@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";

library lib  {



     
    
	function cantor_pair_calculate(uint256 x , uint256 y ) internal pure returns (uint256 ) {
	
	
		uint256 result = ((x+y)*(x+y+1))/2 + y;
	
		return result;
	}

	 function _fetchAmounts(uint256 myshare) internal pure returns(uint256,uint256){
    	
    	(uint256 d0,uint256 d1) = cantor_pair_reverse(myshare);
		
		return ( d0, d1);
    }
    
	 function cantor_pair_reverse (uint256 z ) internal pure returns (uint256 ,uint256 ) {
	
	 	uint256 t =  ( _sqrt(1+ 8 * z ) -1 )  / 2;
		
		uint256 x = t*(t+3)/2 - z;
		
		uint256 y = z - t*(t+1)/2;
	
	 	return (x, y);
	
	 }
	
		
	 // Returns approximate square root of 256-bit value
	  function sqrt(uint256 x) internal pure returns (uint256 y) {
		
	    uint256 z = (x + 1) >> 1;
	    y = x;
	
	    while (z < y) {
	      y = z;
	      z = (x / z + z) >> 1;
	    }
	  }

	function _sqrt(uint256 x) internal pure returns(uint256) {
	  
		if (x == 0)  return 0;   // Avoid zero divide crash    
		
		uint256 z = (x + 1 ) / 2;
		uint256 y = x;
		while(z < y){
		  y = z;
		  z = ( x / z + z ) / 2;
		}
		
		// require(y*y == x , "sqrt");  // cause crash
		
		return y;
	}
       
 /// @dev Rounds tick down towards negative infinity so that it's a multiple
    /// of `tickSpacing`.
    function _floor(int24 tick,int24 tickSpacing) internal pure returns (int24) {
        int24 compressed = tick / tickSpacing;
        if (tick < 0 && tick % tickSpacing != 0) compressed--;
        return compressed * tickSpacing;
    }

    /// @dev Casts uint256 to uint128 with overflow check.
    function _toUint128(uint256 x) internal pure returns (uint128) {
        assert(x <= type(uint128).max);
        return uint128(x);
    }
 
  
  function _validRange(int24 _lower, int24 _upper, int24 _tick, int24 tickSpacing) internal pure {
    	
		
        require(_lower < _upper, "V1");
        require(_lower < _tick , "V2");
        require(_upper > _tick , "V3");
        
        require(_lower >= TickMath.MIN_TICK, "V4");
        require(_upper <= TickMath.MAX_TICK, "V5");

        
        require(_lower % tickSpacing == 0, "V6");
        require(_upper % tickSpacing == 0, "V7");
    }
    
    
	

}