package include

import (
	"../config"

	_ "bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	_ "os"
	"strings"
	"time"

	"strconv"
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
	Strategy1(1, [3]int64{rng, 60, 1})

}

func Rebal(interval int) {

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
			if i > 1000 { // 30 s per check, 1000 checks for 8 hours
				os.Exit(3)
			}

		}

		//testing mode
		time.Sleep(time.Duration(interval) * time.Second)

	}

}

// generate fees by swap in the uniswap pool
func GenFees(t int, acc int, amount int64, sleepSeconds time.Duration) {

	accountId := acc

	zeroForOne := false

	amountX1e18 := config.X1E18(amount)

	for i := 0; i < t; i++ {

		// call swap2 function in pool.go
		Swap2(accountId, amountX1e18, zeroForOne)
		Swap2(accountId, amountX1e18, !zeroForOne)

		time.Sleep(sleepSeconds * time.Second)

	}
}

func check(msg string, e error) {
	if e != nil {
		log.Fatal(e, msg)
	}
}
