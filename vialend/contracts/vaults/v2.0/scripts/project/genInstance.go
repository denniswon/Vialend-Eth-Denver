package include

import (
	"log"
	"math/big"

	//	"time"

	//	factory "../../../../../../../uniswap/v3/deploy/UniswapV3Factory"
	token "viaroot/deploy/Tokens/erc20/deploy/Token"

	//vault "viaroot/deploy/FeeMaker"
	callee "viaroot/deploy/TestUniswapV3Callee"
	VaultFactory "viaroot/deploy/VaultFactory"
	VaultStrategy "viaroot/deploy/VaultStrategy"
	ViaVault "viaroot/deploy/ViaVault"
	cErc20 "viaroot/deploy/cErc20"
	vault "viaroot/deploy/vialendFeeMaker"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func GetInstance3() (*VaultFactory.Api, *VaultStrategy.Api, *ViaVault.Api) {

	A1, err := VaultFactory.NewApi(common.HexToAddress(Network.VaultFactory), Client)
	if err != nil {
		log.Println("VaultFactory Instance err:", err)
	}

	A2, err := VaultStrategy.NewApi(common.HexToAddress(Network.VaultStrat), Client)
	if err != nil {
		log.Println("VaultStrat Instance err:", err)
	}

	A3, err := ViaVault.NewApi(common.HexToAddress(Network.Vault), Client)
	if err != nil {
		log.Println("ViaVault Instance err:", err)
	}

	return A1, A2, A3
}

func GetVaultInstance() *vault.Api {

	instance, err := vault.NewApi(common.HexToAddress(Network.Vault), Client)
	if err != nil {
		log.Fatal("vaultInstance err:", err)
	}
	return instance
}

func GetVaultInstance2(_addr string) *vault.Api {

	instance, err := vault.NewApi(common.HexToAddress(_addr), Client)
	if err != nil {
		log.Fatal("vaultInstance err:", err)
	}
	return instance
}

func GetCalleeInstance() *callee.Api {

	instance, err := callee.NewApi(common.HexToAddress(Network.Callee), Client)
	if err != nil {
		log.Fatal("CalleeInstance err:", err)
	}
	return instance
}

func GetTokenInstance(TokenAddress string) (*token.Api, string, string, uint8, *big.Int) {

	instance, err := token.NewApi(common.HexToAddress(TokenAddress), Client)
	if err != nil {
		log.Fatal("get token Instance,", err)
	}

	name, err := instance.Name(&bind.CallOpts{})

	symbol, err := instance.Symbol(&bind.CallOpts{})

	decimals, err := instance.Decimals(&bind.CallOpts{})

	maxsupply, err := instance.TotalSupply(&bind.CallOpts{})

	return instance, name, symbol, decimals, maxsupply
}

func GetCTokenInstance(Address string) *cErc20.Api {

	instance, err := cErc20.NewApi(common.HexToAddress(Address), Client)
	if err != nil {
		log.Fatal("get token Instance,", err)
	}

	return instance
}
