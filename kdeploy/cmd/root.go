package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	microservice  string
	imageListSize int = 20

	kdeploy = cobra.Command{
		Use:   "kdeploy microservice [image list size]",
		Short: "k[8s]deploy - deploy from the terminal",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Fprintln(os.Stderr, "usage")
				os.Exit(1)
			}
			microservice = args[0]
			if len(args) == 2 {
				imageListSize, _ = strconv.Atoi(args[1])
			}

			fmt.Fprintln(os.Stdout, "kdeploy", microservice, imageListSize)
		},
	}
)

func Execute() {
	if err := kdeploy.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
