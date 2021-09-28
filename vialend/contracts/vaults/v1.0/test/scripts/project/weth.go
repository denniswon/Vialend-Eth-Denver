package include

import (
	"fmt"
	"log"

	weth "../../../../../Tokens/erc20/deploy/WETH9"
	"../config"

	"github.com/ethereum/go-ethereum/common"
)

var WETH = config.Network.LendingContracts.WETH

func DeployWETH() {
	DeployWrappedEther()
}

func Test_weth_deposit() {

	fmt.Println("Env: NetworkId=", config.Networkid, ",client=", config.Network.ProviderUrl[config.ProviderSortId])

	instance := GetWethInstance()

	//weth deposit
	ethAmount := config.X1E18(1)
	config.Auth.Value = ethAmount
	tx, err := instance.Deposit(config.Auth)

	if err != nil {
		log.Fatal("weth deposit err ", err)
	}

	fmt.Println("weth deposit 1 eth tx: ", tx.Hash().Hex())

}

func Test_weth_withdraw() {

	fmt.Println("Env: NetworkId=", config.Networkid, ",client=", config.Network.ProviderUrl[config.ProviderSortId])

	instance := GetWethInstance()

	//weth deposit
	ethAmount := config.X1E18(1)

	tx, err := instance.Withdraw(config.Auth, ethAmount)

	if err != nil {
		log.Fatal("weth withdraw err ", err)
	}

	fmt.Println("weth withdraw 1 eth tx: ", tx.Hash().Hex())

}

func GetWethInstance() *weth.Api {

	instance, err := weth.NewApi(common.HexToAddress(config.Network.LendingContracts.WETH), config.Client)
	if err != nil {
		log.Fatal("get token Instance,", err)
	}

	return instance
}
