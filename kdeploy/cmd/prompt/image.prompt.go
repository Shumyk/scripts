package cmd

import (
	"shumyk/kdeploy/cmd/model"
	util "shumyk/kdeploy/cmd/util"

	"github.com/AlecAivazis/survey/v2"
)

const (
	ImageSelectTitle = "select image to deploy"
	RepoSelectTitle  = "select repo"
)

func ImageSelect(input model.ImageOptions) model.SelectedImage {
	chosenString := prompt(ImageSelectTitle, input.Stringify())
	return model.SelectedImageOf(chosenString)
}

func RepoSelect(repos []string) string {
	return prompt(RepoSelectTitle, repos)
}

func prompt(title string, options []string) (selected string) {
	prompt := &survey.Select{
		Message: title,
		Options: options,
	}
	err := survey.AskOne(prompt, &selected)

	util.Laugh(err)
	util.TerminateOnSigint(selected)
	return
}
