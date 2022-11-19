package cmd

import (
	"os"
	model "shumyk/kdeploy/cmd/model"

	util "shumyk/kdeploy/cmd/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var conf config

// TODO: add statefulsets, gcr url & path, etc
type config struct {
	Previous
}

type Previous map[string][]model.PreviousImage

func (previous Previous) Keys() []string {
	keyMapping := util.ReturnKey[string, []model.PreviousImage]
	return util.MapToSliceMapping(previous, keyMapping)
}

func InitConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".kdeploy")

	// TODO: wrap errors with descriptions and handling
	viper.SafeWriteConfig()
	viper.ReadInConfig()
	viper.Unmarshal(&conf)
}

func SaveConfig(key string, value any) {
	viper.Set(key, value)
	viper.WriteConfig()
}

func GetPrevious() Previous {
	if conf.Previous == nil {
		conf.Previous = make(map[string][]model.PreviousImage)
	}
	return conf.Previous
}
