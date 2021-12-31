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

func Test_weth_deposit(WETH string, accId int, amt int64) {

	fmt.Println("Env: NetworkId=", config.Networkid, ",client=", config.Network.ProviderUrl[config.ProviderSortId])

	fmt.Println("----------------------------------------------")
	fmt.Println(".........wrap .........  ")
	fmt.Println("----------------------------------------------")
	fmt.Println("WETH ADDRESS:", WETH)

	config.ChangeAccount(accId)

	instance := GetWethInstance(common.HexToAddress(WETH))

	//weth deposit
	ethAmount := X1E18(amt)

	config.Auth.Value = ethAmount

	tx, err := instance.Deposit(config.Auth) // value is eth

	if err != nil {
		log.Fatal("weth deposit err ", err)
	}

	fmt.Println("weth deposit 1 eth tx: ", tx.Hash().Hex())
	fmt.Println("wrapped eth amount:", amt, " to: ", config.FromAddress)

	config.ChangeAccount(config.Account)

}

func Test_weth_withdraw(WETH string, accId int, amt int64) {

	fmt.Println("Env: NetworkId=", config.Networkid, ",client=", config.Network.ProviderUrl[config.ProviderSortId])
	fmt.Println("----------------------------------------------")
	fmt.Println(".........unwrap .........  ")
	fmt.Println("----------------------------------------------")
	fmt.Println("WETH ADDRESS:", WETH)

	config.ChangeAccount(accId)

	instance := GetWethInstance(common.HexToAddress(WETH))

	//weth deposit
	ethAmount := X1E18(amt)

	tx, err := instance.Withdraw(config.Auth, ethAmount)

	if err != nil {
		log.Fatal("weth withdraw err ", err)
	}

	fmt.Println("weth withdraw 1 eth tx: ", tx.Hash().Hex())
	fmt.Println("unwrapped weth amount:", amt, " to: ", config.FromAddress)

	config.ChangeAccount(config.Account)

}

func GetWethInstance(WETH common.Address) *weth.Api {

	instance, err := weth.NewApi(WETH, config.Client)
	if err != nil {
		log.Fatal("get token Instance,", err)
	}

	return instance
}
