#v1.0
	contract source code: VialendFeemaker0.1.1.sol
	
addresses: ( goerli  )

	pool 				"0x933EFDc68cB4c1fe4Ef162a87E515768d6f82023", 
	vault 				"0x5A3D3D52aa406AaB48CaDc4318974f8270e7711b", 
	
	tokenA (Weth) 		
		"0x6fD886fd1e728D9386Ba7fE721C856790758aDd9", //tokenA Weth
	tokenB (usdc) 		
		"0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C", //tokenB  usdc
	
	(#token 不是自己发的， 因为要用到compound, 换token用eth在uniswap上swap)
	 	
	Compound C tokens:
			WETH:  "0x6fD886fd1e728D9386Ba7fE721C856790758aDd9", ///on remix solc0.4.12,  injected web3 deployed by 0x2ee9... test admin,
			USDC:  "0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C",
			CUSDC: "0xCEC4a43eBB02f9B80916F1c718338169d6d5C1F0",
			CWBTC: "0x6CE27497A64fFFb5517AA4aeE908b1E7EB63B9fF",
			WBTC:  "0xC04B0d3107736C32e19F1c62b2aF67BE61d63a05",
			CDAI:  "0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb",
			DAI:   "0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60",
			CETH:  "0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF",
		
	privatekey
		governance:	
			2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7
		team: 			
			2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7
		
		
	vault abi:
		C:/Users/xdotk/torukmakto/vialend/contracts/vaults/v1.0/build/ViaLendFeeMaker.abi

	pool abi:
		C:/Users/xdotk/torukmakto/vialend/contracts/vaults/v1.0/build/IUniswapV3Pool.abi
	ctoken abi:
		C:/Users/xdotk/torukmakto/vialend/contracts/vaults/v1.0/build/CErc20.abi
		C:/Users/xdotk/torukmakto/vialend/contracts/vaults/v1.0/build/CEth.abi
	WETH9 abi:	
		C:/Users/xdotk/torukmakto/vialend/contracts/vaults/v1.0/build/WETH9_.abi
		



public functions:
	
		Deposit()		// able to receive any amount of assets
		
		WithDraw()    // able to withdraw any portion of user's share by given number of percentage. i.e. 50  means withdraw 50% of assets in vault
		
		Strategy1()  // remove position from both uniswap and lending pool , reenter with new range
		
		setUniPortionRatio()  // set uniswap liquidity portion , e.g. 10 
		
		setMaxTotalSupply()    // reset max total sepply  
		
		setGovernance()   // set pending governance , need call acceptGovernance afterwords upon successful
		
		acceptGovernance()   // final set governance
		
		setTeam()  // reset team address 
		
		EmergencyBurn()   // emergency button burn all positions from uniswap and lending pool back to vault. 
		
		WhiteHacker()   // withdraw all assets back to governance address



