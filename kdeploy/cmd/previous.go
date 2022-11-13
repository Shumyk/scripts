package cmd

import (
	"time"

	"github.com/spf13/viper"
)

type config struct {
	Previous map[string][]PrevImage
}

type PrevImage struct {
	Tag      string
	Digest   string
	Deployed time.Time
}

func PrevImageOf(tag, digest string) PrevImage {
	return PrevImage{
		Tag:      tag,
		Digest:   digest,
		Deployed: time.Now(),
	}
}

func SavePreviouslyDeployed(tag, digest string) {
	prevImage := PrevImageOf(tag, digest)

	var conf config
	viper.Unmarshal(&conf)
	if conf.Previous == nil {
		conf.Previous = make(map[string][]PrevImage)
	}

	conf.Previous[microservice] = append(conf.Previous[microservice], prevImage)
	viper.Set("previous", conf.Previous)

	viper.WriteConfig()
}
