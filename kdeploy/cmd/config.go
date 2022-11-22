package cmd

import (
	"os"

	. "shumyk/kdeploy/cmd/model"
	. "shumyk/kdeploy/cmd/util"

	"github.com/spf13/viper"
)

var config configuration

func InitConfig() {
	home, err := os.UserHomeDir()
	Laugh(err)

	viper.AddConfigPath(home)
	viper.SetConfigName(".kdeploy")
	viper.SetConfigType("yaml")

	_ = viper.SafeWriteConfig()
	Laugh(viper.ReadInConfig())
	Laugh(viper.Unmarshal(&config))
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
	if config.previous == nil {
		config.previous = make(map[string]PreviousImages)
	}
	return config.previous
}

func Registry() string {
	return config.registry
}

func Repository() string {
	return config.repository
}

func BuildRepository(service string) string {
	return config.repository + service
}
