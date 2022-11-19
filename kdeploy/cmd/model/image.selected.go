package model

import (
	util "shumyk/kdeploy/cmd/util"
	"strings"
)

type SelectedImage struct {
	Tags   []string
	Digest string
}

func SelectedImageOf(value string) (i SelectedImage) {
	selectedImageData := strings.Split(value, util.Divider)
	i.Digest = selectedImageData[1]
	i.Tags = strings.Split(selectedImageData[2], util.Delimiter)
	return
}
