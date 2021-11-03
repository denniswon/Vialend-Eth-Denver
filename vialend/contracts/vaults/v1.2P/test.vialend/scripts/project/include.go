package include

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"../config"
)

var Quiet = false
var Auto = true //auto check pending status

func myPrintln(a ...interface{}) {
	if !Quiet {
		fmt.Println(a)
	}
}

func Readstring(msg string) string {

	fmt.Println(msg)

	if Auto {
		time.Sleep(config.Network.PendingTime * time.Second)
		return ""
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			text = strings.TrimSuffix(strings.TrimSpace(text), " \n")

			return text

		}
	}
}
