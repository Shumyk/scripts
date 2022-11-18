package cmd

import (
	util "shumyk/kdeploy/cmd/util"
	"strings"

	"github.com/AlecAivazis/survey/v2/core"
)

type SelectedImage struct {
	Tags   []string
	Digest string
}

func (i *SelectedImage) WriteAnswer(field string, answer any) error {
	selectedValue := answer.(core.OptionAnswer).Value
	selectedImageData := strings.Split(selectedValue, util.DIVIDER)

	i.Digest = selectedImageData[1]
	i.Tags = strings.Split(selectedImageData[2], util.TAGS_DELIM)

	return nil
}
