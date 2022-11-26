package cmd

import (
	"fmt"
	"os"

	. "shumyk/kdeploy/cmd/model"
	. "shumyk/kdeploy/cmd/util"

	"github.com/spf13/viper"
)

var config configuration

func InitConfig() {
	LoadConfiguration()
	validateVitalConfigs()
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

func validateVitalConfigs() {
	//missingConfigs := make([]Entry, 0, 2)
	if len(config.Registry) == 0 {
		//missingConfigs = append(missingConfigs, EntryOf("registry", "us.gcr.io"))
		printNotFoundVitalConfigError("registry", "us.gcr.io")
	}
	if len(config.Repository) == 0 {
		printNotFoundVitalConfigError("repository", "company-infra/company-")
	}
}

func printNotFoundVitalConfigError(config, configValue string) {
	RedStderr(config, "not found in "+viper.ConfigFileUsed())
	BoringStderr("You can set it using:")
	BoringStderr(fmt.Sprintf("	kdeploy config set %v %v", config, configValue))
	BoringStderr("Or manually editing " + viper.ConfigFileUsed())
	BoringStderr("	vim " + viper.ConfigFileUsed())
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
