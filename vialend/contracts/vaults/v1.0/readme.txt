#v1.0
	contract source code: VialendFeemaker.sol
	


------------------ old versions --------------------

#v0.1.3 
	
	new function strategy0  -- check FeeMaker.sol
	
	用strategy0() 代替原来的rebalance()， 需要计算并带入的参数 ： 
		int24 newLow,
        int24 newHigh,
        int256 swapAmount,
        bool zeroForOne
    参考testVault.go 里的 GetSwapInfo（） 方法获取参数。   

	
To Test: ( goerli  )
	pool 				"0xc4C92691f69fadDd684257E9f5A8d6f9D2c79a93", 
	vault 				"0x31E84D42aB6DEf5Dac84b761b0E5004179e07778", 
	tokenA (eWeth) 		"0x48FCb48bb7F70F399E35d9eC95fd2A614960Dcf8", 
	tokenB (eusdc1) 	"0x6f38602e142D0Bd3BC162f5912535f543D3B73d7", 
	Admin privateKey: 	284b65567176c10bc010345042b1d9852fcc1c42ae4b76317e6da040318fbe7f


	
 external and public functions 

    	deposit  ,  
		widthdraw
		
		swap,  
		strategy0, simple swap
		strategy1, vialending
		
		setMaxTotalSupply
		
	    getBalance0
	    getBalance1
		
		capacity = getTVL / maxTotalSupply 

		getTVL   -- used to be getTotalAmounts  = total in vault + total position in uniswap
		
		getPositionAmounts	  --  pool's token0 token1
		
		getSSLiquidity  -- pool's liquidity
		
		
		collectProtocol		

		sweep
	