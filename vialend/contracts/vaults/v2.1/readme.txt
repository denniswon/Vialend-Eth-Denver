#v2.1

$ ganache-cli -f http://73.153.58.102:8545 -d -m "clutch captain shoe salt awake harvest setup primary inmate ugly among become" -i 999 -u --port 8545 --host "192.168
223"

Available Accounts
==================
(0) 0xa0df350d2637096571F7A701CBc1C5fdE30dF76A (100 ETH)
(1) 0xEC2DD0d0b15D494a58653427246DC076281C377a (100 ETH)
(2) 0x5ACb5DB941E3Fc33E0c0BC80B90114b6CD0249B5 (100 ETH)
(3) 0x04AA0A9d06f3d7De5E0B1f3631cC428A36f2aa93 (100 ETH)
(4) 0xbCc210E4803Bfca8762c76ee5d85bB60263f3e5F (100 ETH)
(5) 0x99c2F835F2415Edb6ECe615Ee9296a232a67633C (100 ETH)
(6) 0x7C7Ec4870Ce62E13340903f1815b3B649C3F2736 (100 ETH)
(7) 0x76EFcC66cF77D7f3493274e8832Aa9FAb205C830 (100 ETH)
(8) 0x524B5428Bc40207574840EbD917C6845043D1f0c (100 ETH)
(9) 0xAdc732c19dca2cb20000898Eb0c4F328bCAC9681 (100 ETH)

Private Keys
==================
(0) 0xb8c1b5c1d81f9475fdf2e334517d29f733bdfa40682207571b12fc1142cbf329
(1) 0x5c2313d8a6b81a83ad1df1bf12a193cbc51d5de84a000db734fd7a05aa63e5a2
(2) 0x2deeef19c7418df1c35425d5b637133305ec425f063a0ea6bc1702559b1e3123
(3) 0xf0a43592054c17fd756f160567a116b93ab3f16a92f7667eb49a4522e3e26cdc
(4) 0x100f17bef675e9532e2bf1ed01b3f0ffa4a64a55a3a109b2a254a69a5603ac0b
(5) 0x191cdb377fd0298d3bff842f50fa3454876cb3473beed299541744c65c9f37bd
(6) 0xc5fd727f2808dfc3b8fdfa5eec64f6212573823770f19a3bf39f197cc4b6a6f8
(7) 0xa05315d1e2eccf42b3e7dd5b47a4a1cf3cc5674d7933a32d617061d46734c7b2
(8) 0x9a4114d070ef076f23b7a42fb4957053bbada99d90036a2bec1fbaff5f0f9d9e
(9) 0x4d528e815fe2170a60b9303e9666e8640fe2a6a1ad84bef2596484083d073715


#v2.0
-2021/12/05

Full script:
	
	deploy ViaAuth, address as viaAdmin 
	
	deploy StratUniComp ( also hidden deploy VaultUniComp)
	
	address of VaultUniComp 
	address of StratUniComp
end script	


	两个合约互相调用， stratedy 合约 创建 viaVault 合约， 
	用go 发布时， 
		solc --optimize --overwrite --abi StratUniComp.sol -o ../build
		solc --optimize --overwrite --bin StratUniComp.sol -o ../build
		abigen --abi=../build/StratUniComp.abi --bin=../build/StratUniComp.bin --pkg=api --out=../deploy/StratUniComp/StratUniComp.go

		solc --optimize --overwrite --abi ViaAuth.sol -o ../build
		solc --optimize --overwrite --bin ViaAuth.sol -o ../build
		abigen --abi=../build/ViaAuth.abi --bin=../build/ViaAuth.bin --pkg=api --out=../deploy/ViaAuth/ViaAuth.go

		
	DeployStratUniComp() 
	
	同时获得StratUniComp 和 ViaVault address
	
	deposit , withdraw 用 viavault 合约
	rebalance  用 StratUniComp 合约
	
	ViaVault 合约重要参数
        address creator;    // strategy address

	StratUniComp 合约重要参数
        address unipool;        // get by uni factory, token0, token1, feetier
        address governance;     // governance of protocol, ownable 创建
	
	StratUniComp 创建合约所需的参数：
		string name;        // strategy name

        address token0;         // underlying token0
        address token1;         // underlying token1
        uint8  feetier;			// uni v3 feetier

        uint32 twapDuration;        // oracle twap durantion
        int24 maxTwapDeviation;      
        uint128 quoteAmount;  		

		uint8 uniPortion;       // uniswap portion ratio
		uint8 compPortion;       // compound portion ratio
       

        address payable weth;       // weth address
        address cToken0;
        address cToken1;
        address cEth;
        

        uint8 creatorFee;		// 0 - 20% of profit

        uint8 protocolFee;	// 0 - 20% of profit of creator's fee

        uint32 threshold;	// initial range

        uint256 vaultCap;	   		// 0: no cap

        uint256 individualCap;	   //  0 : no cap

		
		
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
	
