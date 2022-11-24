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
	LoadConfiguration()
	validateVitalConfigs(cobra)
}

func LoadConfiguration() {
	home, err := os.UserHomeDir()
	Laugh(err)

	viper.AddConfigPath(home)
	viper.SetConfigName(".kdeploy")
	viper.SetConfigType("yaml")

	_ = viper.SafeWriteConfig()
	Laugh(viper.ReadInConfig())
	Laugh(viper.Unmarshal(&config))
}

func validateVitalConfigs(cobra *cobra.Command) {
	if len(config.Registry) == 0 {
		// TODO: stderr
		RedStderr("Registry not found in " + viper.ConfigFileUsed())
		BoringStderr("You can add it using:")
		BoringStderr("	kdeploy config set registry us.gcr.io")
		BoringStderr("Or manually editing " + viper.ConfigFileUsed())
		BoringStderr("	vim " + viper.ConfigFileUsed())
		os.Exit(1)
	}
	if len(config.Repository) == 0 {
		RedStderr("Repository not found in " + viper.ConfigFileUsed())
		BoringStderr("You can add it using:")
		BoringStderr("	kdeploy config set repository company-infra/company-")
		BoringStderr("Or manually editing " + viper.ConfigFileUsed())
		BoringStderr("	vim " + viper.ConfigFileUsed())
		os.Exit(1)
	}
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

func ResolveResourceType(service string) string {
	for _, statefulSet := range config.StatefulSets {
		if statefulSet == service {
			return "statefulsets"
		}
	}
	return "deployments"
}
