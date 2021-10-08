// SPDX-License-Identifier: MIT
pragma solidity >=0.5.0;

import "../interfaces/IWETH9.sol";

library lib  {


	function cantor_pair_calculate(uint256 x , uint256 y ) internal pure returns (uint256 ) {
	
	
		uint256 result = ((x+y)*(x+y+1))/2 + y;
	
		return result;
	}

	// /**
	//  * Return the source integers from a cantor pair integer.
	//  */

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
	
		
		function _sqrt(uint x) internal pure returns(uint) {
		  uint z = (x + 1 ) / 2;
		  uint y = x;
		  while(z < y){
		    y = z;
		    z = ( x / z + z ) / 2;
		  }

	   	  // require(y*y == x , "sqrt");

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
 
  
	

}