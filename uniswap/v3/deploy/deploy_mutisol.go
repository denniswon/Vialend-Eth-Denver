package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	//api "./UniswapV3Factory"
	api "./TestUniswapV3Callee"
)

func main() {

	// ganache 0xfd8a5AE495Df1CA34F90572cb99A33B27173eDe1
	// deployed contract: 0x7ea14B11c22B6e55d881Bda5361284F58F365fAc
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA("e8ef3a782d9002408f2ca6649b5f95b3e5772364a5abe203f1678817b6093ff0")

	//goerli 0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E
	//contract: 0x73728B3377b36D709c0125D4CA2ec30C636f32EB
	//client, err := ethclient.Dial("https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6")
	//	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/68070d464ba04080a428aeef1b9803c6")
	//privateKey, err := crypto.HexToECDSA("2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7")

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
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units
	auth.GasPrice = big.NewInt(2000000000)

	address, tx, instance, err := api.DeployApi(auth, client)
	if err != nil {
		panic(err)
	}

	fmt.Println("contract address:", address.Hex())

	_, _ = instance, tx

}
