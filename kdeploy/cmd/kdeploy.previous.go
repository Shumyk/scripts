package cmd

import (
	prompt "shumyk/kdeploy/cmd/prompt"
	. "shumyk/kdeploy/cmd/util"
)

func KDeployPrevious() {
	previous := GetPrevious()[microservice]
	TerminateOnEmpty(previous, "previous deployments of", microservice, "absent")
	DeployPrevious(previous)
}

func KDeployPreviousWithRegistry() {
	repos := GetPrevious().Keys()
	TerminateOnEmpty(repos, "previous deployments absent")

	microservice = prompt.RepoSelect(repos)
	KDeployPrevious()
}
