package include

import (
	"fmt"
	"log"
	"math/big"
	"time"

	factory "../../../../../../../uniswap/v3/deploy/UniswapV3Factory"
	vault "../../../deploy/FeeMaker"
	"../config"
	"github.com/ethereum/go-ethereum/common"
	/*

		mintcallee "../../uniswap/v3/deploy/TestUniswapV3Callee"
		factory "../../uniswap/v3/deploy/UniswapV3Factory"
		pool "../../uniswap/v3/deploy/UniswapV3Pool"
		token "../../uniswap/v3/deploy/token"
	*/)

func DeployFactory(do bool) *factory.Api {

	if !do {
		return nil
	}

	config.NonceGen()
	address, tx, instance, err := factory.DeployApi(config.Auth, config.Client)
	if err != nil {
		log.Fatal(err)
	}

	config.Network.Factory = address.Hex()

	fmt.Println("factory address:", address.Hex())

	_, _ = instance, tx

	time.Sleep(config.Network.PendingTime * time.Second)
	//	fmt.Println("factory:", config.Network1.Factory)

	return instance
}

func DeployVault(do bool) {

	if !do {
		return
	}

	config.NonceGen()

	pool := GetPoolFromToken(true)
	ttoken := common.HexToAddress(config.Network.BonusToken)
	protocolFee := big.NewInt(10000)
	maxTotalSupply := big.NewInt(1e18).Mul(big.NewInt(1e18), big.NewInt(1e18))

	address, tx, instance, err := vault.DeployApi(config.Auth, config.Client, pool, ttoken, protocolFee, maxTotalSupply)

	if err != nil {
		log.Fatal("deploy vault ", err)
	}

	config.Network.Vault = address.Hex()
	fmt.Println("vault address:", address.Hex())

	time.Sleep(config.Network.PendingTime * time.Second)

	_, _ = instance, tx

}
