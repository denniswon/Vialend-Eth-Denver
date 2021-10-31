package include

import (
	"fmt"
)

var Quiet = false

func myPrintln(a ...interface{}) {
	if !Quiet {
		fmt.Println(a)
	}
}
