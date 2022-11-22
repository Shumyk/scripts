package cmd

import (
	. "shumyk/kdeploy/cmd/model"
	. "shumyk/kdeploy/cmd/util"
)

// Configuration TODO: add gcr url & path, etc
type Configuration struct {
	StatefulSets []string
	Previous     PreviousDeployments
}

type PreviousDeployments map[string]PreviousImages

func (previous PreviousDeployments) Keys() []string {
	keyMapping := ReturnKey[string, PreviousImages]
	return MapToSliceMapping(previous, keyMapping)
}
