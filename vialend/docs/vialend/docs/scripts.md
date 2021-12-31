# Features

## Deploy vault

call factory contract to generate two smart contracts 

	-- strategy address
	-- vault address

go script:

	parameters:
		
		contracts[10]  	-- 10 contract addresses see below 

		vaultCap		-- vault tvl limit in USD or USDC
		individualCap	-- individual capacity in USD or USDC
		uniPortionRate	-- uniswap v3 portion rate, 0 - 100
		compPortionRate -- compound PortionRate , 0 - 100
		creatorFee		-- fee cut for creator
		feetier			-- uniswap pool feetier
		quoteAmount		-- 1 token0 in its decimal. i.e. if token0 is weth, 1e18 is the value

		_stratategyDeployer := Cfg.STRAT_DEPLOYER	-- previously deployed by stratDeployer()
		_vaultDeployer := Cfg.VAULT_DEPLOYER	-- previously deployed by vaultDeplyer()

		contracts := [10]common.Address{
			protocolAddr,		// an address that will take the protocol fees
			reserve,
			WETH,
			CETH,
			CTOKEN0,
			CTOKEN1,
			TokenA,
			TokenB,
			_stratategyDeployer,	
			_vaultDeployer,
		}
		
smart contract method to generate the vault and strategy:

	go:	
	
	factoryContract.CreateVault(Auth,
			contracts,
			vaultCap,
			individualCap,
			uniPortionRate,
			compPortionRate,
			creatorFee,
			feetier,
			quoteAmount
		)		

		
		
				
## View vaults registery

solidity:

	VaultFactory.sol:
	
		struct VaultReg {
			address strategy;
			address vault;
		}
		
		VaultReg[] public vaults;

go:     	

	func ViewVaults()

		count, _ := facotryInstance.GetCount(&bind.CallOpts{}) 
		i := count.Int64()
		for i > 0 {
			i--
			vaults, _ := facotryInstance.Vaults(&bind.CallOpts{}, big.NewInt(i))
	
			myPrintln(i+1, " of ", count)
			myPrintln("strategy:", vaults.Strategy)
			myPrintln("vault:", vaults.Vault)
	
			stat, _ := facotryInstance.GetStat(&bind.CallOpts{}, vaults.Strategy, vaults.Vault)
			myPrintln("stat:", stat)
		}


# TECH DETAILS
	
	
## Compile contracts:

	ViaVault.sol, VaultStrategy.sol, VaultFactory.sol, VaultDeployer.sol, StrategyDeployer.sol
	
	Interface:
		IViaVault.sol, IStrategy.sol, IVaultFactory.sol, IDeployer.sol
		
	Files to modify when changes need to re-deploy. 
		changes in passing parameter to constructor. make sure to check all files above to make proper change accordingly.
		. 
		
	Compile and Deploy:
				
		solc --optimize --overwrite --abi VaultFactory.sol -o ../build
		solc --optimize --overwrite --bin VaultFactory.sol -o ../build
		solc --optimize --overwrite --abi VaultStrategy.sol -o ../build
		solc --optimize --overwrite --bin VaultStrategy.sol -o ../build
		solc --optimize --overwrite --abi ViaVault.sol -o ../build
		solc --optimize --overwrite --bin ViaVault.sol -o ../build
		abigen --abi=../build/ViaVault.abi --bin=../build/ViaVault.bin --pkg=api --out=../deploy/ViaVault/ViaVault.go
		abigen --abi=../build/VaultFactory.abi --bin=../build/VaultFactory.bin --pkg=api --out=../deploy/VaultFactory/VaultFactory.go
		abigen --abi=../build/VaultStrategy.abi --bin=../build/VaultStrategy.bin --pkg=api --out=../deploy/VaultStrategy/VaultStrategy.go
		solc --optimize --overwrite --abi StratDeployer.sol -o ../build
		solc --optimize --overwrite --bin StratDeployer.sol -o ../build
		abigen --abi=../build/StratDeployer.abi --bin=../build/StratDeployer.bin --pkg=api --out=../deploy/StratDeployer/StratDeployer.go
		solc --optimize --overwrite --abi VaultDeployer.sol -o ../build
		solc --optimize --overwrite --bin VaultDeployer.sol -o ../build
		abigen --abi=../build/VaultDeployer.abi --bin=../build/VaultDeployer.bin --pkg=api --out=../deploy/VaultDeployer/VaultDeployer.go

	TroubleShooting.
		FactoryVault() fail. 
			try deploy by the deployer separatedly to nail down whether its the Vault or Strategy Failed. 
			index.go
			`project.DeployVaultByDeployer()`  
			`project.DeployStratByDeployer()`



## Deploy Contracts
	
Solidity:

	VaultDeployer.sol		
	StrategyDeployer.sol
	VaultFactory.sol
	ViaVault.sol

go:
	
	/scripts/index.go  -- go file 
	
		project.DeployVaultDeployer()	-- deploy the vault Deployer
		project.DeployStratDeployer()	-- deploy the strategy Deployer
		project.DeployVaultFactory()	-- Deploy factory
		project.FactoryVault()	-- Call factory.method.createVault() from web3js or 
								-- go client to deploy Strategy and Vault through its deployer

	-- Check Status and Change Status
		s := project.Network.VaultStrat
		v := project.Network.Vault
		project.ChangeStat(s, v, 1)
		project.GetStat(s, v)
		
## GAS COST Estimation

##### 分别deploy Transaction Fee:
	ViaVault  Transaction Fee:	0.004369686 Ether 
	Gas Price:	0.000000002 Ether (2 Gwei)

	VaultStrategy  Transaction Fee:	0.008336874 Ether 
	Gas Price:	0.000000002 Ether (2 Gwei)
	
##### Deploy within SmartContract STRATEGY + VAULT  
	
	STRATEGY + VAULT  Transaction Fee:	0.012671544 Ether ($0.00)
	Gas Price: 0.000000002 Ether (2 Gwei)


## 前端 Web