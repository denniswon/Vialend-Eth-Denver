package main

import (
	"fmt"

	"./config"
	project "./include"
)

type Switcher struct {
	DeployFactory    int
	CreatePool       int
	InitialPool      int
	Deposit          int
	DeployVault      int
	Approve          int
	Withdraw         int
	Rebalance        int
	CreatePosition   int
	IncreasePosition int
	RemovePosition   int
	Swap             int
}

func main() {

	fmt.Println("Env: NetworkId=", config.Networkid, ",client=", config.Network.ProviderUrl[config.ProviderSortId])

	var sw = new(Switcher)
	sw.DeployFactory = 0
	sw.CreatePool = 0
	sw.InitialPool = 0
	sw.DeployVault = 0
	sw.Approve = 0
	sw.Deposit = 1
	sw.Withdraw = 0
	sw.Rebalance = 1

	sw.Swap = -1
	sw.CreatePosition = -1
	sw.IncreasePosition = -1
	sw.RemovePosition = -1

	project.DeployFactory(sw.DeployFactory) ///new factory  for crating new pools if old pool already exists
	//return

	project.CreatePool(sw.CreatePool) /// need to edit networks to setup address of token0 and token1,  fee tier

	project.InitialPool(sw.InitialPool)
	//return

	project.DeployVault(sw.DeployVault) /// deployed by test admin 2, edit networks. token0, token1, fee to get the pool address
	//return

	//project.GetPoolFromToken() /// get the pool address from its  tokens and fee

	project.Approve(sw.Approve)
	//return

	//project.AccountInfo()
	//project.ValutInfo()

	project.Deposit(sw.Deposit, 2000, 100) /// deposit token0 amount * 1e18, token1 amount * 1e6

	//project.AccountInfo()
	//project.ValutInfo()

	project.Withdraw(sw.Withdraw, 10) /// withdraw shares, input number in percentage %

	project.Rebalance(sw.Rebalance) /// make sure Account = 0
	project.PoolInfo()

	project.AccountInfo()
	project.ValutInfo()

	//project.Swap(sw.Swap)
	//project.CreatePosition(sw.CreatePosition)
	//project.IncreasePosition(sw.IncreasePosition)
	//project.RemovePosition(sw.RemovePosition)

	project.CheckPrice(false, 3)

	project.Equation(false, false)

}
