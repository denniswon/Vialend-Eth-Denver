package main

import (
	"fmt"
	"math/big"
	_ "time"

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
	TokenParam       [2]TokenStruct
	TransferToken    int
	TransferParam    [2]TransferStruct
	DeployFactory    int
	CreatePool       int
	InitialPool      int
	MintPool         int
	MintPoolParam    [3]int64
	DeployVault      int
	Deposit          int
	DepositParam     [3]int64
	Withdraw         int
	WithDrawParam    [2]int64
	Rebalance        int
	RebalanceParam   [3]int64
	CreatePosition   int
	IncreasePosition int
	RemovePosition   int
	Swap             int
	CollectFees      int
	Strategy1        int
	Strategy1Param   [3]int64
}

var sw = new(Switcher)

func main() {

	project.Init(-1, -1)

	//	project.Init(0, 0)
	//balance := project.EthBalance(Network.LendingContracts.CETH)
	//balance := project.EthBalanceArb("0x4Ddc2D193948926D02f9B1fE9e1daa0718270ED5")
	// fmt.Println(balance)

	project.Quiet = false

	//networkid, account, protocolfee, uniportion, team address to get fee cut
	//project.DeployVialendFeemaker(3, 1, big.NewInt(10), 90, "0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440")

	//DeployVaultBridge()
	//project.DeployArb()

	// project.SetVaultAddress("0xD0fF8fF803a30C5d7BBDdc797B544E07Ff3458cD", 0)
	// return

	//project.FindPool()
	// v := "0xb102Cd93329d7017Ae83C6E488f00EaB4844CbF2"
	// t := "0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF"
	// v := "0xb102Cd93329d7017Ae83C6E488f00EaB4844CbF2"
	// t := "0xfa5df5372c03d4968d128d624e3afeb61031a777"
	// a := big.NewInt(1e18)
	// project.Sweep(v, t, a)

	// project.EmergencyBurn()
	// return

	//project.DeployVialendFeemaker(3, 1, big.NewInt(10), 100, "0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440")
	//project.Deposit(1, [3]int64{1, 10, 1}, false)
	project.Strategy1(100, 1)
	//project.VaultInfo()
	//project.Withdraw(1, [2]int64{100, 1})
	// project.Withdraw(1, [2]int64{100, 1})
	//project.Withdraw(1, [2]int64{100, 3})

	//project.EmergencyBurn()
	// project.Withdraw(1, [2]int64{100, 3})
	//project.VaultInfo2("0x4aaE0bc3052aD3AB125Ae654f0f2C55Dbd9D6e17")
	project.VaultInfo()
	return
	// // // newVault()

	// project.GetCapital(1)
	// project.GetCapital(3)
	// project.LendingInfo()
	// // project.AccountInfo()
	// project.VaultInfo()
	// // project.PoolInfo()
	// //project.FindPool()
	// // project.GetPool("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", "0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60", 500)
	// return
	//project.Test_weth_deposit("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", 5, 15) // weth address, accountid, amount
	//project.Test_weth_withdraw("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", 3, 15)

	//check token decimals and info
	// fmt.Println(project.GetTokenInstance("0xC04B0d3107736C32e19F1c62b2aF67BE61d63a05"))
	//return

	//project.SetProtocolFee(big.NewInt(10))
	//project.SetUniswapPortionRatio(50)
	//	project.Withdraw(1, [2]int64{100, 4}) // team withdraw
	//project.Deposit(1, [3]int64{1, 10, 1}, false)
	//	project.Deposit(1, [3]int64{2, 1000, 0}, false)
	//project.EmergencyBurn()

	project.Strategy1(1000, 1)
	// project.Strategy1(100, 1)
	//project.AccountInfo()
	// project.Withdraw(1, [2]int64{100, 0})
	//project.Withdraw(1, [2]int64{100, 1})
	//project.Withdraw(1, [2]int64{100, 3})
	project.Alloc(1)
	//project.RemoveCTokens()
	project.VaultInfo()
	return

	// // // project.Strategy1( [3]int64{400, 60, 0})

	// project.AccountInfo()
	//project.Strategy1(500, 1)
	// project.Withdraw(1, [2]int64{100, 0})
	// project.Withdraw(1, [2]int64{100, 1})
	// project.Withdraw(1, [2]int64{100, 3})
	//project.VaultInfo()

	// redeemMyCtoken()
	// return

	//project.Withdraw(1, [2]int64{100, 0})
	//project.Withdraw(1, [2]int64{100, 1})
	//project.Withdraw(1, [2]int64{100, 3})
	//return

	//project.Deposit(1, [3]int64{1, 10, 0}, false)
	// project.Deposit(1, [3]int64{0, 500, 1}, false)
	// return

	// project.AccountInfo()
	// project.VaultInfo()
	// project.Withdraw(1, [2]int64{100, 0}) // team withdraw
	//project.Deposit(1, [3]int64{2, 0, 0}, true)
	//time.Sleep(10 * time.Second)
	// project.AccountInfo()

	project.Strategy1(500, 0)

	project.VaultInfo()

	project.GenFees(4, 2) // swap times, swap account, amount , sleepSeconds

	project.Strategy1(500, 0)

	project.VaultInfo()

	return

	// //project.genFees(4, 5)

	// project.Alloc(0)
	// project.AccountInfo()
	// project.VaultInfo()
	project.Strategy1(500, 1)
	project.Strategy1(200, 1)
	//project.Withdraw(1, [2]int64{100, 0})
	// project.Withdraw(1, [2]int64{100, 1})
	//project.Withdraw(1, [2]int64{100, 3})
	//project.Deposit(1, [3]int64{0, -1, 1}, false)
	project.AccountInfo()
	project.VaultInfo()
	return

}

func newVault() {

	networkId := 3
	acc := 0
	token0 := Network.TokenA
	token1 := Network.TokenB

	feetier := int64(10000) //10000, 3000, 500 //Network.FeeTier

	_protocolfee := big.NewInt(10)
	_quoteAmount := big.NewInt(1e18)
	_uniPortion := 20
	team := "0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440"
	strategy := Indi(0)

	project.VaultGen(networkId, acc, token0, token1, feetier, _protocolfee, _quoteAmount, _uniPortion, team, strategy)

}

func DeployVaultBridge() {

	// vault bridge on goreli 0x033F3C5eAd18496BA462783fe9396CFE751a2342
	// vault admin on goreli  0xb6F0049e37D32dED0ED2FAEeE7b69930FA49A879

	//	project.DeployVaultBridge()

	//project.DeployVaultAdmin()
	// project.AuthAdmin(Network.VaultAdmin, "0xfd8a5AE495Df1CA34F90572cb99A33B27173eDe1")
	// return

	project.SetVaultAddress("0xD0fF8fF803a30C5d7BBDdc797B544E07Ff3458cD", 0)
	// project.SetVaultAddress("0xf231F818a111FE5d2EFf006451689eCBbf5ef159", 1)

	// project.GetVaultAddress(0)
	// project.GetVaultAddress(1)

}

func test_vault() {

	sw.ViewOnly = false

	sw.DeployFactory = 0
	//...manual step... update the new factory addres in networks.go

	sw.DeployToken = 0
	sw.TokenParam[0] = TokenStruct{"e weth 2", "eWETH2", 18, big.NewInt(50000000000000)}
	sw.TokenParam[1] = TokenStruct{"e usdc 2", "eUSDC2", 6, big.NewInt(500000000000000000)}
	//...manual step... update the new token addres in networks.go

	sw.TransferToken = 0
	sw.TransferParam[0] = TransferStruct{0, project.X1E18(900), Network.TokenA, "0xeBb29c07455113c30810Addc123D0D7Cd8637aea"}
	sw.TransferParam[1] = TransferStruct{0, project.X1E18(900), Network.TokenB, "0xeBb29c07455113c30810Addc123D0D7Cd8637aea"}

	sw.CreatePool = 0 // *Note: if token0+token1+fee = pool exists ERROR: createPool VM Exception while processing transaction: revert
	sw.InitialPool = 0
	//...manual step... update the new pool addres in networks.go

	sw.MintPool = 0
	sw.MintPoolParam = [3]int64{1000, 1, 1} // currently hardcoded 1000 * 1e18 as the liquidity,
	// .... manual step.... new vault may apply

	sw.DeployVault = 0
	//...manual step... update the new vault addres in networks.go

	sw.Deposit = 1
	sw.DepositParam = [3]int64{1, 1, 1} // {amount0, amount1 , account}

	sw.Strategy1 = 1
	sw.Strategy1Param = [3]int64{600, 60, 3} // {tick range, tickspacing, account}

	sw.Withdraw = 1
	sw.WithDrawParam = [2]int64{100, 1} // { percent, account}
	// accountid,  amount of shares in percentage %

	sw.Rebalance = 0
	sw.RebalanceParam = [3]int64{10, 60, 3} //[2]int64{22000, 60} // 12000,60   {tick range , tickspacing, account}

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

	if sw.DeployFactory > 0 {
		project.DeployFactory() ///new factory  for crating new pools if old pool already exists
	}

	if sw.DeployToken > 0 {

		token0 := project.DeployToken(sw.TokenParam[0].Name, sw.TokenParam[0].Symbol, sw.TokenParam[0].Decimals, sw.TokenParam[0].MaxTotalSupply)
		token1 := project.DeployToken(sw.TokenParam[1].Name, sw.TokenParam[1].Symbol, sw.TokenParam[1].Decimals, sw.TokenParam[1].MaxTotalSupply)

		//always make token0 = weth = tokenA
		Network.TokenA = token0
		Network.TokenB = token1

	}

	if sw.TransferToken > 0 {

		TokenTransfer(
			sw.TransferParam[0].AccountId,
			sw.TransferParam[0].Amount,
			sw.TransferParam[0].TokenAddress,
			sw.TransferParam[0].ToAddress)

		TokenTransfer(
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

	project.Deposit(sw.Deposit, sw.DepositParam, false) /// deposit token0 amount * 1e18, token1 amount * 1e6

	project.Rebalance(sw.Rebalance, sw.RebalanceParam) /// make sure Account = 0

	project.Withdraw(sw.Withdraw, sw.WithDrawParam) /// withdraw shares, input number in percentage %

	// print all deployed addresses
	for _, i := range InfoString {
		fmt.Println(i)
	}

	project.CheckPrice(false, 3)
	project.Equation(false, false)

}
