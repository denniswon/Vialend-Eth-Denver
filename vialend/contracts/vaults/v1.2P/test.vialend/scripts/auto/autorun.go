package main

import (
	"fmt"

	"../config"
	project "../project"
)

func main() {

	fmt.Println("Env: NetworkId=", config.Networkid, ",client=", config.Network.ProviderUrl[config.ProviderSortId])
	project.Init()

	project.Quiet = false

	justSwap()

	//project.GenFees(40, 5, 1, 5) // swap times, account, , amount , sleepSeconds

	//project.Rebal(20) //interval seconds

	//AutoBoth() // Rebal(30) + GenFees

}

func AutoBoth() {

	go project.GenFees(100, 5, 1, 30) // {swap num of times, account, amount , sleepSeconds}

	project.Rebal(20) //interval seconds between each rebalance

}

func justSwap() {

	accountId := 5

	zeroForOne := true

	amountX1e18 := config.X1E18(1)

	project.Swap2(accountId, amountX1e18, zeroForOne)
	project.Swap2(accountId, amountX1e18, zeroForOne)
	project.Swap2(accountId, amountX1e18, zeroForOne)
	project.Swap2(accountId, amountX1e18, zeroForOne)
	project.Swap2(accountId, amountX1e18, zeroForOne)
	project.Swap2(accountId, amountX1e18, zeroForOne)

}
