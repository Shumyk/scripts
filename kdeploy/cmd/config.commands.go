package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	cmd "shumyk/kdeploy/cmd/util"
)

func runConfig(cmd *cobra.Command, args []string) {
	fmt.Println("config command")
}

func runConfigView(_ *cobra.Command, _ []string) {
	viewBytes, err := yaml.Marshal(config.View())
	cmd.ErrorCheck(err, "Couldn't marshal config file")
	fmt.Println(string(viewBytes))
}

func runConfigSet(cmd *cobra.Command, args []string) {
	fmt.Println("config set command")
	fmt.Println(args[0], "=", args[1])
}

func runConfigEdit(cmd *cobra.Command, args []string) {
	fmt.Println("config edit command")
}
