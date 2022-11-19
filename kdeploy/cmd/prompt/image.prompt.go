package cmd

import (
	"shumyk/kdeploy/cmd/model"
	util "shumyk/kdeploy/cmd/util"

	"github.com/AlecAivazis/survey/v2"
	"github.com/google/go-containerregistry/pkg/v1/google"
)

func PromptImageSelect(tags *google.Tags) model.SelectedImage {
	options := model.ImageOptionsOfTags(tags)
	return prompt(options)
}

func PromptPrevImageSelect(prevs []model.PrevImage) model.SelectedImage {
	options := model.ImageOptionsOfPrevImages(prevs)
	return prompt(options)
}

func prompt(options []model.ImageOption) (s model.SelectedImage) {
	selectedImage := PromptGeneric("select image to deploy", model.Stringify(options))
	return model.SelectedImageOf(selectedImage)
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
