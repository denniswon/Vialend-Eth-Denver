package main

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	mintcallee "../../deploy/TestUniswapV3Callee"
	factory "../../deploy/UniswapV3Factory"
	pool "../../deploy/UniswapV3Pool"
	token "../../deploy/token"
)

type nClient struct {
	clientUrl  string
	factory    string
	callee     string
	privateKey string
	tokenA     string
	tokenB     string
	newOwner   string
}

func main() {

	networks := [...]nClient{

		{ // 0 local
			"http://127.0.0.1:7545",
			"0x42414849A1f13b4d17C2B2eCd2dBfc6124567416", //factory
			"0xE3c433a67e56BD49d93cCA86728C07bE531c2DCc", //callee
			"e8ef3a782d9002408f2ca6649b5f95b3e5772364a5abe203f1678817b6093ff0",
			"0x3b88D0E8B11eb7C5fbC63F1Af1B2795DB1724C59", //tokenA tusdc
			"0xeDFBec53F1DA0995ea493ebB0A8Ff630Bb2f1e23", //tokenB tweth
			"0xeBb29c07455113c30810Addc123D0D7Cd8637aea",
		},

		{ // 1 local user2
			"http://127.0.0.1:7545",
			"0x42414849A1f13b4d17C2B2eCd2dBfc6124567416",                       // factory
			"0xE3c433a67e56BD49d93cCA86728C07bE531c2DCc",                       //callee
			"f804a123dd9876c73cef5d198cce0899e6dfc2f851ed2527b003e11cd5383c54", //(0xeBb29c07455113c30810Addc123D0D7Cd8637aea)
			"0x3b88D0E8B11eb7C5fbC63F1Af1B2795DB1724C59",                       //tokenA tusdc
			"0xeDFBec53F1DA0995ea493ebB0A8Ff630Bb2f1e23",                       //tokenB tweth
			"0xeBb29c07455113c30810Addc123D0D7Cd8637aea",
		},

		{ ///2  goerli admin test 1
			"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"",
			"0xc7853A9E7b602Aafe36b8fb95E4b67a2001FD9C5", //new uniswapv3factory modified requires
			"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7",
			"0x2aDEca523FbBF0937A9419924feAB607Bf599311", //tokenA
			"0xc4AFB13b10f7C49Af721860A188D6443D0fF8747", //tokenB
			"0x4F211267896C4D3f2388025263AC6BD67B0f2C54", //new owner, test user 1
		},

		{"https://rinkeby.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"not yet",
			"",
			"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7",
			"",
			"", ""},
	}

	nid := 1

	doCreatePool := false
	doSetOwner := false
	doApprove := false
	doInitial := false

	doMint := false

	mintAmount := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(200))
	//or SetString("100000000000000000000", 10)

	doSwap := true
	swapAmount := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(20))

	doVaultBurn := false // revert
	doPoolBurn := false  // 'ls' error
	doCollect := false   //

	fmt.Println("Current network: ", networks[nid].clientUrl)

	tokenA := common.HexToAddress(networks[nid].tokenA)
	tokenB := common.HexToAddress(networks[nid].tokenB)
	calleeAddress := common.HexToAddress(networks[nid].callee)

	fmt.Println("token0:", tokenA)
	fmt.Println("token1:", tokenB)

	fee := big.NewInt(3000)

	client, err := ethclient.Dial(networks[nid].clientUrl)

	ContractAddress := networks[nid].factory
	privateKey, err := crypto.HexToECDSA(networks[nid].privateKey)

	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println("msg.sender:", fromAddress)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Value = big.NewInt(0)                 // in wei
	auth.GasLimit = uint64(6721975)            // in gwei
	auth.GasPrice = big.NewInt(8 * 1000000000) // 20 gwei.  1 gwei = 1000000000 wei

	contractAddress := common.HexToAddress(ContractAddress)
	instance, err := factory.NewApi(contractAddress, client)

	if err != nil {
		panic(err)
	}

	if doSetOwner {

		auth.Nonce = nonceGen(client, fromAddress)

		tx, err := instance.SetOwner(auth,
			common.HexToAddress(networks[nid].newOwner),
		)

		if err != nil {
			panic(err)
		}

		//get the transaction hash
		fmt.Println("setowner tx sent: ", tx.Hash().Hex())

	}

	if doCreatePool {

		auth.Nonce = nonceGen(client, fromAddress)

		tx, err := instance.CreatePool(auth,
			tokenA,
			tokenB,
			fee,
		)

		if err != nil {
			panic(err)
		}

		//get the transaction hash
		fmt.Println("createpool tx sent: %s", tx.Hash().Hex())

		readstring("check the address and  press any key when ready... ")
	}

	/*
		tokenA = common.HexToAddress("0x3fF5E22B4be645EF1CCc8C6e32EDe6b35c569AE4")
		tokenB = common.HexToAddress("0xFA5dF5372c03D4968d128D624e3Afeb61031a777")
		fee = big.NewInt(10000)
		contractAddress = common.HexToAddress("0x1f98431c8ad98523631ae4a59f267346ea31f984") //factory address

		fmt.Println("tokenA:", tokenA)
		fmt.Println("tokenB:", tokenB)
		fmt.Println("fee:", fee)
		fmt.Println("factory:", contractAddress)

		//create a new instance for the call

	*/

	newInstance, err := factory.NewApi(contractAddress, client)
	poolAddress, err := newInstance.GetPool(&bind.CallOpts{}, tokenA, tokenB, fee)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("pool:", poolAddress)

	calleeInstance, _ := mintcallee.NewApi(calleeAddress, client)
	if err != nil {
		panic(err)
	}
	poolInstance, err := pool.NewApi(poolAddress, client)
	token0, err := poolInstance.Token0(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	fmt.Println("token0:", token0)

	token0Instance, _ := token.NewApi(token0, client)

	total0, _ := token0Instance.TotalSupply(&bind.CallOpts{})
	fmt.Println("token0 totalsupply:", total0)

	token1, err := poolInstance.Token1(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("token1:", token1)

	token1Instance, _ := token.NewApi(token1, client)

	///initial range
	//	tickLower := big.NewInt(-887220)
	//	tickUpper := big.NewInt(487220)

	tickLower := big.NewInt(126600)
	tickUpper := big.NewInt(206800)

	//tickLower := big.NewInt(-60)
	//tickUpper := big.NewInt(18660)

	//# Set ETH/USDC price to 2000
	inverse := token0 == tokenA

	fmt.Println("token0 is takenA:", inverse)

	if doInitial {

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

		biSqrtPX96, _ := new(big.Int).SetString(s, 10)
		fmt.Println("bigInt.sqrtPriceX96:", biSqrtPX96)
		auth.Nonce = nonceGen(client, fromAddress)

		tx, err := poolInstance.Initialize(auth, biSqrtPX96)

		if err != nil {
			panic(err)
		}
		fmt.Println("initial tx sent:", tx.Hash().Hex())

		readstring("check the address and  press any key when ready... ")

	}

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	fmt.Println("slot0:", slot0)

	rl := (slot0.Tick).Cmp(tickLower)
	ru := (slot0.Tick).Cmp(tickUpper)

	fmt.Println("in range? ", ru == -1 && rl == 1)
	fmt.Println("lower, tick, upper:? ", tickLower, ",", tickUpper, ",", slot0.Tick)

	/*	fmt.Println(`slot0:
		SqrtPriceX96              *big.Int
		Tick                       *big.Int
		ObservationIndex           uint16
		ObservationCardinality     uint16
		ObservationCardinalityNext uint16
		FeeProtocol                uint8
		Unlocked                   bool	`)
	*/
	feeRate, _ := poolInstance.Fee(&bind.CallOpts{})
	fmt.Println("Fee:", feeRate)

	liquidity, _ := poolInstance.Liquidity(&bind.CallOpts{})
	fmt.Println("liquidity:", liquidity)

	if doSwap {

		auth.Nonce = nonceGen(client, fromAddress)

		if doApprove {
			//approve token0 and token1 to be used by pool ?
			var maxToken0, _ = new(big.Int).SetString("90000000000000000000000", 10)
			auth.Nonce = nonceGen(client, fromAddress)
			tx0, err := token0Instance.Approve(auth, calleeAddress, maxToken0)

			if err != nil {
				panic(err)
			}

			var maxToken1, _ = new(big.Int).SetString("9000000000000000000000", 10)
			auth.Nonce = nonceGen(client, fromAddress)
			tx1, err := token1Instance.Approve(auth, calleeAddress, maxToken1)

			if err != nil {
				panic(err)
			}

			fmt.Println("approve amountoken0: %s", tx0.Hash().Hex())
			fmt.Println("approve amountoken1: %s", tx1.Hash().Hex())

			readstring("done... Press any key to continue when ready..........")

		}

		swapN := 10
		for i := 0; i < swapN; i++ {
			auth.Nonce = nonceGen(client, fromAddress)
			auth.GasLimit = uint64(6721975)            // in gwei
			auth.GasPrice = big.NewInt(8 * 1000000000) // 20 gwei.  1 gwei = 1000000000 wei

			amount := swapAmount

			//recipient := fromAddress

			if i < (swapN / 2) {

				_, err := calleeInstance.Swap(auth, poolAddress, false, amount)

				/*
					slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

					_, err := calleeInstance.Swap0ForExact1(auth,
						poolAddress,
						amount,
						recipient,
						slot0.SqrtPriceX96.Sub(slot0.SqrtPriceX96, slot0.Tick))
				*/
				if err != nil {
					panic(err)
				}

			} else {
				_, err := calleeInstance.Swap(auth, poolAddress, true, amount)

				/*
					slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

					_, err := calleeInstance.SwapExact1For0(auth,
						poolAddress,
						amount,
						recipient,
						slot0.SqrtPriceX96.Add(slot0.SqrtPriceX96, slot0.Tick))
				*/
				if err != nil {
					panic(err)
				}

			}

			if err != nil {
				panic(err)
			}

			//sqrtPx962Price := math.Floor((sqrtPriceX96 * sqrtPriceX96) / math.Pow(2, 2*96))
			fmt.Println("slot0.sqrtPriceX96:", slot0.SqrtPriceX96)

			s1 := big.NewInt(0)

			s1.Set(slot0.SqrtPriceX96)

			s1 = s1.Mul(s1, s1)
			fmt.Println("s1:", s1)

			sqrtPx962Price := s1.Rsh(s1, 2*96)
			fmt.Println("slot0.sqrtPriceX96:", slot0.SqrtPriceX96)
			fmt.Println("curent Price;", sqrtPx962Price)

			amount0, _ := token0Instance.BalanceOf(&bind.CallOpts{}, poolAddress)
			amount1, _ := token1Instance.BalanceOf(&bind.CallOpts{}, poolAddress)
			fmt.Println("token0 in pool :", amount0)
			fmt.Println("token1 in pool :", amount1)

			from0, _ := token0Instance.BalanceOf(&bind.CallOpts{}, fromAddress)
			from1, _ := token1Instance.BalanceOf(&bind.CallOpts{}, fromAddress)
			fmt.Println("token0 in from after swap:", from0)
			fmt.Println("token1 in from after swap:", from1)

			slot0, _ := poolInstance.Slot0(&bind.CallOpts{})
			fmt.Println("slot0:", slot0)

			fmt.Println("swap ", amount)

			rl := (slot0.Tick).Cmp(tickLower)
			ru := (slot0.Tick).Cmp(tickUpper)

			fmt.Println("in range? ", ru == -1 && rl == 1)
			fmt.Println("lower, tick, upper:? ", tickLower, ",", tickUpper, ",", slot0.Tick)

			readstring("check the address and  press any key when ready... ")
		}
	}

	if doMint {
		/*
				auth.Nonce = nonceGen(client, fromAddress)

				tx, err := poolInstance.IncreaseObservationCardinalityNext(auth, 100)

				if err != nil {
					panic(err)
				}
				fmt.Println("IncreaseObservationCardinalityNext tx sent:", tx.Hash().Hex())

			readstring("check the address and  press any key when ready... ")
		*/
		fromToken0, _ := token0Instance.BalanceOf(&bind.CallOpts{}, fromAddress)
		fromToken1, _ := token1Instance.BalanceOf(&bind.CallOpts{}, fromAddress)
		fmt.Println("fromtoken0:", fromToken0)
		fmt.Println("fromToken1:", fromToken1)

		if doApprove {
			//approve token0 and token1 to be used by pool ?
			var maxToken0, _ = new(big.Int).SetString("9000000000000000000000000000000", 10)
			auth.Nonce = nonceGen(client, fromAddress)
			tx0, err := token0Instance.Approve(auth, calleeAddress, maxToken0)

			if err != nil {
				panic(err)
			}

			var maxToken1, _ = new(big.Int).SetString("9000000000000000000000000000000", 10)
			auth.Nonce = nonceGen(client, fromAddress)
			tx1, err := token1Instance.Approve(auth, calleeAddress, maxToken1)

			if err != nil {
				panic(err)
			}

			fmt.Println("approve amountoken0: %s", tx0.Hash().Hex())
			fmt.Println("approve amountoken1: %s", tx1.Hash().Hex())

			readstring("done... Press any key to continue when ready..........")
		}
		/// do mint pool
		recipient := fromAddress

		/*
			var data []byte
						hash := sha3.NewLegacyKeccak256() // old sha3.NewKeccak256()
						transferFnSignature := []byte("transfer(address,uint256)")
						hash.Write(transferFnSignature)
					methodID := hash.Sum(nil)[:4]
					fmt.Println("methodID:", hexutil.Encode(methodID)) // 0xa9059cbb

					paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
					fmt.Println("padded address:", hexutil.Encode(paddedAddress)) // 0x00

				paddedAddress := common.LeftPadBytes(fromAddress.Bytes(), 32)
				//data = keccak256(abi.encodePacked(fromAddress))
				data1 := hexutil.Encode(paddedAddress)
				data = append(data, data1...)
		*/

		auth.Nonce = nonceGen(client, fromAddress)
		auth.GasLimit = uint64(6721975)            // in gwei
		auth.GasPrice = big.NewInt(5 * 1000000000) // 20 gwei.  1 gwei = 1000000000 wei

		tx, err := calleeInstance.Mint(auth,
			poolAddress,
			recipient,
			tickLower,
			tickUpper,
			mintAmount)

		if err != nil {
			panic(err)
		}
		fmt.Println("Mint tx sent:", tx.Hash().Hex())

		readstring("check the address and  press any key when ready... ")

	}

	inpool0, _ := token0Instance.BalanceOf(&bind.CallOpts{}, poolAddress)
	inpool1, _ := token1Instance.BalanceOf(&bind.CallOpts{}, poolAddress)
	from0, _ := token0Instance.BalanceOf(&bind.CallOpts{}, fromAddress)
	from1, _ := token1Instance.BalanceOf(&bind.CallOpts{}, fromAddress)
	liquidity, _ = poolInstance.Liquidity(&bind.CallOpts{})
	fmt.Println("liquidity:", liquidity)
	fmt.Println("token0 in pool :", inpool0)
	fmt.Println("token1 in pool :", inpool1)
	fmt.Println("token0 in from:", from0)
	fmt.Println("token1 in from:", from1)

	if doPoolBurn {

		auth.Nonce = nonceGen(client, fromAddress)

		tx, err := poolInstance.Burn(auth,
			tickLower,
			tickUpper,
			liquidity)

		if err != nil {
			panic(err)
		}
		fmt.Println("burn tx sent:", tx.Hash().Hex())

		readstring("check the address and  press any key when ready... ")

		auth.Nonce = nonceGen(client, fromAddress)
		tx, err = poolInstance.Collect(auth,
			fromAddress,
			tickLower,
			tickUpper,
			inpool0,
			inpool1)

		if err != nil {
			panic(err)
		}
		fmt.Println("collect tx sent:", tx.Hash().Hex())

		readstring("check the address and  press any key when ready... ")

		liquidity, _ := poolInstance.Liquidity(&bind.CallOpts{})
		fmt.Println("liquidity:", liquidity)

		inpool0, _ := token0Instance.BalanceOf(&bind.CallOpts{}, poolAddress)
		inpool1, _ := token1Instance.BalanceOf(&bind.CallOpts{}, poolAddress)
		fmt.Println("token0 in pool :", inpool0)
		fmt.Println("token1 in pool :", inpool1)

		from0, _ := token0Instance.BalanceOf(&bind.CallOpts{}, fromAddress)
		from1, _ := token1Instance.BalanceOf(&bind.CallOpts{}, fromAddress)
		fmt.Println("token0 in from:", from0)
		fmt.Println("token1 in from:", from1)

	}

	if doVaultBurn {

		auth.Nonce = nonceGen(client, fromAddress)

		tx, err := calleeInstance.BurnAndCollect(auth,
			poolAddress,
			tickLower,
			tickUpper,
		)

		if err != nil {
			panic(err)
		}
		fmt.Println("burn tx sent:", tx.Hash().Hex())

		readstring("check the address and  press any key when ready... ")

		liquidity, _ := poolInstance.Liquidity(&bind.CallOpts{})
		fmt.Println("liquidity:", liquidity)

		inpool0, _ := token0Instance.BalanceOf(&bind.CallOpts{}, poolAddress)
		inpool1, _ := token1Instance.BalanceOf(&bind.CallOpts{}, poolAddress)
		fmt.Println("token0 in pool :", inpool0)
		fmt.Println("token1 in pool :", inpool1)

		from0, _ := token0Instance.BalanceOf(&bind.CallOpts{}, fromAddress)
		from1, _ := token1Instance.BalanceOf(&bind.CallOpts{}, fromAddress)
		fmt.Println("token0 in from:", from0)
		fmt.Println("token1 in from:", from1)

	}

	if doCollect {

		auth.Nonce = nonceGen(client, fromAddress)

		recipient := fromAddress
		amount0Requested, _ := new(big.Int).SetString("11579208923731619542357098500868790", 10)
		amount1Requested, _ := new(big.Int).SetString("11579208923731619542357098500868790", 10)

		fmt.Println("amount0Requested:", amount0Requested)
		fmt.Println("amount1Requested:", amount1Requested)

		tx, err := poolInstance.Collect(auth,
			recipient,
			tickLower,
			tickUpper,
			amount0Requested,
			amount1Requested,
		)

		if err != nil {
			panic(err)
		}
		fmt.Println("collect tx sent:", tx.Hash().Hex())

	}

}

// first digit is n=0
func nthDigit(i int64, n int64) int64 {
	var quotient big.Int
	quotient.Exp(big.NewInt(10), big.NewInt(n), nil)

	bigI := big.NewInt(i)
	bigI.Div(bigI, &quotient)

	var result big.Int
	result.Mod(bigI, big.NewInt(10))

	return result.Int64()
}

func nonceGen(client bind.ContractBackend, account common.Address) *big.Int {

	nonce, err := client.PendingNonceAt(context.Background(), account)

	if err != nil {
		panic(err)
	}

	return big.NewInt(int64(nonce))

}

func floor(n *big.Rat) *big.Rat {
	f := &big.Rat{}
	z := new(big.Int)
	z.Div(n.Num(), n.Denom())
	f.SetInt(z)
	return f
}

///  return big.NewInt(10 * 1e18) //10 * 10**18
func x1E18(x int64) *big.Int {

	e18, _ := new(big.Int).SetString("1000000000000000000", 10)
	bigx := big.NewInt(x)

	return bigx.Mul(bigx, e18)
}

func readstring(msg string) string {

	fmt.Println("moving next step...........wait for 10 sec")
	time.Sleep(10 * time.Second)
	return "time up"

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(msg)
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.TrimSuffix(strings.TrimSpace(text), " \n")

		return text

	}

}
