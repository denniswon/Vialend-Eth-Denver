package include

import (
	"context"
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

var prevFees0 = new(big.Int)
var prevFees1 = new(big.Int)

func Deposit(do int, param [3]int64, doRebalance bool) {

	if do <= 0 {
		return
	}

	myPrintln("----------------------------------------------")
	myPrintln(".........Deposit.........  ")
	myPrintln("----------------------------------------------")

	config.ChangeAccount(int(param[2]))

	myPrintln("vaultAddress: ", common.HexToAddress(config.Network.Vault))
	myPrintln("TokenA:", config.Network.TokenA)
	myPrintln("TokenB:", config.Network.TokenB)
	myPrintln("fromAddress to deposit: ", config.FromAddress)

	instance, err := vault.NewApi(common.HexToAddress(config.Network.Vault), config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	tokenAInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenA), config.Client)
	if err != nil {
		log.Fatal("tokenAInstance,", err)
	}

	bal0, err := tokenAInstance.BalanceOf(&bind.CallOpts{}, config.FromAddress)
	myPrintln("tokenA in Wallet ", bal0)

	if err != nil {
		log.Fatal("bal0 err ", err)
	}

	tokenBInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenB), config.Client)
	if err != nil {
		log.Fatal("tokenBInstance,", err)
	}

	bal1, err := tokenBInstance.BalanceOf(&bind.CallOpts{}, config.FromAddress)
	myPrintln("tokenB in Wallet ", bal1)

	allowA, _ := tokenAInstance.Allowance(&bind.CallOpts{}, config.FromAddress, common.HexToAddress(config.Network.Vault))
	allowB, _ := tokenBInstance.Allowance(&bind.CallOpts{}, config.FromAddress, common.HexToAddress(config.Network.Vault))

	var maxToken0 = PowX(99999, int(config.Token[0].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)
	var maxToken1 = PowX(99999, int(config.Token[1].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)

	if allowA.Cmp(big.NewInt(0)) <= 0 {

		ApproveToken(common.HexToAddress(config.Network.TokenA), maxToken0, config.Network.Vault)
	}

	if allowB.Cmp(big.NewInt(0)) <= 0 {
		ApproveToken(common.HexToAddress(config.Network.TokenB), maxToken1, config.Network.Vault)
	}

	//  amount0 * 10^decimals
	amount0 := big.NewInt(param[0])
	amount1 := big.NewInt(param[1])

	decimals0 := big.NewInt(int64(config.Token[0].Decimals))
	decimals1 := big.NewInt(int64(config.Token[1].Decimals))
	d0 := new(big.Int).Exp(big.NewInt(10), decimals0, nil)
	d1 := new(big.Int).Exp(big.NewInt(10), decimals1, nil)

	if amount0.Cmp(big.NewInt(-1)) == 0 {
		amount0 = bal0.Div(bal0, d0)
	}
	if amount1.Cmp(big.NewInt(-1)) == 0 {
		amount1 = bal1.Div(bal1, d1)
	}
	amountToken0 := new(big.Int).Mul(amount0, d0)
	amountToken1 := new(big.Int).Mul(amount1, d1)

	// myPrintln("decimals0, decimals1:", decimals0, decimals1)
	// myPrintln("d0, d1:", d0, d1)
	myPrintln("amountToken0 to Vault:", amountToken0)
	myPrintln("amountToken1 to Vault:", amountToken1)

	config.NonceGen()
	tx, err := instance.Deposit(config.Auth,
		amountToken0,
		amountToken1,
		doRebalance,
	)

	if err != nil {
		log.Fatal("deposit err: ", err)
	}

	config.ChangeAccount(config.Account)

	//get the transaction hash
	myPrintln("deposit tx: %s", tx.Hash().Hex())

	//	time.Sleep(config.Network.PendingTime * time.Second)
	//Readstring("deposit sent...wait for pending..next .. ")
	TxConfirm(tx.Hash())

}

func Withdraw(do int, param [2]int64) {

	if do <= 0 {
		return
	}

	myPrintln("----------------------------------------------")
	myPrintln(".........Withdraw.........  ")
	myPrintln("----------------------------------------------")

	/// set new account Auth
	config.ChangeAccount(int(param[1]))

	recipient := config.FromAddress

	myPrintln("Withdraw to  ", recipient)

	myPrintln("vaultAddress: ", common.HexToAddress(config.Network.Vault))

	percent := big.NewInt(param[0])

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

		myPrintln("myshares  = ", myshares, 18)

		///calc required shares to withdraw

		shares := new(big.Int).Mul(myshares, percent)
		shares.Div(shares, big.NewInt(100))
		myPrintln("shares required = ", percent, "% = ", shares)

		/// if required share > available share, set withdraw shares = awailable shares
		if shares.Cmp(myshares) == 1 {
			shares = myshares
		}

		if shares.Cmp(big.NewInt(0)) <= 0 {
			myPrintln("share<=0 ", shares)
			return
		}
	*/
	config.NonceGen()
	tx, err := instance.Withdraw(config.Auth, percent)

	if err != nil {
		log.Fatal("withdraw: ", err)
	}

	/// reset account back
	config.ChangeAccount(config.Account)

	//get the transaction hash
	myPrintln("withdraw tx: ", tx.Hash().Hex())

	//	Readstring("withdraw sent.... wait for pending..next .. ")
	TxConfirm(tx.Hash())

}

func GetSwapInfo(rangeRatio int64) (amount0 float64, amount1 float64, swapAmount float64, zeroForOne bool) {

	poolInstance := GetPoolInstance()

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	tick := slot0.Tick
	sqrtPriceX96 := slot0.SqrtPriceX96

	//myPrintln("tick: ", tick)
	//myPrintln("sqrtPriceX96: ", sqrtPriceX96)

	Totals := GetTVL()

	myPrintln("total locked token0: ", Totals.Total0)
	myPrintln("total locked token1: ", Totals.Total1)

	_, pf := getPrice(sqrtPriceX96, tick)

	min := pf * (1 - float64(rangeRatio)/100)
	max := pf * (1 + float64(rangeRatio)/100)
	x := BigIntToFloat64(Totals.Total0) / math.Pow10(int(config.Token[0].Decimals))
	y := BigIntToFloat64(Totals.Total1) / math.Pow10(int(config.Token[1].Decimals))

	xDecimals, yDecimals := config.Token[0].Decimals, config.Token[1].Decimals

	//myPrintln("pf, min, max, rangeRatio: ", pf, min, max, rangeRatio)

	//myPrintln("pf,min, max, rangeRatio: ", pf, min, max, rangeRatio)

	a, b := getTicks(pf, min, max, float64(xDecimals), float64(yDecimals))
	tickA := math.Round(a/60) * 60
	tickB := math.Round(b/60) * 60

	//myPrintln("tick a b:", tickA, tickB)

	tickLower = big.NewInt(int64(tickA)) //tickA) //big.NewInt(-1140) //
	tickUpper = big.NewInt(int64(tickB)) //tickB) //big.NewInt(840)   //

	//myPrintln("---abminmax:", pf-float64(rangeRatio)/100, pf, a, b, tickLower, tickUpper)
	//os.Exit(3)
	// myPrintln(pf, min, max, x, y)
	// myPrintln(priceFromsqrtP)
	// myPrintln(BigIntToFloat64(Total0))
	// myPrintln(BigIntToFloat64(Total1))

	//amt0, amt1, swapAmount, zeroForOne := getBestAmounts(pf, min, max, x, y)
	return getBestAmounts(pf, min, max, x, y)
}
func EmergencyBurn() {

	myPrintln("----------------------------------------------")
	myPrintln(".........Emergency withdraw , burn all positions and send fund back to users.........  ")
	myPrintln("----------------------------------------------")

	myPrintln("vaultAddress: ", common.HexToAddress(config.Network.Vault))

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

	myPrintln("emergency tx: ", tx.Hash().Hex())

	//Readstring("emergency sent sent.....  wait for pending..next .. white hacker to withdraw ")
	TxConfirm(tx.Hash())

}

/// param0: fullRangeSize,
/// param1: tickspacing,
/// param2: accId
func Strategy1(fullRrange int64, acc int64) {

	myPrintln("----------------------------------------------")
	fmt.Println(".........Strategy1 .........  ")
	myPrintln("----------------------------------------------")

	myPrintln("vaultAddress: ", common.HexToAddress(config.Network.Vault))

	vaultInstance, err := vault.NewApi(common.HexToAddress(config.Network.Vault), config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	totals := GetTVL()
	if totals.Total0.Add(totals.Total0, totals.Total1).Cmp(big.NewInt(0)) == 0 {
		myPrintln("!!!! Vault is empty !!!! ")
		return
	}
	//init ticklow and tickupp
	//GetSwapInfo(param[0])
	poolInstance := GetPoolInstance()
	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})
	tick := slot0.Tick

	hrange := new(big.Int).Div(big.NewInt(fullRrange), big.NewInt(2))
	tickLower = new(big.Int).Sub(tick, hrange)
	tickUpper = new(big.Int).Add(tick, hrange)

	tickSpacing := config.Network.FeeTier / 50
	myPrintln("tickspacing:", tickSpacing)

	if tickSpacing < 10 {
		log.Fatal("wrong tickSpacing = ", tickSpacing)
	}
	//tickSpacing := param[1]
	tickLower.Div(tickLower, big.NewInt(tickSpacing)).Mul(tickLower, big.NewInt(tickSpacing))
	tickUpper.Div(tickUpper, big.NewInt(tickSpacing)).Mul(tickUpper, big.NewInt(tickSpacing))

	///require governance. redo auth
	config.ChangeAccount(int(acc))

	myPrintln("range size:", fullRrange)
	myPrintln("ticklower, TICK,  tickupper in...", tickLower, tick, tickUpper)
	myPrintln("in range? ", tick.Cmp(tickLower) > 0 && tick.Cmp(tickUpper) < 0)

	// set ticklower and tickupper
	//setRange(param)
	i := 0
	for {
		config.NonceGen()

		tx, err := vaultInstance.Strategy1(config.Auth,
			tickLower,
			tickUpper)

		if err != nil {
			//log.Fatal("strateg1 tx err ", err)
			time.Sleep(2 * time.Second)
		} else {
			myPrintln("strategy1 tx: ", tx.Hash().Hex())
			TxConfirm(tx.Hash())
			break
		}

		if i > 10 {
			log.Fatal("strateg1 tx err ", err)
		}

	}

	///require governance. redo auth
	config.ChangeAccount(config.Account)

	//Readstring("Rebalance by Strategy1 sent....... ")

}

func LendingInfo() {
	myPrintln("----------------------------------------------")
	myPrintln(".........Lending pool info .........  ")
	myPrintln("----------------------------------------------")

	checkCTokenBalance("CUSDC", config.Network.LendingContracts.CUSDC)
	checkCTokenBalance("CETH", config.Network.LendingContracts.CETH)
	checkETHBalance()

	vaultInstance := GetVaultInstance()
	lendingamts, _ := vaultInstance.GetCAmounts(&bind.CallOpts{})

	myPrintln("CToken0, Ctoken1:", lendingamts)

}

func checkETHBalance() *big.Int {

	bal, err := config.Client.BalanceAt(context.Background(), common.HexToAddress(config.Network.Vault), nil)
	if err != nil {
		log.Fatal("eth balance err ", err)
	}

	myPrintln("eth balance: ", bal)

	return (bal)

}

func checkCTokenBalance(tokenName string, cTokenAddress string) *big.Int {

	cInstance := GetCTokenInstance(cTokenAddress)

	bal, err := cInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(config.Network.Vault))

	if err != nil {
		log.Fatal(" err ", err)
	}

	myPrintln(tokenName, " balance: ", bal)

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
func Rebalance(do int, param [3]int64) {

	if do <= 0 {
		return
	}

	// myPrintln("----------------------------------------------")
	// myPrintln(".........Rebalance .........  ")
	// myPrintln("----------------------------------------------")

	// myPrintln("vaultAddress: ", common.HexToAddress(config.Network.Vault))

	// vaultInstance, err := vault.NewApi(common.HexToAddress(config.Network.Vault), config.Client)
	// if err != nil {
	// 	log.Fatal("vault.NewApi ", err)
	// }

	// //Swap(1, param[0])
	// _, _, swapAmount, zeroForOne := GetSwapInfo(param[0])

	// ///require governance. redo auth
	// config.ChangeAccount(int(param[2]))

	// config.NonceGen()

	// myPrintln("ticklower tickupper in...", tickLower, tickUpper)
	// // set ticklower and tickupper
	// //setRange(param)
	// tx, err := vaultInstance.Strategy0(config.Auth,
	// 	tickLower,
	// 	tickUpper,
	// 	Float64ToBigInt(swapAmount),
	// 	zeroForOne)

	// if err != nil {
	// 	log.Fatal("rebalance signature err ", err)
	// }

	// ///require governance. redo auth
	// config.ChangeAccount(config.Account)

	// //get the transaction hash
	// myPrintln("rebal tx: ", tx.Hash().Hex())

	// Readstring("Rebalance sent.....  wait for pending..next .. ")

	// totalfee0, _ := vaultInstance.AccruedProtocolFees0(&bind.CallOpts{})
	// totalfee1, _ := vaultInstance.AccruedProtocolFees1(&bind.CallOpts{})

	// fmt.Printf("AccumulatedFee0:{%.4f}\n", totalfee0)
	// fmt.Printf("AccumulatedFee1:{%.4f}\n", totalfee1)

}

/// param0 : fullRangeSize, param1: tickspacing
func Swap(do int, rangeRatio int64, account int64) {

	if do <= 0 {
		return
	}

	myPrintln("----------------------------------------------")
	myPrintln(".........Swap .........  ")
	myPrintln("----------------------------------------------")

	//myPrintln("vaultAddress: ", common.HexToAddress(config.Network.Vault))

	//poolInstance := GetPoolInstance()
	vaultInstance := GetVaultInstance()

	amt0, amt1, swapAmount, zeroForOne := GetSwapInfo(rangeRatio)
	//myPrintln("amount0 , amount1 desired:", amt0, amt1)
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
	config.ChangeAccount(int(account))

	config.NonceGen()

	TokenDecimals := config.Token[0].Decimals

	if !zeroForOne {
		TokenDecimals = config.Token[1].Decimals
	}

	// swapamount * 1e(decimals)  i.e. usdc 1e6 = 10^6
	swapAmountIn := new(big.Int).SetInt64(int64(float64(swapAmount * math.Pow10(int(TokenDecimals)))))

	//myPrintln("zeroForOne:", zeroForOne)
	//myPrintln("swap amount: ", swapAmount)
	//myPrintln("swap amount in:", swapAmountIn)
	//_ = swapAmount

	// cTotal0, err := vaultInstance.GetBalance0(&bind.CallOpts{})
	// cTotal1, err := vaultInstance.GetBalance1(&bind.CallOpts{})
	// myPrintln("current balance0 in vault: ", cTotal0)
	// myPrintln("current balance1 in vault: ", cTotal1)

	tx, err := vaultInstance.Swap(config.Auth,
		swapAmountIn,
		zeroForOne,
		sqrtPriceLimitX96,
	)

	if err != nil {
		log.Fatal("vaultInstance.swap err ", err)
	}

	///require governance. redo auth
	config.ChangeAccount(config.Account)

	//get the transaction hash
	//myPrintln("swap tx: ", tx.Hash().Hex())

	//Readstring("swap tx sent.....  wait for pending..next .. ")
	TxConfirm(tx.Hash())

	// Total0, err := vaultInstance.GetBalance0(&bind.CallOpts{})
	// Total1, err := vaultInstance.GetBalance1(&bind.CallOpts{})
	// myPrintln("new balance0 in vault: ", Total0)
	// myPrintln("new balance1 in vault: ", Total1)

	// myPrintln("swapped amt0: ", Total0.Sub(Total0, cTotal0))
	// myPrintln("swapped amt1: ", Total1.Sub(Total1, cTotal1))

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
	myPrintln(x1r, y1r)

	xr := x / (x + y/p)
	yr := y / (y + x*p)
	myPrintln(xr, yr)
	// 20/2000, 0.9908
	// 20 * 0.9908
	if x*p > y { // trade x for y
		zeroForOne = true

		r := xr - x1r

		swapAmount = math.Abs(x * r)

		amount0 = x - swapAmount

		amount1 = y + swapAmount*p

		//myPrintln("newX=", amount0)
		//myPrintln("newY=", amount1)

	} else if x*p < y { // trade y for x
		zeroForOne = false

		r := xr - x1r

		swapAmount = math.Abs(y * r)

		amount0 = x + swapAmount/p

		amount1 = y - swapAmount

		//myPrintln("newX=", amount0)
		//myPrintln("newY=", amount1)
	}

	//fmt.Printf("newX={%.18f}, newY={%.6f},swapamount={%.18f},zeroForOne={%t}\n", amount0, amount1, swapAmount, zeroForOne)

	return amount0, amount1, swapAmount, zeroForOne
}

/// protocol fees, my earned value, APY
func CheckFees(xPrice *big.Int) {

	myPrintln("----------------------------------------------")
	myPrintln(".........Check Fees .........  ")
	myPrintln("----------------------------------------------")

	myPrintln("vaultAddress: ", common.HexToAddress(config.Network.Vault))

	vaultInstance := GetVaultInstance()

	Fees, _ := vaultInstance.Fees(&bind.CallOpts{})

	myPrintln("U3Fees0, U3Fees1, LcFees0, LcFees1, AccruedProtocla Fees0,1 : {", Fees, "}")

	totalFees0 := new(big.Int).Add(Fees.U3Fees0, Fees.LcFees0)
	totalFees1 := new(big.Int).Add(Fees.U3Fees1, Fees.LcFees1)
	fmt.Println("totalFees0, totalFees1, prev0, prev1, diff0, diff1: {",
		totalFees0, totalFees1, "}",
		"{", prevFees0, prevFees1, "}",
		"{", new(big.Int).Sub(totalFees0, prevFees0), new(big.Int).Sub(totalFees1, prevFees1), "}")

	prevFees0 = totalFees0
	prevFees1 = totalFees1

	// accumulative protocol fees
	myPrintln("accruedProtocolFees0, accruedProtocolFees1 : {", Fees.AccruedProtocolFees0, Fees.AccruedProtocolFees1, "}")

	// // check team share
	// Team, _ := vaultInstance.Team(&bind.CallOpts{})
	// tokenIns, _, _, _, _ := GetTokenInstance(config.Network.Vault)
	// teamShare, _ := tokenIns.BalanceOf(&bind.CallOpts{}, Team)
	// totalShare, _ := tokenIns.TotalSupply(&bind.CallOpts{})
	// myPrintln("Team address: ", Team)

	myPrintln()

	for j, _ := range config.Network.PrivateKey {

		storedAccount, _ := vaultInstance.Accounts(&bind.CallOpts{}, big.NewInt(int64(j)))

		if storedAccount.String() != "0x0000000000000000000000000000000000000000" {

			fmt.Println("\n*My address:", storedAccount)

			myAddress := common.HexToAddress(storedAccount.String())

			myshare, totalshare := CalcShares(myAddress)

			if myshare.Cmp(big.NewInt(0)) == 0 {
				continue
			}

			Assets, _ := vaultInstance.Assetholder(&bind.CallOpts{}, storedAccount)
			myPrintln("*Assetsholder: ", Assets)

			if totalshare.Cmp(big.NewInt(0)) > 0 && myshare.Cmp(big.NewInt(0)) > 0 {

				//my earned value
				myFees0 := new(big.Int).Mul(totalFees0, myshare)
				myFees0.Div(myFees0, totalshare)

				myFees1 := new(big.Int).Mul(totalFees1, myshare)
				myFees1.Div(myFees1, totalshare)

				fmt.Println("my Fees0: {", myFees0, "}")
				fmt.Println("my Fees1: {", myFees1, "}")

				myFeesInToken1 := new(big.Int).Mul(myFees0, xPrice)
				myFeesInToken1 = myFeesInToken1.Add(myFeesInToken1, myFees1)

				ListAccounts, _ := vaultInstance.Assetholder(&bind.CallOpts{}, storedAccount)
				myPrintln("*Assetsholder: ", ListAccounts)

				// calc APY  below
				blockNumber := Assets.Block

				block, err := config.Client.BlockByNumber(context.Background(), blockNumber)
				if err != nil {
					log.Fatal("block ", err)
				}

				// get block info

				timestamp := block.Time()

				myPrintln("deposit block info:", blockNumber, block.Time()) // 1527211625

				header, err := config.Client.HeaderByNumber(context.Background(), nil)
				if err != nil {
					log.Fatal("block header ", err)
				}

				headerblock, err := config.Client.BlockByNumber(context.Background(), header.Number)
				if err != nil {
					log.Fatal("block ", err)
				}

				htimestamp := headerblock.Time()

				timediff := htimestamp - timestamp // diff in seconds

				myPrintln("timediff from now: {", timediff, htimestamp, timestamp, "}")

				fmt.Println("Period (Days):", timediff/60/60/24, "{sec:}", timediff)

				oneyearINsec := big.NewInt(365 * 24 * 60 * 60)
				myPrintln("oneyearINsec ", oneyearINsec)

				deposit0 := Assets.Deposit0
				deposit1 := Assets.Deposit1

				totals := GetTVL()

				myPrintln("totalTVL ", totals)

				//my value locked: mytvl0, mytvl1
				mytvl0 := new(big.Int).Mul(totals.Total0, myshare)
				mytvl0 = new(big.Int).Div(mytvl0, totalshare)

				mytvl1 := new(big.Int).Mul(totals.Total1, myshare)
				mytvl1 = new(big.Int).Div(mytvl1, totalshare)

				fmt.Println("deposit0,deposit1 {", deposit0, deposit1, "}")
				fmt.Println("mytvl0, mytvl1 {", mytvl0, mytvl1, "}")

				fd0 := BigIntToFloat64(deposit0)

				fd1 := BigIntToFloat64(deposit1) * 1e18 / BigIntToFloat64(xPrice)

				fm0 := BigIntToFloat64(mytvl0)

				fm1 := BigIntToFloat64(mytvl1) * 1e18 / BigIntToFloat64(xPrice)

				myPrintln("fm0:", fm0)
				myPrintln("fm1:", fm1)

				fdd := fd0 + fd1
				fmm := fm0 + fm1

				myPrintln("fdd, fmm:", fdd, fmm)

				APY := (fmm - fdd) / float64(timediff) * BigIntToFloat64(oneyearINsec) / fdd

				fmt.Println("APY:", APY)

				myDepositInToken1 := new(big.Int).Mul(deposit0, xPrice)
				myDepositInToken1 = myDepositInToken1.Add(myDepositInToken1, deposit1)

				fAPY := BigIntToFloat64(myFeesInToken1) / float64(timediff) * BigIntToFloat64(oneyearINsec) / BigIntToFloat64(myDepositInToken1) * 100
				fmt.Println("APY by fees/Deposit ", fAPY, "%")

				// timediff, myshare, totalshare, tvl0, tvl1, deposit0, deposit1, usPrice0, usPrice1
				// mytvl0 = tvl0 * myshare/totalshare
				// mytvl1 = tvl1 * myshare/totalshare
				//APY =  ( (mytvl0 - deposit0) + (mytvl1 -deposit1) ) / blocktimediff * oneyearINsec

				//			myPrintln(block.Difficulty().Uint64())       // 3217000136609065
				//			myPrintln("block hash:", block.Hash().Hex()) // 0x9e8751ebb5069389b855bba72d949
				//			blockHashHex := block.Hash().Hex()

			}

		}

		myPrintln()

	}

}

/// formula:
///
// 1. price = pow(1.0001,tick) * (1e(18-6)
///2.  (sqrtPricex96^2 * 1e(decimal0)/1e(decimal1) >> (96*2)
///3.  (sqrtPricex96^2 * 1e(decimal0)/1e(decimal1) / 2^(96*2)
// 4. javascript: JSBI.BigInt(sqrtPriceX96 *sqrtPriceX96* (1e(decimals_token_0))/(1e(decimals_token_1))/JSBI.BigInt(2) ** (JSBI.BigInt(192));
///5 solc: uint(sqrtPriceX96).mul(uint(sqrtPriceX96)).mul(1e(decimalsDiff)) >> (96 * 2);
func getPrice(SqrtPriceX96 *big.Int, tick *big.Int) (*big.Int, float64) {

	var diff uint8
	if config.Token[0].Decimals > config.Token[1].Decimals {
		diff = config.Token[0].Decimals - config.Token[1].Decimals
	} else {
		diff = config.Token[1].Decimals - config.Token[0].Decimals
	}

	myPrintln("decimals0:", config.Token[0].Decimals)
	myPrintln("decimals1:", config.Token[1].Decimals)
	myPrintln("decimals diff:", diff)
	myPrintln("decimals pow10diff:", math.Pow10(int(diff)))

	tick24 := float64(tick.Int64())
	//myPrintln("tick24 ", tick24)

	powTick := math.Pow(1.0001, tick24)
	Price := powTick * float64(math.Pow10(int(diff)))

	PriceBigInt := Float64ToBigInt(Price * math.Pow10(int(config.Token[1].Decimals)))

	//myPrintln("decimals diff:", diff)
	//myPrintln("pow10diff:", float64(math.Pow10(int(diff))))
	//myPrintln("powTick:", powTick)
	//myPrintln("Price in:", Price) // true price: e.g. 0.991213   / 3890.99932
	//os.Exit(3)
	return PriceBigInt, Price
}

func Alloc(accId int) {

	myPrintln("----------------------------------------------")
	myPrintln(".........alloc .........  ")
	myPrintln("----------------------------------------------")

	vaultInstance := GetVaultInstance()

	config.ChangeAccount(accId)
	config.NonceGen()

	tx, err := vaultInstance.Alloc(config.Auth)

	if err != nil {
		log.Fatal("alloc err: ", err)
	}

	myPrintln("alloc tx: %s", tx.Hash().Hex())

	config.ChangeAccount(config.Account)

	//Readstring("alloc sent...wait for pending..next .. ")
	TxConfirm(tx.Hash())

}

func GetTVL() *struct {
	Total0 *big.Int
	Total1 *big.Int
} {

	Totals := new(struct {
		Total0 *big.Int
		Total1 *big.Int
	})

	vaultInstance := GetVaultInstance()
	token0Ins, _, _, _, _ := GetTokenInstance(config.Network.TokenA)
	token1Ins, _, _, _, _ := GetTokenInstance(config.Network.TokenB)

	//implement gettvl
	cHigh, _ := vaultInstance.CHigh(&bind.CallOpts{})
	cLow, _ := vaultInstance.CLow(&bind.CallOpts{})

	uniliqs, _ := vaultInstance.GetPositionAmounts(&bind.CallOpts{}, cLow, cHigh)
	myPrintln("balance in uniswap:  ", uniliqs)

	lendingAmt0, lendingAmt1, exrate0, exrate1 := GetLendingAmounts()
	//lendingAmt0, lendingAmt1, exrate0, exrate1, _ := vaultInstance.GetLendingAmounts(&bind.CallOpts{})
	myPrintln("balance in lending: ", lendingAmt0, lendingAmt1)
	myPrintln("exchange rate: ", exrate0, exrate1)

	clending0, clending1 := vaultInstance.GetCAmounts(&bind.CallOpts{})
	myPrintln("C Amounts in lending: ", clending0, clending1)

	balance0, _ := token0Ins.BalanceOf(&bind.CallOpts{}, common.HexToAddress(config.Network.Vault))
	balance1, _ := token1Ins.BalanceOf(&bind.CallOpts{}, common.HexToAddress(config.Network.Vault))
	myPrintln("balance in vault: ", balance0, balance1)

	Totals.Total0 = balance0.Add(balance0, uniliqs.Amount0).Add(balance0, lendingAmt0)
	Totals.Total1 = balance1.Add(balance1, uniliqs.Amount1).Add(balance1, lendingAmt1)

	return Totals
}

func GetLendingAmounts() (*big.Int, *big.Int, *big.Int, *big.Int) {

	cInstance0 := GetCTokenInstance(config.Network.CTOKEN0)
	cInstance1 := GetCTokenInstance(config.Network.CTOKEN1)

	//implement gettvl
	exchangeRate0, _ := cInstance0.ExchangeRateStored(&bind.CallOpts{})
	exchangeRate1, _ := cInstance1.ExchangeRateStored(&bind.CallOpts{})

	CAmount0 := checkCTokenBalance("CToken0", config.Network.CTOKEN0)
	CAmount1 := checkCTokenBalance("CToken1", config.Network.CTOKEN1)

	pow1018 := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	//	pow106 := new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)
	underlying0 := CAmount0.Mul(CAmount0, exchangeRate0).Div(CAmount0, pow1018) //.Div(CAmount0, pow1018)
	underlying1 := CAmount1.Mul(CAmount1, exchangeRate1).Div(CAmount1, pow1018) //.Div(CAmount1, pow106)

	return underlying0, underlying1, exchangeRate0, exchangeRate1

}

func ApproveToken(tokenAddress common.Address, maxAmount *big.Int, toAddress string) {

	tokenInstance, err := token.NewApi(tokenAddress, config.Client)
	if err != nil {
		log.Fatal("tokenInstance,", err)
	}

	//check allowance
	allow, _ := tokenInstance.Allowance(&bind.CallOpts{}, config.FromAddress, common.HexToAddress(toAddress))

	if allow.Cmp(big.NewInt(0)) > 0 {
		return
	}

	myPrintln("Address to approve: ", toAddress)
	myPrintln("fromAddress: ", config.FromAddress)
	//myPrintln("pooladdress: ", config.Network.Pool)

	config.NonceGen()

	tx, err := tokenInstance.Approve(config.Auth, common.HexToAddress(toAddress), maxAmount)

	if err != nil {
		log.Fatal("tx, ", err)
	}

	myPrintln("approve token tx: ", tx.Hash().Hex())

	//TxConfirm(tx.Hash())
	//Readstring("Approve sent... wait for pending..next .. ")
	TxConfirm(tx.Hash())

}

func Approve(account int) {

	if account < 0 {
		return
	}

	myPrintln("----------------------------------------------")
	myPrintln(".........Approve.........  ")
	myPrintln("----------------------------------------------")

	config.ChangeAccount(account)
	config.NonceGen()

	poolInstance := GetPoolInstance()
	TokenA, _ := poolInstance.Token0(&bind.CallOpts{})

	TokenB, _ := poolInstance.Token1(&bind.CallOpts{})

	myPrintln("tokenA: ", TokenA)
	myPrintln("tokenB: ", TokenB)

	var maxToken0 = PowX(99999, int(config.Token[0].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)
	var maxToken1 = PowX(99999, int(config.Token[1].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)

	ApproveToken(TokenA, maxToken0, config.Network.Vault)
	ApproveToken(TokenB, maxToken1, config.Network.Vault)

	config.ChangeAccount(config.Account)

}

func AccountInfo() {

	myPrintln("----------------------------------------------")
	myPrintln(".........Account Info.........  ")
	myPrintln("----------------------------------------------")

	for i, _ := range config.Network.PrivateKey {

		getAccountInfo(i)
	}

}

func getAccountInfo(accId int) {

	accountAddress := config.GetAddress(accId)

	myPrintln("Account  ----", accId)
	myPrintln("Account address ", accountAddress)

	tokenAInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenA), config.Client)
	if err != nil {
		log.Fatal("tokenAInstance,", err)
	}

	bal, err := tokenAInstance.BalanceOf(&bind.CallOpts{}, accountAddress)
	myPrintln("tokenA in Wallet ", bal)

	tokenBInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenB), config.Client)
	if err != nil {
		log.Fatal("tokenBInstance,", err)
	}

	bal, err = tokenBInstance.BalanceOf(&bind.CallOpts{}, accountAddress)
	myPrintln("tokenB in Wallet ", bal)

	///----------- token in vault

	// vaultInstance, err := vault.NewApi(vaultAddress, config.Client)
	// if err != nil {
	// 	log.Fatal("vault.NewApi ", err)
	// }

	mybal, totalbal := CalcShares(accountAddress)

	if totalbal.Cmp(big.NewInt(0)) > 0 {

		myPrintln("my share / totalSupply ", mybal.Mul(mybal, big.NewInt(100)).Div(mybal, totalbal), "%")
	}

	myPrintln()

}

func CalcShares(myAddress common.Address) (mybal *big.Int, totalbal *big.Int) {

	/// vault as token
	vaultTokenInstance, err := token.NewApi(common.HexToAddress(config.Network.Vault), config.Client)
	if err != nil {
		log.Fatal("vaultTokenInstance,", err)
	}

	mybal, _ = vaultTokenInstance.BalanceOf(&bind.CallOpts{}, myAddress)

	if mybal.Cmp(big.NewInt(0)) > 0 {
		myPrintln("myShares in vault ", mybal)

	}

	totalbal, _ = vaultTokenInstance.TotalSupply(&bind.CallOpts{})
	if totalbal.Cmp(big.NewInt(0)) > 0 {
		myPrintln("totalSupply in vault ", totalbal)
	}
	return mybal, totalbal
}

func VaultInfo() {

	myPrintln("----------------------------------------------")
	myPrintln(".........Vault Info.........  ")
	myPrintln("----------------------------------------------")

	myPrintln("Vault Address:  ", config.Network.Vault)

	vaultInstance := GetVaultInstance()

	poolAddress, err := vaultInstance.PoolAddress(&bind.CallOpts{})
	myPrintln("pool address from vault:", poolAddress)
	if err != nil {
		log.Fatal("poolAddress err ", err)
	}

	///----------- 当前vault 里的 totalSupply
	_CToken0, _ := vaultInstance.CToken0(&bind.CallOpts{})
	_CToken1, _ := vaultInstance.CToken1(&bind.CallOpts{})
	myPrintln("Ctoken0 address:", _CToken0)
	myPrintln("Ctoken1 address:", _CToken1)

	totalSupply, err := vaultInstance.TotalSupply(&bind.CallOpts{})

	myPrintln("totalSupply (total shares in vault) :", totalSupply)
	if err != nil {
		log.Fatal("totalsupply ", err)
	}

	poolInstance := GetPoolInstance()

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	tick := slot0.Tick

	qTickLower, err := vaultInstance.CLow(&bind.CallOpts{})
	qTickUpper, err := vaultInstance.CHigh(&bind.CallOpts{})

	myPrintln("cLow, tick, cHigh  :", qTickLower, tick, qTickUpper)

	fmt.Println("** in range? ", tick.Cmp(qTickLower) > 0 && tick.Cmp(qTickUpper) < 0)

	liquidity, err := vaultInstance.GetSSLiquidity(&bind.CallOpts{}, qTickLower, qTickUpper)

	myPrintln("Current liquidity in pool :", liquidity)

	if err != nil {
		log.Fatal("getssliquidity  ", err)
	}

	xy, err := vaultInstance.GetPositionAmounts(&bind.CallOpts{}, qTickLower, qTickUpper)
	myPrintln("tokenA (x) in pool ", xy.Amount0)
	myPrintln("tokenB (y) in pool ", xy.Amount1)

	///----------- 分别返回两个toten 的总数, also print tokens in vault
	totals := GetTVL()

	fmt.Printf("TVL token0=%d\n", totals.Total0)
	fmt.Printf("TVL token1=%d\n", totals.Total1)
	// myPrintln("decimals0:", int(config.Token[0].Decimals))
	// myPrintln("decimals1:", int(config.Token[1].Decimals))

	uniPortion, _ := vaultInstance.UniPortion(&bind.CallOpts{})
	myPrintln("uniPortionRate:", uniPortion)

	protocolFeeRate, _ := vaultInstance.ProtocolFee(&bind.CallOpts{})
	myPrintln("ProtocolFeeRate:", protocolFeeRate)

	var twapDuration = uint32(2)
	twap, _ := vaultInstance.GetTwap(&bind.CallOpts{}, poolAddress, twapDuration)

	myPrintln("twap:", twap, ",slot.tick:", slot0.Tick)

	var oraclePriceTwap *big.Int
	if twap != nil {
		//	twap = big.NewInt(-192874)
		baseAmount := big.NewInt(1e18)
		baseToken := common.HexToAddress(config.Network.TokenA)
		quoteToken := common.HexToAddress(config.Network.TokenB)

		//calleeInstance := GetCalleeInstance()
		// oraclePriceTwap, _ = calleeInstance.GetQuoteAtTick(&bind.CallOpts{}, twap, baseAmount, baseToken, quoteToken)

		oraclePriceTwap, _ = vaultInstance.GetQuoteAtTick(&bind.CallOpts{}, twap, baseAmount, baseToken, quoteToken)

		fmt.Println("oraclePriceTwap:", oraclePriceTwap)
	}

	SpotPricebigInt, pricefloat64 := getPrice(slot0.SqrtPriceX96, slot0.Tick)
	_ = pricefloat64
	fmt.Println("Spot price big:", SpotPricebigInt)
	fmt.Println("Spot price :", pricefloat64)

	if twap == nil {
		oraclePriceTwap = SpotPricebigInt
	}

	//	sqrtPriceX96 := slot0.SqrtPriceX96
	// uniswapPriceBySqrtP, _ := vaultInstance.GetPriceBySQRTP(&bind.CallOpts{}, sqrtPriceX96)
	// myPrintln("GetPriceBySQRTP:", uniswapPriceBySqrtP)

	CheckFees(oraclePriceTwap)

	///-----------

}

func Init() {

	myPrintln("tokenA:", common.HexToAddress(config.Network.TokenA))
	myPrintln("tokenB:", common.HexToAddress(config.Network.TokenB))
	myPrintln("A<B:", config.Network.TokenA < config.Network.TokenB)

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

func TxConfirm(tx common.Hash) {

	myPrintln("tx: ", tx.Hex())

	tr, err := config.Client.TransactionReceipt(context.Background(), tx)
	for i := 0; i < 20; i++ {
		if err != nil {
			//log.Fatal(err)
			time.Sleep(2 * time.Second)
		} else {
			break
		}

		tr, err = config.Client.TransactionReceipt(context.Background(), tx)

	}

	myPrintln("Receipt -> BlockNumber -> :", tr.BlockNumber)

}

func PrintPrice() {

	poolInstance := GetPoolInstance()
	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	_, pf := getPrice(slot0.SqrtPriceX96, slot0.Tick)

	fmt.Println("Price now:", pf)
}
