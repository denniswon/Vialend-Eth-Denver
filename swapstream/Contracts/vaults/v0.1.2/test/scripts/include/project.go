package include

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"time"

	factory "../../../../../../../uniswap/v3/deploy/UniswapV3Factory"
	pool "../../../../../../../uniswap/v3/deploy/UniswapV3Pool"
	"../config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	/*

		mintcallee "../../uniswap/v3/deploy/TestUniswapV3Callee"
		factory "../../uniswap/v3/deploy/UniswapV3Factory"
		pool "../../uniswap/v3/deploy/UniswapV3Pool"
		token "../../uniswap/v3/deploy/token"
	*/)

func CreatePool(fee int64, do bool) common.Address {

	factoryAddress := common.HexToAddress(config.Network.Factory)

	instance, err := factory.NewApi(factoryAddress, config.Client)
	if err != nil {
		log.Fatal("factory.NewApi ", err)
	}

	tokenA := common.HexToAddress(config.Network.TokenA)
	tokenB := common.HexToAddress(config.Network.TokenB)
	if do {

		config.NonceGen()
		tx, err := instance.CreatePool(config.Auth,
			tokenA,
			tokenB,
			big.NewInt(fee),
		)

		if err != nil {
			log.Fatal("createPool ", err)
		}

		time.Sleep(config.Network.PendingTime * time.Second)

		//get the transaction hash
		_ = tx // reserve
		//fmt.Println("createpool tx sent: %s", tx.Hash().Hex())

	}

	poolAddress, err := instance.GetPool(&bind.CallOpts{}, tokenA, tokenB, big.NewInt(fee))
	if err != nil {
		log.Fatal("getpool", err)
	}
	config.Network.Pool = poolAddress.String()

	fmt.Println("poolAddress:", poolAddress)
	return poolAddress

}

func InitialPool(do bool) {

	if !do {
		return
	}

	//	poolAddress, err := newInstance.GetPool(&bind.CallOpts{}, Network.TokenA, Network.TokenB, fee)

	poolInstance, err := pool.NewApi(common.HexToAddress(config.Network.Pool), config.Client)

	if err != nil {
		log.Fatal(err)
	}

	token0, err := poolInstance.Token0(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	//

	//# Set ETH/USDC price to 2000
	inverse := token0 == common.HexToAddress(config.Network.TokenA)

	fmt.Println("token0 is takenA = ", inverse)

	var price float64
	var sqrtPriceX96 float64

	if inverse { // token0 = usdc
		price = 1e18 / 2000e6 //  1 eth = 2000usdc
	} else { // token0 = eth
		price = 2000e6 / 1e18 //  2000usdc = 1 eth
	}

	fmt.Println("price:", price)

	sqrtPriceX96 = math.Floor(math.Sqrt(price) * (1 << 96))

	fmt.Println("sqrPriceX96:", sqrtPriceX96)

	fmt.Println("sqrtP*96 -> Price ", math.Floor((sqrtPriceX96*sqrtPriceX96)/math.Pow(2, 2*96)), " == Price:", price)

	s := fmt.Sprintf("%.0f", sqrtPriceX96)

	bigSqrtPX96, _ := new(big.Int).SetString(s, 10)
	fmt.Println("bigInt.sqrtPriceX96:", bigSqrtPX96)

	config.NonceGen()

	tx, err := poolInstance.Initialize(config.Auth, bigSqrtPX96)

	if err != nil {
		log.Fatal("initialize err:", err)
	}

	fmt.Println("initial tx sent:", tx.Hash().Hex())

	time.Sleep(config.Network.PendingTime * time.Second)

	//readstring("check the address and  press any key when ready... ")

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	fmt.Println("slot0.0:", slot0.SqrtPriceX96)
	fmt.Println("slot0.1:", slot0.Tick)

	feeRate, _ := poolInstance.Fee(&bind.CallOpts{})
	fmt.Println("Fee:", feeRate)

	liquidity, _ := poolInstance.Liquidity(&bind.CallOpts{})
	fmt.Println("liquidity:", liquidity)

	/*

		///initial range
		tickLower := big.NewInt(-887220)
		tickUpper := big.NewInt(887220)

		//tickLower := big.NewInt(126600)
		//tickUpper := big.NewInt(206800)

		//tickLower := big.NewInt(-60)
		//tickUpper := big.NewInt(18660)

			rl := (slot0.Tick).Cmp(tickLower)
			ru := (slot0.Tick).Cmp(tickUpper)

			fmt.Println("in range? ", ru == -1 && rl == 1)
			fmt.Println("lower, tick, upper:? ", tickLower, ",", tickUpper, ",", slot0.Tick)

			feeRate, _ := poolInstance.Fee(&bind.CallOpts{})
			fmt.Println("Fee:", feeRate)

			liquidity, _ := poolInstance.Liquidity(&bind.CallOpts{})
			fmt.Println("liquidity:", liquidity)
	*/
}
