package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	microservice string
	previous     bool

	kdeploy = cobra.Command{
		Use:   "kdeploy microservice",
		Short: "k[8s]deploy - deploy from the terminal",
		Run:   run,
		Args:  cobra.MaximumNArgs(1),
	}
)

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		if previous {
			KDeployRegistryPrev()
		} else {
			KDeployRegistry()
		}
		return
	}
	microservice = args[0]
	if previous {
		KDeployPrev()
	} else {
		KDeploy()
	}
}

func Execute() {
	err := kdeploy.Execute()
	cobra.CheckErr(err)
}

func init() {
	cobra.OnInitialize(initConfig)
	kdeploy.Flags().BoolVarP(&previous, "previous", "p", false, "deploy previous")
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".kdeploy")

	viper.SafeWriteConfig()
	viper.ReadInConfig()
}
