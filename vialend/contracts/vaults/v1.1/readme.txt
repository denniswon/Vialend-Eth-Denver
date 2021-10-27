#v1.1

-2021/10/08 05:10 
	**
		public pool 改为 internal pool.  may break your front end code. 
		call poolAddress() to get the pool address instead
	
	
contracts source code: 
	
		VialendFeemaker.sol  
		ViaUniswap.sol
		ViaCompound.sol
		ownable.sol
	
addresses: ( goerli  )

	vault 				"0xAda1AC9D1dBFFF1270f0009d8f20bd0332F48113", 

	pool 			"0x933EFDc68cB4c1fe4Ef162a87E515768d6f82023", 
	tokenA (Weth) 	"0x6fD886fd1e728D9386Ba7fE721C856790758aDd9", 
	tokenB (usdc)   "0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C", 
	
	
	onlygovernance   :	
			0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E
			(privatekey) 2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7
			
	onlyteam privatekey  : 			
			0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E
			2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7
	 	
ABIs：
	vault abi:
		/torukmakto/vialend/contracts/vaults/v1.0/build/ViaLendFeeMaker.abi

	pool abi:
		/torukmakto/vialend/contracts/vaults/v1.0/build/IUniswapV3Pool.abi
	ctoken abi:
		/torukmakto/vialend/contracts/vaults/v1.0/build/CErc20.abi
		/torukmakto/vialend/contracts/vaults/v1.0/build/CEth.abi
	WETH9 abi:	
		/torukmakto/vialend/contracts/vaults/v1.0/build/WETH9_.abi
		



public functions:
	
		Deposit()		// able to receive any amount of assets
		
		WithDraw()    // able to withdraw any portion of user's share by given number of percentage. i.e. 50  means withdraw 50% of assets in vault
		
		Strategy1()   // * remove position from both uniswap and lending pool , reenter with new range
		
		setUniPortionRatio()  //* set uniswap liquidity portion , e.g. 10 
		
		setMaxTotalSupply()    //** reset max total sepply  
		
		setGovernance()   // ** set pending governance , need call acceptGovernance afterwords upon successful
		
		acceptGovernance()   //** final set governance
		
		setTeam()  // ** reset team address 
		
		EmergencyBurn()   //** emergency button burn all positions from uniswap and lending pool back to vault and send all assets to users. 
		
		getLendingAmounts() // may remove later 
		

** onlygovernance
* onlyteam



Test:
		/torukmakto/vialend/contracts/vaults/v1.0/test/scripts/index.go
		
