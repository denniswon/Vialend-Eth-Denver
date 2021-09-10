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
	ProviderUrl []string
	Factory     string
	Callee      string
	PrivateKey  []string
	TokenA      string
	TokenB      string
	NewOwner    string
	PendingTime time.Duration
	Pool        string
	BonusToken  string
	Vault       string
	Fee         int64
}

var Networkid = 2 /// 0: mainnet, 1: local, 2: gorlie
var Account = 0
var ProviderSortId = 1
var Auto = true

var Client, err = ethclient.Dial(Networks[Networkid].ProviderUrl[ProviderSortId])
var Auth = GetSignature(Networkid, Account)
var FromAddress common.Address

var Network = Networks[Networkid]

var Networks = [...]Init{

	{ // 0 mainnet
		[]string{""},
		"",           //factory
		"",           //callee
		[]string{""}, //privatekey
		"",           //tokenA tusdc
		"",           //tokenB tweth
		"",           //newOwner
		60,           //pendingtime
		"",           // pool
		"",           // bonus token
		"",           //vault address
		3000,
	},

	{ // 1 local
		[]string{"http://127.0.0.1:7545", "http://127.0.0.1:8545"},
		"0x0c8D15944A4f799D678029523eC1F82c84b85F32", //factory
		"0xE3c433a67e56BD49d93cCA86728C07bE531c2DCc", //callee
		[]string{"e8ef3a782d9002408f2ca6649b5f95b3e5772364a5abe203f1678817b6093ff0",
			"f804a123dd9876c73cef5d198cce0899e6dfc2f851ed2527b003e11cd5383c54"},
		"0x3b88D0E8B11eb7C5fbC63F1Af1B2795DB1724C59", //tokenA tusdc
		"0xeDFBec53F1DA0995ea493ebB0A8Ff630Bb2f1e23", //tokenB tweth
		"0xeBb29c07455113c30810Addc123D0D7Cd8637aea", //newOwner
		10,
		"0xf320395747Ad48f11Ce8eCA67Fcf82bF479CD532", // pool
		"0xD0d1E195c613Cb6eea9308daB69661CAF9760eF9", // bonus token
		"0x5d2c202CA5A4bAB4BB172C278d24eF78e43AD375", //vault address
		3000, // fee
	},

	{ ///2  goerli admin test 1
		[]string{"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"https://goerli.infura.io/v3/06e0f08cb6884c0fac18ff89fd46d131"}, ///

		"0x1F98431c8aD98523631AE4a59f267346ea31F984",
		"0xc7853A9E7b602Aafe36b8fb95E4b67a2001FD9C5", //new uniswapv3 factory modified
		[]string{"284b65567176c10bc010345042b1d9852fcc1c42ae4b76317e6da040318fbe7f", //test admin 2
			"01e8c8df56230b8b6e4ce6371bed124f4f9950c51d64adc581938239724ed5e6",  //test user 2
			"d8cda34b6928af75aff58c60fe9ed3339896b57a13fa88695aa6da7b775cda2a"}, //test admin 3
		"0x48FCb48bb7F70F399E35d9eC95fd2A614960Dcf8", //tokenA eWeth
		"0xFdA9705FdB20E9A633D4283AfbFB4a0518418Af8", //tokenB  eusdc
		"0x4F211267896C4D3f2388025263AC6BD67B0f2C54", //new owner, test user 1
		40, //time pending interval
		"0x3c7fADe1921Bf9D8308D76d7B09cA54839cfF033", //pool tusdc/ tweth 0xBF93aB266Cd9235DaDE543fAd2EeC884D1cCFc0c // 0x3c7fADe1921Bf9D8308D76d7B09cA54839cfF033", eweth/eusdc //pool
		"0x3C3eF6Ad37F107CDd965C4da5f007526B959532f", // tto  token
		"0x5F959B38c7cdCF18276DB0979993Dc674f73458d", //vault address
		3000, // fee
	},
}

func GetSignature(nid int, accId int) *bind.TransactOpts {

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

	if Auto {
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

	if f64 >= math.MaxInt64 || f64 <= math.MinInt64 {
		log.Fatal("f64 is out of int64 range.", err)
	}

	return big.NewInt(int64(math.Floor(f64)))
}
func RoundFloat64ToBigInt(f64 float64) *big.Int {

	if f64 >= math.MaxInt64 || f64 <= math.MinInt64 {
		log.Fatal("f64 is out of int64 range.", err)
	}

	return big.NewInt(int64(math.Round(f64)))
}

func Float64ToBigInt(f64 float64) *big.Int {

	if f64 >= math.MaxInt64 || f64 <= math.MinInt64 {
		log.Fatal("f64 is out of int64 range.", err)
	}

	return big.NewInt(int64(f64))
}

func Pricef(priceInWei *big.Int, decimal int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(priceInWei.String())
	value := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(decimal)))

	//fmt.Println(value) // 25.729324269165216041
	return value
}
