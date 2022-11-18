package cmd

import (
	util "shumyk/kdeploy/cmd/util"

	"github.com/AlecAivazis/survey/v2"
	"github.com/google/go-containerregistry/pkg/v1/google"
)

func PromptImageSelect(tags *google.Tags) SelectedImage {
	options := ImageOptionsOfTags(tags)
	return prompt(options)
}

func PromptPrevImageSelect(prevs []PrevImage) SelectedImage {
	options := ImageOptionsOfPrevImages(prevs)
	return prompt(options)
}

func prompt(options []ImageOption) (s SelectedImage) {
	selectedImage := PromptGeneric("select image to deploy", Stringify(options))
	return SelectedImageOf(selectedImage)
}

func PromptRepo(repos []string) string {
	return PromptGeneric("select repo", repos)
}

func PromptGeneric(title string, options []string) (res string) {
	prompt := selectPrompt(title, options)
	survey.AskOne(prompt, &res)
	util.TerminateOnSigint(res)
	return
}

func selectPrompt(title string, options []string) *survey.Select {
	return &survey.Select{
		Message: title,
		Options: options,
	}
}
