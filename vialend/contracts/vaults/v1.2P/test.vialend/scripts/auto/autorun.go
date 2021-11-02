package main

import (
	"fmt"
	"time"

	"../config"
	project "../project"
)

func main() {

	fmt.Println("Env: NetworkId=", config.Networkid, ",client=", config.Network.ProviderUrl[config.ProviderSortId])
	project.Init()

	project.Quiet = false

	movePrice()

	//GenFees(10, 5, 1, 5) // swap times, swap account, amount , sleepSeconds

	//project.Rebal(20) //interval seconds

	//AutoBoth() // Rebal(30) + GenFees

}

func AutoBoth() {

	go GenFees(10, 5, 1, 30) // {swap num of times, account, amount , sleepSeconds}

	project.Rebal(10, 20) // max rebalance run ,  interval seconds between each rebalance

}

/*
zeroforone = true,
	1. liquidity pool change: 	more eth, less usdc
	2. swapper: 				swap eth for usdc,
	3. trading action: 			sell eth , buy usdc
	4. vault change: 			more eth, less usdc
	5. price change: 			eth price go down

zeroforone = false,
	1. liquidity pool change: 	less eth, more usdc  (less 0, more 1 )
	2. swapper: 				swap usdc for eth,   (swap usdc for eth)
	3. trading action: 			sell usdc , buy eth
	4. vault change: 			less eth, more usdc
	5. price change: 			eth price go up

*/

func movePrice() {

	accountId := 5 // accountid 5 = 0x4F211267896C4D3f2388025263AC6BD67B0f2C54

	amountX1e18 := config.X1E18(1) // 1 weth

	//0x04B1560f4F58612a24cF13531F4706c817E8A5Fe   //weth/usdc
	//0x1738f9aAB1d370a6d0fd56a18f113DbD9e1DCd4e  // weth/dai
	_pool := config.Network.Pool // check networkid

	// eth price go down
	// zeroForOne := true
	// project.Swap2(accountId, amountX1e18, zeroForOne, _pool)
	// project.Swap2(accountId, amountX1e18, zeroForOne, _pool)

	// eth price go up
	zeroForOne := false
	project.Swap2(accountId, amountX1e18, zeroForOne, _pool)
	project.Swap2(accountId, amountX1e18, zeroForOne, _pool)

}

// generate fees by swap in the uniswap pool
func GenFees(t int, acc int, amount int64, sleepSeconds time.Duration) {

	accountId := acc

	zeroForOne := false

	amountX1e18 := config.X1E18(amount)

	//0x04B1560f4F58612a24cF13531F4706c817E8A5Fe   //weth/usdc
	//0x1738f9aAB1d370a6d0fd56a18f113DbD9e1DCd4e  // weth/dai
	_pool := "0x04B1560f4F58612a24cF13531F4706c817E8A5Fe"

	for i := 0; i < t; i++ {

		// call swap2 function in pool.go
		project.Swap2(accountId, amountX1e18, zeroForOne, _pool)
		project.Swap2(accountId, amountX1e18, !zeroForOne, _pool)

		time.Sleep(sleepSeconds * time.Second)

	}
}
