package cmd

import (
	prompt "shumyk/kdeploy/cmd/prompt"
	printer "shumyk/kdeploy/cmd/util"
)

func init() {
	go printer.InitPrinter()
}

func KDeploy() {
	DeployNew()
}

func KDeployWithRegistry() {
	repos := ListRepos()
	selectedRepo := prompt.RepoSelect(repos)
	microservice = selectedRepo
	KDeploy()
}
