package cmd

import (
	"context"
	"github.com/google/go-containerregistry/pkg/authn"
	"os"
	"shumyk/kdeploy/cmd/model"

	util "shumyk/kdeploy/cmd/util"

	"github.com/spf13/viper"
)

var conf config

var (
	ctx  = context.Background()
	auth = authn.DefaultKeychain
)

// TODO: add gcr url & path, etc
type config struct {
	StatefulSets []string
	Previous
}

type Previous map[string]model.PreviousImages

func (previous Previous) Keys() []string {
	keyMapping := util.ReturnKey[string, model.PreviousImages]
	return util.MapToSliceMapping(previous, keyMapping)
}

func InitConfig() {
	home, err := os.UserHomeDir()
	util.Laugh(err)

	viper.AddConfigPath(home)
	viper.SetConfigName(".kdeploy")
	viper.SetConfigType("yaml")

	util.Laugh(viper.SafeWriteConfig())
	util.Laugh(viper.ReadInConfig())
	util.Laugh(viper.Unmarshal(&conf))
}

func SaveConfig(key string, value any) {
	viper.Set(key, value)
	util.Laugh(viper.WriteConfig())
}

func GetPrevious() Previous {
	if conf.Previous == nil {
		conf.Previous = make(map[string]model.PreviousImages)
	}
	return conf.Previous
}
