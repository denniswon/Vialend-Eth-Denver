// SPDX-License-Identifier: MIT

pragma solidity 0.7.6;

interface WETH9_ 
{
	function  balanceOf(address account) external returns(uint);
	function allowance( address owner, address spender) external returns(uint);

	fallback() external payable ;
	 receive() external payable ;

	function deposit() external payable ;
	function withdraw(uint wad) external ;
	function totalSupply() external view returns (uint) ;
	 
	function approve(address guy, uint wad) external returns (bool) ;
	
	function transfer(address dst, uint wad) external returns (bool) ;
	
	function transferFrom(address src, address dst, uint wad) external returns (bool);
}


contract WrapUnwrap
{

	WETH9_ internal WETH;

	constructor (address payable WETHAddr)  {

	    require(WETHAddr != address(0), "WETH is the zero address");
	
	    WETH = WETH9_(WETHAddr);

	}

    function wrap() payable public  {
	    uint256 ETHAmount =msg.value;
	
	    //create WETH from ETH
	    if (msg.value != 0) {
	        WETH.deposit{value:msg.value}();
	    }   
	    require(WETH.balanceOf(address(this))>=ETHAmount,"Ethereum not deposited");
	}


	function unwrap(uint256 Amount) public 
	{
	    address payable sender= msg.sender;
	
	    if (Amount != 0) {
	        WETH.withdraw(Amount);
	        sender.transfer(address(this).balance);
	    }
	}
}