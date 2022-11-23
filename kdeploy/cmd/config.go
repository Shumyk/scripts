package cmd

import (
	"github.com/spf13/cobra"
	"os"

	. "shumyk/kdeploy/cmd/model"
	. "shumyk/kdeploy/cmd/util"

	"github.com/spf13/viper"
)

var config configuration

func InitConfig(cobra *cobra.Command) {
	home, err := os.UserHomeDir()
	Laugh(err)

	viper.AddConfigPath(home)
	viper.SetConfigName(".kdeploy")
	viper.SetConfigType("yaml")

	_ = viper.SafeWriteConfig()
	Laugh(viper.ReadInConfig())
	Laugh(viper.Unmarshal(&config))

	validateVitalConfigs(cobra)
}

func validateVitalConfigs(cobra *cobra.Command) {
	if len(config.Registry) == 0 {
		// TODO: stderr
		Red("Registry not found in " + viper.ConfigFileUsed())
		Boring("You can add it using:")
		Boring("	kdeploy config set registry us.gcr.io")
		Boring("Or manually editing " + viper.ConfigFileUsed())
		Boring("	vim " + viper.ConfigFileUsed())
		os.Exit(1)
	}
	if len(config.Repository) == 0 {
		Red("Repository not found in " + viper.ConfigFileUsed())
		Boring("You can add it using:")
		Boring("	kdeploy config set repository umbrella-infra/umbrella/umbrella")
		Boring("Or manually editing " + viper.ConfigFileUsed())
		Boring("	vim " + viper.ConfigFileUsed())
		os.Exit(1)
	}
	Boring("works")
	os.Exit(0)
	// TODO: implement for registry, repository
}

func SaveConfig(key string, value any) {
	viper.Set(key, value)
	Laugh(viper.WriteConfig())
}

func SaveDeployedImage(tag, digest string) {
	deployedImage := PrevImageOf(tag, digest)
	previous := GetPreviousDeployments()

	previous[microservice] = append(previous[microservice], deployedImage)
	SaveConfig("previous", previous)
}

func GetPreviousDeployments() PreviousDeployments {
	if config.Previous == nil {
		config.Previous = make(map[string]PreviousImages)
	}
	return config.Previous
}

func Registry() string {
	return config.Registry
}

func Repository() string {
	return config.Repository
}

func BuildRepository(service string) string {
	return config.Repository + service
}

func ResolveResourceType() string {
	// TODO: statefulsets from config
	statefulSets := map[string]any{"api-core": struct{}{}}
	if _, ok := statefulSets[microservice]; ok {
		k8sResource = "statefulsets"
	} else {
		k8sResource = "deployments"
	}
}
