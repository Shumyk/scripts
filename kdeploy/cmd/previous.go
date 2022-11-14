package cmd

import (
	"os"

	prompt "shumyk/kdeploy/cmd/prompt"
	print "shumyk/kdeploy/cmd/util"

	"github.com/spf13/viper"
)

var conf config

type config struct {
	Previous
}

type Previous map[string][]prompt.PrevImage

func KDeployPrev() {
	previous := getPrevious()[microservice]
	if len(previous) == 0 {
		print.Red("no available previous deployments of", microservice)
		os.Exit(1)
	}
	DeployPrevious(previous)
}

func KDeployRegistryPrev() {
	prev := getPrevious()
	repos := make([]string, 0, len(prev))
	for k := range prev {
		repos = append(repos, k)
	}
	selectedRepo := prompt.PromptRepo(repos)
	microservice = selectedRepo
	KDeployPrev()
}

func SavePreviouslyDeployed(tag, digest string) {
	prevImage := prompt.PrevImageOf(tag, digest)
	previous := getPrevious()

	previous[microservice] = append(previous[microservice], prevImage)
	viper.Set("previous", previous)
	viper.WriteConfig()
}

func getPrevious() Previous {
	if conf.Previous != nil {
		return conf.Previous
	}
	viper.Unmarshal(&conf)
	if conf.Previous == nil {
		conf.Previous = make(map[string][]prompt.PrevImage)
	}
	return conf.Previous
}
