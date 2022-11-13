package cmd

import (
	"os"

	prompt "shumyk/kdeploy/cmd/prompt"
	print "shumyk/kdeploy/cmd/util"

	"github.com/spf13/viper"
)

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

	kubeConfig := resolveKubeConfig()
	go Metadata(kubeConfig)
	go ClientSet(kubeConfig, make(chan bool))

	selectedImage := prompt.PromptPrevImageSelect(previous)
	print.Purple(selectedImage.Digest, selectedImage.Tags)
	SetImage(&selectedImage)
}

func SavePreviouslyDeployed(tag, digest string) {
	prevImage := prompt.PrevImageOf(tag, digest)
	previous := getPrevious()

	previous[microservice] = append(previous[microservice], prevImage)
	viper.Set("previous", previous)
	viper.WriteConfig()
}

func getPrevious() Previous {
	var conf config
	viper.Unmarshal(&conf)
	if conf.Previous == nil {
		conf.Previous = make(map[string][]prompt.PrevImage)
	}
	return conf.Previous
}
