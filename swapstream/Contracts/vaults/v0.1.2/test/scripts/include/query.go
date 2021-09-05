package include

import (
	"fmt"
	"log"
	"math/big"
	"time"

	pool "../../../../../../../uniswap/v3/deploy/UniswapV3Pool"
	"../config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func CheckPrice(do bool) {

	if !do {
		return
	}

	var network = config.Networks[config.Networkid]

	poolAddress := common.HexToAddress(network.Pool)
	poolInstance, err := pool.NewApi(poolAddress, config.Client)
	if err != nil {
		log.Fatal("pool.NewApi ", err)
	}

	SECONDS_AGO := []uint32{10}

	for i := 0; i < 500; i++ {

		//TickCumulatives                    []*big.Int
		//SecondsPerLiquidityCumulativeX128s []*big.Int
		tickPrices, _ := poolInstance.Observe(&bind.CallOpts{}, SECONDS_AGO)
		before := tickPrices.TickCumulatives
		after := tickPrices.SecondsPerLiquidityCumulativeX128s

		fmt.Println("before, after:", before, ",", after)

		a1 := after[0].Sub(after[0], before[0])
		twap := a1.Div(a1, big.NewInt(int64(SECONDS_AGO[0])))
		slot0, err := poolInstance.Slot0(&bind.CallOpts{})
		if err != nil {
			log.Fatal("slot0 ", err)
		}

		last := slot0.Tick
		fmt.Println(twap, ",", last)
		//f := new(big.Float).SetInt(twap)
		//f64 := Float64(f, 0)
		//f := 10001 / 10000
		//**twap
		//x := 1.0001

		//one2 := big.NewInt(twap).Exp(1.0001, f)

		//fmt.Println("twap\t{twap}\t 1.0001**twap:" 1.0001**twap, twap)
		//print(f"last\t{last}\t{1.0001**last}")
		//print(f"trend\t{last-twap}")
		//print()
		fmt.Println("done")

		time.Sleep(time.Duration(SECONDS_AGO[0]) * time.Second)
	}
}
