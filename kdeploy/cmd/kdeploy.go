package cmd

import prompt "shumyk/kdeploy/cmd/prompt"

func KDeploy() {
	DeployNew()
}

func KDeployWithRegistry() {
	repos := ListRepos()
	selectedRepo := prompt.RepoSelect(repos)
	microservice = selectedRepo
	KDeploy()
}
