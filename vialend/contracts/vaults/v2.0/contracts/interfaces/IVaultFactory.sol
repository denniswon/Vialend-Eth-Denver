// SPDX-License-Identifier: MIT

pragma solidity >=0.8.8;
     
interface IVaultFactory
{ 


    function getAdmin() external view returns(address);
    function getTeam() external view returns(address);
    function setTeam(address _team) external;
  	function changeStat(address _strategy, address _vault, uint _stat) external;
	
	/// get stored vaults array size
	function getCount() external; 	
	
	function getStat(address) external view returns(uint);
	function getStat2(address _strategy, address _vault) external view returns(uint);
	function getPair0(address _addr) external view returns(address);
	
	///return true if a2 and a1 is a strategy/vault pair
	function onlyPair(address a1, address a2) external view returns(bool); 
	

}