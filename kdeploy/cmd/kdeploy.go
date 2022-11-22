package cmd

import prompt "shumyk/kdeploy/cmd/prompt"

func KDeploy() {
	DeployNew()
}

func KDeployWithRegistry() {
	repos := ListRepos()
	microservice = prompt.RepoSelect(repos)
	DeployNew()
}
