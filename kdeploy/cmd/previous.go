package cmd

import (
	"shumyk/kdeploy/cmd/model"
	prompt "shumyk/kdeploy/cmd/prompt"
	util "shumyk/kdeploy/cmd/util"
)

func KDeployPrev() {
	previous := GetPrevious()[microservice]
	util.TerminateOnEmpty(previous, "no available previous deployments of", microservice)
	DeployPrevious(previous)
}

func KDeployPrevWithRegistry() {
	repos := GetPrevious().Keys()
	util.TerminateOnEmpty(repos, "no available previous deployments")

	selectedRepo := prompt.RepoSelect(repos)
	microservice = selectedRepo
	KDeployPrev()
}

func SavePreviouslyDeployed(tag, digest string) {
	prevImage := model.PrevImageOf(tag, digest)
	previous := GetPrevious()

	previous[microservice] = append(previous[microservice], prevImage)
	SaveConfig("previous", previous)
}
