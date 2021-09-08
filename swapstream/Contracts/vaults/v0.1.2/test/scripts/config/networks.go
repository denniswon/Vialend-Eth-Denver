package config

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strings"
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
	Fee         int64
}

var Networkid = 3
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
		3000,
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
		3000, // fee
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
		3000, // fee
	},

	{ ///3  goerli admin test 1
		"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
		"0x1F98431c8aD98523631AE4a59f267346ea31F984",
		"0xc7853A9E7b602Aafe36b8fb95E4b67a2001FD9C5",                       //new uniswapv3 factory modified
		"284b65567176c10bc010345042b1d9852fcc1c42ae4b76317e6da040318fbe7f", //test admin 2
		"0x48FCb48bb7F70F399E35d9eC95fd2A614960Dcf8",                       //tokenA eWeth
		"0xFdA9705FdB20E9A633D4283AfbFB4a0518418Af8",                       //tokenB  eusdc
		"0x4F211267896C4D3f2388025263AC6BD67B0f2C54",                       //new owner, test user 1
		60,
		"0x3c7fADe1921Bf9D8308D76d7B09cA54839cfF033", //pool tusdc/ tweth 0xBF93aB266Cd9235DaDE543fAd2EeC884D1cCFc0c // 0x3c7fADe1921Bf9D8308D76d7B09cA54839cfF033", eweth/eusdc //pool
		"0x3C3eF6Ad37F107CDd965C4da5f007526B959532f", // tto  token
		"0x40014da7c5B87b7701D1B3681138556667DDEC37", //vault address
		3000, // fee
	},
}

/*
eUSDC 0xFdA9705FdB20E9A633D4283AfbFB4a0518418Af8
	eWeth 0x48FCb48bb7F70F399E35d9eC95fd2A614960Dcf8

		"0x3fF5E22B4be645EF1CCc8C6e32EDe6b35c569AE4",                       //tokenA tWeth
		"0xFA5dF5372c03D4968d128D624e3Afeb61031a777",                       //tokenB  tusdc

*/

func GetSignature(nid int) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(Network.PrivateKey)

	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	FromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("signed by ", FromAddress)

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

func Readstring(msg string) string {

	fmt.Println(msg)
	fmt.Println("---------------------")

	auto := true

	if auto {
		time.Sleep(Network.PendingTime * time.Second)
		return ""
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			text = strings.TrimSuffix(strings.TrimSpace(text), " \n")

			return text

		}
	}
}

func X1E18(x int64) *big.Int {

	e18, _ := new(big.Int).SetString("1000000000000000000", 10)
	bigx := big.NewInt(x)

	return bigx.Mul(bigx, e18)
}

func FloorFloat64ToBigInt(f64 float64) *big.Int {

	// This number doesn't exist in the float64 world,
	// just a number to perform the test.

	if f64 >= math.MaxInt64 || f64 <= math.MinInt64 {
		log.Fatal("f64 is out of int64 range.", err)
	}

	return big.NewInt(int64(math.Floor(f64)))
}
