package cmd

import (
	"github.com/spf13/cobra"
	. "shumyk/kdeploy/cmd/util"
)

var (
	previousMode bool

	// TODO: more info here needed
	kdeploy = cobra.Command{
		Use:    "kdeploy [microservice]",
		Short:  "k[8s]deploy - deploy from the terminal",
		Run:    kdeployRun,
		Args:   cobra.MaximumNArgs(1),
		PreRun: InitConfig,
	}
)

func kdeployRun(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		deploySelectingRegistry()
	} else {
		deployMicroservice(args)
	}
}

func deploySelectingRegistry() {
	if previousMode {
		KDeployPreviousWithRegistry()
	} else {
		KDeployWithRegistry()
	}
}

func deployMicroservice(args []string) {
	microservice = args[0]
	if previousMode {
		KDeployPrevious()
	} else {
		KDeploy()
	}
}

func Execute() {
	err := kdeploy.Execute()
	ErrorCheck(err, "Failed to execute kdeploy :|")
}

func init() {
	kdeploy.Flags().BoolVarP(&previousMode, "previous", "p", false, "deploy previous")

	configCmd := cobra.Command{
		Use:              "config [action] [args]...",
		PersistentPreRun: LoadConfiguration,
	}
	configViewCmd := cobra.Command{
		Use:  "view",
		Run:  runConfigView,
		Args: cobra.NoArgs,
	}
	configEditCmd := cobra.Command{
		Use:  "edit",
		Run:  RunConfigEdit,
		Args: cobra.NoArgs,
	}
	configSetCmd := cobra.Command{
		Use: "set [property] [value]",
		// TODO: docs about setting array property types
		Run:  RunConfigSet,
		Args: cobra.ExactArgs(2),
	}
	kdeploy.AddCommand(&configCmd)
	configCmd.AddCommand(&configViewCmd, &configSetCmd, &configEditCmd)
}
