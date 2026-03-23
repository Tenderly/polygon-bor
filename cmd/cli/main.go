package main

import (
	"os"

	"github.com/tenderly/polygon-bor/internal/cli"
	"github.com/tenderly/polygon-bor/params"
)

func main() {
	params.UpdateBorInfo()
	os.Exit(cli.Run(os.Args[1:]))
}
