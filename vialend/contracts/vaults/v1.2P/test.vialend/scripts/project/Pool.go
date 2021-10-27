package include

import (
	"fmt"
	"log"
	"math"
	"math/big"

	factory "../../../../../../../uniswap/v3/deploy/UniswapV3Factory"
	pool "../../../../../../../uniswap/v3/deploy/UniswapV3Pool"

	token "../../../../../Tokens/erc20/deploy/Token"
	"../config"

	swapCallee "../../../deploy/TestUniswapV3Callee"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	/*

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

	config.AddSettingString("pool address:", poolAddress.String())
	config.AddSettingString("pool fee tier:", fee.String())

	config.Readstring("createpool done, wait for pending ... next... ")

	return poolAddress

}

func Swap2(accountId int, swapAmount *big.Int, zeroForOne bool) {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................swap2. ..................")
	fmt.Println("----------------------------------------------")

	poolAddress := common.HexToAddress(config.Network.Pool)
	calleeInstance, err := swapCallee.NewApi(common.HexToAddress(config.Network.Callee), config.Client)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("pool address:", poolAddress)
	fmt.Println("callee address:", config.Network.Callee)

	poolInstance, err := pool.NewApi(poolAddress, config.Client)
	if err != nil {
		log.Fatal("poolInstance err:", err)
	}

	TokenA, _ := poolInstance.Token0(&bind.CallOpts{})
	TokenB, _ := poolInstance.Token1(&bind.CallOpts{})
	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	var maxToken0 = config.X1E18(9e18) //new(big.Int).SetString("900000000000000000000000000000", 10)
	var maxToken1 = config.X1E6(9e6)   // new(big.Int).SetString("900000000000000000000000000000", 10)

	config.ChangeAccount(accountId)

	ApproveToken(TokenA, maxToken0, config.Network.Callee)
	ApproveToken(TokenB, maxToken1, config.Network.Callee)

	recipient := config.FromAddress

	config.NonceGen()

	MIN_SQRT_RATIO := big.NewInt(4295128739)
	MAX_SQRT_RATIO, _ := new(big.Int).SetString("1461446703485210103287273052203988822378723970342", 10)
	sqrtP0for1 := new(big.Int).Add(MIN_SQRT_RATIO, big.NewInt(1))
	sqrtP1for0 := new(big.Int).Sub(MAX_SQRT_RATIO, big.NewInt(1))

	if zeroForOne {

		tx, err := calleeInstance.SwapExact0For1(config.Auth,
			poolAddress,
			swapAmount,
			recipient,
			sqrtP0for1)

		fmt.Println("zeroForOne =", zeroForOne, swapAmount, sqrtP0for1, slot0.SqrtPriceX96, slot0.Tick)

		if err != nil {
			panic(err)
		}

		config.TxConfirm(tx.Hash())

	} else {

		tx, err := calleeInstance.Swap1ForExact0(config.Auth,
			poolAddress,
			swapAmount,
			recipient,
			sqrtP1for0)

		fmt.Println("zeroForOne =", zeroForOne, swapAmount, sqrtP1for0, slot0.SqrtPriceX96, slot0.Tick)

		if err != nil {
			panic(err)
		}

		config.TxConfirm(tx.Hash())
	}

	config.ChangeAccount(config.Account)
}
func InitialPool(do int) {

	if do <= 0 {
		return
	}

	fmt.Println("----------------------------------------------")
	fmt.Println(".........Initialize Pool ..............")
	fmt.Println("----------------------------------------------")

	//	poolAddress, err := newInstance.GetPool(&bind.CallOpts{}, Network.TokenA, Network.TokenB, fee)

	fmt.Println("pool address:", common.HexToAddress(config.Network.Pool))

	poolInstance, err := pool.NewApi(common.HexToAddress(config.Network.Pool), config.Client)

	if err != nil {
		log.Fatal(err)
	}

	token0, err := poolInstance.Token0(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	_, _, symbol, _, _ := GetTokenInstance(token0.String())
	//
	fmt.Println("token0 is ", symbol)

	//# Set ETH/USDC price to 2000 . tokenA eth

	var price float64

	if token0 == common.HexToAddress(config.Network.TokenA) {
		price = 3000.0
	} else if token0 == common.HexToAddress(config.Network.TokenB) {
		price = (1.0 / 3000.0) //  1 eth = 2000usdc
	} else {
		log.Fatal("wrong token address ", token0, config.Network.TokenA, config.Network.TokenB)
	}

	fmt.Println("price:", price)

	sqrtPriceX96 := getSqrtPriceX96(price)

	bigIntSqrtPX96 := config.Float64ToBigInt(sqrtPriceX96)

	fmt.Println("bigInt . sqrtPriceX96:", bigIntSqrtPX96)

	//os.Exit(3)

	config.NonceGen()

	tx, err := poolInstance.Initialize(config.Auth, bigIntSqrtPX96)
	if err != nil {
		log.Fatal("initialize err:", err)
	}
	fmt.Println("initial tx sent:", tx.Hash().Hex())

	config.Readstring("initial pool done, waiting for pending... next ... ")
	/*
		tx, err = poolInstance.IncreaseObservationCardinalityNext(config.Auth, 100)

		fmt.Println("IncreaseObservationCardinalityNext sent:", tx.Hash().Hex())

		config.Readstring(" waiting for pending... next ... ")
	*/
	PoolInfo()

}

func GetPoolFromToken() common.Address {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................Calculate Pool address by token and fee tier ..................")
	fmt.Println("----------------------------------------------")

	factoryAddress := common.HexToAddress(config.Network.Factory)

	factoryInstance, err := factory.NewApi(factoryAddress, config.Client)
	if err != nil {
		log.Fatal("factory.NewApi ", err)
	}

	tokenA := common.HexToAddress(config.Network.TokenA)
	tokenB := common.HexToAddress(config.Network.TokenB)
	fee := big.NewInt(config.Network.FeeTier)

	poolAddress, err := factoryInstance.GetPool(&bind.CallOpts{}, tokenA, tokenB, fee)
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

	token0, _ := poolInstance.Token0(&bind.CallOpts{})
	token1, _ := poolInstance.Token1(&bind.CallOpts{})
	feeRate, _ := poolInstance.Fee(&bind.CallOpts{})

	fmt.Println("Token0 address:", token0)
	fmt.Println("Token1 address:", token1)
	fmt.Println("Fee tier:", feeRate)

	_, name, symbol, decimals, _ := GetTokenInstance(token0.String())
	fmt.Println("Token0 info:", name, symbol, decimals)

	_, name, symbol, decimals, _ = GetTokenInstance(token1.String())
	fmt.Println("Token1 info:", name, symbol, decimals)

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	fmt.Println("slot0.SqrtPriceX96:", slot0.SqrtPriceX96)
	fmt.Println("slot0.Tick:", slot0.Tick)

	price := getPrice(slot0.SqrtPriceX96, slot0.Tick)
	fmt.Println("price ", price)

	sqrtPf := new(big.Float)
	///convert big Int to big Float
	sqrtPf.SetString(slot0.SqrtPriceX96.String())
	///convert big float to float64
	sp64, _ := sqrtPf.Float64()
	// operate float64
	sqrtPx962Price := math.Floor((sp64 * sp64) / math.Pow(2, 2*96))
	///verify the price
	fmt.Println("veryfy sqrtP*96 = ", sqrtPx962Price, ",  getPrice = ", price)

	fmt.Println("Price in decimals: ", price)

	liquidity, _ := poolInstance.Liquidity(&bind.CallOpts{})
	//fmt.Println("liquidity in pool:", config.Pricef(liquidity, int(1e18)))
	fmt.Println("liquidity in pool:", liquidity)

	liquidity0, symbol0 := TokenInfo(common.HexToAddress(config.Network.TokenA), common.HexToAddress(config.Network.Pool))
	liquidity1, symbol1 := TokenInfo(common.HexToAddress(config.Network.TokenB), common.HexToAddress(config.Network.Pool))

	fmt.Println(symbol0, " in pool:", liquidity0)
	fmt.Println(symbol1, " in pool:", liquidity1)

}

func MintPool(liquidity int64, amount0 int64, amount1 int64) {

	fmt.Println("----------------------------------------------")
	fmt.Println(".........MintPool initil Pool from admin account.........  ")
	fmt.Println("----------------------------------------------")

	//# Add some liquidity
	max_tick := int64(887272/60) * 60 // // 60 * 60
	min_tick := max_tick * -1
	_ = min_tick

	config.Auth = config.GetSignature(config.Networkid, 0)

	token0Instance, err := token.NewApi(common.HexToAddress(config.Network.TokenA), config.Client)
	if err != nil {
		log.Fatal("token0Instance,", err)
	}
	var maxToken0, _ = new(big.Int).SetString("900000000000000000000000000000", 10)

	config.NonceGen()
	_, err = token0Instance.Approve(config.Auth, common.HexToAddress(config.Network.Callee), maxToken0)

	if err != nil {
		log.Fatal("token0 approve callee err, ", err)
	}

	token1Instance, err := token.NewApi(common.HexToAddress(config.Network.TokenB), config.Client)
	if err != nil {
		log.Fatal("token1Instance,", err)
	}
	var maxToken1, _ = new(big.Int).SetString("900000000000000000000000000000", 10)

	config.NonceGen()
	_, err = token1Instance.Approve(config.Auth, common.HexToAddress(config.Network.Callee), maxToken1)
	if err != nil {
		log.Fatal("token1 approve callee err ", err)
	}

	config.Readstring("Approves sent.....  wait for pending..next .. ")

	balance0, _ := token0Instance.BalanceOf(&bind.CallOpts{}, config.FromAddress)
	balance1, _ := token1Instance.BalanceOf(&bind.CallOpts{}, config.FromAddress)

	fmt.Println("balance0: ", balance0)
	fmt.Println("balance1: ", balance1)
	//os.Exit(1)

	calleeInstance := GetCalleeInstance()

	fmt.Println("max_tick, min_tick:", max_tick, min_tick)

	config.NonceGen()

	//uint128 amount

	var liquidityAmt *big.Int

	liquidityAmt = config.X1E18(liquidity)

	fmt.Println("liquidityAmt:", liquidityAmt)

	_, err = calleeInstance.Mint(config.Auth,
		common.HexToAddress(config.Network.Pool),
		config.FromAddress,
		big.NewInt(min_tick),
		big.NewInt(max_tick),
		liquidityAmt,
	)

	if err != nil {
		log.Fatal("callee mint err ", err)
	}

	///require governance. redo auth
	config.Auth = config.GetSignature(config.Networkid, config.Account)

	PoolInfo()
}

func TokenInfo(tokenAddress common.Address, owner common.Address) (*big.Int, string) {

	tokenInstance, err := token.NewApi(tokenAddress, config.Client)

	if err != nil {
		log.Fatal("tokenInfo NewApi err ", err)
	}

	balance, _ := tokenInstance.BalanceOf(&bind.CallOpts{}, owner)
	symbol, _ := tokenInstance.Symbol(&bind.CallOpts{})

	return balance, symbol

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
