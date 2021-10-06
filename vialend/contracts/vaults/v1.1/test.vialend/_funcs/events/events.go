package main

import (
	project "../../scripts/project"
)

func main() {

	// NOTE: make sure vault address in network.go is up to date

	//eventname: MyLog, Mylog2, Withdraw, Deposit
	eventname := "Deposit"
	block0 := 5606572

	block1 := 5606575
	project.VaultEvent(eventname, int64(block0), int64(block1))

}

// withdraw 1%	   129684769397849 	 74398713
// withdraw 100% 22841308326813345 7366472947

//1%    [_burnLendingShare   46766727])[]interface {}  {"_burnLendingShare ", 46766727}
//100% [_burnLendingShare  4629906058])[]interface {}{"_burnLendingShare ", 4629906058
