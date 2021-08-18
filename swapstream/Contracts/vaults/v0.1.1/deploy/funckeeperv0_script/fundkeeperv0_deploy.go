package main

import (
	"context"
	"crypto/ecdsa"

	"fmt"
	"math/big"

	api "goblockchain/Practice/SwapStream/swapstream-v0-core/deploy/FundKeeper"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type nClient struct {
	clientUrl  string
	privateKey string
	pool       string
}

func main() {

	networks := [...]nClient{

		{"http://127.0.0.1:7545",
			"e8ef3a782d9002408f2ca6649b5f95b3e5772364a5abe203f1678817b6093ff0",
			"n/a"},

		{"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7",
			"0x616fb433fb74c2ad03613f7b200a9e02df193d0c"},

		{"https://rinkeby.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7",
			"pool"},
	}

	nid := 1
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
	auth.Value = big.NewInt(0)                  // in wei
	auth.GasLimit = uint64(6721975)             // in units
	auth.GasPrice = big.NewInt(20 * 1000000000) // 1 gwei = 1000000000
	/*auth, err := config.Auth(client)
	if err != nil {
		panic(err)
	}
	*/

	pool := common.HexToAddress(networks[nid].pool)
	protocolFee := big.NewInt(10000)
	maxTotalSupply := big.NewInt(1000000000000000000)

	address, tx, instance, err := api.DeployApi(auth, client, pool, protocolFee, maxTotalSupply)
	if err != nil {
		panic(err)
	}

	fmt.Println("contract address:", address.Hex())

	_, _ = instance, tx
}
