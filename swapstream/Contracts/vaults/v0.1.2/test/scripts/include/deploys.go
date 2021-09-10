package include

import (
	"fmt"
	"log"
	"math/big"

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

func DeployFactory(do int) *factory.Api {

	if do <= 0 {
		return nil
	}
	fmt.Println(".......................Deploy Uniswap Factory. ..................")

	config.NonceGen()
	address, tx, instance, err := factory.DeployApi(config.Auth, config.Client)
	if err != nil {
		log.Fatal(err)
	}

	config.Network.Factory = address.Hex()

	_, _ = instance, tx

	config.Readstring("Uniswap Factory deploy done, wait for pending ... next... ")

	fmt.Println("factory address:", address.Hex())

	return instance
}

func DeployVault(do int) {

	if do <= 0 {
		return
	}

	fmt.Println(".......................Deploy Vault ...................")

	///require governance. always use account 0 as the deployer
	config.Auth = config.GetSignature(config.Networkid, 0)

	config.NonceGen()

	pool := GetPoolFromToken()
	ttoken := common.HexToAddress(config.Network.BonusToken)
	protocolFee := big.NewInt(10000)
	maxTotalSupply := big.NewInt(1e18).Mul(big.NewInt(1e18), big.NewInt(1e18))

	address, tx, instance, err := vault.DeployApi(config.Auth, config.Client, pool, ttoken, protocolFee, maxTotalSupply)

	///set auth back to Account
	config.Auth = config.GetSignature(config.Networkid, config.Account)

	if err != nil {
		log.Fatal("deploy vault ", err)
	}

	config.Network.Vault = address.Hex()

	config.Readstring("Vault deploy done, wait for pending ... next... ")

	fmt.Println("vault address:", address.Hex())

	_, _ = instance, tx

}
