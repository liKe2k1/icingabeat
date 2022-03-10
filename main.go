package main

import (
	"os"

	"github.com/like2k1/icingabeat/cmd"

	_ "github.com/like2k1/icingabeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
