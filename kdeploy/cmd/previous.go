package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	// kdeploy.AddCommand(previousCmd)
}

var previousCmd = &cobra.Command{
	Use:   "previous",
	Short: "deploy-previous mode",
	Long: `Quickly redeploy what was before your last deployment.
However, it has goldfish memory - can redeploy only the previous deployment.`,
	Aliases: []string{"p"},
	Run:     runPrevious,
}

func runPrevious(cmd *cobra.Command, args []string) {
	fmt.Fprintln(os.Stdout, "deploying previous")
}
