// SPDX-License-Identifier: MIT

pragma solidity >=0.8.8;

interface IVaultFactory
{ 
    function getAdmin() external view returns(address);
    function getTeam() external view returns(address);
    function setTeam(address _team) external;
  	function changeStat(address _strategy, address _vault, uint _stat) external;
	function getCount() external; /// get stored vaults array size
	function getStat(address, address) external view returns(uint);
	function checkActive(address sORv ) external view returns(bool); 
	

}