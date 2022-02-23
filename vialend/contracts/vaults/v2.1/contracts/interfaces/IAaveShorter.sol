// SPDX-License-Identifier: MIT

pragma solidity >=0.8.8;

interface IAaveShorter {
    function short(
        address _aavePool,
		address _colateral,
        address _shorting,
		uint256 _amount,
        uint256 _shortSize
    ) external returns(bool);
}