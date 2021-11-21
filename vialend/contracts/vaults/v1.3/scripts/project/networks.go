package include

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

var Networkid = 3 /// 0: mainnet, 1: local, 2: local , 3: gorlie, 4: gorlie,  5: goreli , 6: rinkeby
var Account = 0
var ProviderSortId = 0

type TokenStruct struct {
	Name           string
	Symbol         string
	Decimals       uint8
	MaxTotalSupply *big.Int
}

type LendingStruct struct {
	WETH  string
	CETH  string
	DAI   string
	CDAI  string
	ETH   string
	WBTC  string
	CWBTC string
	USDC  string
	CUSDC string
	USDT  string
	CUSDT string
}
type Params struct {
	ProviderUrl []string
	Factory     string
	Callee      string
	PrivateKey  []string
	Governance  string
	TokenA      string
	TokenB      string
	CTOKEN0     string
	CTOKEN1     string
	NewOwner    string
	PendingTime time.Duration
	Pool        string
	SwapRouter  string
	BonusToken  string
	Vault       string
	FeeTier     int64
	VaultBridge string
	VaultAdmin  string

	LendingContracts LendingStruct
}

type Indi int

const (
	Envelope Indi = iota
	Bollinger
	TMA
)

var Token [2]TokenStruct

//var Decimals0, Decimals1 uint8

var Network = Networks[Networkid]

var Client = GetClient(Networkid, ProviderSortId)

var Auth = GetSignature(Networkid, Account)

var FromAddress common.Address

var DEBUG = true

type Info struct {
	Name  string
	Value string
}

var InfoString []Info

var Networks = [...]Params{

	{ // 0 mainnet
		[]string{"https://mainnet.infura.io/v3/68070d464ba04080a428aeef1b9803c6"},
		"0x1F98431c8aD98523631AE4a59f267346ea31F984", //factory
		"", //callee
		[]string{"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7"}, //privatekey fake
		"", // Governance
		"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", // weth  tokenA
		"0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", // usdc 	tokenB
		"0x4Ddc2D193948926D02f9B1fE9e1daa0718270ED5", //ctokenA
		"0x39AA39c021dfbaE8faC545936693aC917d5E7563", //ctokenB
		"", //newOwner
		60, //pendingtime
		"0x8ad599c3A0ff1De082011EFDDc58f1908eb6e6D8", // pool
		"", // swap router
		"", // bonus token
		"", //vault address
		3000,
		"", // VaultBridge
		"", //VaultAdmin
		LendingStruct{
			WBTC:  "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599",
			CWBTC: "0xC11b1268C1A384e55C48c2391d8d480264A3A7F4",
			USDC:  "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			CUSDC: "0x39AA39c021dfbaE8faC545936693aC917d5E7563",
			CETH:  "0x4Ddc2D193948926D02f9B1fE9e1daa0718270ED5",
			DAI:   "0x6B175474E89094C44Da98b954EedeAC495271d0F",
			CDAI:  "0x5d3a536E4D6DbD6114cc1Ead35777bAB948E3643",
			USDT:  "0xdAC17F958D2ee523a2206206994597C13D831ec7",
			CUSDT: "0xf650C3d88D12dB855b8bf7D11Be6C55A4e07dCC9",
		},
	},

	{ // 1 local usdc/usdt
		[]string{"http://127.0.0.1:7545", "http://127.0.0.1:8545"},
		"0x0c8D15944A4f799D678029523eC1F82c84b85F32", //factory
		"0x210FA31C72D9F020D16BF948e54F108D1C688f81", //callee
		[]string{"e8ef3a782d9002408f2ca6649b5f95b3e5772364a5abe203f1678817b6093ff0", // 0xfd8a5AE495Df1CA34F90572cb99A33B27173eDe1
			"f804a123dd9876c73cef5d198cce0899e6dfc2f851ed2527b003e11cd5383c54",  // 0xeBb29c07455113c30810Addc123D0D7Cd8637aea
			"88ca92d92059dea4f8fb4181e97304eb8b2264c8e1e577ba4358aea4699a7c09"}, //0xcCebCE6AF95bfD69EEE193390CCF027e8d47d9e2
		"", // Governance
		"0x59Cd9D486a8fA9b39F715915743997daA12d138e", //tokenB usdt
		"0x9D96eC63f96A4E985e227BF520dD742315AB77c7", //tokenA usdc
		"", //ctoken0
		"", //ctoken1
		"0xeBb29c07455113c30810Addc123D0D7Cd8637aea", //newOwner
		10,
		"0x96Dd142387281a16F72962CCb659cAc67D73d882", // pool
		"", // swap router
		"0xD0d1E195c613Cb6eea9308daB69661CAF9760eF9", // bonus token
		"0xcB0b392e747C101Ed949247730eC3aa6A75E4D3B", //vault address
		3000, // fee
		"",   // VaultBridge
		"0xdB11518974087276048364aef0596637527ddDBd", //VaultAdmin

		LendingStruct{DAI: "ASDF", CDAI: "ASD"},
	},

	{ // 2 local weth/usdc
		[]string{"http://127.0.0.1:7545", "http://127.0.0.1:8545"},
		"0x0c8D15944A4f799D678029523eC1F82c84b85F32", //factory
		"0x210FA31C72D9F020D16BF948e54F108D1C688f81", //callee
		[]string{"e8ef3a782d9002408f2ca6649b5f95b3e5772364a5abe203f1678817b6093ff0",
			"f804a123dd9876c73cef5d198cce0899e6dfc2f851ed2527b003e11cd5383c54"},
		"", // Governance
		"0xB73A78A3C493ACdbA893da9331ff39Fe4E59bFA3", //e weth1
		"0xd8F4E5E1cE1a2961b5fB401B8c2286549607B294", //e usdc1
		"", //ctoken0
		"", //ctoken1
		"0xeBb29c07455113c30810Addc123D0D7Cd8637aea", //newOwner
		10,
		"0xaEbc0569A8Ad476530d765dBE6308842Bd4D699c", // pool
		"", // swap router
		"0xD0d1E195c613Cb6eea9308daB69661CAF9760eF9", // bonus token
		"0x2723f0d5F2E60D1BF686B835e630C55453307eEA", //vault address
		3000, // fee
		"",   // VaultBridge
		"",   //VaultAdmin

		LendingStruct{DAI: "ASDF", CDAI: "ASD"},
	},
	{ ///3  goerli admin test 1
		[]string{"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"https://goerli.infura.io/v3/06e0f08cb6884c0fac18ff89fd46d131"}, ///  provider url

		"0x1F98431c8aD98523631AE4a59f267346ea31F984", //factory
		// "0xd648DB0713965e927963182Dc44D07D122a703ed", //callee
		"0xE97f1488F053251032ef358dE5b4188cD960D413", //callee
		[]string{"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7", //0,  admin 	0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E
			"284b65567176c10bc010345042b1d9852fcc1c42ae4b76317e6da040318fbe7f",  //1,  admin 1  0x6dd19aEB91d1f43C46f0DD74C9E8A92BFe2a3Cd0
			"d8cda34b6928af75aff58c60fe9ed3339896b57a13fa88695aa6da7b775cda2a",  //2,  admin 2  0xD8Dbe65b64428464fFa14DEAbe288b83C240e713
			"2d9e2b4c955159dd8a22faf3cb3074f03cfc182213729224915921daabaa5d6a",  //3, team			0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440
			"01e8c8df56230b8b6e4ce6371bed124f4f9950c51d64adc581938239724ed5e6",  //4,  user 1	0x14792757D21e54453179376c849662dE341797F2
			"67f7046a9f3712d77dab07a843c91d060ab5f27b808ed54d6db1293c7cd5eff3",  //5,  user 2	0x4F211267896C4D3f2388025263AC6BD67B0f2C54
			"a830f08514d29b0d278b251773b2265cd462e02ad14ca016591929d42fb203d1"}, //6 arb01 0x8a01C3E04798D0B6D7423EaFF171932943FB9A8D
		"284b65567176c10bc010345042b1d9852fcc1c42ae4b76317e6da040318fbe7f", //  0x6dd19aEB91d1f43C46f0DD74C9E8A92BFe2a3Cd0"", // Governance

		// "0x48FCb48bb7F70F399E35d9eC95fd2A614960Dcf8", //tokenA eWeth
		// "0x6f38602e142D0Bd3BC162f5912535f543D3B73d7", //tokenB  eusdc

		"0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", // tokenA Weth
		"0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C", //tokenB  usdc
		"0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF", // CETH
		"0xCEC4a43eBB02f9B80916F1c718338169d6d5C1F0", //CUSDC

		"0x4a9c85e96C30EA642728926bC8df23eFC95224dF", //new owner, test user 1
		40, //time pending interval
		"0x04B1560f4F58612a24cF13531F4706c817E8A5Fe", //pool
		"0xe592427a0aece92de3edee1f18e0157c05861564", // uni swap router
		"0x3C3eF6Ad37F107CDd965C4da5f007526B959532f", // team  token

		"0xD0fF8fF803a30C5d7BBDdc797B544E07Ff3458cD", //vault   can delete account

		3000, // fee
		"0x033F3C5eAd18496BA462783fe9396CFE751a2342", // VaultBridge
		"", //VaultAdmin

		LendingStruct{
			WETH:  "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6",
			CETH:  "0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF",
			USDC:  "0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C",
			CUSDC: "0xCEC4a43eBB02f9B80916F1c718338169d6d5C1F0",
			CWBTC: "0x6CE27497A64fFFb5517AA4aeE908b1E7EB63B9fF",
			WBTC:  "0xC04B0d3107736C32e19F1c62b2aF67BE61d63a05",
			CDAI:  "0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb",
			DAI:   "0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60",
		},
	},
	{ ///4  goerli weth / usdc fee tier 0.1%
		[]string{"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"https://goerli.infura.io/v3/06e0f08cb6884c0fac18ff89fd46d131"}, ///  provider url

		"0x1F98431c8aD98523631AE4a59f267346ea31F984", //factory
		//"0xd648DB0713965e927963182Dc44D07D122a703ed", //callee
		"0xE97f1488F053251032ef358dE5b4188cD960D413", //callee

		[]string{
			"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7",  //0,  admin 	0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E
			"284b65567176c10bc010345042b1d9852fcc1c42ae4b76317e6da040318fbe7f",  //1,  admin 1  0x6dd19aEB91d1f43C46f0DD74C9E8A92BFe2a3Cd0
			"d8cda34b6928af75aff58c60fe9ed3339896b57a13fa88695aa6da7b775cda2a",  //2,  admin 2  0xD8Dbe65b64428464fFa14DEAbe288b83C240e713
			"2d9e2b4c955159dd8a22faf3cb3074f03cfc182213729224915921daabaa5d6a",  //3, team			0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440
			"01e8c8df56230b8b6e4ce6371bed124f4f9950c51d64adc581938239724ed5e6",  //4,  user 1	0x14792757D21e54453179376c849662dE341797F2
			"67f7046a9f3712d77dab07a843c91d060ab5f27b808ed54d6db1293c7cd5eff3",  //5,  user 2	0x4F211267896C4D3f2388025263AC6BD67B0f2C54
			"a830f08514d29b0d278b251773b2265cd462e02ad14ca016591929d42fb203d1"}, //6 arb01 0x8a01C3E04798D0B6D7423EaFF171932943FB9A8D

		"284b65567176c10bc010345042b1d9852fcc1c42ae4b76317e6da040318fbe7f", //  0x6dd19aEB91d1f43C46f0DD74C9E8A92BFe2a3Cd0"", // Governance

		"0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", //  Weth
		"0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60", //  DAI
		"0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF", // CETH
		"0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb", //CDAI

		"", //new owner, test user 1
		30, //time pending interval
		"0x1738f9aAB1d370a6d0fd56a18f113DbD9e1DCd4e", //pool  (weth,dai , 500)
		"0xe592427a0aece92de3edee1f18e0157c05861564", // uni swap router  not used
		"0x3C3eF6Ad37F107CDd965C4da5f007526B959532f", // team  token  not used

		"0xb102Cd93329d7017Ae83C6E488f00EaB4844CbF2", //vault

		500, // fee
		"0x033F3C5eAd18496BA462783fe9396CFE751a2342", // VaultBridge
		"", //VaultAdmin

		LendingStruct{
			WETH:  "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", ///on remix solc0.4.12,  injected web3 deployed by 0x2ee9... test admin,
			CETH:  "0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF",
			USDC:  "0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C", // REFERENCE BELOW
			CUSDC: "0xCEC4a43eBB02f9B80916F1c718338169d6d5C1F0",
			CWBTC: "0x6CE27497A64fFFb5517AA4aeE908b1E7EB63B9fF",
			WBTC:  "0xC04B0d3107736C32e19F1c62b2aF67BE61d63a05",
			CDAI:  "0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb",
			DAI:   "0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60",
		},
	},
	{ ///5  goerli weth / usdt fee tier 0.1%
		[]string{"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"https://goerli.infura.io/v3/06e0f08cb6884c0fac18ff89fd46d131"}, ///  provider url

		"0x1F98431c8aD98523631AE4a59f267346ea31F984", //factory
		//"0xd648DB0713965e927963182Dc44D07D122a703ed", //callee
		"0xE97f1488F053251032ef358dE5b4188cD960D413", //callee

		[]string{
			"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7",  //0,  admin 	0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E
			"284b65567176c10bc010345042b1d9852fcc1c42ae4b76317e6da040318fbe7f",  //1,  admin 1  0x6dd19aEB91d1f43C46f0DD74C9E8A92BFe2a3Cd0
			"d8cda34b6928af75aff58c60fe9ed3339896b57a13fa88695aa6da7b775cda2a",  //2,  admin 2  0xD8Dbe65b64428464fFa14DEAbe288b83C240e713
			"2d9e2b4c955159dd8a22faf3cb3074f03cfc182213729224915921daabaa5d6a",  //3, team			0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440
			"01e8c8df56230b8b6e4ce6371bed124f4f9950c51d64adc581938239724ed5e6",  //4,  user 1	0x14792757D21e54453179376c849662dE341797F2
			"67f7046a9f3712d77dab07a843c91d060ab5f27b808ed54d6db1293c7cd5eff3",  //5,  user 2	0x4F211267896C4D3f2388025263AC6BD67B0f2C54
			"a830f08514d29b0d278b251773b2265cd462e02ad14ca016591929d42fb203d1"}, //6 arb01 0x8a01C3E04798D0B6D7423EaFF171932943FB9A8D

		"284b65567176c10bc010345042b1d9852fcc1c42ae4b76317e6da040318fbe7f", //  0x6dd19aEB91d1f43C46f0DD74C9E8A92BFe2a3Cd0"", // Governance

		"0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", //  Weth
		"0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60", //  DAI
		"0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF", // CETH
		"0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb", //CDAI

		"", //new owner, test user 1
		30, //time pending interval
		"0x1738f9aAB1d370a6d0fd56a18f113DbD9e1DCd4e", //pool  (weth,dai , 500)
		"0xe592427a0aece92de3edee1f18e0157c05861564", // uni swap router  not used
		"0x3C3eF6Ad37F107CDd965C4da5f007526B959532f", // team  token  not used

		"0xb102Cd93329d7017Ae83C6E488f00EaB4844CbF2", //vault

		10000, // fee
		"0x033F3C5eAd18496BA462783fe9396CFE751a2342", // VaultBridge
		"", //VaultAdmin

		LendingStruct{
			WETH:  "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", ///on remix solc0.4.12,  injected web3 deployed by 0x2ee9... test admin,
			CETH:  "0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF",
			USDC:  "0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C", // REFERENCE BELOW
			CUSDC: "0xCEC4a43eBB02f9B80916F1c718338169d6d5C1F0",
			CWBTC: "0x6CE27497A64fFFb5517AA4aeE908b1E7EB63B9fF",
			WBTC:  "0xC04B0d3107736C32e19F1c62b2aF67BE61d63a05",
			CDAI:  "0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb",
			DAI:   "0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60",
		},
	},

	{ ///6  rinkeby tester admin
		[]string{"https://rinkeby.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"https://rinkeby.infura.io/v3/06e0f08cb6884c0fac18ff89fd46d131"}, ///  provider url

		"0x1F98431c8aD98523631AE4a59f267346ea31F984", //factory
		"", //callee
		[]string{"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7", //0,  admin 	0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E
			"284b65567176c10bc010345042b1d9852fcc1c42ae4b76317e6da040318fbe7f",  //1,  admin 1  0x6dd19aEB91d1f43C46f0DD74C9E8A92BFe2a3Cd0
			"d8cda34b6928af75aff58c60fe9ed3339896b57a13fa88695aa6da7b775cda2a",  //2,  admin 2  0xD8Dbe65b64428464fFa14DEAbe288b83C240e713
			"2d9e2b4c955159dd8a22faf3cb3074f03cfc182213729224915921daabaa5d6a",  //3, team			0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440
			"01e8c8df56230b8b6e4ce6371bed124f4f9950c51d64adc581938239724ed5e6",  //4,  user 1	0x14792757D21e54453179376c849662dE341797F2
			"67f7046a9f3712d77dab07a843c91d060ab5f27b808ed54d6db1293c7cd5eff3",  //5,  user 2	0x4F211267896C4D3f2388025263AC6BD67B0f2C54
			"a830f08514d29b0d278b251773b2265cd462e02ad14ca016591929d42fb203d1"}, //6 arb01 0x8a01C3E04798D0B6D7423EaFF171932943FB9A8D
		"284b65567176c10bc010345042b1d9852fcc1c42ae4b76317e6da040318fbe7f", //  0x6dd19aEB91d1f43C46f0DD74C9E8A92BFe2a3Cd0"", // Governance

		"0xd606ddFA13914F274CBa3B4B22120eCc8Ba1C67a", //tokenA Weth
		"0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C", //tokenB  usdc
		"0xd6801a1DfFCd0a410336Ef88DeF4320D6DF1883e", //ctoken0
		"0x5B281A6DdA0B271e91ae35DE655Ad301C976edb1", //ctoken1
		"",   //new owner, test user 1
		50,   //time pending interval
		"",   //pool 0x3c7fADe1921Bf9D8308D76d7B09cA54839cfF033", //pool tusdc/ tweth 0xBF93aB266Cd9235DaDE543fAd2EeC884D1cCFc0c // 0x3c7fADe1921Bf9D8308D76d7B09cA54839cfF033", eweth/eusdc //pool
		"",   // swap router
		"",   // tto  token
		"",   //vault address
		3000, // fee
		"",   // VaultBridge
		"",   //VaultAdmin

		LendingStruct{
			WETH:  "0xd606ddFA13914F274CBa3B4B22120eCc8Ba1C67a",
			USDC:  "0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C",
			CUSDC: "0x5B281A6DdA0B271e91ae35DE655Ad301C976edb1",
			CWBTC: "0x0014F450B8Ae7708593F4A46F8fa6E5D50620F96",
			WBTC:  "0x577D296678535e4903D59A4C929B718e1D575e0A",
			CDAI:  "0x6D7F0754FFeb405d23C51CE938289d4835bE3b14",
			DAI:   "0x5592EC0cfb4dbc12D3aB100b257153436a1f0FEa",
			CETH:  "0xd6801a1DfFCd0a410336Ef88DeF4320D6DF1883e",
		},
	},
}

func ChangeAccount(account int) {

	Auth = GetSignature(Networkid, account)
	//fmt.Println("fromAddress changed: ", FromAddress)

}

func TokenTransfer(AccountId int, amount *big.Int, _tokenAddress string, _toAddress string) {

	transferFnSignature := []byte("transfer(address,uint256)")

	tokenAddress := common.HexToAddress(_tokenAddress)
	toAddress := common.HexToAddress(_toAddress)

	hash := sha3.NewLegacyKeccak256() // old sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	privateKey, err := crypto.HexToECDSA(Network.PrivateKey[AccountId])

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	gasLimit := uint64(345607)
	gasPrice := big.NewInt(3100000000)

	nonce, err := Client.PendingNonceAt(context.Background(), fromAddress)
	value := big.NewInt(0)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("sent tx: %s\n", signedTx.Hash().Hex()) // tx

}

func GetAddress(accId int) common.Address {

	///get fromAddress
	privateKey, err := crypto.HexToECDSA(Network.PrivateKey[accId])

	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA)
}

func AddSettingString(name string, value string) {

	InfoString = append(InfoString, Info{name, value})

}

func GetClient(nid int, sortId int) *ethclient.Client {

	Networkid = nid
	Network = Networks[Networkid]

	c, err := ethclient.Dial(Network.ProviderUrl[sortId])
	if err != nil {
		log.Fatal("getclient err:", err)
	}

	return c

}

func GetSignature(nid int, accId int) *bind.TransactOpts {

	Account = accId

	privateKey, err := crypto.HexToECDSA(Network.PrivateKey[accId])

	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	FromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	//fmt.Println("signed by ", FromAddress)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Value = big.NewInt(0)          // in wei
	auth.GasLimit = uint64(6721975)     // in gwei
	auth.GasPrice = big.NewInt(2 * 1e9) // 2*1e9 = 2 gwei.  1 gwei = 1e9 wei

	return auth

}

func NonceGen() {

	nonce, err := Client.PendingNonceAt(context.Background(), FromAddress)

	if err != nil {
		log.Fatal("NonceGen ", err)
	}

	Auth.Nonce = big.NewInt(int64(nonce))

}

func PrivateToPublic(_privateKey string) common.Address {

	privateKey, err := crypto.HexToECDSA(_privateKey)
	if err != nil {
		log.Fatal("HexToECDSA ", err)
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA)
}

func Init(nid int, acc int) {

	if nid == -1 {
		nid = Networkid
	}

	if acc == -1 {
		acc = Account

	}

	Client = GetClient(nid, 0)

	Auth = GetSignature(nid, acc)

	fmt.Println("Env: NetworkId=", Networkid, ",client=", Network.ProviderUrl[ProviderSortId])

	_, Token[0].Name, Token[0].Symbol, Token[0].Decimals, Token[0].MaxTotalSupply = GetTokenInstance(Network.TokenA)
	_, Token[1].Name, Token[1].Symbol, Token[1].Decimals, Token[1].MaxTotalSupply = GetTokenInstance(Network.TokenB)

	myPrintln(Token[0].Symbol, ":", common.HexToAddress(Network.TokenA))
	myPrintln(Token[1].Symbol, ":", common.HexToAddress(Network.TokenB))
	//	myPrintln("A<B:", Network.TokenA < Network.TokenB)

}
