package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	_ "time"

	"../../config"
	project "../../project"
)

/*

usage:
	moveprice bull 1
	moveprice bear 1
*/
func main() {

	fmt.Println("Env: NetworkId=", config.Networkid, ",client=", config.Network.ProviderUrl[config.ProviderSortId])
	project.Init()

	project.Quiet = true

	accoundid := 0
	amount := int64(1)

	var op = 1 // buy

	var tn = 1

	//做price move 和 genfee 要 检查钱包是否有钱 否则FAIL

	if len(os.Args) > 1 {

		if os.Args[1] == "bull" || os.Args[1] == "up" || os.Args[1] == "buy" {
			op = 1
		} else if os.Args[1] == "bear" || os.Args[1] == "down" || os.Args[1] == "sell" {
			op = -1
		}

		if len(os.Args[2]) > 0 {
			amt, err := strconv.Atoi(strings.TrimSpace(os.Args[2]))
			if err != nil {
				log.Fatal("argument err ", err)
			}
			amount = int64(amt)

		}

	}

	project.MovePrice(accoundid, amount, op, tn)

}
