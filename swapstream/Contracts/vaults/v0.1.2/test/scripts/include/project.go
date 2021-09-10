package include

import (
	"fmt"
	"log"
	"math"
	"math/big"

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

func CreatePool(do int) common.Address {
	if do <= 0 {
		return common.HexToAddress("0x0")
	}

	fmt.Println(".......................Create Pool. ..................")

	factoryAddress := common.HexToAddress(config.Network.Factory)

	instance, err := factory.NewApi(factoryAddress, config.Client)
	if err != nil {
		log.Fatal("factory.NewApi ", err)
	}

	tokenA := common.HexToAddress(config.Network.TokenA)
	tokenB := common.HexToAddress(config.Network.TokenB)
	fee := big.NewInt(config.Network.Fee)
	config.NonceGen()
	tx, err := instance.CreatePool(config.Auth,
		tokenA,
		tokenB,
		fee,
	)

	if err != nil {
		log.Fatal("createPool ", err)
	}

	config.Readstring("createpool done, wait for pending ... next getPool... ")

	//get the transaction hash
	_ = tx // reserve
	//fmt.Println("createpool tx sent: %s", tx.Hash().Hex())

	poolAddress, err := instance.GetPool(&bind.CallOpts{}, tokenA, tokenB, fee)
	if err != nil {
		log.Fatal("getpool ", err)
	}

	config.Network.Pool = poolAddress.String()

	config.Readstring("createpool done, wait for pending ... next... ")

	fmt.Println("poolAddress:", poolAddress)

	return poolAddress

}

func GetPoolFromToken() common.Address {

	factoryAddress := common.HexToAddress(config.Network.Factory)

	instance, err := factory.NewApi(factoryAddress, config.Client)
	if err != nil {
		log.Fatal("factory.NewApi ", err)
	}

	tokenA := common.HexToAddress(config.Network.TokenA)
	tokenB := common.HexToAddress(config.Network.TokenB)
	fee := big.NewInt(config.Network.Fee)

	poolAddress, err := instance.GetPool(&bind.CallOpts{}, tokenA, tokenB, fee)
	if err != nil {
		log.Fatal("getpool ", err)
	}

	config.Network.Pool = poolAddress.String()

	fmt.Println("pool address set:", poolAddress)
	fmt.Println("by token A address:", tokenA)
	fmt.Println("by token B address:", tokenB)
	fmt.Println("by fee tier:", fee)

	return poolAddress

}

func InitialPool(do int) {

	if do <= 0 {
		return
	}

	fmt.Println(".........Initialize Pool ..............")

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

	sqrtPriceX96 = getSqrtPriceX96(price)

	fmt.Println("sqrPriceX96:", sqrtPriceX96)

	fmt.Println("sqrtP*96 -> Price ", math.Floor((sqrtPriceX96*sqrtPriceX96)/math.Pow(2, 2*96)), " == Price:", price)

	bigSqrtPX96 := getBigFromFloat64(sqrtPriceX96)
	fmt.Println("bigInt.sqrtPriceX96:", bigSqrtPX96)

	config.NonceGen()

	tx, err := poolInstance.Initialize(config.Auth, bigSqrtPX96)

	if err != nil {
		log.Fatal("initialize err:", err)
	}

	fmt.Println("initial tx sent:", tx.Hash().Hex())

	config.Readstring("initial pool done, waiting for pending... next ... ")

	PoolInfo()

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

func PoolInfo() {

	//	poolAddress, err := newInstance.GetPool(&bind.CallOpts{}, Network.TokenA, Network.TokenB, fee)

	fmt.Println(".............. Pool Info ....... ")

	fmt.Println("pool address:", common.HexToAddress(config.Network.Pool))

	poolInstance, err := pool.NewApi(common.HexToAddress(config.Network.Pool), config.Client)

	if err != nil {
		log.Fatal(err)
	}

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	fmt.Println("slot0.SqrtPriceX96:", slot0.SqrtPriceX96)
	fmt.Println("slot0.Tick:", slot0.Tick)

	feeRate, _ := poolInstance.Fee(&bind.CallOpts{})
	fmt.Println("Fee:", feeRate)

	liquidity, _ := poolInstance.Liquidity(&bind.CallOpts{})
	fmt.Println("liquidity in pool:", liquidity)

}
func getSqrtPriceX96(price float64) float64 {
	return math.Floor(math.Sqrt(price) * (1 << 96))
}

func getBigFromFloat64(v float64) *big.Int {

	s := fmt.Sprintf("%.0f", v)
	b, ok := new(big.Int).SetString(s, 10)
	if !ok {
		log.Fatal("getBigFromFloat64 Err:", ok)
	}
	return b

}
