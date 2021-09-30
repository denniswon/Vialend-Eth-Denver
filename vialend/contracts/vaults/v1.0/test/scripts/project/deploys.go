package include

import (
	"fmt"
	"log"
	"math/big"

	factory "../../../../../../../uniswap/v3/deploy/UniswapV3Factory"
	mocktoken "../../../../../Tokens/erc20/deploy/ERC20fixedSupply"
	weth "../../../../../Tokens/erc20/deploy/WETH9"

	vault "../../../deploy/FeeMaker"
	//	vault "../../../deploy/vialendFeeMaker"
	vialend "../../../deploy/vialendFeeMaker"

	"../config"
	"github.com/ethereum/go-ethereum/common"
	/*

		mintcallee "../../uniswap/v3/deploy/TestUniswapV3Callee"
		factory "../../uniswap/v3/deploy/UniswapV3Factory"
		pool "../../uniswap/v3/deploy/UniswapV3Pool"
		token "../../uniswap/v3/deploy/token"
	*/)

func DeployFactory() *factory.Api {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy Uniswap Factory. ..................")
	fmt.Println("----------------------------------------------")

	config.NonceGen()
	address, tx, instance, err := factory.DeployApi(config.Auth, config.Client)
	if err != nil {
		log.Fatal(err)
	}

	config.Network.Factory = address.Hex()

	_, _ = instance, tx

	config.Readstring("Uniswap Factory deploy done, wait for pending ... next... ")

	fmt.Println("factory address:", address.Hex())

	config.AddSettingString("factory address:", address.Hex())

	return instance
}

func DeployVault() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy Vault ...................")
	fmt.Println("----------------------------------------------")

	///require governance. always use account 0 as the deployer
	config.Auth = config.GetSignature(config.Networkid, 0)

	config.NonceGen()

	pool := GetPoolFromToken()
	ttoken := common.HexToAddress(config.Network.BonusToken)
	protocolFee := big.NewInt(10000)

	maxTotalSupply, ok := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
	if !ok {
		log.Fatal("maxTotalSupply err ")
	}
	var maxTwapDeviation = big.NewInt(20)
	var twapDuration = uint32(2)

	address, tx, instance, err := vault.DeployApi(config.Auth, config.Client, pool, ttoken, protocolFee, maxTotalSupply, maxTwapDeviation, twapDuration)

	///set auth back to Account
	config.Auth = config.GetSignature(config.Networkid, config.Account)

	if err != nil {
		log.Fatal("deploy vault ", err)
	}

	//refresh vault address in networks.go
	config.Network.Vault = address.Hex()
	config.AddSettingString("vault address:", address.Hex())

	fmt.Println("vault address:", address.Hex())

	config.Readstring("Vault deploy done, wait for pending ... next... ")

	_, _ = instance, tx

}

func DeployToken(name string, symbol string, decimals uint8, totalSupply *big.Int) string {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy Token. ..................")
	fmt.Println("----------------------------------------------")

	config.NonceGen()

	address, tx, instance, err := mocktoken.DeployApi(config.Auth, config.Client, name, symbol, decimals, totalSupply)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = instance, tx

	fmt.Println("token address:", address.Hex())
	config.AddSettingString(symbol+" token address:", address.Hex())

	config.Readstring("token deploy done, wait for pending ... next... ")

	return address.Hex()
}

func DeployWrappedEther() *weth.Api {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy Uniswap Factory. ..................")
	fmt.Println("----------------------------------------------")

	config.NonceGen()
	address, tx, instance, err := weth.DeployApi(config.Auth, config.Client)
	if err != nil {
		log.Fatal(err)
	}

	config.Network.LendingContracts.WETH = address.Hex()

	_, _ = instance, tx

	fmt.Println("WETH9 address:", address.Hex())

	config.Readstring("WETH deployed, wait for pending ... next... ")

	return instance
}

func DeployVialendFeemaker(acc int) {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy VialendFeemaker ...................")
	fmt.Println("----------------------------------------------")

	///require governance. always use account 0 as the deployer
	config.Auth = config.GetSignature(config.Networkid, acc)

	config.NonceGen()

	pool := GetPoolFromToken()

	protocolFee := big.NewInt(10000)

	maxTotalSupply, ok := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
	if !ok {
		log.Fatal("maxTotalSupply err ")
	}
	var maxTwapDeviation = big.NewInt(20000)
	var twapDuration = uint32(2)
	var _weth = common.HexToAddress(config.Network.LendingContracts.WETH)
	var _cToken0 = common.HexToAddress(config.Network.LendingContracts.CETH)
	var _cToken1 = common.HexToAddress(config.Network.LendingContracts.CUSDC)
	var _cEth = common.HexToAddress(config.Network.LendingContracts.CETH)

	var _uniPortionRate = uint8(90)

	address, tx, instance, err := vialend.DeployApi(config.Auth, config.Client,
		pool,
		_weth,
		_cToken0,
		_cToken1,
		_cEth,
		protocolFee,
		maxTotalSupply,
		maxTwapDeviation,
		twapDuration,
		_uniPortionRate)

	///set auth back to Account
	config.Auth = config.GetSignature(config.Networkid, config.Account)

	if err != nil {
		log.Fatal("deploy vault ", err)
	}

	//refresh vault address in networks.go
	config.Network.Vault = address.Hex()
	config.AddSettingString("vault address:", address.Hex())

	fmt.Println("vault address:", address.Hex())

	config.Readstring("Vault deploy done, wait for pending ... next... ")

	_, _ = instance, tx

}
