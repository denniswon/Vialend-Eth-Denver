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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	callee "../../deploy/TestUniswapV3Callee"
)

type nClient struct {
	clientUrl string
	Contract  string
}

func main() {

	networks := [...]nClient{

		{ // 0 local
			"ws://localhost:7545",
			"0xF582510cBCa56C79AEaFE6FFb984eA80519753d1",
		},

		{ // 1 test admin
			"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"0xCe1CC3cC667D5B8925962B6a8F8F997aF018A7AC",
		},

		{ //2 test user 1
			"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"0xB0eeD5760749E2B05db16131bc44c55b5E3fE2b5",
		},

		{"https://rinkeby.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"0x623220D8e2C9bEf6cEb22B1895CA84fb38FC36ec",
		},
	}

	nid := 0

	fmt.Println("Current network using: ", networks[nid].clientUrl)
	client, err := ethclient.Dial(networks[nid].clientUrl)
	contract := networks[nid].Contract

	if err != nil {
		log.Fatal("eth dial ERR:", err)

	}

	contractAddress := common.HexToAddress(contract)

	contractAddress = common.HexToAddress("0xF582510cBCa56C79AEaFE6FFb984eA80519753d1")
	//CallEvent(client, contractAddress, "SwapCallback", 500, 500)

	contractAddress = common.HexToAddress("0x606B7f7093aa4df2282763F4e9f714706221838b")
	//	CallEvent(client, contractAddress, "SwapCallback", 502, 502)

	// pool: 0xC7bf0Acc3Fe1D1dc89e06056CeD28a0D0E7d0E26
	contractAddress = common.HexToAddress("0xC7bf0Acc3Fe1D1dc89e06056CeD28a0D0E7d0E26")
	//644 CollectFees  0xC7bf0Acc3Fe1D1dc89e06056CeD28a0D0E7d0E26
	//CallEvent(client, contractAddress, "CollectFees", 644, 644)
	CallEvent(client, contractAddress, "CollectFees", 681, 681)

	// callee: 0xE3c433a67e56BD49d93cCA86728C07bE531c2DCc
	contractAddress = common.HexToAddress("0xE3c433a67e56BD49d93cCA86728C07bE531c2DCc")
	//CallEvent(client, contractAddress, "MintCallback", 676, 680)

}

func CallEvent(client *ethclient.Client, contractAddress common.Address, eventName string, block_begin int64, block_end int64) {

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(block_begin),
		ToBlock:   big.NewInt(block_end),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal("filterlogs ERR:", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(callee.ApiABI)))
	if err != nil {
		log.Fatal("abi.JSON ERR:", err)
	}

	fmt.Println("event name: ", eventName)
	x := uint(0)
	for _, vLog := range logs {
		x++
		fmt.Println("blockhash:", vLog.BlockHash.Hex())
		fmt.Println("blocknumber:", vLog.BlockNumber)
		fmt.Println("txhash:", vLog.TxHash.Hex())

		/// READ TOPICS
		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
			fmt.Println("topic[", i, "]: ", topics[i])
		}

		/// READ EVENT LOG
		ev, err := contractAbi.Unpack(eventName, vLog.Data)
		if err != nil {
			log.Fatal("unpack ERR:", err)
		}

		//Print event data
		fmt.Printf("Unpack event %d: %#v\n", x, ev)

		//readstring("Next.....")

		///read into key - value
		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}

		fmt.Println(string(event.Key[:]))   // foo
		fmt.Println(string(event.Value[:])) // bar

		err = contractAbi.UnpackIntoInterface(&event, eventName, vLog.Data)
		if err != nil {
			//log.Fatal("unpackIntointerface ERR:", err)
			fmt.Println("unpackintointerface ERR:", err)
		}

		fmt.Println(string(event.Key[:]))   // foo
		fmt.Println(string(event.Value[:])) // bar

		fmt.Println("<< \n")
	}

	if x > 0 {
		eventSignature := []byte(eventName + "(bytes32,bytes32)")
		hash := crypto.Keccak256Hash(eventSignature)
		fmt.Println(eventName, " hash: ", hash.Hex())
	}

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
