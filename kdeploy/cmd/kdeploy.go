package cmd

import (
	prompt "shumyk/kdeploy/cmd/prompt"
	. "shumyk/kdeploy/cmd/util"
)

func KDeploy() {
	DeployNew()
}

func KDeployPrevious() {
	previous := GetPrevious()[microservice]
	TerminateOnEmpty(previous, "previous deployments of", microservice, "absent")
	DeployPrevious(previous)
}

func KDeployWithRegistry() {
	repos := ListRepos()
	microservice = prompt.RepoSelect(repos)
	DeployNew()
}

func KDeployPreviousWithRegistry() {
	repos := GetPrevious().Keys()
	TerminateOnEmpty(repos, "previous deployments absent")

	microservice = prompt.RepoSelect(repos)
	KDeployPrevious()
}
