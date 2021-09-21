package main

import (
	"fmt"
	"math/big"

	"./config"
	project "./project"
)

type TransferStruct struct {
	AccountId    int
	Amount       *big.Int
	TokenAddress string
	ToAddress    string
}

type Switcher struct {
	ViewOnly         bool
	ViewEvent        bool
	DeployToken      int
	TokenParam       config.TokenStruct
	TransferToken    int
	TransferParam    TransferStruct
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
	CollectFees      int
}

func main() {

	fmt.Println("Env: NetworkId=", config.Networkid, ",client=", config.Network.ProviderUrl[config.ProviderSortId])

	var sw = new(Switcher)

	sw.ViewEvent = false
	sw.ViewOnly = false

	sw.DeployFactory = 0
	//...hidden step... update the new factory addres in networks.go

	sw.DeployToken = 0
	sw.TokenParam = config.TokenStruct{"aaa usdc", "usdc", 18, big.NewInt(50000000000)}
	//...hidden step... update the new token addres in networks.go

	sw.TransferToken = 0
	sw.TransferParam = TransferStruct{0, config.X1E18(90000), config.Network.TokenB, "0xeBb29c07455113c30810Addc123D0D7Cd8637aea"}

	sw.CreatePool = 0 // *Note: if token0+token1+fee = pool exists ERROR: createPool VM Exception while processing transaction: revert
	sw.InitialPool = 0
	//...hidden step... update the new pool addres in networks.go

	sw.DeployVault = 0
	//...hidden step... update the new vault addres in networks.go

	sw.CollectFees = 0

	sw.Approve = 1

	sw.Deposit = 1
	sw.DepositAmount = [2]int64{100, 100} // amount0, amount1 to deposit

	sw.Swap = 1
	// 1: single swap, 2: multiple swaps

	sw.Withdraw = 0
	sw.WithDrawParam = [2]int64{0, 100}
	// accountid,  amount of shares in percentage %

	sw.Rebalance = 0
	sw.RebalanceParam = [2]int64{1000, 60} //[2]int64{22000, 60} // 12000,60   {full range , tickspacing}

	sw.CreatePosition = -1
	sw.IncreasePosition = -1
	sw.RemovePosition = -1

	project.Init()

	if sw.CollectFees > 0 {

		project.AccountInfo()
		project.VaultInfo(1)
		project.CollectProtocolFees(big.NewInt(100), big.NewInt(100))
		project.AccountInfo()
		project.VaultInfo(1)

		return
	}

	if sw.ViewOnly {

		project.AccountInfo()
		project.VaultInfo(1)
		project.PoolInfo()
		project.GetPoolFromToken() /// get the pool address from its  tokens and fee

		return

	}

	if sw.ViewEvent {
		eventname := "RebalanceLog"
		block0 := 1250
		block1 := 1253
		project.VaultEvent(eventname, int64(block0), int64(block1))
		return
	}

	if sw.DeployFactory > 0 {
		project.DeployFactory() ///new factory  for crating new pools if old pool already exists
		return
	}

	project.CreatePool(sw.CreatePool) /// need to edit networks to setup address of token0 and token1,  fee tier

	project.InitialPool(sw.InitialPool)
	//return

	if sw.DeployToken > 0 {
		project.DeployToken(sw.TokenParam.Name, sw.TokenParam.Symbol, sw.TokenParam.Decimals, sw.TokenParam.MaxTotalSupply)

		return
	}

	if sw.TransferToken > 0 {

		config.TokenTransfer(
			sw.TransferParam.AccountId,
			sw.TransferParam.Amount,
			sw.TransferParam.TokenAddress,
			sw.TransferParam.ToAddress)
		return
	}

	if sw.DeployVault > 0 {
		project.DeployVault() /// deployed by test admin 2, edit networks. token0, token1, fee to get the pool address
		return
	}

	project.Approve(sw.Approve)
	//return

	project.Deposit(sw.Deposit, sw.DepositAmount[0], sw.DepositAmount[1]) /// deposit token0 amount * 1e18, token1 amount * 1e6
	project.VaultInfo(1)

	if sw.Swap == 1 {

		project.VaultInfo(1)

		project.Swap(sw.Swap)

		project.VaultInfo(1)

	} else if sw.Swap == 2 {

		for i := 0; i < 5; i++ {
			fmt.Println("swap y for x:", i)

			project.Deposit(1, 2, 10)

			project.Swap(sw.Swap)

		}

		project.VaultInfo(1)

		for i := 0; i < 5; i++ {
			fmt.Println("swap x for y:", i)

			project.Deposit(1, 10, 2)

			project.Swap(sw.Swap)

		}

		project.VaultInfo(1)

	}

	project.Rebalance(sw.Rebalance, sw.RebalanceParam) /// make sure Account = 0
	project.VaultInfo(sw.Rebalance)

	project.Withdraw(sw.Withdraw, sw.WithDrawParam) /// withdraw shares, input number in percentage %
	project.VaultInfo(sw.Withdraw)

	project.PoolInfo()

	//project.CreatePosition(sw.CreatePosition)
	//project.IncreasePosition(sw.IncreasePosition)
	//project.RemovePosition(sw.RemovePosition)

	project.CheckPrice(false, 3)

	project.Equation(false, false)

}