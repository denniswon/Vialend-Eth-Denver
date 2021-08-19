package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	fundkeepergo "goblockchain/Practice/SwapStream/swapstream-v0-core/deploy/FundKeeper"
)

type nClient struct {
	clientUrl          string
	fundkeeperContract string
	privateKey         string
	from               string
	to                 string
}

func main() {

	networks := [...]nClient{

		{ // 0 local
			"http://127.0.0.1:7545",
			"0x0eE9560D0C3F101E897e82e7179E582Bf8D87a3C",
			"e8ef3a782d9002408f2ca6649b5f95b3e5772364a5abe203f1678817b6093ff0",
			"",
			"0x4F211267896C4D3f2388025263AC6BD67B0f2C54",
		},

		{ // 1 test admin
			"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"0xCe1CC3cC667D5B8925962B6a8F8F997aF018A7AC",
			"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7",
			"",
			"0x4F211267896C4D3f2388025263AC6BD67B0f2C54"},

		{"https://rinkeby.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"0x623220D8e2C9bEf6cEb22B1895CA84fb38FC36ec",
			"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7",
			"",
			"0x4F211267896C4D3f2388025263AC6BD67B0f2C54"},
	}

	blocks := [...][]int64{
		{55, 89},
		{5333538, 5333538},
		{0, 0},
	}

	nid := 1

	fmt.Println("Current network using: ", networks[nid].clientUrl)
	client, err := ethclient.Dial(networks[nid].clientUrl)
	fundkeeperContract := networks[nid].fundkeeperContract

	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(fundkeeperContract)

	CallEvent("Deposit", contractAddress, client, blocks[nid][0], blocks[nid][1])
	CallEvent("Withdraw", contractAddress, client, blocks[nid][0], blocks[nid][1])

}

func CallEvent(eventName string, contractAddress common.Address, client *ethclient.Client, block_begin int64, block_end int64) {

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(block_begin),
		ToBlock:   big.NewInt(block_end),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(fundkeepergo.ApiABI)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("calling ", eventName, " event:......")

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		fmt.Println(vLog.BlockNumber)     // 2394201
		fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		x, _ := contractAbi.Unpack(eventName, vLog.Data)

		//Print event data
		fmt.Printf("%#v\n", x)

		readstring("Next.....")

		/*
			event := struct {
				Key   [32]byte
				Value [32]byte
			}{}

			event, err := contractAbi.Unpack("ItemSet", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("event.Key[:]: ", string(event.Key[:]))     // foo
			fmt.Println("event.Value[:]: ", string(event.Value[:])) // bar

			var topics [4]string
			for i := range vLog.Topics {
				topics[i] = vLog.Topics[i].Hex()
			}

			fmt.Println("topic[0]: ", topics[0])
		*/
	}

	//		eventSignature := []byte("ItemSet(bytes32,bytes32)")
	//	hash := crypto.Keccak256Hash(eventSignature)
	//fmt.Println("itemset hash: ", hash.Hex())

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
