package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	_ "time"

	"../config"
	project "../project"
)

func main() {
	fmt.Println("Env: NetworkId=", config.Networkid, ",client=", config.Network.ProviderUrl[config.ProviderSortId])
	project.Init(-1, -1)

	project.Quiet = true

	iteration := -1 //-1 infinite

	nid := config.Networkid

	if len(os.Args) > 1 {
		if len(os.Args[1]) > 0 {
			_nid, err := strconv.Atoi(strings.TrimSpace(os.Args[1]))
			if err != nil {
				log.Fatal("argument err ", err)
			}
			nid = _nid

		}
	}

	if nid != 3 && nid != 4 {
		log.Fatal("Wrong networkid ", nid)
	}

	project.MonitorVault(nid, 1, iteration, 600)

	//project.MonitorVault(4, 1, iteration, 1000)
	//project.MonitorVault(4, 1, iteration, 1000)

}
