package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	microservice string

	kdeploy = cobra.Command{
		Use:   "kdeploy microservice",
		Short: "k[8s]deploy - deploy from the terminal",
		Run:   run,
		Args:  cobra.ExactArgs(1),
	}
)

func run(cmd *cobra.Command, args []string) {
	microservice = args[0]

	KDeploy()
}

func Execute() {
	if err := kdeploy.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
