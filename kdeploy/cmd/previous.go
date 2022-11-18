package cmd

import (
	prompt "shumyk/kdeploy/cmd/prompt"
	print "shumyk/kdeploy/cmd/util"
)

func KDeployPrev() {
	previous := GetPrevious()[microservice]
	terminateOnEmpty(previous, "no available previous deployments of", microservice)
	DeployPrevious(previous)
}

func KDeployPrevWithRegistry() {
	repos := GetPrevious().Keys()
	terminateOnEmpty(repos, "no available previous deployments")

	selectedRepo := prompt.PromptRepo(repos)
	microservice = selectedRepo
	KDeployPrev()
}

func terminateOnEmpty[T any](args []T, msg ...string) {
	if len(args) == 0 {
		print.Error(msg...)
	}
}

func SavePreviouslyDeployed(tag, digest string) {
	prevImage := prompt.PrevImageOf(tag, digest)
	previous := GetPrevious()

	previous[microservice] = append(previous[microservice], prevImage)
	SaveConfig("previous", previous)
}
