package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	. "shumyk/kdeploy/cmd/util"
)

var (
	previousMode bool

	// TODO: more info here needed
	kdeploy = cobra.Command{
		Use:   "kdeploy [microservice]",
		Short: "k[8s]deploy - deploy from the terminal",
		Run:   kdeployRun,
		Args:  cobra.MaximumNArgs(1),
	}
)

func kdeployRun(cmd *cobra.Command, args []string) {
	// TODO: remove when config commands finished
	fmt.Println("kdeploy main")
	InitConfig(cmd)
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
		Use: "config [action] [args]...",
		Run: runConfig,
	}
	configViewCmd := cobra.Command{
		Use:    "view",
		Run:    runConfigView,
		Args:   cobra.NoArgs,
		PreRun: loadConfig,
	}
	configEditCmd := cobra.Command{
		Use:  "edit",
		Run:  runConfigEdit,
		Args: cobra.NoArgs,
	}
	configSetCmd := cobra.Command{
		Use:  "set [property] [value]",
		Run:  runConfigSet,
		Args: cobra.ExactArgs(2),
	}
	kdeploy.AddCommand(&configCmd)
	configCmd.AddCommand(&configViewCmd, &configSetCmd, &configEditCmd)
}

func loadConfig(_ *cobra.Command, _ []string) {
	LoadConfiguration()
}
