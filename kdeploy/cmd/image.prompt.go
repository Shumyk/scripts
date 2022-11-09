package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/core"
)

const DIVIDER = "     "

type ImageOption struct {
	Created time.Time
	Tags    []string
	Digest  string
}

func (o ImageOption) Stringify() string {
	return fmt.Sprintf(
		"%v%v%v%v%v",
		o.Created.Format("2006-01-02 15:04:05"), DIVIDER,
		strings.TrimPrefix(o.Digest, "sha256:"), DIVIDER,
		strings.Join(o.Tags, ","),
	)
}

type SelectedImage struct {
	Tags   []string
	Digest string
}

func (i *SelectedImage) WriteAnswer(field string, answer interface{}) error {
	selectedValue := answer.(core.OptionAnswer).Value
	selectedImageData := strings.Split(selectedValue, DIVIDER)

	i.Digest = selectedImageData[1]
	i.Tags = strings.Split(selectedImageData[2], ",")

	return nil
}

func PromptImageSelect(options []ImageOption) *SelectedImage {
	prompt := &survey.Select{
		Message: "select image to deploy",
		Options: stringify(options),
	}
	var selectedImage SelectedImage
	survey.AskOne(prompt, &selectedImage)
	return &selectedImage
}

func stringify(options []ImageOption) []string {
	var formattedOptions []string
	for _, option := range options {
		formattedOptions = append(formattedOptions, option.Stringify())
	}
	return formattedOptions
}
