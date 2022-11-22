package cmd

import (
	prompt "shumyk/kdeploy/cmd/prompt"
	. "shumyk/kdeploy/cmd/util"
)

func KDeployPrev() {
	previous := GetPrevious()[microservice]
	TerminateOnEmpty(previous, "no available previous deployments of", microservice)
	DeployPrevious(previous)
}

func KDeployPrevWithRegistry() {
	repos := GetPrevious().Keys()
	TerminateOnEmpty(repos, "no available previous deployments")

	selectedRepo := prompt.RepoSelect(repos)
	microservice = selectedRepo
	KDeployPrev()
}
