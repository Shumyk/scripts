package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	prompt "shumyk/kdeploy/cmd/prompt"

	. "shumyk/kdeploy/cmd/model"
	. "shumyk/kdeploy/cmd/util"

	"github.com/spf13/viper"
)

var config configuration

func InitConfig(_ *cobra.Command, _ []string) {
	LoadConfiguration(nil, nil)
	validateVitalConfigs()
}

func LoadConfiguration(_ *cobra.Command, _ []string) {
	home, err := os.UserHomeDir()
	Laugh(err)

	viper.AddConfigPath(home)
	viper.SetConfigName(".kdeploy")
	viper.SetConfigType("yaml")

	_ = viper.SafeWriteConfig()
	Laugh(viper.ReadInConfig())
	Laugh(viper.Unmarshal(&config))
}

func validateVitalConfigs() {
	if len(config.Registry) == 0 {
		promptAndSaveConfig("registry")
	}
	if len(config.Repository) == 0 {
		promptAndSaveConfig("repository")
	}
}

func promptAndSaveConfig(configName string) {
	RedStderr(configName, " not found in ", viper.ConfigFileUsed())
	configValue, err := prompt.TextInput(configName)
	if err != nil {
		printMissingConfigInfo(configName)
	}
	SetConfigHandling(configName, configValue)
}

func printMissingConfigInfo(config string) {
	BoringStderr("Looks like you ctrl-c input. However, you can set it using:")
	BoringStderr(fmt.Sprintf("	kdeploy config set %v <value>", config))
	BoringStderr("Or manually editing:")
	BoringStderr("	kdeploy config edit")
	os.Exit(1)
}

func SetConfig(key string, value any) error {
	viper.Set(key, value)
	return viper.WriteConfig()
}

func SetConfigHandling(key string, value any) {
	ErrorCheck(SetConfig(key, value), "Could not set config")
}

func SaveDeployedImage(tag, digest string) {
	deployedImage := PrevImageOf(tag, digest)
	previous := GetPreviousDeployments()

	previous[microservice] = append(previous[microservice], deployedImage)
	Laugh(SetConfig("previous", previous))
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
