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

func Deposit(do bool, amount0 int64, amount1 int64) {

	if !do {
		return
	}
	fmt.Println("Deposit.........  ")

	fmt.Println("vaultAddress: ", vaultAddress)
	fmt.Println("fromAddress: ", config.FromAddress)
	fmt.Println("TokenA:", config.Network.TokenA)
	fmt.Println("TokenB:", config.Network.TokenB)

	instance, err := vault.NewApi(vaultAddress, config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	//test deposit
	amountToken0 := big.NewInt(1e18).Mul(big.NewInt(1e18), big.NewInt(amount0))
	amountToken1 := big.NewInt(1e6).Mul(big.NewInt(1e6), big.NewInt(amount1))

	tokenAInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenA), config.Client)
	if err != nil {
		log.Fatal("tokenAInstance,", err)
	}

	bal, err := tokenAInstance.BalanceOf(&bind.CallOpts{}, config.FromAddress)
	fmt.Println("tokenA amount ", bal)

	tokenBInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenB), config.Client)
	if err != nil {
		log.Fatal("tokenBInstance,", err)
	}

	bal, err = tokenBInstance.BalanceOf(&bind.CallOpts{}, config.FromAddress)
	fmt.Println("tokenB amount ", bal)

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
	fmt.Println("deposit tx sent: %s", tx.Hash().Hex())

	//	time.Sleep(config.Network.PendingTime * time.Second)
	config.Readstring("wait for pending..next .. ")

}
func Withdraw(do bool, percent float64) {

	if !do {
		return
	}
	fmt.Println("Withdraw.........  ")

	fmt.Println("vaultAddress: ", vaultAddress)

	instance, err := vault.NewApi(vaultAddress, config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	ashares, err := instance.BalanceOf(&bind.CallOpts{}, config.FromAddress)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	fmt.Println("shares awailable = ", ashares)
	//shares := new(big.Int).Mul(ashares, config.FloorFloat64ToBigInt(percent))
	shares := new(big.Int).Div(ashares, big.NewInt(2))

	fmt.Println("shares required = ", shares)

	//time.Sleep(config.Network.PendingTime * time.Second)
	if shares.Cmp(ashares) == 1 {
		shares = ashares
	}

	shareHolderAddress := config.FromAddress
	config.NonceGen()
	tx, err := instance.Withdraw(config.Auth,
		shares,
		big.NewInt(0),
		big.NewInt(0),
		shareHolderAddress,
	)

	if err != nil {
		log.Fatal("withdraw: ", err)
	}

	//get the transaction hash
	fmt.Println("withdraw tx sent: ", tx.Hash().Hex())

	config.Readstring("wait for pending..next .. ")

}
func Rebalance(do bool) {

	if !do {
		return
	}

	fmt.Println("vaultAddress: ", vaultAddress)

	///require governance. redo auth
	config.Account = 0
	config.GetSignature(config.Networkid)

	vaultInstance, err := vault.NewApi(vaultAddress, config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	poolInstance, err := pool.NewApi(common.HexToAddress(config.Network.Pool), config.Client)

	if err != nil {
		log.Fatal("poolInstance err:", err)
	}

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	fmt.Println("slot0.Tick:", slot0.Tick)
	fmt.Println("slot0.SqrtPriceX96:", slot0.SqrtPriceX96)
	fmt.Println("getPrice:", getPrice(slot0.SqrtPriceX96, slot0.Tick))

	rangeThreadhold := big.NewInt(60 * 100)
	big0 := big.NewInt(0)

	if slot0.Tick.Cmp(big0) < 0 {
		//rangeThreadhold = rangeThreadhold * -1
	}

	tickLower := new(big.Int).Sub(slot0.Tick, rangeThreadhold)
	tickUpper := new(big.Int).Add(slot0.Tick, rangeThreadhold)

	fmt.Println("tickLower:", tickLower)
	fmt.Println("tickUpper:", tickUpper)

	tickLower = new(big.Int).Div(tickLower, big.NewInt(60))
	tickLower = new(big.Int).Mul(tickLower, big.NewInt(60))

	tickUpper = new(big.Int).Div(tickUpper, big.NewInt(60))
	tickUpper = new(big.Int).Mul(tickUpper, big.NewInt(60))

	fmt.Println("tickLower % tickspacing:", tickLower)
	fmt.Println("tickUpper % tickspacing", tickUpper)

	/// lower cannot be smaller than upper
	if tickLower.Cmp(tickUpper) >= 0 {
		log.Fatal("Error: tickLower >= tickUpper")
	}

	config.NonceGen()
	tx, err := vaultInstance.Rebalance(config.Auth,
		tickLower,
		tickUpper,
		big.NewInt(0),
	)

	if err != nil {
		log.Fatal("rebalance err ", err)
	}

	//get the transaction hash
	fmt.Println("rebal tx sent: %s", tx.Hash().Hex())

	config.Readstring("rebalance set wait for pending..next .. ")

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
func GetTotalAmounts(do bool) *struct {
	Total0 *big.Int
	Total1 *big.Int
} {
	if !do {
		return nil
	}

	instance, err := vault.NewApi(vaultAddress, config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	totals, err := instance.GetTotalAmounts(&bind.CallOpts{})
	fmt.Printf("tvl token0=%d, tvl token1=%d \n", totals.Total0, totals.Total1)
	return &totals

}

func Approve(do bool) {

	if !do {
		return
	}

	fmt.Println("Approve.........  ")

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
	var maxToken0, _ = new(big.Int).SetString("90000000000000000000000", 10)
	tx0, err := token0Instance.Approve(config.Auth, vaultAddress, maxToken0)

	if err != nil {
		log.Fatal("tx0, ", err)
	}

	fmt.Println("approve amountoken0: %s", tx0.Hash().Hex())

	time.Sleep(config.Network.PendingTime * time.Second)
	//	readstring("done... Press any key to continue when ready..........")

	token1Instance, err := token.NewApi(TokenB, config.Client)
	if err != nil {
		log.Fatal("token0Instance,", err)
	}
	config.NonceGen()
	//var amountToken0, _ = new(big.Int).SetString("90000000000000000000000", 10)
	var maxToken1, _ = new(big.Int).SetString("90000000000000000000000", 10)

	tx1, err := token1Instance.Approve(config.Auth, vaultAddress, maxToken1)

	if err != nil {
		log.Fatal("tx1, ", err)
	}

	fmt.Println("approve amountoken1: %s", tx1.Hash().Hex())

	config.Readstring("Approve, wait for pending..next .. ")

}

func AccountInfo(do bool) {

	if !do {
		return
	}

	fmt.Println("Account.........  ")

	tokenAInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenA), config.Client)
	if err != nil {
		log.Fatal("tokenAInstance,", err)
	}

	bal, err := tokenAInstance.BalanceOf(&bind.CallOpts{}, config.FromAddress)
	fmt.Println("tokenA in Wallet ", bal)

	tokenBInstance, err := token.NewApi(common.HexToAddress(config.Network.TokenB), config.Client)
	if err != nil {
		log.Fatal("tokenBInstance,", err)
	}

	bal, err = tokenBInstance.BalanceOf(&bind.CallOpts{}, config.FromAddress)
	fmt.Println("tokenB in Wallet ", bal)

}
func ValutInfo(do bool) {
	if !do {
		return
	}

	fmt.Println("VaultInfo.........  ")

	instance, err := vault.NewApi(vaultAddress, config.Client)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	///----------- 当前vault 里的 totalSupply
	totalSupply, err := instance.TotalSupply(&bind.CallOpts{})
	fmt.Println("totalSupply (total shares in vault) :", totalSupply)

	///----------- 当前uniswap里的total liquidity
	tickLower := big.NewInt(-199740)
	tickUpper := big.NewInt(-187740)
	liquidity, err := instance.GetSSLiquidity(&bind.CallOpts{}, tickLower, tickUpper)
	fmt.Println("Current liquidity  :", liquidity)

	///----------- fee0 的总数
	accumulateFees0, _ := instance.AccumulateFees0(&bind.CallOpts{})
	fmt.Println("accumulateFees0  :", accumulateFees0)

	///----------- fee1 的总数
	accumulateFees1, _ := instance.AccumulateFees1(&bind.CallOpts{})
	fmt.Println("accumulateFees1  :", accumulateFees1)

	///----------- 分别返回两个toten 的总数
	GetTotalAmounts(true)

	///-----------

	///----------- account x 的 shares
	tokenInstance, err := token.NewApi(vaultAddress, config.Client)
	if err != nil {
		log.Fatal("tokenInstance,", err)
	}

	bal, err := tokenInstance.BalanceOf(&bind.CallOpts{}, config.FromAddress)
	fmt.Println("shares in vault for Account ", config.Account)
	fmt.Println(bal)

}
