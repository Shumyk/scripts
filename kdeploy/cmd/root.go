package cmd

import (
	"fmt"
	"os"
	cmd "shumyk/kdeploy/cmd/prompt"
	"time"

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
		Args:  cobra.ExactArgs(1),
	}
)

func run(cmd *cobra.Command, args []string) {
	microservice = args[0]
	if previous {
		fmt.Println("deploy previous")
	} else {
		deployedImage := KDeploy()
		SavePreviouslyDeployed(PrevImageOf(deployedImage))
	}
}

type config struct {
	Previous map[string][]PrevImage
}

type PrevImage struct {
	Tag      string
	Digest   string
	Deployed time.Time
}

func PrevImageOf(i cmd.SelectedImage) PrevImage {
	return PrevImage{
		Tag:      i.Tags[0],
		Digest:   i.Digest,
		Deployed: time.Now(),
	}
}

func SavePreviouslyDeployed(i PrevImage) {
	var conf config
	viper.Unmarshal(&conf)
	if conf.Previous == nil {
		conf.Previous = make(map[string][]PrevImage)
	}

	conf.Previous[microservice] = append(conf.Previous[microservice], i)
	viper.Set("previous", conf.Previous)

	viper.WriteConfig()
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

	if err = viper.ReadInConfig(); err == nil {
		fmt.Println("using config file:", viper.ConfigFileUsed())
	}
}
