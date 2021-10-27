// SPDX-License-Identifier: MIT

pragma solidity >=0.4.21 <= 0.8.6;

/// @title 
interface IFundKeeper {
	
    function deposit(
        uint256,
        uint256,
        uint256,
        uint256,
        address
    )
        external
        returns (
            uint256,
            uint256,
            uint256
        );


   
    function withdraw(
        uint256,
        uint256,
        uint256,
        address
    ) external returns (uint256, uint256);


    function getTotalAmounts() external view returns (uint256, uint256);
}
