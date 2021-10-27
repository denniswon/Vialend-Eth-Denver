package main

import (
	"math/big"

	config "../../scripts/config"
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
	TransferParam    []TransferStruct
	DeployFactory    int
	CreatePool       int
	InitialPool      int
	MintPool         int
	MintPoolParam    [3]int64
	DeployVault      int
	Deposit          int
	DepositParam     [4]int64
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

	// sw.DeployToken = 0
	// sw.TokenParam[0] = config.TokenStruct{"e weth 2", "eWETH2", 18, big.NewInt(50000000000000)}
	// sw.TokenParam[1] = config.TokenStruct{"e usdc 2", "eUSDC2", 6, big.NewInt(500000000000000000)}

	// transfer weth, usdc
	sw.TransferParam = append(sw.TransferParam, TransferStruct{
		1,
		config.X1E18(5),
		"0x6fD886fd1e728D9386Ba7fE721C856790758aDd9", //weth
		"0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E"})

	sw.TransferParam = append(sw.TransferParam, TransferStruct{
		1,
		config.X1E6(1),
		"0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C", //usdc
		"0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E"})

	transfer()
}

func transfer() {

	for i := 0; i < len(sw.TransferParam); i++ {
		config.TokenTransfer(
			sw.TransferParam[i].AccountId,
			sw.TransferParam[i].Amount,
			sw.TransferParam[i].TokenAddress,
			sw.TransferParam[i].ToAddress)
	}
}
