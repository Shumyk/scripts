package cmd

import (
	"github.com/spf13/cobra"
	. "shumyk/kdeploy/cmd/util"
)

var (
	previousMode bool

	// TODO: more info here needed
	kdeploy = cobra.Command{
		Use:   "kdeploy microservice",
		Short: "k[8s]deploy - deploy from the terminal",
		Run:   run,
		Args:  cobra.MaximumNArgs(1),
	}
)

func run(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		if previousMode {
			KDeployPreviousWithRegistry()
		} else {
			KDeployWithRegistry()
		}
	} else {
		microservice = args[0]
		if previousMode {
			KDeployPrevious()
		} else {
			KDeploy()
		}
	}
}

func Execute() {
	err := kdeploy.Execute()
	ErrorCheck(err, "Failed to execute kdeploy :|")
}

func init() {
	cobra.OnInitialize(InitConfig)
	kdeploy.Flags().BoolVarP(&previousMode, "previous", "p", false, "deploy previous")
}
