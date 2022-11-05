package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var kdeploy = cobra.Command{
	Use:   "kdeploy",
	Short: "k[8s]deploy - deploy from the terminal",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("kdeploy yay")
	},
}

func Execute() {
	if err := kdeploy.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
