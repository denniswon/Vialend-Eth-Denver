package main

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"fmt"
	"math"

	token "goblockchain/Practice/UniswapV3/uniswap-v3-core/deploy/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type nClient struct {
	clientUrl  string
	privateKey string
}

func main() { ///-----------

	networks := [...]nClient{

		{ //0
			"http://127.0.0.1:7545",
			"e8ef3a782d9002408f2ca6649b5f95b3e5772364a5abe203f1678817b6093ff0"},

		{ //1 test admin
			"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7"},

		{ //2 test user 1
			"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"67f7046a9f3712d77dab07a843c91d060ab5f27b808ed54d6db1293c7cd5eff3"},

		{ //3 test user 2
			"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"01e8c8df56230b8b6e4ce6371bed124f4f9950c51d64adc581938239724ed5e6"},

		{ //4
			"https://rinkeby.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7"},
	}

	nid := 0

	client, err := ethclient.Dial(networks[nid].clientUrl)
	privateKey, err := crypto.HexToECDSA(networks[nid].privateKey)

	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)                 // in wei
	auth.GasLimit = uint64(6721975)            // in units
	auth.GasPrice = big.NewInt(8 * 1000000000) // 1 gwei = 1000000000

	token0Address := common.HexToAddress("0x034A156169e7f3239eDC318747f305d7aA161d8B")
	//	token1Address := common.HexToAddress("0xFA5dF5372c03D4968d128D624e3Afeb61031a777")

	tokenInstance, err := token.NewApi(token0Address, client)
	if err != nil {
		panic(err)
	}

	fmt.Println("network:", networks[nid].clientUrl)

	fmt.Println("from address:", fromAddress)

	fmt.Println("token address:", token0Address)

	totalSupply, _ := tokenInstance.TotalSupply(&bind.CallOpts{})
	fmt.Println("totalsupply:", totalSupply)

	decimals, _ := tokenInstance.Decimals(&bind.CallOpts{})
	fmt.Println("decimals:", decimals)
	name, _ := tokenInstance.Name(&bind.CallOpts{})
	fmt.Println("name:", name)
	symbol, _ := tokenInstance.Symbol(&bind.CallOpts{})
	fmt.Println("symbol:", symbol)

	bal, _ := tokenInstance.BalanceOf(&bind.CallOpts{}, fromAddress)
	//convert to float64

	balanceOf := new(big.Float).SetInt(bal)

	res := new(big.Float).Quo(balanceOf, big.NewFloat(math.Pow(10, float64(decimals))))

	fmt.Println("balanceOf fromAddress:", balanceOf)
	fmt.Println("adjusted balanceOf fromAddress:", res)

}
