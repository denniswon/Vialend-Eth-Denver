package include

import (
	_ "bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	_ "os"
	"strings"
	"time"

	"strconv"

	"../config"
)

const (
	Buy  = 1
	Sell = -1
)

/// read file
/// parse file
/// get range
/// calc range
/// trigger Strategy1

func ReadSignal() int64 {

	// read signal file
	content, _ := ioutil.ReadFile("C:\\Program Files (x86)\\MetaTrader 4 - 1\\tester\\files\\_Signal.txt")

	// trim any space or hard cariage
	s := strings.TrimSpace(string(content))

	// parse string to int64
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	return i

}

func doRebal(rng int64) {

	fmt.Println("Rebalance Triggered , new range:", rng)
	Strategy1(rng, 1)

}

func Rebal(max int, interval int) {

	i := 0
	oldrange := int64(0)

	for {

		newRange := ReadSignal()
		i++

		fmt.Println("Reading range:........(", i, ")", "{", newRange, "}")

		if newRange > 0 {

			// only do rebalance when new range is detected
			if newRange != oldrange {
				oldrange = newRange
				doRebal(newRange)
			}

			// in testing mode, in case something went wrong
			if i > max { // 30 s per check, 1000 checks for 8 hours
				os.Exit(3)
			}

		}

		//testing mode
		time.Sleep(time.Duration(interval) * time.Second)

	}

}

func check(msg string, e error) {
	if e != nil {
		log.Fatal(e, msg)
	}
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
func MovePrice(dir int) {

	accountId := 0 // accountid 5 = 0x4F211267896C4D3f2388025263AC6BD67B0f2C54

	amountX1e18 := X1E18(1) // 1 weth

	//0x04B1560f4F58612a24cF13531F4706c817E8A5Fe   //weth/usdc
	//0x1738f9aAB1d370a6d0fd56a18f113DbD9e1DCd4e  // weth/dai
	_pool := config.Network.Pool // check networkid

	var zeroForOne bool
	if dir < 0 {
		// eth price go down
		zeroForOne = true
	}
	if dir > 0 {
		// eth price go up
		zeroForOne = false
	}

	Swap2(accountId, amountX1e18, zeroForOne, _pool)

}

// generate fees by swap in the uniswap pool
func GenFees(t int, sleepSeconds time.Duration) {

	for i := 0; i < t; i++ {

		MovePrice(Sell)
		MovePrice(Buy)

		time.Sleep(sleepSeconds * time.Second)
	}
}
