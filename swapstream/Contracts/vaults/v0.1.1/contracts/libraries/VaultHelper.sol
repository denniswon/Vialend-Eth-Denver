// SPDX-License-Identifier: Unlicense

pragma solidity >0.5.0 < 0.9.0;

contract VaultHelper {
	
	function gas() returns(uint _gasLimit, uint _gasPrice) {
		
	}
	

  function getValueByOwner(address _owner) external view returns(uint[] memory) {
    
    uint[] memory result = new uint[](ownerValueCount[_owner]);
    
    return result;
    
  }	
}

