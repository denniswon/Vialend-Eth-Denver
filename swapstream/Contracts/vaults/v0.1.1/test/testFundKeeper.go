package main

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	fundkeepergo "../deploy/FundKeeper"
)

type nClient struct {
	clientUrl          string
	fundkeeperContract string
	privateKey         string
	from               string
	to                 string
}

func main() {
	/* ///test go big string
	var prime1, _ = new(big.Int).SetString("21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)
	var prime2, _ = new(big.Int).SetString("21888242871839275222246405745257275088696311157297823662689037894645226208583", 10)

	var product = new(big.Int)
	product.Mul(prime1, prime2)

	fmt.Printf("Prime1: %v\n", prime1)
	fmt.Printf("Prime2: %v\n", prime2)
	fmt.Printf("Product: %v\n", product)
	return
	*/
	networks := [...]nClient{

		{ //0
			"http://127.0.0.1:7545",
			"0x0eE9560D0C3F101E897e82e7179E582Bf8D87a3C",
			"e8ef3a782d9002408f2ca6649b5f95b3e5772364a5abe203f1678817b6093ff0",
			"",
			"0x4F211267896C4D3f2388025263AC6BD67B0f2C54"},

		{ //1 test admin
			"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"0xeDaC99A7AE93F6EA3bc23b985553D77eEF7C0009",
			"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7",
			"",
			"0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E"},

		{ //2 test user 1
			"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"0xeDaC99A7AE93F6EA3bc23b985553D77eEF7C0009",
			"67f7046a9f3712d77dab07a843c91d060ab5f27b808ed54d6db1293c7cd5eff3",
			"",
			"0x4F211267896C4D3f2388025263AC6BD67B0f2C54"},

		{ //3 test user 2
			"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"0x553eE61a6AD3f5Fc5974EfA3DBfD54A33A5bb7E0",
			"01e8c8df56230b8b6e4ce6371bed124f4f9950c51d64adc581938239724ed5e6",
			"",
			"0x14792757D21e54453179376c849662dE341797F2"},

		{ //4
			"https://rinkeby.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"not yet",
			"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7",
			"",
			"0x4F211267896C4D3f2388025263AC6BD67B0f2C54"},
	}

	nid := 1
	doApprove := false

	doDeposit := false

	doWithDraw := false

	doWithDrawAll := false

	token0Address := common.HexToAddress("0xFA5dF5372c03D4968d128D624e3Afeb61031a777")
	token1Address := common.HexToAddress("0x3fF5E22B4be645EF1CCc8C6e32EDe6b35c569AE4")

	contractAddress := common.HexToAddress(networks[nid].fundkeeperContract)

	to := common.HexToAddress(networks[nid].to)

	privateKey, err := crypto.HexToECDSA(networks[nid].privateKey)

	fmt.Println("Current network: ", networks[nid].clientUrl)

	client, err := ethclient.Dial(networks[nid].clientUrl)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("from:", fromAddress)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Value = big.NewInt(0)                  // in wei
	auth.GasLimit = uint64(6721975)             // in units
	auth.GasPrice = big.NewInt(19 * 1000000000) // 1 gwei = 1000000000 wei

	fmt.Println("fundkeeper instance:", contractAddress)
	fmt.Println("token0 instance:", token0Address)
	fmt.Println("token1 instance:", token1Address)

	instance, err := fundkeepergo.NewApi(contractAddress, client)
	if err != nil {
		panic(err)
	}

	token0Instance, err := fundkeepergo.NewApi(token0Address, client)
	if err != nil {
		panic(err)
	}

	token1Instance, err := fundkeepergo.NewApi(token1Address, client)
	if err != nil {
		panic(err)
	}

	callfunc_greet(instance)
	callfunc_hello(instance)
	callfunc_plus(instance, big.NewInt(12), big.NewInt(2))

	myLiquidity(instance)

	var nonce uint64

	if doApprove {

		//new nonce for approve token0
		nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal("nonce:", err)
		}
		auth.Nonce = big.NewInt(int64(nonce))

		var maxToken0, _ = new(big.Int).SetString("90000000000000000000000", 10)
		tx0, err := token0Instance.Approve(auth, contractAddress, maxToken0)

		if err != nil {
			panic(err)
		}

		fmt.Println("approve amountoken0: %s", tx0.Hash().Hex())

		readstring("done... Press any key to continue when ready..........")

		// new nonce for approve token1
		nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}
		auth.Nonce = big.NewInt(int64(nonce))

		//var amountToken0, _ = new(big.Int).SetString("90000000000000000000000", 10)
		var maxToken1, _ = new(big.Int).SetString("90000000000000000000000", 10)

		tx1, err := token1Instance.Approve(auth, contractAddress, maxToken1)

		if err != nil {
			panic(err)
		}

		fmt.Println("approve amountoken1: %s", tx1.Hash().Hex())

		readstring("done... Press any key to continue when ready..........")

	}

	if doDeposit {

		// new nonce for deposit
		nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}
		auth.Nonce = big.NewInt(int64(nonce))

		//amountToken0 := big.NewInt(10 * 1000000000000000000)  //overflow error
		//amountToken1 := big.NewInt(200 * 1000000000000000000)  // overlow error

		//var amountToken0, _ = new(big.Int).SetString("100000000000000000000", 10)
		//var amountToken1, _ = new(big.Int).SetString("200000000000000000000", 10)
		amountToken0 := big.NewInt(1 * 1000000000000000000)
		amountToken1 := big.NewInt(2 * 1000000000000000000)
		amount0Min := big.NewInt(1)
		amount1Min := big.NewInt(1)

		callfunc_deposit(instance, auth,
			amountToken0, //
			amountToken1, //
			amount0Min,
			amount1Min,
			to,
		)
		readstring("done... Press any key to continue when ready..........")

	}
	if doWithDraw {

		// new nonce for withdraw
		nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}
		auth.Nonce = big.NewInt(int64(nonce))

		shares := big.NewInt(2000000000000000000)
		amount0Min := big.NewInt(1000000000000000000)
		amount1Min := big.NewInt(1000000000000000000)

		callfunc_withdraw(instance, auth,
			shares,
			amount0Min,
			amount1Min,
			to,
		)

		readstring("done... Press any key to continue when ready..........")

	}

	if doWithDrawAll {

		//amounts, _ := instance.GetTotalAmounts(&bind.CallOpts{})
		//get total amount of tokens in vault
		total0, total1 := getTotalAmounts(instance)

		//get total shares in vault
		totalSupply, _ := instance.TotalSupply(&bind.CallOpts{})

		// new nonce for withdraw
		nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}
		auth.Nonce = big.NewInt(int64(nonce))

		shares := totalSupply
		amount0Min := total0
		amount1Min := total1

		callfunc_withdraw(instance, auth,
			shares,
			amount0Min,
			amount1Min,
			to,
		)

		readstring("done... Press any key to continue when ready..........")

	}

	getTotalAmounts(instance)

}

func callfunc_greet(instance *fundkeepergo.Api) {

	message := "Hello eric"
	reply, err := instance.Greet(&bind.CallOpts{}, message)
	if err != nil {
		panic(err)
	}

	fmt.Println("greeting:", reply)

}

func getTotalAmounts(instance *fundkeepergo.Api) (*big.Int, *big.Int) {

	amounts, _ := instance.GetTotalAmounts(&bind.CallOpts{})

	fmt.Println("total amount: %s,%s", amounts.Total0, amounts.Total1)

	return amounts.Total0, amounts.Total1

}

func callfunc_hello(instance *fundkeepergo.Api) {

	reply, err := instance.Hello(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)

}

func callfunc_plus(instance *fundkeepergo.Api, x *big.Int, y *big.Int) {

	reply, err := instance.Plus(&bind.CallOpts{}, x, y)
	if err != nil {
		panic(err)
	}

	fmt.Println(x, "+", y, "=", reply)

}

func callfunc_deposit(instance *fundkeepergo.Api, auth *bind.TransactOpts,
	amount0Desired *big.Int,
	amount1Desired *big.Int,
	amount0Min *big.Int,
	amount1Min *big.Int,
	to common.Address,
) {

	tx, err := instance.Deposit(auth,
		amount0Desired,
		amount1Desired,
		amount0Min,
		amount1Min,
		to,
	)

	if err != nil {
		panic(err)
	}

	//get the transaction hash
	fmt.Println("deposit tx sent: %s", tx.Hash().Hex())

}

func callfunc_withdraw(instance *fundkeepergo.Api, auth *bind.TransactOpts,
	shares *big.Int,
	amount0Min *big.Int,
	amount1Min *big.Int,
	to common.Address,
) {

	tx, err := instance.Withdraw(auth,
		shares,
		amount0Min,
		amount1Min,
		to,
	)

	if err != nil {
		panic(err)
	}

	//get the transaction hash
	fmt.Println("withdraw tx sent: %s", tx.Hash().Hex())

}

func myLiquidity(instance *fundkeepergo.Api) {

	account := common.HexToAddress("0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E")

	myShares, err := instance.BalanceOf(&bind.CallOpts{}, account)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("my liquidity: %v\n", myShares)

}

func readstring(msg string) string {

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
