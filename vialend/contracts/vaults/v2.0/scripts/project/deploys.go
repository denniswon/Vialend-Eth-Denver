package include

import (
	"fmt"
	"log"
	"math/big"
	"reflect"

	AB "viaroot/deploy/AB"
	vault "viaroot/deploy/FeeMaker"
	StratDeployer "viaroot/deploy/StratDeployer"
	callee "viaroot/deploy/TestUniswapV3Callee"
	mocktoken "viaroot/deploy/Tokens/erc20/deploy/ERC20fixedSupply"
	weth "viaroot/deploy/Tokens/erc20/deploy/WETH9"
	factory "viaroot/deploy/UniswapV3Factory"
	VaultDeployer "viaroot/deploy/VaultDeployer"
	VaultFactory "viaroot/deploy/VaultFactory"
	VaultStrategy "viaroot/deploy/VaultStrategy"
	ViaVault "viaroot/deploy/ViaVault"

	//	VaultStrategy "viaroot/deploy/VaultStrategy"
	//	ViaVault "viaroot/deploy/ViaVault"
	arb "viaroot/deploy/arb"
	_ "viaroot/deploy/cether"
	admin "viaroot/deploy/vaultAdmin"
	bridge "viaroot/deploy/vaultBridge"
	vialend "viaroot/deploy/vialendFeeMaker"

	//ViaVault "viaroot/deploy/ViaVault"

	//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

	fmt.Println("factory address:", address.Hex())

	Readstring("Uniswap Factory deploy done, wait for pending ... next... ")

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

	address, tx, instance, err := callee.DeployApi(Auth, Client)
	if err != nil {
		log.Fatal(err)
	}

	TxConfirm(tx.Hash())

	Network.Callee = address.Hex()

	_, _ = instance, tx

	myPrintln("callee address:", address)

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

func DeployVialendFeemaker(networkid int, acc int, _protocolfee *big.Int, _uniPortion int, _quoteAmount *big.Int, team string) {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy VialendFeemaker ...................")
	fmt.Println("----------------------------------------------")

	if networkid != -1 {
		Networkid = networkid
		Network = Networks[networkid]
	}

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

	var quoteAmount = _quoteAmount // big.NewInt(1e18) /// make sure it's the token0 amount for oracle price quote.

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

	Readstring("Vault deployed, pending process... next... ")

}

func DeployAB() *AB.B {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy AB. ..................")
	fmt.Println("----------------------------------------------")

	NonceGen()

	address, tx, instance, err := AB.DeployB(Auth, Client)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = instance, tx

	fmt.Println("B address:", address.Hex())

	//Readstring("Uniswap Factory deploy done, wait for pending ... next... ")

	TxConfirm(tx.Hash())

	return instance
}

func DeployVaultFactory() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy VaultFactory. ..................")
	fmt.Println("----------------------------------------------")

	NonceGen()

	address, tx, instance, err := VaultFactory.DeployApi(Auth, Client, common.HexToAddress("0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440"))

	if err != nil {
		log.Fatal(err)
	}

	_, _ = instance, tx

	Network.VaultFactory = address.Hex()
	fmt.Println("VaultFactory address:", address.Hex())
	Cfg.Contracts.VAULT = address.Hex()

	//Readstring("Uniswap Factory deploy done, wait for pending ... next... ")
	TxConfirm(tx.Hash())

}

func DeployVaultDeployer() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy VaultDeployer. ..................")
	fmt.Println("----------------------------------------------")

	NonceGen()

	address, tx, instance, err := VaultDeployer.DeployApi(Auth, Client)

	if err != nil {
		log.Fatal(err)
	}

	_, _, _ = address, instance, tx

	Cfg.Contracts.VAULT_DEPLOYER = address.Hex()

	//Readstring("Uniswap Factory deploy done, wait for pending ... next... ")
	TxConfirm(tx.Hash())
}

func DeployVaultByDeployer() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy Vault by Deployer. ..................")
	fmt.Println("----------------------------------------------")

	NonceGen()

	myPrintln("DEPLOYER:", Cfg.Contracts.VAULT_DEPLOYER)

	instance, err := VaultDeployer.NewApi(common.HexToAddress(Cfg.Contracts.VAULT_DEPLOYER), Client)
	if err != nil {
		log.Fatal("vault deployer Instance err:", err)
	}

	viaAdmin := FromAddress
	token0 := common.HexToAddress(Network.TokenA)
	token1 := common.HexToAddress(Network.TokenB)
	vaultCap, _ := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
	individualCap, _ := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
	factory := common.HexToAddress(Cfg.Contracts.VAULT_FACTORY)

	tx, err := instance.DeployVault(Auth, factory, viaAdmin, token0, token1, vaultCap, individualCap)

	//Readstring("Uniswap Factory deploy done, wait for pending ... next... ")
	TxConfirm(tx.Hash())

}

func DeployVaultByGo() string {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy Vault by Go. ..................")
	fmt.Println("----------------------------------------------")

	NonceGen()

	viaAdmin := FromAddress
	token0 := common.HexToAddress(Network.TokenA)
	token1 := common.HexToAddress(Network.TokenB)
	vaultCap, _ := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
	individualCap, _ := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
	factory := common.HexToAddress(Cfg.Contracts.VAULT_FACTORY)

	address, tx, instance, err := ViaVault.DeployApi(Auth, Client, factory, viaAdmin, token0, token1, vaultCap, individualCap)

	if err != nil {
		log.Fatal("vault deploy by go err:", err)
	}
	_ = instance

	//Readstring("Uniswap Factory deploy done, wait for pending ... next... ")
	TxConfirm(tx.Hash())

	Cfg.Contracts.VAULT = address.Hex()

	return (Cfg.Contracts.VAULT)

}

func DeployStratByDeployer() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy Strat by Deployer. ..................")
	fmt.Println("----------------------------------------------")

	NonceGen()

	_deployer := Cfg.Contracts.STRAT_DEPLOYER

	myPrintln("DEPLOYER:", _deployer)

	deployerInstance, err := StratDeployer.NewApi(common.HexToAddress(_deployer), Client)
	if err != nil {
		log.Println("strat by deployer Instance err:", err)
	}

	//	pool := FindPool() //tokens based on network selection
	vaultCap, _ := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
	individualCap, _ := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)

	var quoteAmount = big.NewInt(1e18) /// make sure it's the token0 amount for oracle price quote.
	var uniPortionRate = uint8(30)
	var compPortionRate = uint8(70)

	networkid := uint8(5) // 5: goerli
	VaultFactory := common.HexToAddress(Network.VaultFactory)
	protocolAddr := common.HexToAddress("0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440")
	creatorFee := uint8(10)
	vault := common.HexToAddress(Cfg.Contracts.VAULT)

	contracts := [11]common.Address{
		protocolAddr,
		FromAddress,
		common.HexToAddress(Network.LendingContracts.WETH),
		common.HexToAddress(Network.LendingContracts.CETH),
		common.HexToAddress(Network.CTOKEN0),
		common.HexToAddress(Network.CTOKEN1),
		common.HexToAddress(Network.TokenA),
		common.HexToAddress(Network.TokenB),
		vault,
		FromAddress,
		VaultFactory,
	}

	_ = vaultCap
	_ = individualCap

	myPrintln("VaultFactory:", VaultFactory)
	myPrintln("Network:", networkid)
	tx, err := deployerInstance.DeployStrategy(Auth,
		contracts,
		uniPortionRate,
		compPortionRate,
		creatorFee,
		big.NewInt(Network.FeeTier),
		quoteAmount)

	if err != nil {
		log.Fatal(err)
	}

	TxConfirm(tx.Hash())
	//WriteINI([]string{VaultFactory.Hex(), stratAddress, vaultAddress.Hex()})

}

func DeployStratByGoStruct() string {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy Strat by Go struct. ..................")
	fmt.Println("----------------------------------------------")

	NonceGen()

	var params VaultStrategy.UniCompParam
	params.Uni3Factory = common.HexToAddress(Network.Factory)
	params.VaultFactory = common.HexToAddress(Cfg.Contracts.VAULT_FACTORY)
	params.Protocol = common.HexToAddress("0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440")
	params.Creator = FromAddress
	params.CETH = common.HexToAddress(Network.LendingContracts.WETH)
	params.WETH = common.HexToAddress(Network.LendingContracts.CETH)
	params.CToken0 = common.HexToAddress(Network.CTOKEN0)
	params.CToken1 = common.HexToAddress(Network.CTOKEN1)
	params.Token0 = common.HexToAddress(Network.TokenA)
	params.Token1 = common.HexToAddress(Network.TokenB)
	params.Token0Decimals = Token[0].Decimals
	params.Token1Decimals = Token[1].Decimals
	params.VaultCap, _ = new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
	params.IndividualCap, _ = new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
	params.UniPortionRate = uint8(30)
	params.CompPortionRate = uint8(70)
	params.FeeTier = big.NewInt(Network.FeeTier)
	params.TwapDuration = uint32(5)
	params.MaxTwapDeviation = big.NewInt(2000)
	params.ProtocolFeeRate = uint8(10)
	params.MotivatorFeeRate = uint8(5)

	address, tx, _, err := VaultStrategy.DeployApi(Auth, Client, params)
	if err != nil {
		log.Fatal("deployAPI panic>>> ", err)
	}

	myPrintln("strat address:", address.Hex())

	TxConfirm(tx.Hash())

	Cfg.Contracts.VAULT_STRATEGY = address.Hex()

	return (address.Hex())
	//WriteINI([]string{VaultFactory.Hex(), stratAddress, vaultAddress.Hex()})

}

// func DeployStratByGo() string {

// 	fmt.Println("----------------------------------------------")
// 	fmt.Println(".......................Deploy Strat by Go. ..................")
// 	fmt.Println("----------------------------------------------")

// 	NonceGen()

// 	//	pool := FindPool() //tokens based on network selection
// 	vaultCap, _ := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
// 	individualCap, _ := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)

// 	var quoteAmount = big.NewInt(1e18) /// make sure it's the token0 amount for oracle price quote.
// 	var uniPortionRate = uint8(30)
// 	var compPortionRate = uint8(70)

// 	protocolAddr := common.HexToAddress("0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440")
// 	creatorFee := uint8(10)
// 	vault := Cfg.Contracts.VAULT

// 	contracts := [11]common.Address{
// 		protocolAddr,
// 		FromAddress,
// 		common.HexToAddress(Network.LendingContracts.WETH),
// 		common.HexToAddress(Network.LendingContracts.CETH),
// 		common.HexToAddress(Network.CTOKEN0),
// 		common.HexToAddress(Network.CTOKEN1),
// 		common.HexToAddress(Network.TokenA),
// 		common.HexToAddress(Network.TokenB),
// 		common.HexToAddress(vault),
// 		FromAddress,
// 		common.HexToAddress(Cfg.Contracts.VAULT_FACTORY),
// 	}

// 	_ = vaultCap
// 	_ = individualCap

// 	address, tx, _, err := VaultStrategy.DeployApi(Auth, Client,
// 		contracts,
// 		uniPortionRate,
// 		compPortionRate,
// 		creatorFee,
// 		big.NewInt(Network.FeeTier),
// 		quoteAmount)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	myPrintln("strat address:", address.Hex())

// 	TxConfirm(tx.Hash())

// 	Cfg.Contracts.VAULT_STRATEGY = address.Hex()

// 	return (address.Hex())
// 	//WriteINI([]string{VaultFactory.Hex(), stratAddress, vaultAddress.Hex()})

// }

func FactoryVault() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Factory Vault. ..................")
	fmt.Println("----------------------------------------------")

	NonceGen()

	factoryInstance, _, _ := GetInstance3()

	//	pool := FindPool() //tokens based on network selection
	vaultCap, _ := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
	individualCap, _ := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)

	var quoteAmount = big.NewInt(1e18) /// make sure it's the token0 amount for oracle price quote.
	var uniPortionRate = uint8(30)
	var compPortionRate = uint8(70)

	networkid := uint8(5) // 5: goerli
	VaultFactory := common.HexToAddress(Network.VaultFactory)
	protocolAddr := common.HexToAddress("0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440")
	creatorFee := uint8(10)
	feetier := Network.FeeTier
	//	vault := common.HexToAddress(Network.Vault)

	_stratD := Cfg.Contracts.STRAT_DEPLOYER
	_vaultD := Cfg.Contracts.VAULT_DEPLOYER
	myPrintln("strategy deployer:", _stratD)
	myPrintln("vault deployer:", _vaultD)

	contracts := [10]common.Address{
		protocolAddr,
		protocolAddr,
		common.HexToAddress(Network.LendingContracts.WETH),
		common.HexToAddress(Network.LendingContracts.CETH),
		common.HexToAddress(Network.CTOKEN0),
		common.HexToAddress(Network.CTOKEN1),
		common.HexToAddress(Network.TokenA),
		common.HexToAddress(Network.TokenB),
		common.HexToAddress(_stratD),
		common.HexToAddress(_vaultD),
	}

	_ = vaultCap
	_ = individualCap

	myPrintln("VaultFactory:", VaultFactory)
	myPrintln("Network:", networkid)
	tx, err := factoryInstance.CreateVault(Auth,
		contracts,
		vaultCap,
		individualCap,
		uniPortionRate,
		compPortionRate,
		creatorFee,
		big.NewInt(feetier),
		quoteAmount)

	if err != nil {
		log.Fatal(err)
	}

	TxConfirm(tx.Hash())

	Readstring("via vaults deploys done, wait for pending ... next... ")

	lastone, err := factoryInstance.GetCount(&bind.CallOpts{})
	vaults, err := factoryInstance.Vaults(&bind.CallOpts{}, lastone.Sub(lastone, big.NewInt(1)))
	myPrintln("new Vaults:", vaults)

	Cfg.Contracts.VAULT_STRATEGY = vaults.Strategy.Hex()
	Cfg.Contracts.VAULT = vaults.Vault.Hex()

	// Cfg.Contracts.VAULT_STRATEGY = "new strat3"
	// Cfg.Contracts.VAULT = "new vault3"

	//WriteINI([]string{VaultFactory.Hex(), stratAddress, vaultAddress.Hex()})

}

func DeployStratDeployer() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Deploy StratDeployer. ..................")
	fmt.Println("----------------------------------------------")

	NonceGen()

	address, tx, instance, err := StratDeployer.DeployApi(Auth, Client)

	if err != nil {
		log.Fatal(err)
	}

	_, _, _ = address, instance, tx

	Cfg.Contracts.STRAT_DEPLOYER = address.Hex()

	//Readstring("Uniswap Factory deploy done, wait for pending ... next... ")
	TxConfirm(tx.Hash())
}

// func DeployVaultStrategy() {
// 	fmt.Println("----------------------------------------------")
// 	fmt.Println(".......................Deploy  VaultStrategy and ViaVault. ..................")
// 	fmt.Println("----------------------------------------------")

// 	NonceGen()

// 	pool := FindPool() //tokens based on network selection
// 	vaultCap, _ := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
// 	individualCap, _ := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)

// 	var quoteAmount = big.NewInt(1e18) /// make sure it's the token0 amount for oracle price quote.
// 	var uniPortionRate = uint8(30)
// 	var compPortionRate = uint8(70)

// 	networkid := uint8(5) // 5: goerli
// 	VaultFactory := common.HexToAddress(Network.VaultFactory)
// 	protocolAddr := common.HexToAddress("0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440")
// 	creatorFee := uint8(10)
// 	vault := common.HexToAddress(Network.Vault)

// 	contracts := [7]common.Address{
// 		protocolAddr,
// 		pool,
// 		common.HexToAddress(Network.LendingContracts.WETH),
// 		common.HexToAddress(Network.LendingContracts.CETH),
// 		common.HexToAddress(Network.CTOKEN0),
// 		common.HexToAddress(Network.CTOKEN1),
// 		vault,
// 	}

// 	_ = vaultCap
// 	_ = individualCap

// 	myPrintln("VaultFactory:", VaultFactory)
// 	myPrintln("Network:", networkid)
// 	address, tx, instanceStrat, err := VaultStrategy.DeployApi(Auth, Client,
// 		contracts,
// 		uniPortionRate,
// 		compPortionRate,
// 		creatorFee,
// 		quoteAmount)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	TxConfirm(tx.Hash())

// 	stratAddress := address.Hex()
// 	fmt.Println("StratUniComp address:", stratAddress)

// 	Readstring2("strat uni comp deploy done, wait for pending ... next deploy viaVault... ", false)

// 	vaultAddress, _ := instanceStrat.Vault(&bind.CallOpts{})
// 	fmt.Println("viaVault Address:", vaultAddress)

// 	stat, _ := instanceStrat.Stat(&bind.CallOpts{})
// 	fmt.Println("vault status:", stat)

// 	if IsZeroAddress(vaultAddress) {
// 		log.Fatal(">>>>FAILED......vaultAddress  is 0: ", vaultAddress)
// 	}

// 	WriteINI([]string{VaultFactory.Hex(), stratAddress, vaultAddress.Hex()})

// }

// func DeployViaVault() {
// 	fmt.Println("----------------------------------------------")
// 	fmt.Println(".......................Deploy  ViaVault. ..................")
// 	fmt.Println("----------------------------------------------")

// 	NonceGen()

// 	vaultCap, _ := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
// 	individualCap, _ := new(big.Int).SetString("9999999999999999999999999999999999999999", 10)
// 	admin := FromAddress
// 	strategy := common.HexToAddress(Network.VaultStrat)
// 	token0 := common.HexToAddress(Network.TokenA)
// 	token1 := common.HexToAddress(Network.TokenB)

// 	address, tx, _, err := ViaVault.DeployApi(Auth, Client,
// 		admin,
// 		token0,
// 		token1,
// 		vaultCap,
// 		individualCap,
// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	TxConfirm(tx.Hash())

// 	if IsZeroAddress(address) {
// 		log.Fatal(">>>>FAILED......vaultAddress  is 0: ", address)
// 	}

// }

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}
