package cmd

import (
	"fmt"
	prompt "shumyk/kdeploy/cmd/prompt"
	printer "shumyk/kdeploy/cmd/util"
)

func init() {
	go printer.InitPrinter()
}

func KDeploy() {
	DeployNew()
}

func KDeployRegistry() {
	repos := ListRepos()
	selectedRepo := prompt.PromptRepo(repos)
	fmt.Println("selected repo:", selectedRepo)
	microservice = selectedRepo
	KDeploy()
}
