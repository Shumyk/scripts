package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func runConfig(cmd *cobra.Command, args []string) {
	fmt.Println("config command")
}

func runConfigView(_ *cobra.Command, _ []string) {
	fmt.Println("config view command")
}

func runConfigSet(cmd *cobra.Command, args []string) {
	fmt.Println("config set command")
	fmt.Println(args[0], "=", args[1])
}

func runConfigEdit(cmd *cobra.Command, args []string) {
	fmt.Println("config edit command")
}
