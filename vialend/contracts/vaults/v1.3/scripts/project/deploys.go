package include

import (
	"fmt"
	"log"
	"math/big"

	weth "viaroot/Tokens/erc20/deploy/WETH9"
	vault "viaroot/deploy/FeeMaker"
	callee "viaroot/deploy/TestUniswapV3Callee"
	factory "viaroot/deploy/UniswapV3Factory"
	arb "viaroot/deploy/arb"
	_ "viaroot/deploy/cether"
	admin "viaroot/deploy/vaultAdmin"
	bridge "viaroot/deploy/vaultBridge"
	vialend "viaroot/deploy/vialendFeeMaker"

	mocktoken "viaroot/Tokens/erc20/deploy/ERC20fixedSupply"

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

	NonceGen()
	address, tx, instance, err := factory.DeployApi(Auth, Client)
	if err != nil {
		log.Fatal(err)
	}

	Network.Factory = address.Hex()

	_, _ = instance, tx

	Readstring("Uniswap Factory deploy done, wait for pending ... next... ")

	fmt.Println("factory address:", address.Hex())

	AddSettingString("factory address:", address.Hex())

	return instance
}

func DeployVault() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy Vault ...................")
	fmt.Println("----------------------------------------------")

	///require governance. always use account 0 as the deployer
	Auth = GetSignature(Networkid, 0)

	NonceGen()

	pool := FindPool()
	// ttoken := common.HexToAddress(Network.BonusToken)
	protocolFee := big.NewInt(10000)

	maxTotalSupply, ok := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
	if !ok {
		log.Fatal("maxTotalSupply err ")
	}
	var maxTwapDeviation = big.NewInt(20)
	var twapDuration = uint32(2)

	address, tx, instance, err := vault.DeployApi(Auth, Client,
		pool,
		protocolFee,
		maxTotalSupply,
		maxTwapDeviation,
		twapDuration)

	if err != nil {
		log.Fatal("deploy vault ", err)
	}

	///set auth back to Account
	Auth = GetSignature(Networkid, Account)

	//refresh vault address in networks.go
	Network.Vault = address.Hex()
	AddSettingString("vault address:", address.Hex())

	fmt.Println("vault address:", address.Hex())

	Readstring("Vault deploy done, wait for pending ... next... ")

	_, _ = instance, tx

}

func DeployToken(name string, symbol string, decimals uint8, totalSupply *big.Int) string {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy Token. ..................")
	fmt.Println("----------------------------------------------")

	NonceGen()

	address, tx, instance, err := mocktoken.DeployApi(Auth, Client, name, symbol, decimals, totalSupply)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = instance, tx

	fmt.Println("token address:", address.Hex())
	AddSettingString(symbol+" token address:", address.Hex())

	Readstring("token deploy done, wait for pending ... next... ")

	return address.Hex()
}

func DeployWrappedEther() *weth.Api {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy WETH . ..................")
	fmt.Println("----------------------------------------------")

	NonceGen()
	address, tx, instance, err := weth.DeployApi(Auth, Client)
	if err != nil {
		log.Fatal(err)
	}

	Network.LendingContracts.WETH = address.Hex()

	_, _ = instance, tx

	fmt.Println("WETH9 address:", address.Hex())

	Readstring("WETH deployed, wait for pending ... next... ")

	return instance
}

func DeployVialendFeemaker(networkid int, acc int, _protocolfee *big.Int, _uniPortion int, team string) {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy VialendFeemaker ...................")
	fmt.Println("----------------------------------------------")

	Networkid = networkid
	Network = Networks[networkid]
	///require governance. always use account 0 as the deployer
	Auth = GetSignature(Networkid, acc)

	NonceGen()

	pool := FindPool() //tokens based on network selection

	protocolFee := _protocolfee

	maxTotalSupply, ok := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
	if !ok {
		log.Fatal("maxTotalSupply err ")
	}
	var maxTwapDeviation = big.NewInt(20000)
	var twapDuration = uint32(2)

	var quoteAmount = big.NewInt(1e18) /// make sure it's the token0 amount for oracle price quote.

	var _weth = common.HexToAddress(Network.LendingContracts.WETH)
	var _cEth = common.HexToAddress(Network.LendingContracts.CETH)
	var _cToken0 = common.HexToAddress(Network.CTOKEN0)
	var _cToken1 = common.HexToAddress(Network.CTOKEN1)

	var uniPortionRate = uint8(_uniPortion)

	fmt.Println(".......................Deploy vault ...................")

	address, tx, instance, err := vialend.DeployApi(Auth, Client,
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
	Network.Vault = address.Hex()
	AddSettingString("vault address:", address.Hex())

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
	strategy Indi) {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Vault Gen ...................")
	fmt.Println("----------------------------------------------")

	temp := Networkid
	Networkid = networkId

	///require governance. always use account 0 as the deployer
	Auth = GetSignature(Networkid, acc)

	NonceGen()

	pool := GetPool(token0, token1, feetier) //tokens based on network selection

	protocolFee := _protocolfee

	maxTotalSupply, ok := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
	if !ok {
		log.Fatal("maxTotalSupply err ")
	}
	var maxTwapDeviation = big.NewInt(20000)
	var twapDuration = uint32(2)
	var _weth = common.HexToAddress(Network.LendingContracts.WETH)
	var _cEth = common.HexToAddress(Network.LendingContracts.CETH)
	var _cToken0 = common.HexToAddress(Network.CTOKEN0)
	var _cToken1 = common.HexToAddress(Network.CTOKEN1)

	var uniPortionRate = uint8(_uniPortion)

	address, tx, instance, err := vialend.DeployApi(Auth, Client,
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

	Networkid = temp
	Auth = GetSignature(Networkid, Account)

	if err != nil {
		log.Fatal("deploy vault ", err)
	}

	//refresh vault address in networks.go

	Network.Vault = address.Hex()
	AddSettingString("vault address:", address.Hex())

	fmt.Println("vault address:", address.Hex())

	Readstring("Vault deploy done, wait for pending ... next... ")

	_, _ = instance, tx

}

func DeployCallee() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy Test Callee . ..................")
	fmt.Println("----------------------------------------------")

	NonceGen()
	address, tx, instance, err := callee.DeployApi(Auth, Client)
	if err != nil {
		log.Fatal(err)
	}

	Network.Callee = address.Hex()

	_, _ = instance, tx

	fmt.Println("callee address:", address.Hex())

	Readstring("Callee deploy done, wait for pending ... next... ")

}

func DeployVaultBridge() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy VaultBridge . ..................")
	fmt.Println("----------------------------------------------")

	NonceGen()

	address, tx, instance, err := bridge.DeployApi(Auth, Client)

	if err != nil {
		log.Fatal(err)
	}

	Network.VaultBridge = address.Hex()

	_, _ = instance, tx

	fmt.Println("vaultbridge address:", address.Hex())

}

func DeployVaultAdmin() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy VaultAdmin . ..................")
	fmt.Println("----------------------------------------------")

	NonceGen()
	/*
			local address test in remix
		   ["0xfd8a5AE495Df1CA34F90572cb99A33B27173eDe1","0xeBb29c07455113c30810Addc123D0D7Cd8637aea","0xcCebCE6AF95bfD69EEE193390CCF027e8d47d9e2"]

	*/
	admins := []common.Address{
		PrivateToPublic(Network.PrivateKey[0]),
		PrivateToPublic(Network.PrivateKey[1]),
		PrivateToPublic(Network.PrivateKey[2]),
	}
	address, tx, instance, err := admin.DeployApi(Auth, Client, admins)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = instance, tx

	fmt.Println("vaultAdmin address:", address.Hex())

	TxConfirm(tx.Hash())

}

func DeployArb() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy arb . ..................")
	fmt.Println("----------------------------------------------")

	NonceGen()

	address, tx, instance, err := arb.DeployApi(Auth, Client)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = instance, tx

	fmt.Println("arb address:", address.Hex())

	TxConfirm(tx.Hash())

}
