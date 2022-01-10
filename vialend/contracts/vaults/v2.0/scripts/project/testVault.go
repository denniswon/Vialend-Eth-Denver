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
	token "viaroot/deploy/Tokens/erc20/deploy/Token"

	//vault "viaroot/deploy/FeeMaker"

	arb "viaroot/deploy/arb"
	admin "viaroot/deploy/vaultAdmin"
	bridge "viaroot/deploy/vaultBridge"
	vault "viaroot/deploy/vialendFeeMaker"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// global variables
var tickLower = new(big.Int)
var tickUpper = new(big.Int)

var qTickLower = new(big.Int) // for query only, monitor
var qTickUpper = new(big.Int) // for query only, monitor

var prevFees0 = new(big.Int)
var prevFees1 = new(big.Int)

func Deposit(_amount0 int, _amount1 int, acc int) {

	myPrintln("----------------------------------------------")
	myPrintln(".........Deposit.........  ")
	myPrintln("----------------------------------------------")

	ChangeAccount(acc)

	myPrintln("vaultAddress: ", common.HexToAddress(Network.Vault))
	myPrintln("TokenA:", Network.TokenA)
	myPrintln("TokenB:", Network.TokenB)
	myPrintln("fromAddress to deposit: ", FromAddress)

	// instance, err := vault.NewApi(common.HexToAddress(Network.Vault), Client)
	// if err != nil {
	// 	log.Fatal("vault.NewApi ", err)
	// }

	_, _, instance := GetInstance3()

	tokenAInstance, err := token.NewApi(common.HexToAddress(Network.TokenA), Client)
	if err != nil {
		log.Fatal("tokenAInstance,", err)
	}

	bal0, err := tokenAInstance.BalanceOf(&bind.CallOpts{}, FromAddress)
	myPrintln("tokenA in Wallet ", bal0)

	if err != nil {
		log.Fatal("bal0 err ", err)
	}

	tokenBInstance, err := token.NewApi(common.HexToAddress(Network.TokenB), Client)
	if err != nil {
		log.Fatal("tokenBInstance,", err)
	}

	bal1, err := tokenBInstance.BalanceOf(&bind.CallOpts{}, FromAddress)
	myPrintln("tokenB in Wallet ", bal1)
	var maxToken0 = PowX(99999, int(Token[0].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)
	var maxToken1 = PowX(99999, int(Token[1].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)

	ApproveToken(common.HexToAddress(Network.TokenA), maxToken0, Network.Vault)
	ApproveToken(common.HexToAddress(Network.TokenB), maxToken1, Network.Vault)

	//  amount0 * 10^decimals
	amount0 := big.NewInt(int64(_amount0))
	amount1 := big.NewInt(int64(_amount1))

	// not multi 10^decimals
	amountToken0 := amount0
	amountToken1 := amount1

	myPrintln("amountToken0 to Vault:", amountToken0)
	myPrintln("amountToken1 to Vault:", amountToken1)

	NonceGen()
	tx, err := instance.Deposit(Auth,
		amountToken0,
		amountToken1,
	)

	if err != nil {
		log.Fatal("deposit err: ", err)
	}

	ChangeAccount(Account)

	//get the transaction hash
	myPrintln("deposit tx: %s", tx.Hash().Hex())

	//	time.Sleep(Network.PendingTime * time.Second)
	//Readstring("deposit sent...wait for pending..next .. ")
	TxConfirm(tx.Hash())

}

func WithdrawPending(pct int64, acc int64) {

	myPrintln("----------------------------------------------")
	myPrintln(".........Withdraw pending with percentage.........  ")
	myPrintln("----------------------------------------------")

	ChangeAccount(int(acc))

	NonceGen()
	vaultInstance := GetVaultInstance()
	tx, err := vaultInstance.WithdrawPending(Auth, uint8(pct))

	if err != nil {
		log.Fatal("withdraw pending: ", err)
	}

	/// reset account back
	ChangeAccount(Account)

	//	Readstring("withdraw sent.... wait for pending..next .. ")
	TxConfirm(tx.Hash())

}

func Withdraw(_percent int, acc int) {

	myPrintln("----------------------------------------------")
	myPrintln(".........Withdraw.........  ")
	myPrintln("----------------------------------------------")

	/// set new account Auth
	ChangeAccount(acc)

	recipient := FromAddress

	myPrintln("Withdraw to  ", recipient)

	myPrintln("vaultAddress: ", common.HexToAddress(Network.Vault))

	percent := uint8(_percent)

	// instance, err := vault.NewApi(common.HexToAddress(Network.Vault), Client)
	// if err != nil {
	// 	log.Fatal("vault.NewApi ", err)
	// }
	_, _, instance := GetInstance3()

	///get account's available shares
	myshares, err := instance.BalanceOf(&bind.CallOpts{}, recipient)
	if err != nil {
		log.Fatal("balance of myshare ", err)
	}

	if myshares.Cmp(big.NewInt(0)) == 0 {
		myPrintln("share==0 ", myshares)
		return
	}

	NonceGen()
	tx, err := instance.Withdraw(Auth, percent)

	if err != nil {
		log.Fatal("withdraw: ", err)
	}

	/// reset account back
	ChangeAccount(Account)

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
	x := BigIntToFloat64(Totals.Total0) / math.Pow10(int(Token[0].Decimals))
	y := BigIntToFloat64(Totals.Total1) / math.Pow10(int(Token[1].Decimals))

	xDecimals, yDecimals := Token[0].Decimals, Token[1].Decimals

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

func Sweep(vaultAddr string, tokenAddr string, amount *big.Int) {
	myPrintln("----------------------------------------------")
	myPrintln(".........sweep.........  ")
	myPrintln("----------------------------------------------")

	vaultInstance := GetVaultInstance2(vaultAddr)

	myPrintln("vaultAddress: ", vaultAddr)

	tx, err := vaultInstance.Sweep(Auth, common.HexToAddress(tokenAddr), amount)

	if err != nil {
		log.Fatal("sweep err ", err)
	}

	TxConfirm(tx.Hash())

}

func EmergencyBurn() {

	myPrintln("----------------------------------------------")
	myPrintln(".........Emergency pull , burn all positions and send fund back to vault.........  ")
	myPrintln("----------------------------------------------")

	myPrintln("vaultAddress: ", common.HexToAddress(Network.Vault))

	_, _, vaultInstance := GetInstance3()

	// tickLower, err := vaultInstance.CLow(&bind.CallOpts{})
	// tickUpper, err := vaultInstance.CHigh(&bind.CallOpts{})
	// liquidity, err := vaultInstance.GetSSLiquidity(&bind.CallOpts{}, qTickLower, qTickUpper)

	// lendingAmount0 := checkCTokenBalance("CETH", Network.LendingContracts.CETH)
	// lendingAmount1 := checkCTokenBalance("CUSDC", Network.LendingContracts.CUSDC)

	tx, err := vaultInstance.EmergencyBurn(Auth)

	if err != nil {
		log.Fatal("emergency tx err ", err)
	}

	myPrintln("emergency tx: ", tx.Hash().Hex())

	//Readstring("emergency sent sent.....  wait for pending..next .. white hacker to withdraw ")
	TxConfirm(tx.Hash())

}

func EmergencyWithdraw(acc int) {
	myPrintln("----------------------------------------------")
	myPrintln(".........Emergency withdraw .........  ")
	myPrintln("----------------------------------------------")
	_, _, vaultInstance := GetInstance3()

	ChangeAccount(acc)

	tx, _ := vaultInstance.EmergencyWithdraw(Auth)

	ChangeAccount(Account)

	TxConfirm(tx.Hash())

}

/// param0: fullRangeSize,
/// param1: tickspacing,
/// param2: accId
func Strategy1(_range int64, acc int64) {

	myPrintln("----------------------------------------------")
	fmt.Println(".........Strategy1 .........  ")
	myPrintln("----------------------------------------------")

	myPrintln("vaultAddress: ", common.HexToAddress(Network.Vault))

	vaultInstance, err := vault.NewApi(common.HexToAddress(Network.Vault), Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	//init ticklow and tickupp
	//GetSwapInfo(param[0])
	poolInstance := GetPoolInstance()
	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})
	tick := slot0.Tick

	hrange := new(big.Int).Div(big.NewInt(_range), big.NewInt(2))
	tickLower = new(big.Int).Sub(tick, hrange)
	tickUpper = new(big.Int).Add(tick, hrange)

	tickSpacing := Network.FeeTier / 50 // ie 3000/50= 60, 500/50 = 10, 10000/50 = 200
	myPrintln("tickspacing:", tickSpacing)

	if tickSpacing < 10 {
		log.Fatal("wrong tickSpacing = ", tickSpacing)
	}
	//tickSpacing := param[1]
	tickLower.Div(tickLower, big.NewInt(tickSpacing)).Mul(tickLower, big.NewInt(tickSpacing))
	tickUpper.Div(tickUpper, big.NewInt(tickSpacing)).Mul(tickUpper, big.NewInt(tickSpacing))

	///require governance. redo auth
	ChangeAccount(int(acc))

	myPrintln("range size:", _range)
	myPrintln("ticklower, TICK,  tickupper in...", tickLower, tick, tickUpper)
	myPrintln("in range? ", tick.Cmp(tickLower) > 0 && tick.Cmp(tickUpper) < 0)

	// set ticklower and tickupper
	//setRange(param)
	i := 0
	for {
		NonceGen()

		tx, err := vaultInstance.Strategy1(Auth,
			tickLower,
			tickUpper)

		if err != nil {
			fmt.Println("strateg1 tx err , trying one more time in 2 seconds..", err)
			time.Sleep(2 * time.Second)
		} else {
			//myPrintln("strategy1 tx: ", tx.Hash().Hex())
			TxConfirm(tx.Hash())
			break
		}

		if i > 10 {
			log.Fatal("strateg1 tx err ", err)
		}

	}

	///require governance. redo auth
	ChangeAccount(Account)

	fmt.Println()

	//Readstring("Rebalance by Strategy1 sent....... ")

}

func LendingInfo() {
	myPrintln("----------------------------------------------")
	myPrintln(".........Lending pool info .........  ")
	myPrintln("----------------------------------------------")

	myPrintln(".........Compound info .........  ")
	checkCTokenBalance(Network.Vault, "CUSDC", Network.LendingContracts.CUSDC)
	checkCTokenBalance(Network.Vault, "CETH", Network.LendingContracts.CETH)
	checkETHBalance()

	_, stratInstance, _ := GetInstance3()
	CAmounts, _ := stratInstance.GetCAmounts(&bind.CallOpts{})
	myPrintln("CToken0, Ctoken1:", CAmounts)

	exchangeRateStored, _ := GetCTokenInstance(Network.CTOKEN0).ExchangeRateStored(&bind.CallOpts{})

	//dem := int(18 + int(Token[0].Decimals) - 8)
	ctoken0Underlying := BigIntToFloat64(CAmounts.Amount0) * (BigIntToFloat64(exchangeRateStored) / 1e18) /// 1 * BigIntToFloat64(PowX(10, dem)))
	myPrintln("exchangeRateStored0:", exchangeRateStored)
	myPrintln("ctoken0Underlying", ctoken0Underlying)

	exchangeRateStored, _ = GetCTokenInstance(Network.CTOKEN1).ExchangeRateStored(&bind.CallOpts{})
	//dem = int(18 + int(Token[1].Decimals) - 8)
	ctoken1Underlying := BigIntToFloat64(CAmounts.Amount1) * (BigIntToFloat64(exchangeRateStored) / 1e18) /// 1 * BigIntToFloat64(PowX(10, dem)))
	myPrintln("exchangeRateStored1:", exchangeRateStored)
	myPrintln("ctoken1Underlying", ctoken1Underlying)

	myPrintln("counter check with GetCompAmounts() from contract:")
	GetCompAmounts()

	myPrintln("wbtc info")
	exchangeRateStored, _ = GetCTokenInstance(Network.LendingContracts.CWBTC).ExchangeRateStored(&bind.CallOpts{})
	cwbtcAmount := float64(141008860)
	cwbtcUnderlying := cwbtcAmount * (BigIntToFloat64(exchangeRateStored) / 1e18) /// 1 * BigIntToFloat64(PowX(10, dem)))
	myPrintln("exchangeRateStored0:", exchangeRateStored)
	myPrintln(cwbtcAmount, "CWBTC =  ", cwbtcUnderlying, "WBTC")

}

func checkETHBalance() *big.Int {

	bal := EthBalance(Network.Vault)

	myPrintln("eth balance: ", bal)

	return (bal)

}

func checkCTokenBalance(who string, tokenName string, cTokenAddress string) *big.Int {

	cInstance := GetCTokenInstance(cTokenAddress)

	bal, err := cInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(who))

	if err != nil {
		log.Fatal(" err ", err)
	}

	//	myPrintln(tokenName, " balance: ", bal)

	return (bal)

}

/// param0 : fullRangeSize, param1: account
func Rebalance(_range int, acc int) {

	myPrintln("----------------------------------------------")
	myPrintln(".........Rebalance New.........  ")
	myPrintln("----------------------------------------------")

	_, stratInstance, _ := GetInstance3()

	//init ticklow and tickupp
	//GetSwapInfo(param[0])
	poolInstance := GetPoolInstance()
	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})
	tick := slot0.Tick

	hrange := new(big.Int).Div(big.NewInt(int64(_range)), big.NewInt(2))
	tickLower = new(big.Int).Sub(tick, hrange)
	tickUpper = new(big.Int).Add(tick, hrange)

	tickSpacing := Network.FeeTier / 50 // ie 3000/50= 60, 500/50 = 10, 10000/50 = 200
	myPrintln("tickspacing:", tickSpacing)

	if tickSpacing < 10 {
		log.Fatal("wrong tickSpacing = ", tickSpacing)
	}
	//tickSpacing := param[1]
	tickLower.Div(tickLower, big.NewInt(tickSpacing)).Mul(tickLower, big.NewInt(tickSpacing))
	tickUpper.Div(tickUpper, big.NewInt(tickSpacing)).Mul(tickUpper, big.NewInt(tickSpacing))

	///require governance. redo auth
	ChangeAccount(int(acc))

	myPrintln("range size:", _range)
	myPrintln("ticklower, TICK,  tickupper in...", tickLower, tick, tickUpper)
	myPrintln("in range? ", tick.Cmp(tickLower) > 0 && tick.Cmp(tickUpper) < 0)

	// set ticklower and tickupper
	//setRange(param)
	NonceGen()

	tx, err := stratInstance.Rebalance(Auth,
		tickLower,
		tickUpper)

	if err != nil {
		log.Fatal("Rebalance err:", err)
	}

	TxConfirm(tx.Hash())

	///require governance. redo auth
	ChangeAccount(Account)

	fmt.Println()

	//Readstring("Rebalance by Strategy1 sent....... ")

}

func Stat2Str(stat uint64) string {
	if stat == 0 {
		return ("FAIL")
	} else {
		return ("SUCCESS")
	}

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

func OraclePrice() (twapPrice *big.Int, spotPrice *big.Int) {

	vaultInstance := GetVaultInstance()

	var twapDuration = uint32(2)

	poolAddress := Network.Pool

	twap, _ := vaultInstance.GetTwap(&bind.CallOpts{}, common.HexToAddress(poolAddress), twapDuration)

	Sleep(100)

	poolInstance := GetPoolInstance()

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})
	Sleep(100)
	//	tick := slot0.Tick

	myPrintln(",slot.tick:", slot0.Tick)
	myPrintln(",slot0.SqrtPriceX96:", slot0.SqrtPriceX96)

	spotPrice, pricefloat64 := getPrice(slot0.SqrtPriceX96, slot0.Tick)

	fmt.Println("Spot price bigint:", spotPrice)
	fmt.Println("Spot price :", pricefloat64)

	if twap != nil {
		//	twap = big.NewInt(-192874)
		baseAmount := big.NewInt(1e8)
		baseToken := common.HexToAddress(Network.TokenA)
		quoteToken := common.HexToAddress(Network.TokenB)

		//myPrintln(" twap, baseAmount, baseToken, quoteToken", twap, baseAmount, baseToken, quoteToken)

		calleeInstance := GetCalleeInstance()

		twapQuatebyCallee, _ := calleeInstance.GetQuoteAtTick(&bind.CallOpts{}, twap, baseAmount, baseToken, quoteToken)
		Sleep(200)

		twapPrice, _ = vaultInstance.GetQuoteAtTick(&bind.CallOpts{}, twap, baseAmount, baseToken, quoteToken)
		Sleep(200)

		fmt.Println("twap Price:", twapPrice)
		fmt.Println("twapQuatebyCallee:", twapQuatebyCallee)
	}

	myPrintln("twap:", twap)

	if twap == nil {
		twapPrice = spotPrice
	}

	return
}

/// protocol fees, my earned value, APY
func CheckFees() {

	myPrintln("----------------------------------------------")
	myPrintln(".........Check Fees .........  ")
	myPrintln("----------------------------------------------")

	myPrintln("vaultAddress: ", common.HexToAddress(Network.Vault))

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
	// tokenIns, _, _, _, _ := GetTokenInstance(Network.Vault)
	// teamShare, _ := tokenIns.BalanceOf(&bind.CallOpts{}, Team)
	// totalShare, _ := tokenIns.TotalSupply(&bind.CallOpts{})
	// myPrintln("Team address: ", Team)

	myPrintln()

	return

	_, xPrice := OraclePrice()

	// crashed on forked mainnet
	for j, _ := range Network.PrivateKey {
		Sleep(2000)
		storedAccount, _ := vaultInstance.Accounts(&bind.CallOpts{}, big.NewInt(int64(j)))
		Sleep(1000)
		if storedAccount.String() != "0x0000000000000000000000000000000000000000" {

			fmt.Println("\n*My address:", storedAccount)

			myAddress := common.HexToAddress(storedAccount.String())

			myshare, totalshare := CalcShares(myAddress)

			if myshare.Cmp(big.NewInt(0)) == 0 {
				continue
			}

			Assets, _ := vaultInstance.Assetholder(&bind.CallOpts{}, storedAccount)
			Sleep(1000)
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
				Sleep(1000)
				myPrintln("*Assetsholder: ", ListAccounts)

				// calc APY  below
				blockNumber := Assets.Block

				block, err := Client.BlockByNumber(context.Background(), blockNumber)

				if err != nil {
					log.Fatal("block ", err)
				}
				Sleep(1000)

				// get block info

				timestamp := block.Time()

				myPrintln("deposit block info:", blockNumber, block.Time()) // 1527211625

				header, err := Client.HeaderByNumber(context.Background(), nil)
				if err != nil {
					log.Fatal("block header ", err)
				}
				Sleep(1000)

				headerblock, err := Client.BlockByNumber(context.Background(), header.Number)
				if err != nil {
					log.Fatal("block ", err)
				}
				Sleep(1000)

				htimestamp := headerblock.Time()

				timediff := htimestamp - timestamp // diff in seconds

				myPrintln("timediff from now: {", timediff, htimestamp, timestamp, "}")

				fmt.Println("Period (Days):", timediff/60/60/24, "{sec:}", timediff)

				oneyearINsec := big.NewInt(365 * 24 * 60 * 60)
				myPrintln("oneyearINsec ", oneyearINsec)

				deposit0 := Assets.Deposit0
				deposit1 := Assets.Deposit1

				totals := GetTVL()
				Sleep(1000)

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

	myPrintln("decimals0:", Token[0].Decimals)
	myPrintln("decimals1:", Token[1].Decimals)

	decimalDiff := int(Token[0].Decimals) - int(Token[1].Decimals)
	myPrintln("decimals diff:", decimalDiff)

	tick24 := float64(tick.Int64())
	//myPrintln("tick24 ", tick24)

	powTick := math.Pow(1.0001, tick24)

	tickPrice := powTick * float64(math.Pow10(int(decimalDiff)))
	PriceBigInt := Float64ToBigInt(tickPrice * math.Pow10(int(Token[1].Decimals)))
	myPrintln("pricebigint", PriceBigInt)

	sqrtPf := new(big.Float)
	///convert big Int to big Float
	sqrtPf.SetString(SqrtPriceX96.String())
	///convert big float to float64
	sp64, _ := sqrtPf.Float64()
	// operate float64

	sqrtPx962Price := (sp64 * sp64) * math.Pow10(int(decimalDiff)) / math.Pow(2, 192)
	myPrintln("counter check by sqrtPrice: ", sqrtPx962Price)
	myPrintln("sqrtPriceX96:", sp64)
	myPrintln("pow10(dif)", math.Pow10(int(decimalDiff)))

	//myPrintln("counter check price with sqrtPx96 ^ 2 >> 192 = ", sqrtPx962Price)
	return PriceBigInt, tickPrice
}

func GetVaults(n int) {
	fin, _, _ := GetInstance3()
	vaults, _ := fin.Vaults(&bind.CallOpts{}, big.NewInt(int64(n)))
	myPrintln(vaults)

}

func SetVaults() {
	fin, _, _ := GetInstance3()
	vaults, _ := fin.Vaults(&bind.CallOpts{}, big.NewInt(0))
	myPrintln(vaults)

}

// func Alloc(accId int) {

// 	myPrintln("----------------------------------------------")
// 	myPrintln(".........alloc .........  ")
// 	myPrintln("----------------------------------------------")

// 	//vaultInstance := GetVaultInstance()
// 	_, vaultInstance, _ := GetInstance3()

// 	ChangeAccount(accId)
// 	NonceGen()

// 	tx, err := vaultInstance.Alloc(Auth)

// 	if err != nil {
// 		log.Fatal("alloc err: ", err)
// 	}

// 	myPrintln("alloc tx: %s", tx.Hash().Hex())

// 	ChangeAccount(Account)

// 	//Readstring("alloc sent...wait for pending..next .. ")
// 	TxConfirm(tx.Hash())

// }
func GetTotalSupply() {
	_, _, vaultInstance := GetInstance3()

	totalsupply, _ := vaultInstance.TotalSupply(&bind.CallOpts{})
	myPrintln("totalsupply:", totalsupply)

}

func GetTVL2(vaultAddr string) *struct {
	Total0 *big.Int
	Total1 *big.Int
} {

	Totals := new(struct {
		Total0 *big.Int
		Total1 *big.Int
	})

	_, stratInstance, _ := GetInstance3()
	token0Ins, _, _, _, _ := GetTokenInstance(Network.TokenA)
	token1Ins, _, _, _, _ := GetTokenInstance(Network.TokenB)

	//implement gettvl
	cHigh, _ := stratInstance.CHigh(&bind.CallOpts{})
	cLow, _ := stratInstance.CLow(&bind.CallOpts{})

	uniliqs, _ := stratInstance.GetUniAmounts(&bind.CallOpts{}, cLow, cHigh)
	myPrintln("liquidity in uniswap:  ", uniliqs)

	lendingAmt0, lendingAmt1, exrate0, exrate1 := GetLendingAmounts(Network.VaultStrat)
	myPrintln("balance in Compound: ", lendingAmt0, lendingAmt1)
	myPrintln("exchange rate: ", exrate0, exrate1)

	// clending0, clending1 := stratInstance.GetCAmounts(&bind.CallOpts{})
	// myPrintln("C Amounts in lending: ", clending0, clending1)

	sbalance0, _ := token0Ins.BalanceOf(&bind.CallOpts{}, common.HexToAddress(Network.VaultStrat))
	sbalance1, _ := token1Ins.BalanceOf(&bind.CallOpts{}, common.HexToAddress(Network.VaultStrat))
	myPrintln("balance in strat: ", sbalance0, sbalance1)

	vbalance0, _ := token0Ins.BalanceOf(&bind.CallOpts{}, common.HexToAddress(Network.Vault))
	vbalance1, _ := token1Ins.BalanceOf(&bind.CallOpts{}, common.HexToAddress(Network.Vault))
	myPrintln("balance in vault: ", vbalance0, vbalance1)

	balance0 := new(big.Int).Add(sbalance0, vbalance0)
	balance1 := new(big.Int).Add(sbalance1, vbalance1)

	Totals.Total0 = balance0.Add(balance0, uniliqs.Amount0).Add(balance0, lendingAmt0)
	Totals.Total1 = balance1.Add(balance1, uniliqs.Amount1).Add(balance1, lendingAmt1)

	myPrintln("tvl: ", Totals.Total0, Totals.Total1)

	return Totals
}

func GetTVL() *struct {
	Total0 *big.Int
	Total1 *big.Int
} {

	return (GetTVL2(Network.Vault))
}

func GetLendingAmounts(vaultAddr string) (*big.Int, *big.Int, *big.Int, *big.Int) {

	cInstance0 := GetCTokenInstance(Network.CTOKEN0)
	cInstance1 := GetCTokenInstance(Network.CTOKEN1)

	//implement gettvl
	exchangeRate0, _ := cInstance0.ExchangeRateStored(&bind.CallOpts{})
	exchangeRate1, _ := cInstance1.ExchangeRateStored(&bind.CallOpts{})

	CAmount0 := checkCTokenBalance(vaultAddr, "CToken0", Network.CTOKEN0)
	CAmount1 := checkCTokenBalance(vaultAddr, "CToken1", Network.CTOKEN1)

	pow1018 := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	//	pow106 := new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)
	underlying0 := CAmount0.Mul(CAmount0, exchangeRate0).Div(CAmount0, pow1018) //.Div(CAmount0, pow1018)
	underlying1 := CAmount1.Mul(CAmount1, exchangeRate1).Div(CAmount1, pow1018) //.Div(CAmount1, pow106)

	return underlying0, underlying1, exchangeRate0, exchangeRate1

}

func ApproveToken(tokenAddress common.Address, maxAmount *big.Int, toAddress string) {

	tokenInstance, err := token.NewApi(tokenAddress, Client)
	if err != nil {
		log.Fatal("tokenInstance,", err)
	}

	//check allowance
	allow, _ := tokenInstance.Allowance(&bind.CallOpts{}, FromAddress, common.HexToAddress(toAddress))

	if allow.Cmp(big.NewInt(0)) > 0 {
		return
	}

	myPrintln("APPROVE:")
	myPrintln("to address be approved: ", toAddress)
	myPrintln("fromAddress: ", FromAddress)
	myPrintln("Allowance amount:", allow)

	NonceGen()

	tx, err := tokenInstance.Approve(Auth, common.HexToAddress(toAddress), maxAmount)

	if err != nil {
		log.Fatal("token approve tx, ", err)
	}

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

	ChangeAccount(account)
	NonceGen()

	poolInstance := GetPoolInstance()
	TokenA, _ := poolInstance.Token0(&bind.CallOpts{})

	TokenB, _ := poolInstance.Token1(&bind.CallOpts{})

	myPrintln("tokenA: ", TokenA)
	myPrintln("tokenB: ", TokenB)

	var maxToken0 = PowX(99999, int(Token[0].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)
	var maxToken1 = PowX(99999, int(Token[1].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)

	ApproveToken(TokenA, maxToken0, Network.Vault)
	ApproveToken(TokenB, maxToken1, Network.Vault)

	ChangeAccount(Account)

}

func AccountInfo() {

	myPrintln("----------------------------------------------")
	myPrintln(".........Account Info.........  ")
	myPrintln("----------------------------------------------")

	for i, _ := range Network.PrivateKey {

		MyAccountInfo(i)
	}

}

func MyAccountInfo(accId int) {

	accountAddress := GetAddress(accId)

	myPrintln("Account  ----", accId)
	myPrintln("Account address ", accountAddress)

	myPrintln("Eth balance:", EthBalance(accountAddress.String()))

	tokenAInstance, err := token.NewApi(common.HexToAddress(Network.TokenA), Client)
	if err != nil {
		log.Fatal("tokenAInstance,", err)
	}

	bal, err := tokenAInstance.BalanceOf(&bind.CallOpts{}, accountAddress)
	myPrintln("tokenA in Wallet ", bal)

	tokenBInstance, err := token.NewApi(common.HexToAddress(Network.TokenB), Client)
	if err != nil {
		log.Fatal("tokenBInstance,", err)
	}

	bal, err = tokenBInstance.BalanceOf(&bind.CallOpts{}, accountAddress)
	myPrintln("tokenB in Wallet ", bal)

	///----------- token in vault

	// vaultInstance, err := vault.NewApi(vaultAddress, Client)
	// if err != nil {
	// 	log.Fatal("vault.NewApi ", err)
	// }

	mybal, totalbal := CalcShares(accountAddress)

	if totalbal.Cmp(big.NewInt(0)) > 0 {

		myPrintln("my share / totalSupply ", mybal.Mul(mybal, big.NewInt(100)).Div(mybal, totalbal), "%")
	}

	vaultInstance := GetVaultInstance()
	Assets, _ := vaultInstance.Assetholder(&bind.CallOpts{}, accountAddress)

	myPrintln("*Assetsholder: ", Assets)

	myPrintln()

}

func CalcShares(myAddress common.Address) (mybal *big.Int, totalbal *big.Int) {

	/// vault as token
	vaultTokenInstance, err := token.NewApi(common.HexToAddress(Network.Vault), Client)
	if err != nil {
		log.Fatal("vaultTokenInstance,", err)
	}

	mybal, _ = vaultTokenInstance.BalanceOf(&bind.CallOpts{}, myAddress)
	Sleep(1000)

	//if mybal.Cmp(big.NewInt(0)) > 0 {
	myPrintln("myShares in vault ", mybal)

	//}

	totalbal, _ = vaultTokenInstance.TotalSupply(&bind.CallOpts{})
	Sleep(1000)
	//if totalbal.Cmp(big.NewInt(0)) > 0 {
	myPrintln("totalSupply in vault ", totalbal)
	//}

	return mybal, totalbal
}

func VaultInfo() {
	VaultInfo2(Network.Vault)
}

func VaultInfo2(vaultAddr string) {

	myPrintln("----------------------------------------------")
	myPrintln(".........Vault Info.........  ")
	myPrintln("----------------------------------------------")

	//vaultInstance := GetVaultInstance2(vaultAddr)
	_, _, vaultInstance := GetInstance3()

	myPrintln("Vault Address:  ", vaultAddr)

	//poolAddress := Network.Pool
	//get ctoken address
	// _CToken0Addr, _ := vaultInstance.CToken0(&bind.CallOpts{})
	// _CToken1Addr, _ := vaultInstance.CToken1(&bind.CallOpts{})
	// myPrintln("Ctoken0 address:", _CToken0Addr)
	// myPrintln("Ctoken1 address:", _CToken1Addr)

	totalSupply, err := vaultInstance.TotalSupply(&bind.CallOpts{})

	myPrintln("totalSupply (total shares in vault) :", totalSupply)
	if err != nil {
		log.Fatal("totalsupply ", err)
	}

	Sleep(100)
	poolInstance := GetPoolInstance()

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})
	Sleep(100)
	tick := slot0.Tick

	// qTickLower, err := vaultInstance.CLow(&bind.CallOpts{})
	// Sleep(100)
	// qTickUpper, err := vaultInstance.CHigh(&bind.CallOpts{})
	// Sleep(100)
	// myPrintln("cLow, tick, cHigh  :", qTickLower, tick, qTickUpper)

	fmt.Println("** in range? ", tick.Cmp(qTickLower) > 0 && tick.Cmp(qTickUpper) < 0)

	// liquidity, err := vaultInstance.GetSSLiquidity(&bind.CallOpts{}, qTickLower, qTickUpper)
	// Sleep(100)
	// myPrintln("Current liquidity in pool :", liquidity)

	// if err != nil {
	// 	log.Fatal("getssliquidity  ", err)
	// }

	// xy, err := vaultInstance.GetPositionAmounts(&bind.CallOpts{}, qTickLower, qTickUpper)
	// Sleep(100)
	// myPrintln("tokenA (x) in pool ", xy.Amount0)
	// myPrintln("tokenB (y) in pool ", xy.Amount1)

	///----------- 分别返回两个toten 的总数, also print tokens in vault
	// totals := GetTVL2(vaultAddr)

	// fmt.Printf("TVL token0=%d\n", totals.Total0)
	// fmt.Printf("TVL token1=%d\n", totals.Total1)
	// myPrintln("decimals0:", int(Token[0].Decimals))
	// myPrintln("decimals1:", int(Token[1].Decimals))

	// uniPortion, _ := vaultInstance.UniPortion(&bind.CallOpts{})
	// Sleep(100)
	// myPrintln("uniPortionRate:", uniPortion)

	// protocolFeeRate, _ := vaultInstance.ProtocolFee(&bind.CallOpts{})
	// Sleep(100)
	// myPrintln("ProtocolFeeRate:", protocolFeeRate)

	//	sqrtPriceX96 := slot0.SqrtPriceX96
	// uniswapPriceBySqrtP, _ := vaultInstance.GetPriceBySQRTP(&bind.CallOpts{}, sqrtPriceX96)
	// myPrintln("GetPriceBySQRTP:", uniswapPriceBySqrtP)

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

	tr, err := Client.TransactionReceipt(context.Background(), tx)
	for i := 0; i < 20; i++ {
		if err != nil {
			//log.Fatal(err)
			time.Sleep(2 * time.Second)
		} else {
			break
		}
		tr, err = Client.TransactionReceipt(context.Background(), tx)
	}

	myPrintln("BlockNumber:", tr.BlockNumber)
	myPrintln("Status:", Stat2Str(tr.Status))
	myPrintln("CumulativeGasUsed:", tr.CumulativeGasUsed)
	myPrintln("GasUsed:", tr.GasUsed)

	//	if 	Sleep(5000)

	//	myPrintln("ContractAddress:", tr.ContractAddress)

}

func SetVaultAddress(_address string, ind int64) {

	instance, err := bridge.NewApi(common.HexToAddress(Network.VaultBridge), Client)

	if err != nil {
		log.Fatal("vaultbridgeInstance err:", err)
	}

	tx, err := instance.SetAddress(Auth, common.HexToAddress(_address), big.NewInt(ind))

	if err != nil {
		log.Fatal("setVaultBridge err: ", err)
	}

	TxConfirm(tx.Hash())

}

func GetVaultAddress(ind int64) {

	instance, err := bridge.NewApi(common.HexToAddress(Network.VaultBridge), Client)

	if err != nil {
		log.Fatal("vaultBridgeInstance err:", err)
	}

	vaultAddress, err := instance.GetAddress(&bind.CallOpts{}, big.NewInt(ind))

	if err != nil {
		log.Fatal("getVaultBridge err: ", err)
	}

	fmt.Println("vaultAddress:", vaultAddress)

}

func AuthAdmin(vaultAddr string, _admin string) {

	instance, err := admin.NewApi(common.HexToAddress(vaultAddr), Client)

	if err != nil {
		log.Fatal("vaultAdmin Instance err:", err)
	}

	exists, err := instance.AuthAdmin(&bind.CallOpts{}, common.HexToAddress(_admin))

	if err != nil {
		log.Fatal("auth admin  err: ", err)
	}

	fmt.Println("auth? ", exists)

}

func GetArbInstance() *arb.Api {

	arbAddress := "0xa9712b5e7C1537936Ba383B2632455A02D9d49B6"

	instance, err := arb.NewApi(common.HexToAddress(arbAddress), Client)
	if err != nil {
		log.Fatal("arb Instance err -- :", err)
	}
	return instance

}

func EthBalanceArb(_addr string) *big.Int {

	instance := GetArbInstance()

	bal, err := instance.EthBalance(&bind.CallOpts{}, common.HexToAddress(_addr))
	if err != nil {
		log.Fatal("eth balance vis Arb instance err:    ", err)
	}

	return bal

}

func ERC20Balance(_erc20 string, _owner string) {

	erc20Instance, name, symbol, decimals, _ := GetTokenInstance(_erc20)
	_ = name
	_ = decimals
	balance, _ := erc20Instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(_owner))
	myPrintln(symbol, " in Wallet ", balance)

}
