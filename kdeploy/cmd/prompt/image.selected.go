package cmd

import (
	"strings"

	"github.com/AlecAivazis/survey/v2/core"
)

const DIVIDER = "     "

type SelectedImage struct {
	Tags   []string
	Digest string
}

func (i *SelectedImage) IsEmpty() bool {
	return len(i.Digest) == 0
}

func (i *SelectedImage) WriteAnswer(field string, answer any) error {
	selectedValue := answer.(core.OptionAnswer).Value
	selectedImageData := strings.Split(selectedValue, DIVIDER)

	i.Digest = selectedImageData[1]
	i.Tags = strings.Split(selectedImageData[2], ",")

	return nil
}
