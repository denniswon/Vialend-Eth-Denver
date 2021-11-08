#v1.2P
-2021/11/07

	contract source /vialend/contracts/vaults/v1.2P/contracts/vaultBridge/vaultBridge.sol
	
	contract address 0x033F3C5eAd18496BA462783fe9396CFE751a2342
	abi: /vialend/contracts/vaults/v1.2P/contracts/vaultBridge/vaultBridge.abi
	
	public method: getGetVaultAddress(int index)
		@index  
			0: weth/usdc
			1: weth/dai
		return: address	
	
	

-2021/11/04 
	--  CToken0 and CToken1 address public
	--  new paramter {quoteAmount = 1e18 } required to deploy vault    check DeployVialendFeemaker() 
	
	-- new vaults for pairs:
	
	vault weth/usdc:  0x4aaE0bc3052aD3AB125Ae654f0f2C55Dbd9D6e17
	vault weth/dai:  0x35938d9b221238BBcE1F9b5196FFeE0f87E22D26



-- 2021 / 10 / 30 

(weth/usdc 0.3% )   vault	0x6F520a253EC8f4d0B745649a5C02bB7a5201d70b 

(weth/dai 0.05% )	vault	"0x522f6c4C073A86787F5D8F676795290973498929"  



-2021/10/22 

		0x6F520a253EC8f4d0B745649a5C02bB7a5201d70b  //vault
		"0x04B1560f4F58612a24cF13531F4706c817E8A5Fe", //pool 0.3%


		0x8c2CFFE9e0BFa86Fea2753C1ffb756da32c6e8bB	// vault (by template 发布)
		0xe979387E6dAD7D4a92F9aC88e42C6e6461DB8b64  // (pool 1% )


		"0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", // token0 Weth
		"0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C", //token1  Usdc



-- Changes:

	function deposit(
        uint256 amountToken0,		// deposit amount of token0
        uint256 amountToken1,		// deposit amount of token1
        bool doRebalence			// whether do rebalance or not
    ) 
	
	function setProtocolFee(uint256 _protocolFee) external onlyGovernance
	
		


-- vault deploy note:

	pool := GetPoolFromToken()
	protocolFee := big.NewInt(10000)
	maxTotalSupply, ok := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
	var maxTwapDeviation = big.NewInt(20000)
	var twapDuration = uint32(2)
	var _weth = common.HexToAddress(config.Network.LendingContracts.WETH)
	var _cToken0 = common.HexToAddress(config.Network.LendingContracts.CETH)
	var _cToken1 = common.HexToAddress(config.Network.LendingContracts.CUSDC)
	var _cEth = common.HexToAddress(config.Network.LendingContracts.CETH)
	
