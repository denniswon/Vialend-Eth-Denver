#v1.0
	contract source code: VialendFeemaker.sol
	
addresses: ( goerli  )

	pool 				"0x933EFDc68cB4c1fe4Ef162a87E515768d6f82023", 
	vault 				"0xCF61ef07AA57089C7431599f239f7Dc144D31863", 
	tokenA (Weth) 	"0x6fD886fd1e728D9386Ba7fE721C856790758aDd9", 
	tokenB (usdc)   "0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C", 
	(#NOTE 这俩token 不是咱们自己发的， 换token用eth在uniswap上swap)
	
	privatekey   （ metamask import）
		onlygovernance:	
			2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7
		onlyteam: 			
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
		
		EmergencyBurn()   //** emergency button burn all positions from uniswap and lending pool back to vault. 
		
		WhiteHacker()   //** withdraw all assets back to governance address

** onlygovernance
* onlyteam



Test:
		/torukmakto/vialend/contracts/vaults/v1.0/test/scripts/index.go
		
