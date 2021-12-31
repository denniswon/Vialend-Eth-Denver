// SPDX-License-Identifier: MIT
pragma solidity >=0.8.8;

contract UniCompFees  {

	uint256 internal uFees0;	// uniswap fee of token0 for the current position
	uint256 internal uFees1;	// uniswap fee of token1 for the current position
	uint256 internal lFees0;	// lending fee of token0 for the current position
	uint256 internal lFees1;	// lending fee of token1 for the current position

	struct FeeStruct {
		uint256 U3Fees0;				// uni v3 fees 0
		uint256 U3Fees1;				// uni v3 fees 1
		uint256 LcFees0;				// compound fees0
		uint256 LcFees1;				// compound fees1
	    uint256 AccruedProtocolFees0;		// for view
	    uint256 AccruedProtocolFees1;		// for view
	}


 	FeeStruct public Fees;
 
}