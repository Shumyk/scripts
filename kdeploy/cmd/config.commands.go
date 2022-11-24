package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	cmd "shumyk/kdeploy/cmd/util"
)

func runConfig(cmd *cobra.Command, args []string) {
	fmt.Println("config command")
}

func runConfigView(_ *cobra.Command, _ []string) {
	configFilePath := viper.GetViper().ConfigFileUsed()
	configFileBytes, err := os.ReadFile(configFilePath)
	cmd.ErrorCheck(err, "Couldn't read config file:", configFilePath)
	fmt.Println(string(configFileBytes))
}

func runConfigSet(cmd *cobra.Command, args []string) {
	fmt.Println("config set command")
	fmt.Println(args[0], "=", args[1])
}

func runConfigEdit(cmd *cobra.Command, args []string) {
	fmt.Println("config edit command")
}
