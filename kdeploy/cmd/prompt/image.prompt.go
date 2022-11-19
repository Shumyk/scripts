package cmd

import (
	"shumyk/kdeploy/cmd/model"
	util "shumyk/kdeploy/cmd/util"

	"github.com/AlecAivazis/survey/v2"
)

func ImageSelect(input []model.ImageOption) model.SelectedImage {
	chosenString := prompt(
		"select image to deploy",
		model.Stringify(input),
	)
	return model.SelectedImageOf(chosenString)
}

func RepoSelect(repos []string) string {
	return prompt("select repo", repos)
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
