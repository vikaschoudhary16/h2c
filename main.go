package main

import (
	"fmt"
	"os"

	"github.com/fstab/h2c/cli"
)

func main() {
	msg, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		os.Exit(-1)
	} else if msg != "" {
		fmt.Println(msg)
	}
}
