package cmd

import (
	"os"

	prompt "shumyk/kdeploy/cmd/prompt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var conf config

// TODO: add statefulsets, gcr url & path, etc
type config struct {
	Previous
}

type Previous map[string][]prompt.PrevImage

func (p Previous) Keys() []string {
	keys := make([]string, len(p))
	position := 0
	for key := range p {
		keys[position] = key
		position++
	}
	return keys
}

func InitConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".kdeploy")

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
		conf.Previous = make(map[string][]prompt.PrevImage)
	}
	return conf.Previous
}
