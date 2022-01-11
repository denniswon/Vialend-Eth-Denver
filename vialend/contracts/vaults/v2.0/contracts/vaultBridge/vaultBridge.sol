// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.4.0 <0.9.0;
/*
solc --optimize --overwrite --abi vaultBridge.sol -o ../../build/
solc --optimize --overwrite --bin vaultBridge.sol -o ../../build/

abigen --abi=../../build/vaultBridge.abi --bin=../../build/vaultBridge.bin --pkg=api --out=../../deploy/vaultBridge/vaultBridge.go

abigen --abi=../../build/vaultAdmin.abi --bin=../../build/vaultAdmin.bin --pkg=api --out=../../deploy/vaultAdmin/vaultAdmin.go

*/

contract VaultBridgeOwnable {
	address internal owner  ;


   modifier onlyOwner {
         require(msg.sender == owner, "owner");
        _;
    }
 	
}

/*
	VaultBridge 
	store vault address indexed by a mapping var
		0: weth/usdc vault
		1: weth/dai  vault
*/

contract VaultBridge is VaultBridgeOwnable {


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

    

}


/*
	VaultAdmin
	store admin addresses
*/
contract VaultAdmin is VaultBridgeOwnable {


    address[] private admins;
    
    mapping (address => uint) private addrIndex;

  	constructor(address[] memory _admins )  {
		owner = msg.sender;
		
		for (uint i=0; i< _admins.length; i++) {
			
			admins.push(_admins[i]) ;
			
			addrIndex[_admins[i]] = i+1;   // index can not be 0
			
		}
	}
        
	function setAdmin(address _addr ) external onlyOwner {
			
		if (! exist( _addr)  ) {
		 admins.push(_addr);
		 addrIndex[_addr] = admins.length;
		}
	}

	function delAdmin(address _addr ) external onlyOwner {
			
		require( addrIndex[_addr] > 0, "invalid address");

		uint index = addrIndex[_addr]-1;

		delete admins[ index ] ;	
		
		uint newIndex = admins.length-1;

		admins[index] = admins[newIndex];   // move last to fill the deleted slot

		admins.pop();
		
		addrIndex[admins[index]] = newIndex+1;  // update the address Index mapping
		
		delete addrIndex[_addr]; // delete the address mapping 
		

	}


// 	function getIndex(address _addr ) public view returns (uint ){
		
// 		return(  addrIndex[_addr] );
		
// 	}

	function exist(address _addr ) private view returns (bool){
		
		bool exists = addrIndex[_addr] > 0 ;
		
		return (exists);
		
	}
	
	function authAdmin(address _admin) external view returns (  bool  ) {
		return exist(_admin) ;
	}

    

}