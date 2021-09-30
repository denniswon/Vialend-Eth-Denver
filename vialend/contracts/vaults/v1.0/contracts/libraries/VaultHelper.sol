// SPDX-License-Identifier: MIT
pragma solidity >=0.5.0;

library VaultHelper  {

   
	function _hint2(string memory msg1, uint v1, uint v2, uint v3,string memory msg2) public view returns (string memory)	{
		
		string memory sv1;
		string memory sv2;
		string memory sv3;
		
		if (v1 > 0 )
			sv1 = uint2str(v1);
		
		if (v2 > 0 )
			sv2 = uint2str(v2);

		if (v3 > 0 )
			sv3 = uint2str(v3);
			
		if (true)  // #debug, turn to false when launch
			return append( msg1 , ",", append(uint2str(v1),",", uint2str(v2),",",uint2str(v3)), ",", msg2) ;
		else
			return msg1;
	}
	
	

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


	function append(string memory a1, string  memory a2, string memory  a3, string memory  a4, string memory  a5) public view returns (string memory) {
	
	    return string(abi.encodePacked(a1, a2, a3, a4, a5));
	
	}

}