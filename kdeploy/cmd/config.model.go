package cmd

import (
	. "shumyk/kdeploy/cmd/model"
	. "shumyk/kdeploy/cmd/util"
)

type configuration struct {
	registry     string
	repository   string
	statefulSets []string
	previous     PreviousDeployments
}

type PreviousDeployments map[string]PreviousImages

func (previous PreviousDeployments) Keys() []string {
	keyMapping := ReturnKey[string, PreviousImages]
	return MapToSliceMapping(previous, keyMapping)
}
