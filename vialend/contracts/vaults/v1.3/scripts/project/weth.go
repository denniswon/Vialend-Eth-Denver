package include

import (
	"fmt"
	"log"

	weth "viaroot/deploy/Tokens/erc20/deploy/WETH9"

	"github.com/ethereum/go-ethereum/common"
)

var WETH = Network.LendingContracts.WETH

func DeployWETH() {
	DeployWrappedEther()
}

func Test_weth_deposit(WETH string, accId int, amt int64) {

	fmt.Println("Env: NetworkId=", Networkid, ",client=", Network.ProviderUrl[ProviderSortId])

	fmt.Println("----------------------------------------------")
	fmt.Println(".........wrap .........  ")
	fmt.Println("----------------------------------------------")
	fmt.Println("WETH ADDRESS:", WETH)

	ChangeAccount(accId)

	instance := GetWethInstance(common.HexToAddress(WETH))

	//weth deposit
	ethAmount := X1E18(amt)

	Auth.Value = ethAmount

	tx, err := instance.Deposit(Auth) // value is eth

	if err != nil {
		log.Fatal("weth deposit err ", err)
	}

	fmt.Println("weth deposit 1 eth tx: ", tx.Hash().Hex())
	fmt.Println("wrapped eth amount:", amt, " to: ", FromAddress)

	ChangeAccount(Account)

}

func Test_weth_withdraw(WETH string, accId int, amt int64) {

	fmt.Println("Env: NetworkId=", Networkid, ",client=", Network.ProviderUrl[ProviderSortId])
	fmt.Println("----------------------------------------------")
	fmt.Println(".........unwrap .........  ")
	fmt.Println("----------------------------------------------")
	fmt.Println("WETH ADDRESS:", WETH)

	ChangeAccount(accId)

	instance := GetWethInstance(common.HexToAddress(WETH))

	//weth deposit
	ethAmount := X1E18(amt)

	tx, err := instance.Withdraw(Auth, ethAmount)

	if err != nil {
		log.Fatal("weth withdraw err ", err)
	}

	fmt.Println("weth withdraw 1 eth tx: ", tx.Hash().Hex())
	fmt.Println("unwrapped weth amount:", amt, " to: ", FromAddress)

	ChangeAccount(Account)

}

func GetWethInstance(WETH common.Address) *weth.Api {

	instance, err := weth.NewApi(WETH, Client)
	if err != nil {
		log.Fatal("get token Instance,", err)
	}

	return instance
}
