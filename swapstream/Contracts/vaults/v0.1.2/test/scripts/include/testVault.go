package include

import (
	"fmt"
	"log"
	"math/big"
	"time"

	//	"time"

	//	factory "../../../../../../../uniswap/v3/deploy/UniswapV3Factory"
	pool "../../../../../../../uniswap/v3/deploy/UniswapV3Pool"
	token "../../../../../Tokens/erc20/deploy/Token"
	vault "../../../deploy/FeeMaker"
	"../config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var vaultAddress = common.HexToAddress(config.Network.Vault)
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

	fmt.Println("vaultAddress: ", vaultAddress)
	fmt.Println("fromAddress: ", config.FromAddress)
	fmt.Println("TokenA:", config.Network.TokenA)
	fmt.Println("TokenB:", config.Network.TokenB)

	instance, err := vault.NewApi(vaultAddress, config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	tokenAInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenA), config.Client)
	if err != nil {
		log.Fatal("tokenAInstance,", err)
	}

	bal, err := tokenAInstance.BalanceOf(&bind.CallOpts{}, config.FromAddress)
	fmt.Println("tokenA in Wallet ", config.Pricef(bal, int(config.Decimals0)))

	if err != nil {
		log.Fatal("bal err ", err)
	}

	tokenBInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenB), config.Client)
	if err != nil {
		log.Fatal("tokenBInstance,", err)
	}

	bal, err = tokenBInstance.BalanceOf(&bind.CallOpts{}, config.FromAddress)
	fmt.Println("tokenB in Wallet ", config.Pricef(bal, int(config.Decimals1)))

	//  amount0 * 10^decimals
	d0 := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(config.Decimals0)), nil)
	amountToken0 := new(big.Int).Mul(d0, big.NewInt(int64(amount0)))
	fmt.Println("d0:", d0)

	d1 := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(config.Decimals1)), nil)
	fmt.Println("d1:", d1)
	amountToken1 := new(big.Int).Mul(d1, big.NewInt(int64(amount1)))

	fmt.Println("tokenA to Vault:", config.Pricef(amountToken0, int(config.Decimals0)))
	fmt.Println("tokenB to Vault:", config.Pricef(amountToken1, int(config.Decimals1)))

	config.NonceGen()

	tx, err := instance.Deposit(config.Auth,
		amountToken0,
		amountToken1,
		config.FromAddress,
	)

	if err != nil {
		log.Fatal("deposit err: ", err)
	}

	//get the transaction hash
	fmt.Println("deposit tx: %s", tx.Hash().Hex())

	//	time.Sleep(config.Network.PendingTime * time.Second)
	config.Readstring("deposit sent...wait for pending..next .. ")

	VaultInfo(1)

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

	fmt.Println("vaultAddress: ", vaultAddress)

	percent := param[1]

	instance, err := vault.NewApi(vaultAddress, config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	///get account's available shares
	ashares, err := instance.BalanceOf(&bind.CallOpts{}, config.FromAddress)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	fmt.Println("shares available = ", config.Pricef(ashares, 18))

	///calc required shares to withdraw

	shares := new(big.Int).Div(new(big.Int).Mul(ashares, big.NewInt(percent)), big.NewInt(100))
	fmt.Println("shares required = ", percent, "% = ", config.Pricef(shares, 18))

	/// if required share > available share, set withdraw shares = awailable shares
	if shares.Cmp(ashares) == 1 {
		shares = ashares
	}

	if shares.Cmp(big.NewInt(0)) <= 0 {
		fmt.Println("share<=0 ", shares)
		return
	}

	config.NonceGen()
	tx, err := instance.Withdraw(config.Auth,
		shares,
		big.NewInt(0),
		big.NewInt(0),
		recipient,
	)

	/// reset account back
	config.Auth = config.GetSignature(config.Networkid, config.Account)

	if err != nil {
		log.Fatal("withdraw: ", err)
	}

	//get the transaction hash
	fmt.Println("withdraw tx: ", tx.Hash().Hex())

	config.Readstring("withdraw sent.... wait for pending..next .. ")

}

/// param0 : fullRangeSize, param1: tickspacing
func Rebalance(do int, param [2]int64) {

	if do <= 0 {
		return
	}

	fmt.Println("----------------------------------------------")
	fmt.Println(".........Rebalance .........  ")
	fmt.Println("----------------------------------------------")

	fmt.Println("vaultAddress: ", vaultAddress)

	vaultInstance, err := vault.NewApi(vaultAddress, config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	poolInstance, err := pool.NewApi(common.HexToAddress(config.Network.Pool), config.Client)

	if err != nil {
		log.Fatal("poolInstance err:", err)
	}

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	//fmt.Println("slot0.Tick:", slot0.Tick)
	//fmt.Println("slot0.SqrtPriceX96:", slot0.SqrtPriceX96)
	fmt.Println("Current Price from sqrtPriceX96: ", config.Pricef(getPrice(slot0.SqrtPriceX96, slot0.Tick), 6))

	setRange(slot0.Tick, param)

	///require governance. redo auth
	config.Auth = config.GetSignature(config.Networkid, 0)

	config.NonceGen()

	tx, err := vaultInstance.Rebalance(config.Auth,
		tickLower,
		tickUpper,
		big.NewInt(0),
	)

	///require governance. redo auth
	config.Auth = config.GetSignature(config.Networkid, config.Account)

	if err != nil {
		log.Fatal("rebalance err ", err)
	}

	//get the transaction hash
	fmt.Println("rebal tx: ", tx.Hash().Hex())

	config.Readstring("Rebalance sent.....  wait for pending..next .. ")

	PoolInfo()

}

func setRange(tick *big.Int, param [2]int64) {

	fullRange := param[0]
	tickSpacing := param[1]

	/// rangehtehrehold = fullrange / 2
	rangeThreadhold := new(big.Int).Div(big.NewInt(fullRange), big.NewInt(2))
	//	big0 := big.NewInt(0)

	/// ticklower = current tick - rangeThread
	tickLower = new(big.Int).Sub(tick, rangeThreadhold)
	/// tickupper = current tick + rangeThread
	tickUpper = new(big.Int).Add(tick, rangeThreadhold)

	fmt.Println("tick Upper:", tickUpper)
	fmt.Println("current tick:", tick)
	fmt.Println("tick Lower:", tickLower)

	/// must be dividable by tickspacing exactly.
	tickLower = new(big.Int).Div(tickLower, big.NewInt(tickSpacing))
	tickLower = new(big.Int).Mul(tickLower, big.NewInt(tickSpacing))

	tickUpper = new(big.Int).Div(tickUpper, big.NewInt(tickSpacing))
	tickUpper = new(big.Int).Mul(tickUpper, big.NewInt(tickSpacing))

	qTickLower = tickLower
	qTickUpper = tickUpper

	fmt.Printf("tickspacing %d:\n", tickSpacing)
	fmt.Printf("tickLower adjust by tickspacing %d:\n", tickLower)
	fmt.Printf("tickUpper adjust by tickspacing %d:\n", tickUpper)

	/// lower cannot be smaller than upper
	if tickLower.Cmp(tickUpper) >= 0 {
		log.Fatal("Error: tickLower >= tickUpper")
	}

}

/// formula:
///
/// (sqrtPricex96^2 * 1e18) >> (96*2)
/// solc: uint(sqrtPriceX96).mul(uint(sqrtPriceX96)).mul(1e18) >> (96 * 2);
func getPrice(SqrtPriceX96 *big.Int, tick *big.Int) *big.Int {

	//fmt.Println("test sqrtP * tick: ", SqrtPriceX96.Add(SqrtPriceX96, tick))
	//SqrtPriceX96 = new(big.Int).Add(SqrtPriceX96, tick)
	//fmt.Println("after sqrtP * tick:", SqrtPriceX96)

	sqrpricePow2 := new(big.Int).Mul(SqrtPriceX96, SqrtPriceX96)
	oPrice := new(big.Int).Mul(sqrpricePow2, big.NewInt(1e18))
	Price := new(big.Int).Rsh(oPrice, 96*2)
	return Price
}

func GetTotalAmounts() *struct {
	Total0 *big.Int
	Total1 *big.Int
} {

	vaultInstance, err := vault.NewApi(vaultAddress, config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	totals, err := vaultInstance.GetTotalAmounts(&bind.CallOpts{})
	fmt.Printf("tvl token0=%f, tvl token1=%f \n",
		config.Pricef(totals.Total0, int(config.Decimals0)),
		config.Pricef(totals.Total1, int(config.Decimals1)))
	return &totals

}

func Approve(do int) {

	if do <= 0 {
		return
	}

	fmt.Println("----------------------------------------------")
	fmt.Println(".........Approve.........  ")
	fmt.Println("----------------------------------------------")

	fmt.Println("vaultAddress: ", vaultAddress)
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

	var maxToken0, _ = new(big.Int).SetString("900000000000000000000000000000", 10)
	tx0, err := token0Instance.Approve(config.Auth, vaultAddress, maxToken0)

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
	var maxToken1, _ = new(big.Int).SetString("900000000000000000000000000000", 10)

	tx1, err := token1Instance.Approve(config.Auth, vaultAddress, maxToken1)

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

	getAccountInfo(0)
	getAccountInfo(1)

}

func getAccountInfo(accId int) {

	accountAddress := config.GetAddress(accId)

	fmt.Println("Account  ", accId)
	fmt.Println("Account address ", accountAddress)

	tokenAInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenA), config.Client)
	if err != nil {
		log.Fatal("tokenAInstance,", err)
	}

	bal, err := tokenAInstance.BalanceOf(&bind.CallOpts{}, accountAddress)
	fmt.Println("tokenA in Wallet ", config.Pricef(bal, int(config.Decimals0)))

	tokenBInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenB), config.Client)
	if err != nil {
		log.Fatal("tokenBInstance,", err)
	}

	bal, err = tokenBInstance.BalanceOf(&bind.CallOpts{}, accountAddress)
	fmt.Println("tokenB in Wallet ", config.Pricef(bal, int(config.Decimals1)))

	///----------- token in vault

	// vaultInstance, err := vault.NewApi(vaultAddress, config.Client)
	// if err != nil {
	// 	log.Fatal("vault.NewApi ", err)
	// }

	/// vault as token
	vaultTokenInstance, err := token.NewApi(vaultAddress, config.Client)
	if err != nil {
		log.Fatal("vaultTokenInstance,", err)
	}

	decimalsShare, _ := vaultTokenInstance.Decimals(&bind.CallOpts{})
	bal, err = vaultTokenInstance.BalanceOf(&bind.CallOpts{}, accountAddress)
	fmt.Println("shares in vault ", config.Pricef(bal, int(decimalsShare)))

}

func VaultInfo(do int64) {

	if do <= 0 {
		return
	}
	fmt.Println("----------------------------------------------")
	fmt.Println(".........Vault Info.........  ")
	fmt.Println("----------------------------------------------")

	fmt.Println("Vault Address:  ", config.Network.Vault)

	instance, err := vault.NewApi(vaultAddress, config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	///----------- 当前vault 里的 totalSupply
	totalSupply, err := instance.TotalSupply(&bind.CallOpts{})
	fmt.Println("totalSupply (total shares in vault) :", totalSupply)
	if err != nil {
		log.Fatal("totalsupply ", err)
	}
	///----------- 当前uniswap里的total liquidity
	//tickLower := big.NewInt(-199740)
	//tickUpper := big.NewInt(-187740)
	//tickLower := big.NewInt(194280)
	//tickUpper := big.NewInt(206280)

	liquidity, err := instance.GetSSLiquidity(&bind.CallOpts{}, qTickLower, qTickUpper)
	fmt.Println("Current liquidity  :", liquidity)
	if err != nil {
		log.Fatal("getssliquidity  ", err)
	}
	///----------- fee0 的总数
	accumulateFees0, _ := instance.AccumulateFees0(&bind.CallOpts{})
	fmt.Println("accumulateFees0  :", accumulateFees0)

	///----------- fee1 的总数
	accumulateFees1, _ := instance.AccumulateFees1(&bind.CallOpts{})
	fmt.Println("accumulateFees1  :", accumulateFees1)

	///----------- 分别返回两个toten 的总数
	GetTotalAmounts()

	///-----------

	///----------- account x 的 shares
	tokenInstance, err := token.NewApi(vaultAddress, config.Client)
	if err != nil {
		log.Fatal("tokenInstance,", err)
	}

	accAddress0 := config.GetAddress(0)
	bal, err := tokenInstance.BalanceOf(&bind.CallOpts{}, accAddress0)
	fmt.Printf("shares in vault for Account %d\n", 0)
	if err != nil {
		log.Fatal("account 0 share balanceOf ", err)
	}
	if bal.Cmp(big.NewInt(0)) > 0 {
		fmt.Println(bal, " ,  ", new(big.Int).Div(new(big.Int).Mul(bal, big.NewInt(100)), totalSupply), "% of total ", totalSupply)
	}

	accAddress1 := config.GetAddress(1)
	bal, err = tokenInstance.BalanceOf(&bind.CallOpts{}, accAddress1)
	fmt.Printf("shares in vault for Account %d\n", 1)
	if err != nil {
		log.Fatal("account 1 share balanceOf ", err)
	}
	if bal.Cmp(big.NewInt(0)) > 0 {
		fmt.Println(bal, " ,  ", new(big.Int).Div(new(big.Int).Mul(bal, big.NewInt(100)), totalSupply), "% of total ", totalSupply)
	}

	/// token 0 and token1 balance in vault
	tokenAInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenA), config.Client)
	if err != nil {
		log.Fatal("tokenAInstance,", err)
	}

	bal, err = tokenAInstance.BalanceOf(&bind.CallOpts{}, vaultAddress)
	if err != nil {
		log.Fatal("tokenA balanceof vaule ", err)
	}
	fmt.Println("tokenA in vault ", bal)

	tokenBInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenB), config.Client)
	if err != nil {
		log.Fatal("tokenBInstance,", err)
	}

	bal, err = tokenBInstance.BalanceOf(&bind.CallOpts{}, vaultAddress)
	if err != nil {
		log.Fatal("tokenB balanceof vaule ", err)
	}
	fmt.Println("tokenB in vault ", bal)

	/// tokens in pool

	poolInstance, err := pool.NewApi(common.HexToAddress(config.Network.Pool), config.Client)

	if err != nil {
		log.Fatal("poolInstance err:", err)
	}

	slot0, err := poolInstance.Slot0(&bind.CallOpts{})
	if err != nil {
		log.Fatal("slot0 err:", err)
	}
	/*
		x := liquidity / slot0.SqrtPriceX96
		y := liquidity * slot0.SqrtPriceX96
	*/

	lf := new(big.Float)
	lf.SetString(liquidity.String())
	p := getPrice(slot0.SqrtPriceX96, slot0.Tick)
	pf := new(big.Float)
	pf.SetString(p.String())

	sqrtP := new(big.Float).Sqrt(pf)

	fmt.Println("sqrtP ", sqrtP)
	fmt.Println("liquidity ", lf)

	x := new(big.Float).Quo(lf, sqrtP)
	y := new(big.Float).Mul(lf, sqrtP)

	fmt.Println("tokenA (x) in pool ", x)

	fmt.Println("tokenB (y) in pool ", y)

	//fmt.("y/x = p ??",  new(big.Float).Quo(y, x), "=",pf )

	//fmt.("x ratio = ",  new(big.Float).Quo(new(big.Float).Mul(x, pf), x), "=",pf )

}

func Init() {

	tokenAInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenA), config.Client)
	if err != nil {
		log.Fatal("tokenAInstance,", err)
	}
	config.Decimals0, _ = tokenAInstance.Decimals(&bind.CallOpts{})

	tokenBInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenB), config.Client)
	if err != nil {
		log.Fatal("tokenBInstance,", err)
	}

	config.Decimals1, _ = tokenBInstance.Decimals(&bind.CallOpts{})

}
