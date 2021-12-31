# Gas setting on test net 

	auth.GasLimit = 6721975     -- in gwei
	auth.GasPrice = 2gwei   --  1 gwei = 1e9 wei,  2 gwei = 2000000000 wei


## basic cost based on test net settings: (mainnet may need 50x )
	1. (template) deploy vault deployer			2210275
	2. (template) deploy strategy deployer		4623309
	3. (template) deploy factory:  				 699464	
	4. factory create strategy + vault :    	5944202	
	
							total gas:  = 	13,477,250
							


----------------------------------------------
.......................Deploy VaultDeployer. ..................
----------------------------------------------
	[tx:  0x8b4601bf538eba392adc9fbd477e65bd8b1d01959437db1ce769d3f5430c6272]
	[BlockNumber: 6116609]
	[Status: SUCCESS]
	[CumulativeGasUsed: 2344928]
	[GasUsed: 2210275]
	
----------------------------------------------
.......................Deploy StratDeployer. ..................
----------------------------------------------
	[tx:  0x486b2991ef883ec7a6530812eb4626930d6d3df7c522ba05f7c6cbbcfbac64b9]
	[BlockNumber: 6116610]
	[Status: SUCCESS]
	[CumulativeGasUsed: 5023843]
	[GasUsed: 4623309]

----------------------------------------------
.......................Factory Vault. -- deploy strategy and vault together by factory
----------------------------------------------
	[tx:  0xd4a32863e3fdb0c2939ad7c2ca4a7dec4793c77d88f60a4e86b331e003c99753]
	[BlockNumber: 6116612]
	[Status: SUCCESS]
	[CumulativeGasUsed: 5965202]
	[GasUsed: 5944202]


----------------------------------------------
.......................Deploy VaultFactory. ..................
----------------------------------------------
	[tx:  0xf727a6f86fee8f68dd90d950d1ac830ed27833dff9ba6c0a375ed5261dbefae8]
	[BlockNumber: 6116611]
	[Status: SUCCESS]
	[CumulativeGasUsed: 925668]
	[GasUsed: 699464]


----------------------------------------------
.........Deposit.........  
----------------------------------------------
	[Aprollowance ] -- Approve
	[tx:  0x287653450a35182a501e3de3f803fdd9ad9c58a288ad8f6f854fa47a34a1a9d3]
	[BlockNumber: 6116621]
	[Status: SUCCESS]
	[CumulativeGasUsed: 242523]
	[GasUsed: 46111]

	[deposit tx: %s 0x6f6943cddb053f22832bfa735e9c93ba5e338db6063481f43d08d05b31a6a71b]
	[BlockNumber: 6116623]
	[Status: SUCCESS]
	[CumulativeGasUsed: 367752]
	[GasUsed: 283752]


----------------------------------------------
.........Withdraw.........  
----------------------------------------------
	[withdraw tx:  0x51e8abe94334ab0de3fa74b4af6a1bfc7d4893b88d8f971ef4d4b2c903edf3dc]
	[BlockNumber: 6116627]
	[Status: SUCCESS]
	[CumulativeGasUsed: 313156]
	[GasUsed: 203367]


----------------------------------------------
.........Rebalance New.........  
----------------------------------------------
	[tx:  0x203f733ddfa633dd67cddad5c446b3f9463f9eb6dd2a192600f9aaf6201fa15e]
	[BlockNumber: 6116624]
	[Status: SUCCESS]
	[CumulativeGasUsed: 1471608]
	[GasUsed: 495838]


----------------------------------------------
.........Emergency withdraw , burn all positions and send fund back to users.........  
----------------------------------------------
	[emergency tx:  0x08836b1449533f46a558be4e0f76249f370abe7a1a8b54d05a0dfbfc0efe8f09]
	[BlockNumber: 6117238]
	[Status: SUCCESS]
	[CumulativeGasUsed: 1273971]
	[GasUsed: 375454]

if fail:

	[emergency tx:  0x5d9857637d0fe99f5e674a43fbc1228f9b6019be70b99435dd264807aa41fc99]
	[BlockNumber: 6116451]
	[Status: FAIL]
	[CumulativeGasUsed: 29393]
	[GasUsed: 29393]
	
	
	
# Compare with single deploy


	1. deploy vault 				2210275
	2. deploy strategy 				4623309
	3. deploy factory:  			 699464	
	4. register to factory:  		 299464	
	
					total gas:  = 7,832,512


----------------------------------------------
.......................Deploy Strat by Go. ..................
----------------------------------------------
	[strat address: 0x7ffE86c23117A9809941D5f693C9fCb9aBB3a998]
	[tx:  0x3d2cfd3af756a5095beef1b9d67f16e4a7437de2fcff1053b09543dc31ca3f1e]
	[BlockNumber: 6116446]
	[Status: SUCCESS]
	[CumulativeGasUsed: 4199822]
	[GasUsed: 4157822]
	

----------------------------------------------
.......................Deploy Vault by Go. ..................
----------------------------------------------
	[tx:  0x03b2e33fb798c5c325f65cd2bdedb26489f9a00db012d5ebed6e1dd296bd351f]
	[BlockNumber: 6114598]
	[Status: SUCCESS]
	[CumulativeGasUsed: 2562972]
	[GasUsed: 2341105]


----------------------------------------------
.......................Deploy VaultFactory. ..................
----------------------------------------------
	VaultFactory address: 0xa93a4e530AdB28adB4BED52eD189B0d19C7553d7
	[tx:  0x2f3e91964f6df531d5f02f9acf848e41735ff86aedca86f120ee7cc6e84b8252]
	[BlockNumber: 6114601]
	[Status: SUCCESS]
	[CumulativeGasUsed: 1111133]
	[GasUsed: 625925]
		
		
		