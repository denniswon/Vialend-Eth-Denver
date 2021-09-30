package main

import (
	project "../../scripts/project"
)

func main() {

	eventname := "MyLog"
	block0 := 5584506
	block1 := 5584506
	project.VaultEvent(eventname, int64(block0), int64(block1))
	return
}
