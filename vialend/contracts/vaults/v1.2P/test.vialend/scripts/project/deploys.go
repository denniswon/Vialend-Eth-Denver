package include

import (
	"fmt"
	"log"
	"math/big"

	mocktoken "../../../../../Tokens/erc20/deploy/ERC20fixedSupply"
	weth "../../../../../Tokens/erc20/deploy/WETH9"

	vault "../../../deploy/FeeMaker"
	//	vault "../../../deploy/vialendFeeMaker"
	factory "../../../deploy/UniswapV3Factory"
	vialend "../../../deploy/vialendFeeMaker"

	callee "../../../deploy/TestUniswapV3Callee"
	"../config"
	"github.com/ethereum/go-ethereum/common"
	/*

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

	Readstring("Uniswap Factory deploy done, wait for pending ... next... ")

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

	pool := FindPool()
	// ttoken := common.HexToAddress(config.Network.BonusToken)
	protocolFee := big.NewInt(10000)

	maxTotalSupply, ok := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
	if !ok {
		log.Fatal("maxTotalSupply err ")
	}
	var maxTwapDeviation = big.NewInt(20)
	var twapDuration = uint32(2)

	address, tx, instance, err := vault.DeployApi(config.Auth, config.Client,
		pool,
		protocolFee,
		maxTotalSupply,
		maxTwapDeviation,
		twapDuration)

	if err != nil {
		log.Fatal("deploy vault ", err)
	}

	///set auth back to Account
	config.Auth = config.GetSignature(config.Networkid, config.Account)

	//refresh vault address in networks.go
	config.Network.Vault = address.Hex()
	config.AddSettingString("vault address:", address.Hex())

	fmt.Println("vault address:", address.Hex())

	Readstring("Vault deploy done, wait for pending ... next... ")

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

	Readstring("token deploy done, wait for pending ... next... ")

	return address.Hex()
}

func DeployWrappedEther() *weth.Api {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy WETH . ..................")
	fmt.Println("----------------------------------------------")

	config.NonceGen()
	address, tx, instance, err := weth.DeployApi(config.Auth, config.Client)
	if err != nil {
		log.Fatal(err)
	}

	config.Network.LendingContracts.WETH = address.Hex()

	_, _ = instance, tx

	fmt.Println("WETH9 address:", address.Hex())

	Readstring("WETH deployed, wait for pending ... next... ")

	return instance
}

func DeployVialendFeemaker(networkid int, acc int, _protocolfee *big.Int, _uniPortion int, team string) {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy VialendFeemaker ...................")
	fmt.Println("----------------------------------------------")

	config.Networkid = networkid
	config.Network = config.Networks[networkid]
	///require governance. always use account 0 as the deployer
	config.Auth = config.GetSignature(config.Networkid, acc)

	config.NonceGen()

	pool := FindPool() //tokens based on network selection

	protocolFee := _protocolfee

	maxTotalSupply, ok := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
	if !ok {
		log.Fatal("maxTotalSupply err ")
	}
	var maxTwapDeviation = big.NewInt(20000)
	var twapDuration = uint32(2)

	var quoteAmount = big.NewInt(1e18) /// make sure it's the token0 amount for oracle price quote.

	var _weth = common.HexToAddress(config.Network.LendingContracts.WETH)
	var _cEth = common.HexToAddress(config.Network.LendingContracts.CETH)
	var _cToken0 = common.HexToAddress(config.Network.CTOKEN0)
	var _cToken1 = common.HexToAddress(config.Network.CTOKEN1)

	var uniPortionRate = uint8(_uniPortion)

	fmt.Println(".......................Deploy vault ...................")

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
		quoteAmount,
		uniPortionRate,
		common.HexToAddress(team))

	if err != nil {
		log.Fatal("deploy vault ", err)
	}

	//refresh vault address in networks.go
	config.Network.Vault = address.Hex()
	config.AddSettingString("vault address:", address.Hex())

	fmt.Println("vault address:", address.Hex())

	_, _ = instance, tx

	Readstring("Vault deploy done, wait for pending ... next... ")

}

func VaultGen(
	networkId int,
	acc int,
	token0 string,
	token1 string,
	feetier int64,
	_protocolfee *big.Int,
	quoteAmount *big.Int,
	_uniPortion int,
	team string,
	strategy config.Indi) {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Vault Gen ...................")
	fmt.Println("----------------------------------------------")

	temp := config.Networkid
	config.Networkid = networkId

	///require governance. always use account 0 as the deployer
	config.Auth = config.GetSignature(config.Networkid, acc)

	config.NonceGen()

	pool := GetPool(token0, token1, feetier) //tokens based on network selection

	protocolFee := _protocolfee

	maxTotalSupply, ok := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
	if !ok {
		log.Fatal("maxTotalSupply err ")
	}
	var maxTwapDeviation = big.NewInt(20000)
	var twapDuration = uint32(2)
	var _weth = common.HexToAddress(config.Network.LendingContracts.WETH)
	var _cEth = common.HexToAddress(config.Network.LendingContracts.CETH)
	var _cToken0 = common.HexToAddress(config.Network.CTOKEN0)
	var _cToken1 = common.HexToAddress(config.Network.CTOKEN1)

	var uniPortionRate = uint8(_uniPortion)

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
		quoteAmount,
		uniPortionRate,
		common.HexToAddress(team))

	///set auth back to Account

	config.Networkid = temp
	config.Auth = config.GetSignature(config.Networkid, config.Account)

	if err != nil {
		log.Fatal("deploy vault ", err)
	}

	//refresh vault address in networks.go

	config.Network.Vault = address.Hex()
	config.AddSettingString("vault address:", address.Hex())

	fmt.Println("vault address:", address.Hex())

	Readstring("Vault deploy done, wait for pending ... next... ")

	_, _ = instance, tx

}

func DeployCallee() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy Test Callee . ..................")
	fmt.Println("----------------------------------------------")

	config.NonceGen()
	address, tx, instance, err := callee.DeployApi(config.Auth, config.Client)
	if err != nil {
		log.Fatal(err)
	}

	config.Network.Callee = address.Hex()

	_, _ = instance, tx

	fmt.Println("callee address:", address.Hex())

	Readstring("Callee deploy done, wait for pending ... next... ")

}
