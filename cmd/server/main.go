package main

import (
	"github.com/RejwankabirHamim/money-transfer-workflow-iwf/cmd/server/iwf"
	"os"
)

func main() {
	app := iwf.BuildCLI()
	app.Run(os.Args)
}
