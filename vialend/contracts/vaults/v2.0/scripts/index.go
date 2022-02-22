package main

import (
	"fmt"
	"math/big"
	_ "time"

	project "viaroot/scripts/project"
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
	TokenParam       [2]project.TokenStruct
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

	// project.AuthAdmin(project.Network.VaultAdmin, "0xb6F0049e37D32dED0ED2FAEeE7b69930FA49A879")
	// return

	// project.SendTestTokens()
	// return

	project.ConfigParser()
	defer project.ConfigWrite()

	project.Init(-1, -1)

	//project.DeployVaultBridge()  // deploy and manually update config.json
	// setVaultBridge()
	//	return

	//project.DeployVaultFactory()
	//return

	//#reload strategy

	//// newStratAddr := project.DeployStratByGo()
	// newStratAddr := project.DeployStratByGoStruct()
	// project.ConfigWrite()
	// project.Register(newStratAddr, project.Network.Vault)
	// project.ChangeStat(newStratAddr, project.Network.Vault, 1)
	// return

	//#reload vault
	// newVaultAddr := project.DeployVaultByGo()
	// project.ConfigWrite()
	// project.Register(project.Network.VaultStrat, newVaultAddr)
	// project.ChangeStat(project.Network.VaultStrat, newVaultAddr, 1)

	//#reload strategy & vault
	// newStratAddr := project.DeployStratByGoStruct()
	// newVaultAddr := project.DeployVaultByGo()
	// project.ConfigWrite()
	// project.Register(newStratAddr, newVaultAddr)
	// project.ChangeStat(newStratAddr, newVaultAddr, 1)
	// return

	//# Just Register/Register strategy & vault
	// project.Register(project.Network.VaultStrat, project.Network.Vault)
	// project.ChangeStat(project.Network.VaultStrat, project.Network.Vault, 1)
	//return

	project.Init(-1, -1)

	// project.IsContract("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	// return

	//### Events

	// blockfrom := 6178203
	// blockend := 6178262
	// project.Events("PendingWithdraw", blockfrom, blockend)
	// project.Events("Withdraw", blockfrom, blockend)
	// project.Events("Deposit", blockfrom, blockend)
	// project.Events("MintFees", blockfrom, blockend)

	// return

	// project.DeployVaultByDeployer()
	// project.DeployStratByDeployer()
	// return

	//# WEB deploy
	// project.DeployVaultDeployer()
	// project.DeployStratDeployer()
	// // return
	// project.FactoryVault()

	// project.ConfigWrite()
	// project.Init(-1, -1)

	// // return
	s := project.Network.VaultStrat
	v := project.Network.Vault
	// project.ChangeStat(s, v, 1)
	project.GetStat(s, v)

	// project.CheckActive("0x2a8179A7893d00B33D2d9DBe9F0e4bBf2Cb97DE7")
	// project.Rebalance(400, 0)
	return

	// project.FactoryPublicList()
	// project.ViewVaults()
	// // project.CheckStatus(project.Network.Vault, 1) // CheckVaultStatus()
	// project.ViaVaultPublicList()
	// //project.ViaStratUniCompPublicList()

	// // project.GetTwap()
	// // project.GetPriceStratCall()

	// return

	// project.SetTwapduration(5)
	// project.GetPriceStratCall()
	// project.GetTwap()
	// project.GetTickPrice()
	// return

	// project.MyAccountInfo(0)

	// project.MyAccountInfo(1)

	// //### ViaVaultInfo

	//project.GetTVL()
	// project.GetTotalAmounts()
	// project.GetCompAmounts()
	//project.LendingInfo()
	//return

	// //project.DeployVaultFactory()
	// //project.DeployVaultStrategy()
	//project.ChangeStat("0xd029FDcEB5B0E971e675D7f2766188e5F8ccEeE9", "0xD06d9CF030401a72476B9e10AA47a94CE5f3798E", 2)
	// //project.CheckVaultStatus()
	//return

	// project.Deposit(1e16, 1e6, 0)
	// //project.Withdraw(100, 0)
	// project.Rebalance(500, 0)
	// // return
	// //project.EmergencyCall() //by admin
	// project.CallFunds() // by admin, check to make sure all funds are back to vaults.
	// // // //project.ChangeStat(4)		//by admin to change the status to abandoned and only with draw allowed
	// // project.EmergencyWithdraw(0)
	// //project.EmergencyWithdraw(1)
	// project.Withdraw(100, 0)
	// // project.Withdraw(3)
	// return

	// s := "0x9a94272446f0c119E1006935c9E6D6fEB6c206f4"
	// v := "0x6E09167c444AAbe5cD49Cff5Af16B15E33096e6C"
	//project.FactoryVault()

	//project.Init(-1, -1)
	//project.DeployStratByDeployer()
	//	project.DeployVaultStrategy()
	//project.DeployViaVault()
	// //	project.DeployStratUniComp()
	//project.Sleep(5000)

	//project.SetPortionRatio(90, 100)
	//project.SetTwapduration(10)
	//return

	project.Init(-1, -1)
	project.Quiet = false

	//#### WETH/USDC Test
	//project.Deposit(1e16, 1e6, 0)
	// project.Deposit(2e16, 2e6, 1)

	project.Withdraw(100, 0)
	// // project.Deposit(2e17, 2e18, 1)
	// //project.Withdraw(100, 0)
	project.GetTVL()
	// // project.GetTotalAmounts()
	// // project.GetCompAmounts()

	// //project.MoveFunds()
	// project.Rebalance(400, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

	// // project.GetTVL()
	// // project.GetTotalAmounts()
	// // project.GetCompAmounts()
	// //project.LendingInfo()

	// // //project.Rebalance(400, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
	// // //project.Rebalance(600, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
	// // project.Deposit(1e17, 1e6, 0)
	// // // //project.Alloc(0)
	// // //project.Rebalance(400, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
	// // project.Rebalance(600, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

	// // project.Deposit(2e17, 2e6, 1)
	// // project.Rebalance(600, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

	// // // project.EmergencyBurn() // vault calls strategy.callFunds...alloc/removepositions/transferfunds
	// // // project.EmergencyWithdraw(0)
	// // // project.EmergencyWithdraw(1)

	// project.Withdraw(100, 0)
	// project.Withdraw(100, 1)
	// project.Withdraw(100, 2)
	// project.Withdraw(100, 3)
	// project.Rebalance(400, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

	// // // //project.VaultInfo()
	// // // project.MyAccountInfo(0)
	// // // project.MyAccountInfo(1)
	// // project.GetTVL()
	// // project.GetTotalSupply()
	// // project.GetTotalAmounts()
	// // project.GetCompAmounts()

	return

	//#### WETH/DAI Test
	//project.SetTwapduration(10)
	project.Deposit(1e17, 1e18, 0)
	project.Deposit(2e17, 2e18, 1)
	// project.Withdraw(100, 0)
	// project.Withdraw(100, 1)
	// project.Deposit(2e17, 2e18, 1)
	// //project.Withdraw(100, 0)
	//project.GetTVL()
	// project.GetTotalAmounts()
	// project.GetCompAmounts()

	//project.MoveFunds()
	project.Rebalance(400, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
	project.Withdraw(50, 0)

	project.GetTVL()
	// project.GetTotalAmounts()
	// project.GetCompAmounts()

	//project.Rebalance(400, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
	// //project.Rebalance(600, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
	project.Deposit(1e17, 1e18, 0)
	// // //project.Alloc(0)
	// //project.Rebalance(400, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
	// project.Rebalance(600, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

	// project.Deposit(2e17, 2e18, 1)
	// project.Rebalance(600, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

	// // project.EmergencyWithdraw(0)
	// // project.EmergencyWithdraw(1)

	project.Withdraw(100, 0)
	project.Withdraw(100, 1)
	project.Withdraw(100, 2)
	project.Withdraw(100, 3)
	project.Rebalance(400, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

	// // //project.VaultInfo()
	// // project.MyAccountInfo(0)
	// // project.MyAccountInfo(1)
	project.GetTVL()
	// project.GetTotalSupply()
	// project.GetTotalAmounts()
	// project.GetCompAmounts()

	return

	// fmt.Println(project.Network.ViaFactory)
	// fmt.Println(project.Network.VaultStrat)
	// fmt.Println(project.Network.Vault)

	// project.Network.ViaFactory = os.Getenv("ViaFactory")
	// project.Network.VaultStrat = os.Getenv("VaultStrat")
	// project.Network.Vault = os.Getenv("Vault")
	// project.Quiet = false

	//project.GetViaVaultAddress()

	//project.DeployFactory()
	//BuildAll()

	//balance := project.EthBalance(Network.LendingContracts.CETH)
	//balance := project.EthBalanceArb("0x4Ddc2D193948926D02f9B1fE9e1daa0718270ED5")
	//project.BlockNumber()

	//project.TransferEth("1b280901929b5cd52f362b544072b66bfe29a9396db485a23da7de9f485512b0", project.X1E18(1), "0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E")
	// project.VaultInfo()
	// return
	// mainnet fork: weth/usdc

	//project.FindPool()
	//project.GetPool("0xc18360217d8f7ab5e7c516566761ea12ce7f9d72", "0xdac17f958d2ee523a2206206994597c13d831ec7", 3000)
	//return
	// //project.EthBalance("0xB7A41b27af9Ed23F65E36f9d92d287327c4D997d")
	// //return
	// project.DeployCallee()
	// project.DeployVialendFeemaker(-1, 0, big.NewInt(10), 50, big.NewInt(1e6), "0x5ACb5DB941E3Fc33E0c0BC80B90114b6CD0249B5")
	// project.EthBalance("0xa0df350d2637096571F7A701CBc1C5fdE30dF76A")
	// project.Wrap(project.Network.TokenB, 0, 30)
	// project.EthBalance("0xa0df350d2637096571F7A701CBc1C5fdE30dF76A")
	//return
	// project.EthBalance("0xEC2DD0d0b15D494a58653427246DC076281C377a")
	// project.Swap(0, 10000, 2, 1)
	// project.ERC20Balance(project.Network.TokenA, "0xa0df350d2637096571F7A701CBc1C5fdE30dF76A")
	// project.ERC20Balance(project.Network.TokenB, "0xa0df350d2637096571F7A701CBc1C5fdE30dF76A")
	// return

	// project.Deposit(1e17, 1e18, 0)
	// project.Deposit(2e17, 2e18, 1)
	// // project.VaultInfo()
	// project.MoveFunds()
	// project.Rebalance(400, 0)
	// //project.Alloc(0)
	// //project.EmergencyBurn()
	// project.Withdraw(100, 1)
	// project.Withdraw(100, 0)
	// //project.VaultInfo()
	// return

	// project.Deposit(1, [3]int64{4000, 10, 0}, false)
	// project.Deposit(1, [3]int64{4000, 10, 1}, false)
	// project.VaultInfo()
	//project.Strategy1(800, 0)
	//project.VaultInfo()

	// project.Withdraw(1, [2]int64{100, 0})
	// project.Withdraw(1, [2]int64{100, 1})

	// project.MyAccountInfo(1)
	// project.VaultInfo()

	// project.CheckFees()

	//project.SetProtocolFee(big.NewInt(10))
	//project.SetUniswapPortionRatio(50)

	// project.EmergencyBurn()

	// return

	// goerli test data
	// project.DeployVialendFeemaker(-1, 0, big.NewInt(10), 100, big.NewInt(1e18), "0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440")
	// project.Sleep(1000)

	// project.Deposit(1, [3]int64{3, 3, 0}, false)
	// project.Sleep(1000)

	// project.Deposit(1, [3]int64{2, 2, 1}, false)
	// project.Sleep(1000)

	// // // //	 project.Deposit(1, [3]int64{4000, 10, 1}, false)
	// // // project.VaultInfo()
	//project.Strategy1(800, 0)
	//project.Alloc(0)
	// project.Sleep(5000)
	// project.VaultInfo()
	// project.CheckFees()
	// project.GetLendingAmounts(project.Network.Vault)
	//return
	//project.SetProtocolFee(big.NewInt(0))
	//project.WithdrawPending(100, 0)
	//project.Sleep(1000)
	// project.WithdrawPending(0, 1)
	// project.Sleep(1000)

	//project.Strategy1(800, 0)
	// project.Sleep(1000)
	// project.Withdraw(1, [2]int64{100, 1})

	//project.VaultInfo()

	// project.MyAccountInfo(0)
	// project.Sleep(1000)

	// project.MyAccountInfo(1)
	// project.Sleep(1000)

	//project.Withdraw(1, [2]int64{100, 0})
	// project.Withdraw(1, [2]int64{100, 1})

	// project.MyAccountInfo(0)
	// project.Sleep(1000)

	//project.Withdraw(1, [2]int64{100, 1})

	//	project.MyAccountInfo(1)
	//project.VaultInfo()
	// project.CheckFees()

	return

	//networkid, account, protocolfee, uniportion, team address to get fee cut
	//project.DeployVialendFeemaker(3, 1, big.NewInt(10), 90, "0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440")

	//project.DeployArb()
	// project.DeployCallee()
	// return

	//project.SetVaultAddress("0xBC6b6e273171C428d85cDdB23D344a8400B48441", 2)
	// return

	// project.GetPool("0xC04B0d3107736C32e19F1c62b2aF67BE61d63a05", "0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60", 10000)

	// return

	//project.FindPool()
	// v := "0xb102Cd93329d7017Ae83C6E488f00EaB4844CbF2"
	// t := "0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF"
	// v := "0xb102Cd93329d7017Ae83C6E488f00EaB4844CbF2"
	// t := "0xfa5df5372c03d4968d128d624e3afeb61031a777"
	// a := big.NewInt(1e18)
	// project.Sweep(v, t, a)

	// project.EmergencyBurn()
	//project.VaultInfo2("0xb102Cd93329d7017Ae83C6E488f00EaB4844CbF2")
	//return

	//nid := int(0)
	//acc := int(0)

	//project.PrintPrice()
	//project.Wrap(project.Network.TokenA, 0, 10)

	//project.DeployVialendFeemaker(-1, 1, big.NewInt(10), 100, big.NewInt(1e8), "0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440")
	//project.DeployVialendFeemaker(-1, 1, big.NewInt(10), 100, big.NewInt(1e18), "0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440")
	//project.Deposit(1, [3]int64{1, 10, 0}, false)
	//project.Strategy1(800, 1)
	//project.VaultInfo()
	//project.Withdraw(1, [2]int64{100, 0})
	//project.Withdraw(1, [2]int64{100, 1})
	//project.Withdraw(1, [2]int64{100, 3})

	//project.EmergencyBurn()
	project.SetProtocolFee(big.NewInt(0))
	//project.Withdraw(1, [2]int64{100, 3})
	//project.VaultInfo2("0x4aaE0bc3052aD3AB125Ae654f0f2C55Dbd9D6e17")
	//project.MyAccountInfo(0)
	//project.PoolInfo()
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

	//	project.Withdraw(1, [2]int64{100, 4}) // team withdraw
	//project.Deposit(1, [3]int64{1, 10, 1}, false)
	//	project.Deposit(1, [3]int64{2, 1000, 0}, false)
	//project.EmergencyBurn()

	//	project.Strategy1(1000, 1)
	// project.Strategy1(100, 1)
	//project.AccountInfo()
	// project.Withdraw(1, [2]int64{100, 0})
	//project.Withdraw(1, [2]int64{100, 1})
	//project.Withdraw(1, [2]int64{100, 3})
	//	project.Alloc(1)  // removed public as internal method
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
	token0 := project.Network.TokenA
	token1 := project.Network.TokenB

	feetier := int64(10000) //10000, 3000, 500 //Network.FeeTier

	_protocolfee := big.NewInt(10)
	_quoteAmount := big.NewInt(1e18)
	_uniPortion := 20
	team := "0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440"
	strategy := project.Indi(0)

	project.VaultGen(networkId, acc, token0, token1, feetier, _protocolfee, _quoteAmount, _uniPortion, team, strategy)

}

func setVaultBridge() {

	// vault bridge v1 on goreli 0x033F3C5eAd18496BA462783fe9396CFE751a2342
	// vault bridge v2 on goreli 0x428EeA0B87f8E0f5653155057f58aaaBb667A3ec

	//	project.DeployVaultBridge()

	//project.DeployVaultAdmin()
	// project.AuthAdmin(Network.VaultAdmin, "0xfd8a5AE495Df1CA34F90572cb99A33B27173eDe1")
	// return

	project.SetVaultAddress("0xBDa573F33c18c69Cda004d5035e44Dd4635f69d1", 0)
	project.SetVaultAddress("0x45aE8C6868F068d2e4AC774106aA86C8489E6E60", 1)
	// project.SetVaultAddress("0xf231F818a111FE5d2EFf006451689eCBbf5ef159", 1)

	// project.GetVaultAddress(0)
	// project.GetVaultAddress(1)

}

func BuildAll() {

	sw.ViewOnly = false

	sw.DeployFactory = 0
	//...manual step... update the new factory addres in networks.go

	sw.DeployToken = 1
	sw.TokenParam[0] = project.TokenStruct{"f weth 1", "fWETH1", 18, big.NewInt(50000000000000)}
	sw.TokenParam[1] = project.TokenStruct{"f usdc 1", "fUSDC1", 6, big.NewInt(500000000000000000)}
	//...manual step... update the new token addres in networks.go

	sw.TransferToken = 0
	sw.TransferParam[0] = TransferStruct{0, project.X1E18(900), project.Network.TokenA, "0xeBb29c07455113c30810Addc123D0D7Cd8637aea"}
	sw.TransferParam[1] = TransferStruct{0, project.X1E18(900), project.Network.TokenB, "0xeBb29c07455113c30810Addc123D0D7Cd8637aea"}

	sw.CreatePool = 1 // *Note: if token0+token1+fee = pool exists ERROR: createPool VM Exception while processing transaction: revert
	sw.InitialPool = 1
	//...manual step... update the new pool addres in networks.go

	sw.MintPool = 1
	sw.MintPoolParam = [3]int64{1000, 1, 1} // currently hardcoded 1000 * 1e18 as the liquidity,
	// .... manual step.... new vault may apply

	sw.DeployVault = 1
	//...manual step... update the new vault addres in networks.go

	sw.Deposit = 0
	sw.DepositParam = [3]int64{1, 1, 1} // {amount0, amount1 , account}

	sw.Strategy1 = 0
	sw.Strategy1Param = [3]int64{600, 60, 3} // {tick range, tickspacing, account}

	sw.Withdraw = 0
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
		project.Network.TokenA = token0
		project.Network.TokenB = token1

	}

	if sw.TransferToken > 0 {

		project.TokenTransfer(
			sw.TransferParam[0].AccountId,
			sw.TransferParam[0].Amount,
			sw.TransferParam[0].TokenAddress,
			sw.TransferParam[0].ToAddress)

		project.TokenTransfer(
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

	//	project.Deposit(sw.DepositParam, false) /// deposit token0 amount * 1e18, token1 amount * 1e6

	//	project.Rebalance(sw.Rebalance, sw.RebalanceParam) /// make sure Account = 0

	//	project.Withdraw(sw.Withdraw, sw.WithDrawParam) /// withdraw shares, input number in percentage %

	// print all deployed addresses
	for _, i := range project.InfoString {
		fmt.Println(i)
	}

	project.CheckPrice(false, 3)
	project.Equation(false, false)

}
