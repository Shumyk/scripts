package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/google/go-containerregistry/pkg/v1/google"
)

func PromptImageSelect(tags *google.Tags) (s SelectedImage) {
	options := ImageOptions(tags)

	prompt := &survey.Select{
		Message: "select image to deploy",
		Options: Stringify(options),
	}
	survey.AskOne(prompt, &s)

	terminateOnSigint(&s)
	return
}

func terminateOnSigint(selected *SelectedImage) {
	if selected.IsEmpty() {
		fmt.Fprintln(os.Stdout, "heh, ctrl+C combination was gently pressed. see you")
		os.Exit(0)
	}
}
