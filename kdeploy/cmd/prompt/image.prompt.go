package cmd

import (
	"shumyk/kdeploy/cmd/model"
	util "shumyk/kdeploy/cmd/util"

	"github.com/AlecAivazis/survey/v2"
)

type PromptInput[T any] struct {
	Data           T
	ToImageOptions func(T) []model.ImageOption
}

func ImageSelect[T any](input PromptInput[T]) model.SelectedImage {
	options := input.ToImageOptions(input.Data)
	chosenString := doPrompt(
		"select image to deploy",
		model.Stringify(options),
	)
	return model.SelectedImageOf(chosenString)
}

func RepoSelect(repos []string) string {
	return doPrompt("select repo", repos)
}

func doPrompt(title string, options []string) (res string) {
	prompt := &survey.Select{
		Message: title,
		Options: options,
	}
	err := survey.AskOne(prompt, &res)

	util.Laugh(err)
	util.TerminateOnSigint(res)
	return
}
