package include

/*
DeployWETH
Test_weth_deposit   wrap
Test_weth_withdraw	unwrap

TokenTransfer	erc20

*/
import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	weth "viaroot/deploy/Tokens/erc20/deploy/WETH9"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

var WETH = Network.LendingContracts.WETH

func DeployWETH() {
	DeployWrappedEther()
}

func Wrap(WETH string, accId int, amt int64) {

	fmt.Println("Env: NetworkId=", Networkid, ",EthClient=", Network.ProviderUrl[ProviderSortId])

	fmt.Println("----------------------------------------------")
	fmt.Println(".........wrap .........  ")
	fmt.Println("----------------------------------------------")
	fmt.Println("WETH ADDRESS:", WETH)

	ChangeAccount(accId)

	instance := GetWethInstance(common.HexToAddress(WETH))

	//weth deposit
	ethAmount := X1E18(amt)

	Auth.Value = ethAmount

	NonceGen()
	tx, err := instance.Deposit(Auth) // value is eth

	if err != nil {
		log.Fatal("weth deposit err ", err)
	}

	fmt.Println("wrapped eth amount:", amt, " to: ", FromAddress)

	TxConfirm(tx.Hash())

	ChangeAccount(Account)

}

func UnWrap(WETH string, accId int, amt int64) {

	fmt.Println("Env: NetworkId=", Networkid, ",EthClient=", Network.ProviderUrl[ProviderSortId])
	fmt.Println("----------------------------------------------")
	fmt.Println(".........unwrap .........  ")
	fmt.Println("----------------------------------------------")
	fmt.Println("WETH ADDRESS:", WETH)

	ChangeAccount(accId)

	instance := GetWethInstance(common.HexToAddress(WETH))

	//weth deposit
	ethAmount := X1E18(amt)

	tx, err := instance.Withdraw(Auth, ethAmount)

	if err != nil {
		log.Fatal("weth withdraw err ", err)
	}

	fmt.Println("weth withdraw 1 eth tx: ", tx.Hash().Hex())
	fmt.Println("unwrapped weth amount:", amt, " to: ", FromAddress)

	ChangeAccount(Account)

}

func GetWethInstance(WETH common.Address) *weth.Api {

	instance, err := weth.NewApi(WETH, EthClient)
	if err != nil {
		log.Fatal("get token Instance,", err)
	}

	return instance
}

// AccountId, from and sign tx
func TokenTransfer(AccountId int, amount *big.Int, _tokenAddress string, _toAddress string) {

	transferFnSignature := []byte("transfer(address,uint256)")

	tokenAddress := common.HexToAddress(_tokenAddress)
	toAddress := common.HexToAddress(_toAddress)

	hash := sha3.NewLegacyKeccak256() // old sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	privateKey, err := crypto.HexToECDSA(Network.PrivateKey[AccountId])

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	gasLimit := uint64(345607)
	gasPrice := big.NewInt(3100000000)

	nonce, err := EthClient.PendingNonceAt(context.Background(), fromAddress)
	value := big.NewInt(0)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := EthClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("sent tx: %s\n", signedTx.Hash().Hex()) // tx
	//TxConfirm(signedTx.Hash())

}

func TransferEth(_fromKey string, value *big.Int, _toAddress string) {

	EthClient := EthClient
	privateKey, err := crypto.HexToECDSA(_fromKey)
	//privateKey, err := crypto.HexToECDSA(Network.PrivateKey[1])

	if err != nil {
		log.Fatal("privatekey err", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := EthClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("nonce err", err)
	}

	gasLimit := uint64(21000) // in units
	gasPrice, err := EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("gasprice err", err)
	}

	toAddress := common.HexToAddress(_toAddress)
	var data []byte
	value = new(big.Int).Sub(value, new(big.Int).Mul(gasPrice, big.NewInt(int64(gasLimit))))
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := EthClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal("chain error", err)
	}

	fmt.Println("chainID:", chainID)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal("signed err:", err)
	}

	err = EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Println("sendtransaction err:", err)
	}

	fmt.Println("eth for ", value, " has been sent. %s", signedTx.Hash().Hex())
	//TxConfirm(signedTx.Hash())

}

func getBalance(tokenAddress string, tokenHolder string) (string, *big.Int) {

	//cInstance := GetCTokenInstance(cTokenAddress)
	instance, _, symbol, _, _ := GetTokenInstance(tokenAddress)

	//	bal, err := cDaiInstance.BalanceOf(&bind.CallOpts{}, FromAddress)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(tokenHolder))

	if err != nil {
		log.Fatal("getTokenBalance err ", err)
	}

	Sleep(100)
	return symbol, bal

}

func SendTestTokens() {
	env := strings.Split(os.Getenv("VIALEND_COLLECTOR"), ":")
	collector := env[0]
	collectorkey := env[1]

	tokens := []string{
		Network.LendingContracts.WETH,
		Network.LendingContracts.WBTC,
		Network.LendingContracts.USDC,
		Network.LendingContracts.DAI,
	}

	amounts := []int{
		2,
		2,
		200,
		2000,
	}

	recipients := []string{
		"0x8ee95fe2DB1e3f7FAACCdEd1cbCc237267EB4a00",
		"0x5D2Dd8F38a74294DD1FAF7B8c806446A3379f330",
		"0x80f80F7D8E947031a583054090955E5DE431d6D3",
		"0x4336B32C22dae8E207c93A4d105467F845AD3e8d",
	}

	for i, recipient := range recipients {
		myPrintln(i, "--", recipient)

		//send ether
		//TransferEth(collectorkey, big.NewInt(1e18), recipient)
		//Sleep(5000)

		//send erc20
		for j, token := range tokens {
			if token == "" {
				continue
			}
			myPrintln(j, "--", token)
			TokenTransferAmount(collectorkey, collector, recipient, token, amounts[j])
			Sleep(5000)
		}
		//break
	}
}

func CollectTokens() {
	env := strings.Split(os.Getenv("VIALEND_COLLECTOR"), ":")
	collector := env[0]
	collectorkey := env[1]

	tokens := []string{
		Network.LendingContracts.WETH,
		Network.LendingContracts.WBTC,
		Network.LendingContracts.USDC,
		Network.LendingContracts.DAI,
		Network.LendingContracts.USDT,
	}

	froms := []string{
		"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7", //0,  admin 	0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E
		"284b65567176c10bc010345042b1d9852fcc1c42ae4b76317e6da040318fbe7f", //1,  admin 1  0x6dd19aEB91d1f43C46f0DD74C9E8A92BFe2a3Cd0
		"d8cda34b6928af75aff58c60fe9ed3339896b57a13fa88695aa6da7b775cda2a", //2,  admin 2  0xD8Dbe65b64428464fFa14DEAbe288b83C240e713
		"2d9e2b4c955159dd8a22faf3cb3074f03cfc182213729224915921daabaa5d6a", //3, team			0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440
		"01e8c8df56230b8b6e4ce6371bed124f4f9950c51d64adc581938239724ed5e6", //4,  user 1	0x14792757D21e54453179376c849662dE341797F2
		"67f7046a9f3712d77dab07a843c91d060ab5f27b808ed54d6db1293c7cd5eff3", //5,  user 2	0x4F211267896C4D3f2388025263AC6BD67B0f2C54
		"a830f08514d29b0d278b251773b2265cd462e02ad14ca016591929d42fb203d1", //6 arb01 0x8a01C3E04798D0B6D7423EaFF171932943FB9A8D
	}

	for i, key := range froms {
		_, publicAddress := GetSignatureByKey(key)
		myPrintln(i, "--", publicAddress)

		_ = collectorkey
		//TransferEth(collectorkey, big.NewInt(1e17), publicAddress)
		//Sleep(5000)

		for j, token := range tokens {
			if token != "" {
				myPrintln("--", j, "----", token)
				//TokenTransferAll(key, publicAddress, collector, token)
				balance, _ := EthClient.BalanceAt(context.Background(), common.HexToAddress(publicAddress), nil)
				if balance.Cmp(big.NewInt(0)) > 0 {
					TransferEth(key, balance, collector)
					Sleep(3000)
				}

			}
		}

		// Sleep(3000)
		// balance, _ := EthClient.BalanceAt(context.Background(), common.HexToAddress(publicAddress), nil)
		// myPrintln("eth balance left:", balance)

		// TransferEth(key, balance, collector)

		// Sleep(3000)

		//break
	}

}

// AccountId, from and sign tx
func TokenTransferAll(_key string, _from string, _toAddress string, _tokenAddress string) {
	tokenIns, _, _, _, _ := GetTokenInstance(_tokenAddress)
	balance, _ := tokenIns.BalanceOf(&bind.CallOpts{}, common.HexToAddress(_from))

	if balance.Cmp(big.NewInt(0)) == 0 {
		return
	}

	amount := balance
	myPrintln("balance:", balance)

	transferFnSignature := []byte("transfer(address,uint256)")

	tokenAddress := common.HexToAddress(_tokenAddress)
	toAddress := common.HexToAddress(_toAddress)

	hash := sha3.NewLegacyKeccak256() // old sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	privateKey, err := crypto.HexToECDSA(_key)

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	gasLimit := uint64(61000)
	gasPrice := big.NewInt(2 * 1e9) //big.NewInt(2100000000)

	//gasLimit := uint64(21000) // in units
	//gasPrice, err := EthClient.SuggestGasPrice(context.Background())

	nonce, err := EthClient.PendingNonceAt(context.Background(), fromAddress)
	value := big.NewInt(0)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := EthClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {

		log.Println("sending error:", err.Error())

	} else {

		fmt.Printf("sent tx: %s\n", signedTx.Hash().Hex()) // tx
	}

}

func TokenTransferAmount(_key string, _from string, _toAddress string, _tokenAddress string, _amount int) {

	tokenIns, _, _, decimals, _ := GetTokenInstance(_tokenAddress)
	balance, _ := tokenIns.BalanceOf(&bind.CallOpts{}, common.HexToAddress(_from))

	if balance.Cmp(big.NewInt(0)) == 0 {
		return
	}

	amount := PowX(int64(_amount), int(decimals))

	if balance.Cmp(amount) < 0 {
		myPrintln("fund is not enough", balance, amount)
		return
	}
	transferFnSignature := []byte("transfer(address,uint256)")

	tokenAddress := common.HexToAddress(_tokenAddress)
	toAddress := common.HexToAddress(_toAddress)

	hash := sha3.NewLegacyKeccak256() // old sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	privateKey, err := crypto.HexToECDSA(_key)

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	gasLimit := uint64(61000)
	gasPrice := big.NewInt(2 * 1e9) //big.NewInt(2100000000)

	//gasLimit := uint64(21000) // in units
	//gasPrice, err := EthClient.SuggestGasPrice(context.Background())

	nonce, err := EthClient.PendingNonceAt(context.Background(), fromAddress)
	value := big.NewInt(0)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := EthClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {

		log.Println("sending error:", err.Error())

	} else {

		fmt.Printf("sent tx: %s\n", signedTx.Hash().Hex()) // tx
	}

}
