// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0;

import "./@OpenZeppelin/contracts/token/ERC20/ERC20.sol";



contract ERC20fixedSupply is ERC20 {

    constructor (
    	string memory name,
		string memory symbol,
		uint8 decimals,
		uint totalSupply
		
    )  ERC20(name, symbol) {
		_mint(msg.sender, totalSupply * (10** decimals)  );

    }
}