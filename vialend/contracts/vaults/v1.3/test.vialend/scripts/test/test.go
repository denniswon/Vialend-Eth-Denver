package main

import (
	"fmt"

	project "../project"
)

func main() {
	c := project.Client

	fmt.Println("we have a connection")
	_ = c // we'll use this in the upcoming sections
}
