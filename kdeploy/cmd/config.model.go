package cmd

import (
	. "shumyk/kdeploy/cmd/model"
	. "shumyk/kdeploy/cmd/util"
)

type configuration struct {
	Registry     string
	Repository   string
	StatefulSets []string
	Previous     PreviousDeployments
}

type PreviousDeployments map[string]PreviousImages

func (previous PreviousDeployments) Keys() []string {
	keyMapping := ReturnKey[string, PreviousImages]
	return MapToSliceMapping(previous, keyMapping)
}
