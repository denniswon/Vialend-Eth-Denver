package main

import (
	project "../../scripts/project"
)

func main() {

	// NOTE: make sure vault address in network.go is up to date

	//eventname: MyLog, MyLog2, Withdraw, Deposit
	eventname := "MyLog"
	block0 := 5802871
	block1 := 5802871

	project.VaultEvent(eventname, int64(block0), int64(block1))
	//project.TestCalleeEvent(eventname, int64(block0), int64(block1))
}

// withdraw 1%	   129684769397849 	 74398713
// withdraw 100% 22841308326813345 7366472947

//1%    [_burnLendingShare   46766727])[]interface {}  {"_burnLendingShare ", 46766727}
//100% [_burnLendingShare  4629906058])[]interface {}{"_burnLendingShare ", 4629906058
