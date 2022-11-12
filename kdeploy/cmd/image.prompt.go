package cmd

import (
	"github.com/AlecAivazis/survey/v2"
)

func PromptImageSelect(options []ImageOption) (s SelectedImage) {
	prompt := &survey.Select{
		Message: "select image to deploy",
		Options: Stringify(options),
	}
	survey.AskOne(prompt, &s)
	return
}
