package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	microservice string
	previous     bool

	kdeploy = cobra.Command{
		Use:   "kdeploy microservice",
		Short: "k[8s]deploy - deploy from the terminal",
		Run:   run,
		Args:  cobra.ExactArgs(1),
	}
)

func run(cmd *cobra.Command, args []string) {
	microservice = args[0]
	if previous {
		fmt.Println("deploy previous")
	} else {
		KDeploy()
	}
}

func Execute() {
	if err := kdeploy.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	kdeploy.Flags().BoolVarP(&previous, "previous", "p", false, "deploy previous")
}
