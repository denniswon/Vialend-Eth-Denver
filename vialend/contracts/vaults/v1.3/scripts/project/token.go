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

	weth "viaroot/deploy/Tokens/erc20/deploy/WETH9"

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

	fmt.Println("Env: NetworkId=", Networkid, ",client=", Network.ProviderUrl[ProviderSortId])

	fmt.Println("----------------------------------------------")
	fmt.Println(".........wrap .........  ")
	fmt.Println("----------------------------------------------")
	fmt.Println("WETH ADDRESS:", WETH)

	ChangeAccount(accId)

	instance := GetWethInstance(common.HexToAddress(WETH))

	//weth deposit
	ethAmount := X1E18(amt)

	Auth.Value = ethAmount

	tx, err := instance.Deposit(Auth) // value is eth

	if err != nil {
		log.Fatal("weth deposit err ", err)
	}

	fmt.Println("weth deposit 1 eth tx: ", tx.Hash().Hex())
	fmt.Println("wrapped eth amount:", amt, " to: ", FromAddress)

	ChangeAccount(Account)

}

func UnWrap(WETH string, accId int, amt int64) {

	fmt.Println("Env: NetworkId=", Networkid, ",client=", Network.ProviderUrl[ProviderSortId])
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

	instance, err := weth.NewApi(WETH, Client)
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

	nonce, err := Client.PendingNonceAt(context.Background(), fromAddress)
	value := big.NewInt(0)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("sent tx: %s\n", signedTx.Hash().Hex()) // tx

}

func TransferEth(_fromKey string, value *big.Int, _toAddress string) {

	client := Client
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
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("nonce err", err)
	}

	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("gasprice err", err)
	}

	toAddress := common.HexToAddress(_toAddress)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("chain error", err)
	}

	fmt.Println("chainID:", chainID)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal("signed err:", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("sendtransaction err", err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
