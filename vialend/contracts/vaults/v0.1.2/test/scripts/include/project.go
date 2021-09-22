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

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Create Pool. ..................")
	fmt.Println("----------------------------------------------")

	factoryAddress := common.HexToAddress(config.Network.Factory)

	instance, err := factory.NewApi(factoryAddress, config.Client)
	if err != nil {
		log.Fatal("factory.NewApi ", err)
	}

	tokenA := common.HexToAddress(config.Network.TokenA)
	tokenB := common.HexToAddress(config.Network.TokenB)
	fee := big.NewInt(config.Network.FeeTier)
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
	fmt.Println("poolAddress:", poolAddress)

	config.Readstring("createpool done, wait for pending ... next... ")

	return poolAddress

}

func InitialPool(do int) {

	if do <= 0 {
		return
	}

	fmt.Println("----------------------------------------------")
	fmt.Println(".........Initialize Pool ..............")
	fmt.Println("----------------------------------------------")

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
		price = 1e18 / 2000e18 //  1 eth = 2000usdc
	} else { // token0 = eth
		price = 1e18 / 1e18 //  2000usdc = 1 eth
	}

	fmt.Println("price:", price)

	sqrtPriceX96 = getSqrtPriceX96(price)

	//fmt.Println("sqrPriceX96:", sqrtPriceX96)

	fmt.Println("sqrtP*96 -> Price ", math.Floor((sqrtPriceX96*sqrtPriceX96)/math.Pow(2, 2*96)), " == Price:", price)

	bigSqrtPX96 := getBigFromFloat64(sqrtPriceX96)
	//fmt.Println("bigInt.sqrtPriceX96:", bigSqrtPX96)

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

func GetPoolFromToken() common.Address {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Calculate Pool address by token and fee tier ..................")
	fmt.Println("----------------------------------------------")

	factoryAddress := common.HexToAddress(config.Network.Factory)

	instance, err := factory.NewApi(factoryAddress, config.Client)
	if err != nil {
		log.Fatal("factory.NewApi ", err)
	}

	tokenA := common.HexToAddress(config.Network.TokenA)
	tokenB := common.HexToAddress(config.Network.TokenB)
	fee := big.NewInt(config.Network.FeeTier)

	poolAddress, err := instance.GetPool(&bind.CallOpts{}, tokenA, tokenB, fee)
	if err != nil {
		log.Fatal("getpool ", err)
	}

	config.Network.Pool = poolAddress.String()

	fmt.Println("token A address:", tokenA)
	fmt.Println("token B address:", tokenB)
	fmt.Println("fee tier:", fee)
	fmt.Println("pool address calc:", poolAddress)

	return poolAddress

}

func PoolInfo() {

	//	poolAddress, err := newInstance.GetPool(&bind.CallOpts{}, Network.TokenA, Network.TokenB, fee)

	fmt.Println("----------------------------------------------")
	fmt.Println(".............. Pool Info ....... ")
	fmt.Println("----------------------------------------------")

	fmt.Println("pool address:", common.HexToAddress(config.Network.Pool))

	poolInstance, err := pool.NewApi(common.HexToAddress(config.Network.Pool), config.Client)

	if err != nil {
		log.Fatal(err)
	}

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	fmt.Println("slot0.SqrtPriceX96:", slot0.SqrtPriceX96)
	fmt.Println("slot0.Tick:", slot0.Tick)

	price := getPrice(slot0.SqrtPriceX96, slot0.Tick)
	fmt.Println("getPrice ", price)

	sqrtPf := new(big.Float)
	///convert big Int to big Float
	sqrtPf.SetString(slot0.SqrtPriceX96.String())
	///convert big float to float64
	sp64, _ := sqrtPf.Float64()
	// operate float64
	sqrtPx962Price := math.Floor((sp64 * sp64) / math.Pow(2, 2*96))
	///verify the price
	fmt.Println("veryfy sqrtP*96 = ", sqrtPx962Price, ",  getPrice = ", price)

	fmt.Println("Price in decimals: ", config.Pricef(price, int(config.Decimals1)))

	feeRate, _ := poolInstance.Fee(&bind.CallOpts{})
	fmt.Println("Fee tier:", feeRate)

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
