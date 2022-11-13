package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	microservice  string
	imageListSize int

	kdeploy = cobra.Command{
		Use:   "kdeploy microservice",
		Short: "k[8s]deploy - deploy from the terminal",
		Run:   run,
	}
)

func run(cmd *cobra.Command, args []string) {
	validateArgs(args)
	microservice = args[0]

	fmt.Fprintln(os.Stdout, "kdeploy", microservice, imageListSize)
	KDeploy()
}

func validateArgs(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "usage")
		os.Exit(1)
	}
}

func Execute() {
	if err := kdeploy.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	kdeploy.PersistentFlags().IntVarP(&imageListSize, "list", "l", 20, "amount of images to fetch")
}
