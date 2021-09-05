package config

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Init struct {
	ClientUrl   string
	Factory     string
	Callee      string
	PrivateKey  string
	TokenA      string
	TokenB      string
	NewOwner    string
	PendingTime time.Duration
	Pool        string
	BonusToken  string
	Vault       string
}

var Networkid = 2
var Client, err = ethclient.Dial(Networks[Networkid].ClientUrl)
var Auth = GetSignature(Networkid)
var FromAddress common.Address

var Network = Networks[Networkid]

var Networks = [...]Init{

	{ // 0 mainnet
		"",
		"", //factory
		"", //callee
		"", //privatekey
		"", //tokenA tusdc
		"", //tokenB tweth
		"", //newOwner
		60, //pendingtime
		"", // pool
		"", // bonus token
		"", //vault address
	},

	{ // 1 local
		"http://127.0.0.1:7545",
		"0x42414849A1f13b4d17C2B2eCd2dBfc6124567416", //factory
		"0xE3c433a67e56BD49d93cCA86728C07bE531c2DCc", //callee
		"e8ef3a782d9002408f2ca6649b5f95b3e5772364a5abe203f1678817b6093ff0",
		"0x3b88D0E8B11eb7C5fbC63F1Af1B2795DB1724C59", //tokenA tusdc
		"0xeDFBec53F1DA0995ea493ebB0A8Ff630Bb2f1e23", //tokenB tweth
		"0xeBb29c07455113c30810Addc123D0D7Cd8637aea", //newOwner
		10,
		"0x606B7f7093aa4df2282763F4e9f714706221838b", // pool
		"0xD0d1E195c613Cb6eea9308daB69661CAF9760eF9", // bonus token
		"0xF2510458eaf4dC13Fa80bef475D857F6775b38C5", //vault address
	},

	{ // 2 local user2
		"http://127.0.0.1:7545",
		"0xDd9A27A42493AAebBf4d46Af476fd045acD467f5",                       // factory
		"0xE3c433a67e56BD49d93cCA86728C07bE531c2DCc",                       //callee
		"f804a123dd9876c73cef5d198cce0899e6dfc2f851ed2527b003e11cd5383c54", //(0xeBb29c07455113c30810Addc123D0D7Cd8637aea)
		"0x3b88D0E8B11eb7C5fbC63F1Af1B2795DB1724C59",                       //tokenA tusdc
		"0xeDFBec53F1DA0995ea493ebB0A8Ff630Bb2f1e23",                       //tokenB tweth
		"0xeBb29c07455113c30810Addc123D0D7Cd8637aea",                       //new owner
		10,
		"0x606B7f7093aa4df2282763F4e9f714706221838b", //pool
		"0xD0d1E195c613Cb6eea9308daB69661CAF9760eF9", // bonus token
		"0xF2510458eaf4dC13Fa80bef475D857F6775b38C5", //vault address
	},

	{ ///2  goerli admin test 1
		"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
		"",
		"0xc7853A9E7b602Aafe36b8fb95E4b67a2001FD9C5", //new uniswapv3factory modified requires
		"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7",
		"0x2aDEca523FbBF0937A9419924feAB607Bf599311", //tokenA
		"0xc4AFB13b10f7C49Af721860A188D6443D0fF8747", //tokenB
		"0x4F211267896C4D3f2388025263AC6BD67B0f2C54", //new owner, test user 1
		30,
		"", //pool
		"0x3C3eF6Ad37F107CDd965C4da5f007526B959532f", // tto  token
		"", //vault address
	},
}

func GetSignature(nid int) *bind.TransactOpts {

	network := Networks[nid]

	privateKey, err := crypto.HexToECDSA(network.PrivateKey)

	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	FromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Value = big.NewInt(0)                 // in wei
	auth.GasLimit = uint64(6721975)            // in gwei
	auth.GasPrice = big.NewInt(8 * 1000000000) // 20 gwei.  1 gwei = 1000000000 wei

	return auth

}

func NonceGen() {

	nonce, err := Client.PendingNonceAt(context.Background(), FromAddress)

	if err != nil {
		log.Fatal("NonceGen ", err)
	}

	Auth.Nonce = big.NewInt(int64(nonce))

}
