package main

import (
	"fmt"
	_ "time"

	"../config"
	project "../project"
)

func main() {

	fmt.Println("Env: NetworkId=", config.Networkid, ",client=", config.Network.ProviderUrl[config.ProviderSortId])
	project.Init()

	project.Quiet = true

	// project.MovePrice(project.Sell)
	// project.MovePrice(project.Buy)

	project.GenFees(14, 5) // swap times

	//project.Rebal(20) //interval seconds

	//AutoBoth() // Rebal(30) + GenFees

}

func AutoBoth() {

	go project.GenFees(10, 10) // {swap num of times, account, amount , sleepSeconds}

	project.Rebal(10, 20) // max rebalance run ,  interval seconds between each rebalance

}
