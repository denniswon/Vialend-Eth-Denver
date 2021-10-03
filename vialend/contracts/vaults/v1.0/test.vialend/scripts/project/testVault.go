package include

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"time"

	//	"time"

	//	factory "../../../../../../../uniswap/v3/deploy/UniswapV3Factory"
	pool "../../../../../../../uniswap/v3/deploy/UniswapV3Pool"
	token "../../../../../Tokens/erc20/deploy/Token"
	cErc20 "../../../deploy/cErc20"

	//vault "../../../deploy/FeeMaker"
	callee "../../../deploy/TestUniswapV3Callee"
	vault "../../../deploy/vialendFeeMaker"
	"../config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var tickLower = new(big.Int)
var tickUpper = new(big.Int)

var qTickLower = new(big.Int) // for query only
var qTickUpper = new(big.Int) // for query only

func Deposit(do int, amount0 int64, amount1 int64) {

	if do <= 0 {
		return
	}
	fmt.Println("----------------------------------------------")
	fmt.Println(".........Deposit.........  ")
	fmt.Println("----------------------------------------------")

	fmt.Println("vaultAddress: ", common.HexToAddress(config.Network.Vault))
	fmt.Println("fromAddress: ", config.FromAddress)
	fmt.Println("TokenA:", config.Network.TokenA)
	fmt.Println("TokenB:", config.Network.TokenB)

	instance, err := vault.NewApi(common.HexToAddress(config.Network.Vault), config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	tokenAInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenA), config.Client)
	if err != nil {
		log.Fatal("tokenAInstance,", err)
	}

	bal, err := tokenAInstance.BalanceOf(&bind.CallOpts{}, config.FromAddress)
	fmt.Println("tokenA in Wallet ", bal)

	if err != nil {
		log.Fatal("bal err ", err)
	}

	tokenBInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenB), config.Client)
	if err != nil {
		log.Fatal("tokenBInstance,", err)
	}

	bal, err = tokenBInstance.BalanceOf(&bind.CallOpts{}, config.FromAddress)
	fmt.Println("tokenB in Wallet ", bal)

	//  amount0 * 10^decimals
	d0 := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(config.Token[0].Decimals)), nil)
	amountToken0 := new(big.Int).Mul(d0, big.NewInt(int64(amount0)))
	//	fmt.Println("amount0 * 10^decimals: ", d0)

	d1 := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(config.Token[1].Decimals)), nil)
	//	fmt.Println("amount1 * 10^decimals: ", d1)
	amountToken1 := new(big.Int).Mul(d1, big.NewInt(int64(amount1)))

	fmt.Println("amountToken0 to Vault:", amountToken0)
	fmt.Println("amountToken1 to Vault:", amountToken1)

	config.NonceGen()

	tx, err := instance.Deposit(config.Auth,
		amountToken0,
		amountToken1,
	)

	if err != nil {
		log.Fatal("deposit err: ", err)
	}

	//get the transaction hash
	fmt.Println("deposit tx: %s", tx.Hash().Hex())

	//	time.Sleep(config.Network.PendingTime * time.Second)
	config.Readstring("deposit sent...wait for pending..next .. ")

}

func Withdraw(do int, param [2]int64) {

	if do <= 0 {
		return
	}

	fmt.Println("----------------------------------------------")
	fmt.Println(".........Withdraw.........  ")
	fmt.Println("----------------------------------------------")

	/// set new account Auth
	config.Auth = config.GetSignature(config.Networkid, int(param[0]))
	recipient := config.FromAddress

	fmt.Println("Withdraw to  ", recipient)

	fmt.Println("vaultAddress: ", common.HexToAddress(config.Network.Vault))

	percent := big.NewInt(param[1])

	instance, err := vault.NewApi(common.HexToAddress(config.Network.Vault), config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}
	/*
		///get account's available shares
		myshares, err := instance.BalanceOf(&bind.CallOpts{}, config.FromAddress)
		if err != nil {
			log.Fatal("balance of myshare ", err)
		}

		fmt.Println("myshares  = ", myshares, 18)

		///calc required shares to withdraw

		shares := new(big.Int).Mul(myshares, percent)
		shares.Div(shares, big.NewInt(100))
		fmt.Println("shares required = ", percent, "% = ", shares)

		/// if required share > available share, set withdraw shares = awailable shares
		if shares.Cmp(myshares) == 1 {
			shares = myshares
		}

		if shares.Cmp(big.NewInt(0)) <= 0 {
			fmt.Println("share<=0 ", shares)
			return
		}
	*/
	config.NonceGen()
	tx, err := instance.Withdraw(config.Auth, percent)

	if err != nil {
		log.Fatal("withdraw: ", err)
	}

	/// reset account back
	config.Auth = config.GetSignature(config.Networkid, config.Account)

	//get the transaction hash
	fmt.Println("withdraw tx: ", tx.Hash().Hex())

	config.Readstring("withdraw sent.... wait for pending..next .. ")

}

func GetSwapInfo(rangeRatio int64) (amount0 float64, amount1 float64, swapAmount float64, zeroForOne bool) {

	poolInstance := GetPoolInstance()

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	tick := slot0.Tick
	sqrtPriceX96 := slot0.SqrtPriceX96

	//fmt.Println("tick: ", tick)
	//fmt.Println("sqrtPriceX96: ", sqrtPriceX96)

	Totals := GetTVL()

	fmt.Println("total locked token0: ", Totals.Total0)
	fmt.Println("total locked token1: ", Totals.Total1)

	pf := getPrice(sqrtPriceX96, tick)

	min := pf * (1 - float64(rangeRatio)/100)
	max := pf * (1 + float64(rangeRatio)/100)
	x := config.BigIntToFloat64(Totals.Total0) / math.Pow10(int(config.Token[0].Decimals))
	y := config.BigIntToFloat64(Totals.Total1) / math.Pow10(int(config.Token[1].Decimals))

	xDecimals, yDecimals := config.Token[0].Decimals, config.Token[1].Decimals

	//fmt.Println("pf, min, max, rangeRatio: ", pf, min, max, rangeRatio)

	//fmt.Println("pf,min, max, rangeRatio: ", pf, min, max, rangeRatio)

	a, b := getTicks(pf, min, max, float64(xDecimals), float64(yDecimals))
	tickA := math.Round(a/60) * 60
	tickB := math.Round(b/60) * 60

	//fmt.Println("tick a b:", tickA, tickB)

	tickLower = big.NewInt(int64(tickA)) //tickA) //big.NewInt(-1140) //
	tickUpper = big.NewInt(int64(tickB)) //tickB) //big.NewInt(840)   //

	//fmt.Println("---abminmax:", pf-float64(rangeRatio)/100, pf, a, b, tickLower, tickUpper)
	//os.Exit(3)
	// fmt.Println(pf, min, max, x, y)
	// fmt.Println(priceFromsqrtP)
	// fmt.Println(config.BigIntToFloat64(Total0))
	// fmt.Println(config.BigIntToFloat64(Total1))

	//amt0, amt1, swapAmount, zeroForOne := getBestAmounts(pf, min, max, x, y)
	return getBestAmounts(pf, min, max, x, y)
}
func EmergencyBurn() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".........Emergency & whitehacker.........  ")
	fmt.Println("----------------------------------------------")

	fmt.Println("vaultAddress: ", common.HexToAddress(config.Network.Vault))

	vaultInstance, err := vault.NewApi(common.HexToAddress(config.Network.Vault), config.Client)
	if err != nil {
		log.Fatal("emergency vault.NewApi ", err)
	}

	// tickLower, err := vaultInstance.CLow(&bind.CallOpts{})
	// tickUpper, err := vaultInstance.CHigh(&bind.CallOpts{})
	// liquidity, err := vaultInstance.GetSSLiquidity(&bind.CallOpts{}, qTickLower, qTickUpper)

	// lendingAmount0 := checkCTokenBalance("CETH", config.Network.LendingContracts.CETH)
	// lendingAmount1 := checkCTokenBalance("CUSDC", config.Network.LendingContracts.CUSDC)

	tx, err := vaultInstance.EmergencyBurn(config.Auth)

	if err != nil {
		log.Fatal("emergency tx err ", err)
	}

	fmt.Println("emergency tx: ", tx.Hash().Hex())

	config.Readstring("emergency sent sent.....  wait for pending..next .. white hacker to withdraw ")

}

func Whitehacker() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".........Whitehacker.........  ")
	fmt.Println("----------------------------------------------")

	fmt.Println("vaultAddress: ", common.HexToAddress(config.Network.Vault))
	vaultInstance, err := vault.NewApi(common.HexToAddress(config.Network.Vault), config.Client)
	if err != nil {
		log.Fatal("whitehacker vault.NewApi ", err)
	}

	tx, err := vaultInstance.WhiteHacker(config.Auth)
	if err != nil {
		log.Fatal("white hacker tx err ", err)
	}
	fmt.Println("whitehacker tx: ", tx.Hash().Hex())

	config.Readstring("whitehacker sent.....  ..bye...")

}

/// param0 : fullRangeSize, param1: tickspacing
func Strategy1(do int, param [3]int64) {

	if do <= 0 {
		return
	}

	fmt.Println("----------------------------------------------")
	fmt.Println(".........Strategy1 .........  ")
	fmt.Println("----------------------------------------------")

	fmt.Println("vaultAddress: ", common.HexToAddress(config.Network.Vault))

	vaultInstance, err := vault.NewApi(common.HexToAddress(config.Network.Vault), config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	//init ticklow and tickupp
	//GetSwapInfo(param[0])
	poolInstance := GetPoolInstance()
	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})
	tick := slot0.Tick

	hrange := new(big.Int).Div(big.NewInt(param[0]), big.NewInt(2))
	tickLower = new(big.Int).Sub(tick, hrange)
	tickUpper = new(big.Int).Add(tick, hrange)

	tickSpacing := param[1]
	tickLower.Div(tickLower, big.NewInt(tickSpacing)).Mul(tickLower, big.NewInt(tickSpacing))
	tickUpper.Div(tickUpper, big.NewInt(tickSpacing)).Mul(tickUpper, big.NewInt(tickSpacing))

	///require governance. redo auth
	config.Auth = config.GetSignature(config.Networkid, int(param[2]))

	config.NonceGen()

	fmt.Println("ticklower tickupper in...", tickLower, tickUpper)
	// set ticklower and tickupper
	//setRange(param)
	tx, err := vaultInstance.Strategy1(config.Auth,
		tickLower,
		tickUpper)

	if err != nil {
		log.Fatal("strateg1 tx err ", err)
	}

	///require governance. redo auth
	config.Auth = config.GetSignature(config.Networkid, config.Account)

	//get the transaction hash
	fmt.Println("strategy1 tx: ", tx.Hash().Hex())

	config.Readstring("strategy1 sent.....  wait for pending..next .. ")

	totalfee0, _ := vaultInstance.AccruedProtocolFees0(&bind.CallOpts{})
	totalfee1, _ := vaultInstance.AccruedProtocolFees1(&bind.CallOpts{})

	fmt.Printf("AccumulatedFee0:{%.4f}\n", totalfee0)
	fmt.Printf("AccumulatedFee1:{%.4f}\n", totalfee1)

}

func LendingInfo() {
	fmt.Println("----------------------------------------------")
	fmt.Println(".........Lending pool info .........  ")
	fmt.Println("----------------------------------------------")

	checkCTokenBalance("CUSDC", config.Network.LendingContracts.CUSDC)
	checkCTokenBalance("CETH", config.Network.LendingContracts.CETH)
	checkETHBalance()

	vaultInstance := GetVaultInstance()
	lendingamts, _ := vaultInstance.GetCAmounts(&bind.CallOpts{})

	fmt.Println("CToken0, Ctoken1:", lendingamts)

}

func checkETHBalance() *big.Int {

	Instance := GetVaultInstance()

	bal, err := Instance.GetETHBalance(&bind.CallOpts{})

	if err != nil {
		log.Fatal("ethbalance err ", err)
	}

	fmt.Println("eth balance: ", bal)

	return (bal)

}

func checkCTokenBalance(tokenName string, cTokenAddress string) *big.Int {

	cInstance := GetCTokenInstance(cTokenAddress)

	bal, err := cInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(config.Network.Vault))

	if err != nil {
		log.Fatal(" err ", err)
	}

	fmt.Println(tokenName, " balance: ", bal)

	return (bal)

}

func GetCTokenInstance(Address string) *cErc20.Api {

	instance, err := cErc20.NewApi(common.HexToAddress(Address), config.Client)
	if err != nil {
		log.Fatal("get token Instance,", err)
	}

	return instance
}

/// param0 : fullRangeSize, param1: tickspacing
func Rebalance(do int, param [2]int64) {

	if do <= 0 {
		return
	}

	fmt.Println("----------------------------------------------")
	fmt.Println(".........Rebalance .........  ")
	fmt.Println("----------------------------------------------")

	fmt.Println("vaultAddress: ", common.HexToAddress(config.Network.Vault))

	vaultInstance, err := vault.NewApi(common.HexToAddress(config.Network.Vault), config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	//Swap(1, param[0])
	_, _, swapAmount, zeroForOne := GetSwapInfo(param[0])

	///require governance. redo auth
	config.Auth = config.GetSignature(config.Networkid, 0)

	config.NonceGen()

	fmt.Println("ticklower tickupper in...", tickLower, tickUpper)
	// set ticklower and tickupper
	//setRange(param)
	tx, err := vaultInstance.Strategy0(config.Auth,
		tickLower,
		tickUpper,
		config.Float64ToBigInt(swapAmount),
		zeroForOne)

	if err != nil {
		log.Fatal("rebalance signature err ", err)
	}

	///require governance. redo auth
	config.Auth = config.GetSignature(config.Networkid, config.Account)

	//get the transaction hash
	fmt.Println("rebal tx: ", tx.Hash().Hex())

	config.Readstring("Rebalance sent.....  wait for pending..next .. ")

	totalfee0, _ := vaultInstance.AccruedProtocolFees0(&bind.CallOpts{})
	totalfee1, _ := vaultInstance.AccruedProtocolFees1(&bind.CallOpts{})

	fmt.Printf("AccumulatedFee0:{%.4f}\n", totalfee0)
	fmt.Printf("AccumulatedFee1:{%.4f}\n", totalfee1)

}

/// param0 : fullRangeSize, param1: tickspacing
func Swap(do int, rangeRatio int64) {

	if do <= 0 {
		return
	}

	fmt.Println("----------------------------------------------")
	fmt.Println(".........Swap .........  ")
	fmt.Println("----------------------------------------------")

	//fmt.Println("vaultAddress: ", common.HexToAddress(config.Network.Vault))

	//poolInstance := GetPoolInstance()
	vaultInstance := GetVaultInstance()

	amt0, amt1, swapAmount, zeroForOne := GetSwapInfo(rangeRatio)
	//fmt.Println("amount0 , amount1 desired:", amt0, amt1)
	_, _ = amt0, amt1

	/// sqrtPriceLimitX96 currently hardcoded in the contract
	/// zeroForOne ? TickMath.MIN_SQRT_RATIO + 1 : TickMath.MAX_SQRT_RATIO - 1;
	MIN_SQRT_RATIO := big.NewInt(4295128739)
	MAX_SQRT_RATIO, _ := new(big.Int).SetString("1461446703485210103287273052203988822378723970342", 10)

	sqrtPriceLimitX96 := new(big.Int).Sub(MAX_SQRT_RATIO, big.NewInt(1))

	if zeroForOne {
		sqrtPriceLimitX96 = new(big.Int).Add(MIN_SQRT_RATIO, big.NewInt(1))
	}

	///require governance. redo auth
	config.Auth = config.GetSignature(config.Networkid, 0)

	config.NonceGen()

	TokenDecimals := config.Token[0].Decimals

	if !zeroForOne {
		TokenDecimals = config.Token[1].Decimals
	}

	// swapamount * 1e(decimals)  i.e. usdc 1e6 = 10^6
	swapAmountIn := new(big.Int).SetInt64(int64(float64(swapAmount * math.Pow10(int(TokenDecimals)))))

	//fmt.Println("zeroForOne:", zeroForOne)
	//fmt.Println("swap amount: ", swapAmount)
	//fmt.Println("swap amount in:", swapAmountIn)
	//_ = swapAmount

	cTotal0, err := vaultInstance.GetBalance0(&bind.CallOpts{})
	cTotal1, err := vaultInstance.GetBalance1(&bind.CallOpts{})
	fmt.Println("current balance0 in vault: ", cTotal0)
	fmt.Println("current balance1 in vault: ", cTotal1)

	_, err = vaultInstance.Swap(config.Auth,
		swapAmountIn,
		zeroForOne,
		sqrtPriceLimitX96,
	)

	if err != nil {
		log.Fatal("vaultInstance.swap err ", err)
	}

	///require governance. redo auth
	config.Auth = config.GetSignature(config.Networkid, config.Account)

	//get the transaction hash
	//fmt.Println("swap tx: ", tx.Hash().Hex())

	config.Readstring("swap tx sent.....  wait for pending..next .. ")

	Total0, err := vaultInstance.GetBalance0(&bind.CallOpts{})
	Total1, err := vaultInstance.GetBalance1(&bind.CallOpts{})
	fmt.Println("new balance0 in vault: ", Total0)
	fmt.Println("new balance1 in vault: ", Total1)

	fmt.Println("swapped amt0: ", Total0.Sub(Total0, cTotal0))
	fmt.Println("swapped amt1: ", Total1.Sub(Total1, cTotal1))

}

func getBestAmounts(p float64, a float64, b float64, x float64, y float64) (amount0 float64, amount1 float64, swapAmount float64, zeroForOne bool) {

	sp := math.Sqrt(p) //p * *0.5
	sa := math.Sqrt(a) //a * *0.5
	sb := math.Sqrt(b) //b * *0.5
	// calculate the initial liquidity
	L := get_liquidity(x, y, sp, sa, sb)

	P1 := p
	sp1 := math.Sqrt(P1) // P1 * *0.5
	x1 := calculate_x(L, sp1, sb)
	y1 := calculate_y(L, sp1, sa)

	//fmt.Printf("x1={%.4f}\ny1={%.4f}\n", x1, y1)

	x1r := x1 / (x1 + y1/p)
	y1r := y1 / (y1 + x1*p)
	fmt.Println(x1r, y1r)

	xr := x / (x + y/p)
	yr := y / (y + x*p)
	fmt.Println(xr, yr)
	// 20/2000, 0.9908
	// 20 * 0.9908
	if x*p > y { // trade x for y
		zeroForOne = true

		r := xr - x1r

		swapAmount = math.Abs(x * r)

		amount0 = x - swapAmount

		amount1 = y + swapAmount*p

		//fmt.Println("newX=", amount0)
		//fmt.Println("newY=", amount1)

	} else if x*p < y { // trade y for x
		zeroForOne = false

		r := xr - x1r

		swapAmount = math.Abs(y * r)

		amount0 = x + swapAmount/p

		amount1 = y - swapAmount

		//fmt.Println("newX=", amount0)
		//fmt.Println("newY=", amount1)
	}

	//fmt.Printf("newX={%.18f}, newY={%.6f},swapamount={%.18f},zeroForOne={%t}\n", amount0, amount1, swapAmount, zeroForOne)

	return amount0, amount1, swapAmount, zeroForOne
}

///
func CollectProtocolFees(feeRate0 *big.Int, feeRate1 *big.Int) {

	fmt.Println("----------------------------------------------")
	fmt.Println(".........Collect Protocol fees .........  ")
	fmt.Println("----------------------------------------------")

	fmt.Println("vaultAddress: ", common.HexToAddress(config.Network.Vault))
	fmt.Println("PoolAddress: ", config.Network.Pool)

	vaultInstance, err := vault.NewApi(common.HexToAddress(config.Network.Vault), config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	totalfee0, _ := vaultInstance.AccruedProtocolFees0(&bind.CallOpts{})
	totalfee1, _ := vaultInstance.AccruedProtocolFees1(&bind.CallOpts{})
	fmt.Println("accumulatedFee0:", totalfee0)
	fmt.Println("accumulatedFee1:", totalfee1)

	to := config.FromAddress

	amount0 := totalfee0.Mul(totalfee0, feeRate0.Div(feeRate0, big.NewInt(100)))
	amount1 := totalfee1.Mul(totalfee1, feeRate1.Div(feeRate1, big.NewInt(100)))

	//get the transaction hash
	fmt.Println("collecting: ", amount0, amount1)

	///require governance. redo auth
	config.Auth = config.GetSignature(config.Networkid, 0)
	config.NonceGen()

	tx, err := vaultInstance.CollectProtocol(config.Auth,
		amount0,
		amount1,
		to,
	)

	if err != nil {
		log.Fatal("collect protocol fee err ", err)
	}

	///require governance. redo auth
	config.Auth = config.GetSignature(config.Networkid, config.Account)

	//get the transaction hash
	fmt.Println("collect protocol tx: ", tx.Hash().Hex())

	config.Readstring("collect fee sent.....  wait for pending..next .. ")

}

/// formula:
///
// 1. price = pow(1.0001,tick) * (1e(18-6)
///2.  (sqrtPricex96^2 * 1e(decimal0)/1e(decimal1) >> (96*2)
///3.  (sqrtPricex96^2 * 1e(decimal0)/1e(decimal1) / 2^(96*2)
// 4. javascript: JSBI.BigInt(sqrtPriceX96 *sqrtPriceX96* (1e(decimals_token_0))/(1e(decimals_token_1))/JSBI.BigInt(2) ** (JSBI.BigInt(192));
///5 solc: uint(sqrtPriceX96).mul(uint(sqrtPriceX96)).mul(1e(decimalsDiff)) >> (96 * 2);
func getPrice(SqrtPriceX96 *big.Int, tick *big.Int) float64 {

	/*
		sqrpricePow2 := new(big.Int).Mul(SqrtPriceX96, SqrtPriceX96)

		diff := config.Token[0].Decimals - config.Token[1].Decimals
		eDecimals := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(diff)), nil)
		oPrice := new(big.Int).Mul(sqrpricePow2, eDecimals)

		fmt.Println(eDecimals, oPrice)
		//Price := new(big.Int).Rsh(oPrice, 96*2)
		rsh192 := new(big.Int).Exp(big.NewInt(2), big.NewInt(192), nil)
		Price := new(big.Int).Div(oPrice, rsh192)
	*/

	// price = pow(1.0001,tick) * (1e(18-6) )
	var diff uint8
	if config.Token[0].Decimals > config.Token[1].Decimals {
		diff = config.Token[0].Decimals - config.Token[1].Decimals
	} else {
		diff = config.Token[1].Decimals - config.Token[0].Decimals
	}

	// fmt.Println("decimals0:", config.Token[0].Decimals)
	// fmt.Println("decimals1:", config.Token[1].Decimals)
	// fmt.Println("decimals diff:", diff)
	// fmt.Println("decimals pow10diff:", math.Pow10(int(diff)))

	tick24 := float64(tick.Int64())
	//fmt.Println("tick24 ", tick24)

	powTick := math.Pow(1.0001, tick24)
	Price := powTick * float64(math.Pow10(int(diff)))

	//fmt.Println("decimals diff:", diff)
	//fmt.Println("pow10diff:", float64(math.Pow10(int(diff))))
	//fmt.Println("powTick:", powTick)
	//fmt.Println("Price in:", Price) // true price: e.g. 0.991213   / 3890.99932
	//os.Exit(3)
	return Price
}

func GetTVL() *struct {
	Total0 *big.Int
	Total1 *big.Int
} {

	vautlAddress := config.Network.Vault

	vaultInstance, err := vault.NewApi(common.HexToAddress(vautlAddress), config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	totals, err := vaultInstance.GetTVL(&bind.CallOpts{})

	return &totals
}

func Approve(do int) {

	if do <= 0 {
		return
	}

	fmt.Println("----------------------------------------------")
	fmt.Println(".........Approve.........  ")
	fmt.Println("----------------------------------------------")

	fmt.Println("vaultAddress: ", common.HexToAddress(config.Network.Vault))
	fmt.Println("fromAddress: ", config.FromAddress)
	fmt.Println("pooladdress: ", config.Network.Pool)

	/*	instance, err := vault.NewApi(vaultAddress, config.Client)
		if err != nil {
			log.Fatal("vault.NewApi ", err)
		}
	*/

	poolInstance, err := pool.NewApi(common.HexToAddress(config.Network.Pool), config.Client)
	if err != nil {
		log.Fatal("pool.NewApi ", err)
	}

	TokenA, _ := poolInstance.Token0(&bind.CallOpts{})

	TokenB, _ := poolInstance.Token1(&bind.CallOpts{})

	fmt.Println("tokenA: ", TokenA)
	fmt.Println("tokenB: ", TokenB)

	token0Instance, err := token.NewApi(TokenA, config.Client)
	if err != nil {
		log.Fatal("token0Instance,", err)
	}

	config.NonceGen()

	var maxToken0 = config.X1E18(9e18) //new(big.Int).SetString("900000000000000000000000000000", 10)
	tx0, err := token0Instance.Approve(config.Auth, common.HexToAddress(config.Network.Vault), maxToken0)

	if err != nil {
		log.Fatal("tx0, ", err)
	}

	fmt.Println("approve token0 tx: ", tx0.Hash().Hex())

	time.Sleep(config.Network.PendingTime * time.Second)
	//	readstring("done... Press any key to continue when ready..........")

	token1Instance, err := token.NewApi(TokenB, config.Client)
	if err != nil {
		log.Fatal("token0Instance,", err)
	}
	config.NonceGen()
	//var amountToken0, _ = new(big.Int).SetString("90000000000000000000000", 10)
	var maxToken1 = config.X1E18(9e18) // new(big.Int).SetString("900000000000000000000000000000", 10)

	tx1, err := token1Instance.Approve(config.Auth, common.HexToAddress(config.Network.Vault), maxToken1)

	if err != nil {
		log.Fatal("tx1, ", err)
	}

	fmt.Println("approve token1 tx: ", tx1.Hash().Hex())

	config.Readstring("Approve sent... wait for pending..next .. ")

}

func AccountInfo() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".........Account Info.........  ")
	fmt.Println("----------------------------------------------")

	for i, _ := range config.Network.PrivateKey {
		getAccountInfo(i)
	}

}

func getAccountInfo(accId int) {

	accountAddress := config.GetAddress(accId)

	fmt.Println("Account  ----", accId)
	fmt.Println("Account address ", accountAddress)

	tokenAInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenA), config.Client)
	if err != nil {
		log.Fatal("tokenAInstance,", err)
	}

	bal, err := tokenAInstance.BalanceOf(&bind.CallOpts{}, accountAddress)
	fmt.Println("tokenA in Wallet ", bal)

	tokenBInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenB), config.Client)
	if err != nil {
		log.Fatal("tokenBInstance,", err)
	}

	bal, err = tokenBInstance.BalanceOf(&bind.CallOpts{}, accountAddress)
	fmt.Println("tokenB in Wallet ", bal)

	///----------- token in vault

	// vaultInstance, err := vault.NewApi(vaultAddress, config.Client)
	// if err != nil {
	// 	log.Fatal("vault.NewApi ", err)
	// }

	/// vault as token
	vaultTokenInstance, err := token.NewApi(common.HexToAddress(config.Network.Vault), config.Client)
	if err != nil {
		log.Fatal("vaultTokenInstance,", err)
	}

	mybal, _ := vaultTokenInstance.BalanceOf(&bind.CallOpts{}, accountAddress)
	fmt.Println("myShares in vault ", mybal)

	totalbal, _ := vaultTokenInstance.TotalSupply(&bind.CallOpts{})

	fmt.Println("totalSupply in vault ", totalbal)

	if totalbal.Cmp(big.NewInt(0)) > 0 {

		fmt.Println("my share / totalSupply ", mybal.Div(mybal, totalbal).Mul(mybal, big.NewInt(100)), "%")
	}

	fmt.Println()

}

func VaultInfo(do int) {

	if do <= 0 {
		return
	}
	fmt.Println("----------------------------------------------")
	fmt.Println(".........Vault Info.........  ")
	fmt.Println("----------------------------------------------")

	fmt.Println("Vault Address:  ", config.Network.Vault)

	vaultInstance := GetVaultInstance()

	///----------- 当前vault 里的 totalSupply
	totalSupply, err := vaultInstance.TotalSupply(&bind.CallOpts{})
	fmt.Println("totalSupply (total shares in vault) :", totalSupply)
	if err != nil {
		log.Fatal("totalsupply ", err)
	}

	poolInstance := GetPoolInstance()

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	tick := slot0.Tick
	//price := getPrice(slot0.SqrtPriceX96, tick)

	qTickLower, err := vaultInstance.CLow(&bind.CallOpts{})
	qTickUpper, err := vaultInstance.CHigh(&bind.CallOpts{})

	fmt.Println("cLow, tick, cHigh  :", qTickLower, tick, qTickUpper)

	liquidity, err := vaultInstance.GetSSLiquidity(&bind.CallOpts{}, qTickLower, qTickUpper)

	fmt.Println("Current liquidity in pool :", liquidity)

	if err != nil {
		log.Fatal("getssliquidity  ", err)
	}

	xy, err := vaultInstance.GetPositionAmounts(&bind.CallOpts{}, qTickLower, qTickUpper)
	fmt.Println("tokenA (x) in pool ", xy.Amount0)
	fmt.Println("tokenB (y) in pool ", xy.Amount1)

	///-----------accumulated fee0 的总数
	accumulateFees0, _ := vaultInstance.AccumulateUniswapFees0(&bind.CallOpts{})
	fmt.Println("accumulateFees0  :", accumulateFees0)
	///-----------accumulated fee1 的总数
	accumulateFees1, _ := vaultInstance.AccumulateUniswapFees1(&bind.CallOpts{})
	fmt.Println("accumulateFees1  :", accumulateFees1)

	totalfee0, _ := vaultInstance.AccruedProtocolFees0(&bind.CallOpts{})
	totalfee1, _ := vaultInstance.AccruedProtocolFees1(&bind.CallOpts{})
	fmt.Println("Protocol fee0:", totalfee0)
	fmt.Println("Protocol fee1:", totalfee1)

	///----------- 分别返回两个toten 的总数
	totals := GetTVL()
	fmt.Printf("TVL token0=%d\n", totals.Total0)
	fmt.Printf("TVL token1=%d\n", totals.Total1)
	// fmt.Println("decimals0:", int(config.Token[0].Decimals))
	// fmt.Println("decimals1:", int(config.Token[1].Decimals))

	///-----------

}

func SetUniswapPortionRatio(ratio uint8) {

	vaultInstance := GetVaultInstance()

	tx, err := vaultInstance.SetUniPortionRatio(config.Auth, ratio)

	if err != nil {
		log.Fatal("vaultInstance. setuniportion err ", err)
	}

	fmt.Println("setuniportionp tx: ", tx.Hash().Hex())

	config.Readstring("tx sent.....  wait for pending..next .. ")

}

func Init() {

	fmt.Println("tokenA:", common.HexToAddress(config.Network.TokenA))
	fmt.Println("tokenB:", common.HexToAddress(config.Network.TokenB))
	fmt.Println("A<B:", config.Network.TokenA < config.Network.TokenB)

	_, config.Token[0].Name, config.Token[0].Symbol, config.Token[0].Decimals, config.Token[0].MaxTotalSupply = GetTokenInstance(config.Network.TokenA)
	_, config.Token[1].Name, config.Token[1].Symbol, config.Token[1].Decimals, config.Token[1].MaxTotalSupply = GetTokenInstance(config.Network.TokenB)

}

func GetPoolInstance() *pool.Api {

	instance, err := pool.NewApi(common.HexToAddress(config.Network.Pool), config.Client)
	if err != nil {
		log.Fatal("poolInstance err:", err)
	}
	return instance

}

func GetVaultInstance() *vault.Api {

	instance, err := vault.NewApi(common.HexToAddress(config.Network.Vault), config.Client)
	if err != nil {
		log.Fatal("vaultInstance err:", err)
	}
	return instance
}

func GetCalleeInstance() *callee.Api {

	instance, err := callee.NewApi(common.HexToAddress(config.Network.Callee), config.Client)
	if err != nil {
		log.Fatal("CalleeInstance err:", err)
	}
	return instance
}

func GetTokenInstance(TokenAddress string) (*token.Api, string, string, uint8, *big.Int) {

	instance, err := token.NewApi(common.HexToAddress(TokenAddress), config.Client)
	if err != nil {
		log.Fatal("get token Instance,", err)
	}

	name, err := instance.Name(&bind.CallOpts{})

	symbol, err := instance.Symbol(&bind.CallOpts{})

	decimals, err := instance.Decimals(&bind.CallOpts{})

	maxsupply, err := instance.TotalSupply(&bind.CallOpts{})

	return instance, name, symbol, decimals, maxsupply
}

func get_liquidity_0(x float64, sa float64, sb float64) float64 {
	return x * sa * sb / (sb - sa)
}

func get_liquidity_1(y float64, sa float64, sb float64) float64 {
	return y / (sb - sa)
}

func get_liquidity(x float64, y float64, sp float64, sa float64, sb float64) float64 {
	var liquidity, liquidity0, liquidity1 float64
	if sp <= sa {
		liquidity = get_liquidity_0(x, sa, sb)
	} else if sp < sb {
		liquidity0 = get_liquidity_0(x, sp, sb)
		liquidity1 = get_liquidity_1(y, sa, sp)
		liquidity = math.Min(liquidity0, liquidity1)
	} else {
		liquidity = get_liquidity_1(y, sa, sb)
	}
	return liquidity
}

func calculate_x(L float64, sp float64, sb float64) float64 {
	return L * (sb - sp) / (sp * sb)
}
func calculate_y(L float64, sp float64, sa float64) float64 {
	return L * (sp - sa)
}

func getTicks(p float64, a float64, b float64, xDecimals float64, yDecimals float64) (float64, float64) {

	//calc tick  p(i) = 1.0001i

	diffDecimals := math.Pow(10, xDecimals-yDecimals)

	// log(p , 1.0001)  ==  log(p)/ log(1.0001)
	tick := math.Log(p/diffDecimals) / math.Log(1.0001)
	tickLower := math.Log(a/diffDecimals) / math.Log(1.0001)
	tickUpper := math.Log(b/diffDecimals) / math.Log(1.0001)

	fmt.Printf("tick={%.f}\n", tick)
	fmt.Printf("tickLower={%.f}\n", tickLower)
	fmt.Printf("tickUpper={%.f}\n", tickUpper)
	fmt.Printf("\n")

	return tickLower, tickUpper

}
