package cmd

import (
	"github.com/spf13/cobra"
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

func run(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		if previous {
			KDeployPrevWithRegistry()
		} else {
			KDeployWithRegistry()
		}
	} else {
		microservice = args[0]
		if previous {
			KDeployPrev()
		} else {
			KDeploy()
		}
	}
}

func Execute() {
	err := kdeploy.Execute()
	cobra.CheckErr(err)
}

func init() {
	cobra.OnInitialize(InitConfig)
	kdeploy.Flags().BoolVarP(&previous, "previous", "p", false, "deploy previous")
}
