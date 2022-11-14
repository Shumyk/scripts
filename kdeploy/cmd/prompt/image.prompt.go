package cmd

import (
	"os"

	printer "shumyk/kdeploy/cmd/util"

	"github.com/AlecAivazis/survey/v2"
	"github.com/google/go-containerregistry/pkg/v1/google"
)

func PromptImageSelect(tags *google.Tags) SelectedImage {
	options := ImageOptions(tags)
	return prompt(options)
}

func PromptPrevImageSelect(prevs []PrevImage) SelectedImage {
	options := PrevImageToOptions(prevs)
	return prompt(options)
}

func prompt(options []ImageOption) (s SelectedImage) {
	prompt := &survey.Select{
		Message: "select image to deploy",
		Options: Stringify(options),
	}
	survey.AskOne(prompt, &s)

	terminateOnSigint(&s)
	return
}

func PromptRepo(repos []string) (res string) {
	prompt := &survey.Select{
		Message: "select repo",
		Options: repos,
	}
	survey.AskOne(prompt, &res)
	return
}

func terminateOnSigint(selected *SelectedImage) {
	if selected.IsEmpty() {
		printer.Purple("heh, ctrl+C combination was gently pressed. see you")
		os.Exit(0)
	}
}
