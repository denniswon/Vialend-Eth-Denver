package main

import (
	"fmt"

	"./config"
	project "./include"
)

func main() {

	fmt.Println("Env: NetworkId=", config.Networkid, ",client=", config.Network.ClientUrl)

	///new factory  for crating new pools if old pool already exists
	project.DeployFactory(false)

	/// if not creating new pool then use the network1.factory
	fee := int64(3000)
	project.CreatePool(fee, false)

	project.InitialPool(false)

	project.DeployVault(false)

	project.Approve(false)
	project.Deposit(false)

	project.GetTotalAmounts(true)

	project.Rebalance(true)
	/*
		project.Withdraw(false)
		project.Rebalance(false)
		project.Swap(false)
	*/
	project.CheckPrice(false)

}
