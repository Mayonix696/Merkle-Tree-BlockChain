package main

import (
	"os"

	"github.com/Mayonix696/blockchain/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
