package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	token "goblockchain/Practice/tokens/erc20/deploy/token"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type nClient struct {
	clientUrl    string
	TokenAddress string
	block0       int64
	block1       int64
}

// LogTransfer ..
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func main() {
	networks := [...]nClient{

		{"http://127.0.0.1:7545",
			"",
			0,
			0,
		},

		{"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"0xc4AFB13b10f7C49Af721860A188D6443D0fF8747",
			5216055, 5316229,
		},
		{"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"0x2aDEca523FbBF0937A9419924feAB607Bf599311",
			5315055, 5316229,
		},
	}

	nid := 2

	fmt.Println("Current network using: ", networks[nid].clientUrl)

	client, err := ethclient.Dial(networks[nid].clientUrl)

	if err != nil {
		panic(err)
	}

	// 0x Protocol (ZRX) token address
	//	contractAddress := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	contractAddress := common.HexToAddress(networks[nid].TokenAddress)
	fmt.Println("token address: ", contractAddress)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(networks[nid].block0),
		ToBlock:   big.NewInt(networks[nid].block1),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		panic(err)
	}

	fmt.Println(logs)
	contractAbi, err := abi.JSON(strings.NewReader(string(token.ApiABI)))
	if err != nil {
		panic(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")

			var transferEvent LogTransfer

			//err := contractAbi.Unpack(&transferEvent, "Transfer", vLog.Data)
			_, err := contractAbi.Unpack("Transfer", vLog.Data)
			if err != nil {
				panic(err)
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())

		case logApprovalSigHash.Hex():
			fmt.Printf("Log Name: Approval\n")

			var approvalEvent LogApproval

			//err := contractAbi.Unpack(&approvalEvent, "Approval", vLog.Data)
			_, err := contractAbi.Unpack("Approval", vLog.Data)
			if err != nil {
				panic(err)
			}

			approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
			fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())
		}

		fmt.Printf("\n\n")
	}
}
