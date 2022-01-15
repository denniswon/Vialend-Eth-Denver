# Vialend Solidity Contracts


### VaultFactory.sol

### VaultStrategy.sol

### ViaVault.sol

### VaultBridge

`contract address: `

	- v1.x
		0xb6F0049e37D32dED0ED2FAEeE7b69930FA49A879
	
	- v2.x
		0x428EeA0B87f8E0f5653155057f58aaaBb667A3ec

`getAddress`  

	- function getAddress (uint ind) external view returns( address) 
	- inputs 
		0: factory
		x: vault address 
		
		ie:
		1: weth/dai 0.05%
		2: weth/usdc 0.3%
		....
		
	- returns
		vault address 
			or 
		address(0): invalid input number
	
`setAddress` 

	- function setAddress (address newAddress,  uint ind ) external onlyOwner  
	- Set address 

`setPermit`

	- function setPermit (address addr,  uint level ) external onlyOwner  
	- Set admin 
			1: admin
			0/other: non-admin
			
`getPermit` - 

	- function getPermit(address addr ) external view returns(uint)  
	- get admin access level 
	- returns
		1: if address is admin
		0: non-admin 
	
