package include

import (
	"fmt"
	"log"
	"math/big"
	"time"

	//	"time"

	//	factory "../../../../../../../uniswap/v3/deploy/UniswapV3Factory"
	vault "../../../deploy/FeeMaker"
	"../config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var vaultAddress = common.HexToAddress(config.Network.Vault)

func Deposit(do bool) {

	if !do {
		return
	}

	fmt.Println("vaultAddress: ", vaultAddress)

	instance, err := vault.NewApi(vaultAddress, config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	//test deposit
	amountToken0 := big.NewInt(1e18).Mul(big.NewInt(1e18), big.NewInt(9))
	amountToken1 := big.NewInt(1e18).Mul(big.NewInt(1e18), big.NewInt(1))

	config.NonceGen()
	tx, err := instance.Deposit(config.Auth,
		amountToken0,
		amountToken1,
		config.FromAddress,
	)

	if err != nil {
		log.Fatal("deposit err: ", err)
	}

	//get the transaction hash
	fmt.Println("deposit tx sent: %s", tx.Hash().Hex())

	time.Sleep(config.Network.PendingTime * time.Second)

}

func Rebalance(do bool) {

	if !do {
		return
	}

	fmt.Println("vaultAddress: ", vaultAddress)

	instance, err := vault.NewApi(vaultAddress, config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	totals := GetTotalAmounts(true)
	balance0 := totals.Total0
	balance1 := totals.Total1

	poolToken0, poolToken1 := big.NewInt(0), big.NewInt(0) // getPoolAmounts()

	allToken0 := balance0.Add(balance0, poolToken0)
	allToken1 := balance1.Add(balance1, poolToken1)

	config.NonceGen()

	tx, err := instance.Rebalance(auth,
		tickLower,
		tickUpper,
		zeroForOne > 0, //swap token0 for token 1
		swapAmount,
		sqrtPriceLimitX96,
	)

	if err != nil {
		panic(err)
	}

	//get the transaction hash
	fmt.Println("rebal tx sent: %s", tx.Hash().Hex())

	time.Sleep(config.Network.PendingTime * time.Second)

}

func GetTotalAmounts(do bool) *struct {
	Total0 *big.Int
	Total1 *big.Int
} {
	if !do {
		return nil
	}
	instance, err := vault.NewApi(vaultAddress, config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	totals, err := instance.GetTotalAmounts(&bind.CallOpts{})
	fmt.Printf("token0=%d,token1=%d \n", totals.Total0.Div(totals.Total0, big.NewInt(1e18)), totals.Total1.Div(totals.Total1, big.NewInt(1e18)))
	return &totals

}

func Approve(do bool) {

	if !do {
		return
	}

	fmt.Println("vaultAddress: ", vaultAddress)

	/*	instance, err := vault.NewApi(vaultAddress, config.Client)
		if err != nil {
			log.Fatal("vault.NewApi ", err)
		}
	*/
	token0Instance, err := vault.NewApi(common.HexToAddress(config.Network.TokenA), config.Client)
	if err != nil {
		log.Fatal("token0Instance,", err)
	}

	config.NonceGen()
	var maxToken0, _ = new(big.Int).SetString("90000000000000000000000", 10)
	tx0, err := token0Instance.Approve(config.Auth, vaultAddress, maxToken0)

	if err != nil {
		log.Fatal("tx0, ", err)
	}

	fmt.Println("approve amountoken0: %s", tx0.Hash().Hex())

	time.Sleep(config.Network.PendingTime * time.Second)
	//	readstring("done... Press any key to continue when ready..........")

	token1Instance, err := vault.NewApi(common.HexToAddress(config.Network.TokenB), config.Client)
	if err != nil {
		log.Fatal("token0Instance,", err)
	}
	config.NonceGen()
	//var amountToken0, _ = new(big.Int).SetString("90000000000000000000000", 10)
	var maxToken1, _ = new(big.Int).SetString("90000000000000000000000", 10)

	tx1, err := token1Instance.Approve(config.Auth, vaultAddress, maxToken1)

	if err != nil {
		log.Fatal("tx1, ", err)
	}

	fmt.Println("approve amountoken1: %s", tx1.Hash().Hex())

	time.Sleep(config.Network.PendingTime * time.Second)

}
