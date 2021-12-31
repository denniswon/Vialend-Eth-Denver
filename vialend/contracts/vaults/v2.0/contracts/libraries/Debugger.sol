 // SPDX-License-Identifier: MIT
pragma solidity >=0.5.0;

library Debugger  {

	function uint2str(uint i) public view returns (string memory ){
	    if (i == 0) return "0";
	    uint j = i;
	    uint length;
	    while (j != 0){
	        length++;
	        j /= 10;
	    }
	    bytes memory bstr = new bytes(length);
	    uint k = length - 1;
	    while (i != 0){
	        bstr[k--] = byte(uint8(48 + i % 10));
	        i /= 10;
	    }
	    return string(bstr);
	}


	function append2(string memory a1, string  memory a2) public view returns (string memory) {
	
	    return string(abi.encodePacked(a1, a2));
	
	}

}