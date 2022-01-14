# Vialend Solidity Contracts


### VaultFactory.sol

### VaultStrategy.sol

### ViaVault.sol

### VaultBridge

`function getAddress(uint ind) external view returns( address)` - get vault address by number
	 	
	0: factory
	1: weth/dai 0.05%
	2: weth/usdc 0.3%
	
`setAddress(address newAddress,  uint ind ) external onlyOwner ` - Set address 

`setPermit(address addr,  uint level ) external onlyOwner ` - Set admin by level 

	1: admin
	0/other: non-admin
	
`getPermit(address addr ) external onlyOwner ` - get admin permission

	1: admin
	0/other: non-admin

