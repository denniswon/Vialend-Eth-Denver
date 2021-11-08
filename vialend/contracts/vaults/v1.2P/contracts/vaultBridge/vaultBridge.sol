// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.4.0 <0.9.0;
/*
solc --optimize --overwrite --abi vaultBridge.sol -o ./
solc --optimize --overwrite --bin vaultBridge.sol -o ./
abigen --abi=vaultBridge.abi --bin=vaultBridge.bin --pkg=api --out=vaultBridge.go

*/


contract VaultBridge {

	address private owner  ;
    mapping(uint  => address) public vaults;


  	constructor()  {
		owner = msg.sender;
	}
        
    function getAddress(uint ind) external view returns( address) {
    	return vaults[ind];
    }

    function setAddress(address newAddress,  uint ind ) external onlyOwner {
		vaults[ind] = newAddress;
    	
    }


    modifier onlyOwner {
         require(msg.sender == owner, "owner");
        _;
    }
    

}