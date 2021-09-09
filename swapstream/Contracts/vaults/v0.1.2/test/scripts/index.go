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

	/// edit networks. token0, token1, fee
	project.CreatePool(false)

	project.InitialPool(false)

	//project.MintPool(fee, false)

	/// edit networks. token0, token1, fee to get the pool address
	//project.DeployVault(true)
	//return

	/// print the pool address and token address
	//project.GetPoolFromToken(false)
	//return

	project.Approve(false)
	//return
	/// deposit token0 amount * 1e18, token1 amount * 1e6
	project.AccountInfo(true)
	project.ValutInfo(true)

	project.Deposit(true, 1, 5000)

	//project.AccountInfo(true)
	//project.ValutInfo(true)

	/// withdraw shares
	//project.Withdraw(true, 1)

	project.Rebalance(true)

	project.AccountInfo(true)
	project.ValutInfo(true)

	/*
		project.Swap(false)
	*/
	project.CheckPrice(false, 3)

	project.Equation(false, false)

}
