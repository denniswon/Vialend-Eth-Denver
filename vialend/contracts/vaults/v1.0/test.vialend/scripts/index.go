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
	DeployToken      int
	TokenParam       [2]config.TokenStruct
	TransferToken    int
	TransferParam    [2]TransferStruct
	DeployFactory    int
	CreatePool       int
	InitialPool      int
	MintPool         int
	MintPoolParam    [3]int64
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
	Strategy1        int
	Strategy1Param   [3]int64
}

func main() {

	// get some weth
	//project.Test_weth_deposit(30)
	//project.Test_weth_withdraw(3)
	//return

	//project.DeployVialendFeemaker(3)
	//return

	//project.SetUniswapPortionRatio(1)

	// for i := 0; i < 10; i++ {
	// 	test_vault()
	// }
	// // //return

	//project.EmergencyBurn()
	//project.Whitehacker()
	//return

	test_vault()

	project.LendingInfo()
	project.AccountInfo()
	project.VaultInfo(1)
	project.PoolInfo()
	//project.GetPoolFromToken()
}

func test_vault() {
	fmt.Println("Env: NetworkId=", config.Networkid, ",client=", config.Network.ProviderUrl[config.ProviderSortId])

	var sw = new(Switcher)

	sw.ViewOnly = false

	sw.DeployFactory = 0
	//...manual step... update the new factory addres in networks.go

	sw.DeployToken = 0
	sw.TokenParam[0] = config.TokenStruct{"e weth 2", "eWETH2", 18, big.NewInt(50000000000000)}
	sw.TokenParam[1] = config.TokenStruct{"e usdc 2", "eUSDC2", 6, big.NewInt(500000000000000000)}
	//...manual step... update the new token addres in networks.go

	sw.TransferToken = 0
	sw.TransferParam[0] = TransferStruct{0, config.X1E18(900), config.Network.TokenA, "0xeBb29c07455113c30810Addc123D0D7Cd8637aea"}
	sw.TransferParam[1] = TransferStruct{0, config.X1E18(900), config.Network.TokenB, "0xeBb29c07455113c30810Addc123D0D7Cd8637aea"}

	sw.CreatePool = 0 // *Note: if token0+token1+fee = pool exists ERROR: createPool VM Exception while processing transaction: revert
	sw.InitialPool = 0
	//...manual step... update the new pool addres in networks.go

	sw.MintPool = 0
	sw.MintPoolParam = [3]int64{1000, 1, 1} // currently hardcoded 1000 * 1e18 as the liquidity,
	// .... manual step.... new vault may apply

	sw.DeployVault = 0
	//...manual step... update the new vault addres in networks.go

	sw.Approve = 0

	sw.Deposit = 0
	sw.DepositAmount = [2]int64{1, 100} // {amount0, amount1 }

	sw.Strategy1 = 1
	sw.Strategy1Param = [3]int64{600, 60, 3} // {range, tickspacing, account}

	sw.Withdraw = 1
	sw.WithDrawParam = [2]int64{3, 100} // {account, percent}
	// accountid,  amount of shares in percentage %

	sw.Rebalance = 0
	sw.RebalanceParam = [2]int64{10, 60} //[2]int64{22000, 60} // 12000,60   {full range , tickspacing}

	sw.Swap = 0

	// 1: single swap, 2: multiple swaps
	// swapAmount, _ := new(big.Int).SetString("85175185371092425157", 10) // 85 * 1e18
	// zeroForOne := false
	//swapAmount, _ := new(big.Int).SetString("139190665697301284354", 10) // 139 * 1e18
	//zeroForOne := true

	sw.CollectFees = 0

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

	if sw.DeployFactory > 0 {
		project.DeployFactory() ///new factory  for crating new pools if old pool already exists
	}

	if sw.DeployToken > 0 {

		token0 := project.DeployToken(sw.TokenParam[0].Name, sw.TokenParam[0].Symbol, sw.TokenParam[0].Decimals, sw.TokenParam[0].MaxTotalSupply)
		token1 := project.DeployToken(sw.TokenParam[1].Name, sw.TokenParam[1].Symbol, sw.TokenParam[1].Decimals, sw.TokenParam[1].MaxTotalSupply)

		//always make token0 = weth = tokenA
		config.Network.TokenA = token0
		config.Network.TokenB = token1

	}

	if sw.TransferToken > 0 {

		config.TokenTransfer(
			sw.TransferParam[0].AccountId,
			sw.TransferParam[0].Amount,
			sw.TransferParam[0].TokenAddress,
			sw.TransferParam[0].ToAddress)

		config.TokenTransfer(
			sw.TransferParam[1].AccountId,
			sw.TransferParam[1].Amount,
			sw.TransferParam[1].TokenAddress,
			sw.TransferParam[1].ToAddress)

	}

	if sw.CreatePool > 0 {
		project.CreatePool(sw.CreatePool) /// need to edit networks to setup address of token0 and token1,  fee tier
	}
	if sw.InitialPool > 0 {
		project.InitialPool(sw.InitialPool)
	}

	if sw.MintPool > 0 {
		project.MintPool(sw.MintPoolParam[0], sw.MintPoolParam[1], sw.MintPoolParam[2])
		//os.Exit(0)
	}

	if sw.DeployVault > 0 {
		project.DeployVault() /// deployed by test admin 2, edit networks. token0, token1, fee to get the pool address
	}

	project.Approve(sw.Approve)
	//return

	project.Deposit(sw.Deposit, sw.DepositAmount[0], sw.DepositAmount[1]) /// deposit token0 amount * 1e18, token1 amount * 1e6

	if sw.Swap == 1 {

		project.Swap(sw.Swap, sw.RebalanceParam[0])

	} else if sw.Swap == 2 {

		for i := 0; i < 5; i++ {
			fmt.Println("swap y for x:", i)

			project.Deposit(1, 2, 1000)

			project.Swap(sw.Swap, sw.RebalanceParam[0])

		}

		for i := 0; i < 5; i++ {
			fmt.Println("swap x for y:", i)

			project.Deposit(1, 1, 5000)

			project.Swap(sw.Swap, sw.RebalanceParam[0])

		}

	}

	project.Strategy1(sw.Strategy1, sw.Strategy1Param)

	project.Rebalance(sw.Rebalance, sw.RebalanceParam) /// make sure Account = 0

	project.Withdraw(sw.Withdraw, sw.WithDrawParam) /// withdraw shares, input number in percentage %

	// print all deployed addresses
	for _, i := range config.InfoString {
		fmt.Println(i)
	}

	//project.CreatePosition(sw.CreatePosition)
	//project.IncreasePosition(sw.IncreasePosition)
	//project.RemovePosition(sw.RemovePosition)

	project.CheckPrice(false, 3)

	project.Equation(false, false)

}
