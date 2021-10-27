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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

type TokenStruct struct {
	Name           string
	Symbol         string
	Decimals       uint8
	MaxTotalSupply *big.Int
}

type LendingStruct struct {
	WETH  string
	DAI   string
	CDAI  string
	ETH   string
	CETH  string
	WBTC  string
	CWBTC string
	USDC  string
	CUSDC string
}
type Init struct {
	ProviderUrl      []string
	Factory          string
	Callee           string
	PrivateKey       []string
	TokenA           string
	TokenB           string
	NewOwner         string
	PendingTime      time.Duration
	Pool             string
	BonusToken       string
	Vault            string
	FeeTier          int64
	LendingContracts LendingStruct
}

var Networkid = 3 /// 0: mainnet, 1: local, 2: local , 3: gorlie, 4: rinkeby
var Account = 3
var ProviderSortId = 0
var Auto = true

var Token [2]TokenStruct

//var Decimals0, Decimals1 uint8

var Client, err = ethclient.Dial(Networks[Networkid].ProviderUrl[ProviderSortId])
var Auth = GetSignature(Networkid, Account)
var FromAddress common.Address

var Network = Networks[Networkid]

var DEBUG = true

type Info struct {
	Name  string
	Value string
}

var InfoString []Info

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
		LendingStruct{
			WBTC:  "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599",
			CWBTC: "0xC11b1268C1A384e55C48c2391d8d480264A3A7F4",
			USDC:  "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			CUSDC: "0x39AA39c021dfbaE8faC545936693aC917d5E7563",
			CETH:  "0x4Ddc2D193948926D02f9B1fE9e1daa0718270ED5",
		},
	},

	{ // 1 local usdc/usdt
		[]string{"http://127.0.0.1:7545", "http://127.0.0.1:8545"},
		"0x0c8D15944A4f799D678029523eC1F82c84b85F32", //factory
		"0x210FA31C72D9F020D16BF948e54F108D1C688f81", //callee
		[]string{"e8ef3a782d9002408f2ca6649b5f95b3e5772364a5abe203f1678817b6093ff0",
			"f804a123dd9876c73cef5d198cce0899e6dfc2f851ed2527b003e11cd5383c54"},
		"0x59Cd9D486a8fA9b39F715915743997daA12d138e", //tokenB usdt
		"0x9D96eC63f96A4E985e227BF520dD742315AB77c7", //tokenA usdc
		"0xeBb29c07455113c30810Addc123D0D7Cd8637aea", //newOwner
		10,
		"0x96Dd142387281a16F72962CCb659cAc67D73d882", // pool
		"0xD0d1E195c613Cb6eea9308daB69661CAF9760eF9", // bonus token
		"0xcB0b392e747C101Ed949247730eC3aa6A75E4D3B", //vault address
		3000, // fee
		LendingStruct{DAI: "ASDF", CDAI: "ASD"},
	},

	{ // 2 local weth/usdc
		[]string{"http://127.0.0.1:7545", "http://127.0.0.1:8545"},
		"0x0c8D15944A4f799D678029523eC1F82c84b85F32", //factory
		"0x210FA31C72D9F020D16BF948e54F108D1C688f81", //callee
		[]string{"e8ef3a782d9002408f2ca6649b5f95b3e5772364a5abe203f1678817b6093ff0",
			"f804a123dd9876c73cef5d198cce0899e6dfc2f851ed2527b003e11cd5383c54"},
		"0xB73A78A3C493ACdbA893da9331ff39Fe4E59bFA3", //e weth1
		"0xd8F4E5E1cE1a2961b5fB401B8c2286549607B294", //e usdc1
		"0xeBb29c07455113c30810Addc123D0D7Cd8637aea", //newOwner
		10,
		"0xaEbc0569A8Ad476530d765dBE6308842Bd4D699c", // pool
		"0xD0d1E195c613Cb6eea9308daB69661CAF9760eF9", // bonus token
		"0x2723f0d5F2E60D1BF686B835e630C55453307eEA", //vault address
		3000, // fee
		LendingStruct{DAI: "ASDF", CDAI: "ASD"},
	},
	{ ///3  goerli admin test 1
		[]string{"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"https://goerli.infura.io/v3/06e0f08cb6884c0fac18ff89fd46d131"}, ///  provider url

		"0x1F98431c8aD98523631AE4a59f267346ea31F984", //factory
		"0xE97f1488F053251032ef358dE5b4188cD960D413", //callee
		[]string{"284b65567176c10bc010345042b1d9852fcc1c42ae4b76317e6da040318fbe7f", //test admin 2
			"01e8c8df56230b8b6e4ce6371bed124f4f9950c51d64adc581938239724ed5e6",  //test user 2
			"d8cda34b6928af75aff58c60fe9ed3339896b57a13fa88695aa6da7b775cda2a",  //test admin 3
			"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7"}, // test admin
		// "0x48FCb48bb7F70F399E35d9eC95fd2A614960Dcf8", //tokenA eWeth
		// "0x6f38602e142D0Bd3BC162f5912535f543D3B73d7", //tokenB  eusdc
		"0x6fD886fd1e728D9386Ba7fE721C856790758aDd9", //tokenA Weth
		"0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C", //tokenB  usdc
		"0x4F211267896C4D3f2388025263AC6BD67B0f2C54", //new owner, test user 1
		40, //time pending interval
		"0x933EFDc68cB4c1fe4Ef162a87E515768d6f82023", //pool 0xc4C92691f69fadDd684257E9f5A8d6f9D2c79a93"(elf), //pool tusdc/ tweth 0xBF93aB266Cd9235DaDE543fAd2EeC884D1cCFc0c // pool(eee usdc/eth) 0x3c7fADe1921Bf9D8308D76d7B09cA54839cfF033"
		"0x3C3eF6Ad37F107CDd965C4da5f007526B959532f", // tto  token
		"0x58b008535dc1d06dbae22201d3f10b79c80f9dd4", // withheld eth in lending -> 0xAda1AC9D1dBFFF1270f0009d8f20bd0332F48113", // latest worked 0xF2444F64C72BA17948A2adB800a1Df4E55036FBa", //vault address

		3000, // fee
		LendingStruct{
			WETH:  "0x6fD886fd1e728D9386Ba7fE721C856790758aDd9", ///on remix solc0.4.12,  injected web3 deployed by 0x2ee9... test admin,
			USDC:  "0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C",
			CUSDC: "0xCEC4a43eBB02f9B80916F1c718338169d6d5C1F0",
			CWBTC: "0x6CE27497A64fFFb5517AA4aeE908b1E7EB63B9fF",
			WBTC:  "0xC04B0d3107736C32e19F1c62b2aF67BE61d63a05",
			CDAI:  "0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb",
			DAI:   "0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60",
			CETH:  "0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF",
		},
	},
	{ ///4  rinkeby tester admin
		[]string{"https://rinkeby.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"https://rinkeby.infura.io/v3/06e0f08cb6884c0fac18ff89fd46d131"}, ///  provider url

		"", //factory
		"", //callee
		[]string{"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7"},
		"",   //tokenA eWeth
		"",   //tokenB  eusdc
		"",   //new owner, test user 1
		40,   //time pending interval
		"",   //pool 0x3c7fADe1921Bf9D8308D76d7B09cA54839cfF033", //pool tusdc/ tweth 0xBF93aB266Cd9235DaDE543fAd2EeC884D1cCFc0c // 0x3c7fADe1921Bf9D8308D76d7B09cA54839cfF033", eweth/eusdc //pool
		"",   // tto  token
		"",   //vault address
		3000, // fee
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
	//fmt.Println("signed by ", FromAddress)

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

func Pricef(priceInWei *big.Int, decimal int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(priceInWei.String())
	value := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(decimal)))

	//	fmt.Println(value) // 25.729324269165216041
	return value
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

func Float64ToBigInt(val float64) *big.Int {

	//price = int(sqrt(price) * (1 << 96))
	newNum := big.NewRat(1, 1)
	newNum.SetFloat64(val)
	bigf := newNum.FloatString(0)

	//fmt.Println("val, bigf:", val,  bigf)
	//os.Exit(3)

	bigI, ok := new(big.Int).SetString(bigf, 10)
	if !ok {
		log.Fatal("float64 to bigInt err ", val, bigI)
	}

	return bigI

}

func BigFloatToBigInt(bigval *big.Float) *big.Int {

	result := new(big.Int)
	bigval.Int(result)

	return result
}

func BigIntToFloat64(bigN *big.Int) float64 {

	bigF, _ := new(big.Float).SetString(bigN.String())

	f, _ := bigF.Float64()

	return float64(f)
}

func AddSettingString(name string, value string) {

	InfoString = append(InfoString, Info{name, value})

}
