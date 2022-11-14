package cmd

import (
	printer "shumyk/kdeploy/cmd/util"
)

func init() {
	go printer.InitPrinter()
}

func KDeploy() {
	DeployNew()
}
