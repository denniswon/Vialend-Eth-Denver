package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	_ "time"

	project "viaroot/scripts/project"
)

func main() {
	project.Init(-1, -1)

	project.Quiet = true

	iteration := -1 //-1 infinite

	nid := project.Networkid
	acc := 1

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

	rng := 600 // check real range

	project.MonitorVault(nid, acc, iteration, rng)

	//project.MonitorVault(4, 1, iteration, 1000)
	//project.MonitorVault(4, 1, iteration, 1000)

}
