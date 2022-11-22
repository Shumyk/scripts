package cmd

import (
	"os"
	"shumyk/kdeploy/cmd/model"

	. "shumyk/kdeploy/cmd/util"

	"github.com/spf13/viper"
)

var conf config

// TODO: add gcr url & path, etc
type config struct {
	StatefulSets []string
	Previous
}

type Previous map[string]model.PreviousImages

func (previous Previous) Keys() []string {
	keyMapping := ReturnKey[string, model.PreviousImages]
	return MapToSliceMapping(previous, keyMapping)
}

func InitConfig() {
	home, err := os.UserHomeDir()
	Laugh(err)

	viper.AddConfigPath(home)
	viper.SetConfigName(".kdeploy")
	viper.SetConfigType("yaml")

	Laugh(viper.SafeWriteConfig())
	Laugh(viper.ReadInConfig())
	Laugh(viper.Unmarshal(&conf))
}

func SavePreviouslyDeployed(tag, digest string) {
	prevImage := model.PrevImageOf(tag, digest)
	previous := GetPrevious()

	previous[microservice] = append(previous[microservice], prevImage)
	SaveConfig("previous", previous)
}

func SaveConfig(key string, value any) {
	viper.Set(key, value)
	Laugh(viper.WriteConfig())
}

func GetPrevious() Previous {
	if conf.Previous == nil {
		conf.Previous = make(map[string]model.PreviousImages)
	}
	return conf.Previous
}
