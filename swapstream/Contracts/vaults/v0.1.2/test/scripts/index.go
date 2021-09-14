package main

import (
	"fmt"

	"./config"
	project "./include"
)

type Switcher struct {
	ViewOnly         bool
	DeployFactory    int
	CreatePool       int
	InitialPool      int
	DeployVault      int
	Approve          int
	Deposit          int
	DepositAmount    [2]int64
	Withdraw         int
	WithDrawParam    [2]int64
	Rebalance        int
	RebalanceParam   [2]int64
	CreatePosition   int
	IncreasePosition int
	RemovePosition   int
	Swap             int
}

func main() {

	fmt.Println("Env: NetworkId=", config.Networkid, ",client=", config.Network.ProviderUrl[config.ProviderSortId])

	//config.TokenTransfer(0, config.X1E18(10000), config.Network.TokenA, "0xeBb29c07455113c30810Addc123D0D7Cd8637aea")
	//config.TokenTransfer(0, config.X1E18(10000), config.Network.TokenB, "0xeBb29c07455113c30810Addc123D0D7Cd8637aea")
	//return

	var sw = new(Switcher)

	sw.DeployFactory = 0
	sw.CreatePool = 0 // *Note: if token0+token1+fee = pool exists ERROR: createPool VM Exception while processing transaction: revert
	sw.InitialPool = 0
	sw.DeployVault = 0
	sw.Approve = 0

	sw.Deposit = 0
	sw.DepositAmount = [2]int64{2000000, 2000000} // amount0, amount1 to deposit

	sw.Withdraw = 1
	sw.WithDrawParam = [2]int64{0, 100} // accountid,  amount of shares in percentage %

	sw.Rebalance = 0
	sw.RebalanceParam = [2]int64{2000, 60} //[2]int64{22000, 60} // 12000,60   {full range , tickspacing}

	sw.Swap = -1
	sw.CreatePosition = -1
	sw.IncreasePosition = -1
	sw.RemovePosition = -1

	sw.ViewOnly = false

	project.Init()

	if sw.ViewOnly {

		project.AccountInfo()
		project.VaultInfo(1)
		project.PoolInfo()
		project.GetPoolFromToken() /// get the pool address from its  tokens and fee

		return

	}

	project.DeployFactory(sw.DeployFactory) ///new factory  for crating new pools if old pool already exists
	//return

	project.CreatePool(sw.CreatePool) /// need to edit networks to setup address of token0 and token1,  fee tier

	project.InitialPool(sw.InitialPool)
	//return

	project.DeployVault(sw.DeployVault) /// deployed by test admin 2, edit networks. token0, token1, fee to get the pool address
	//return

	project.Approve(sw.Approve)
	//return

	//project.VaultInfo(1)

	project.Deposit(sw.Deposit, sw.DepositAmount[0], sw.DepositAmount[1]) /// deposit token0 amount * 1e18, token1 amount * 1e6

	project.VaultInfo(sw.DepositAmount[0])

	project.Withdraw(sw.Withdraw, sw.WithDrawParam) /// withdraw shares, input number in percentage %

	project.VaultInfo(sw.WithDrawParam[0])

	project.Rebalance(sw.Rebalance, sw.RebalanceParam) /// make sure Account = 0

	project.VaultInfo(sw.RebalanceParam[0])

	//project.PoolInfo()

	//project.Swap(sw.Swap)
	//project.CreatePosition(sw.CreatePosition)
	//project.IncreasePosition(sw.IncreasePosition)
	//project.RemovePosition(sw.RemovePosition)

	project.CheckPrice(false, 3)

	project.Equation(false, false)

}
