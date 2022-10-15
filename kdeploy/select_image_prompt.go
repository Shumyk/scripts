package main

import (
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

const DIVIDER = "     "
const SPACE = " "
const SEPARATOR = "|"

func main() {
	selectedImageTmpFilename := os.Args[1]
	imagesInfoArgs := os.Args[2:]
	imagesInfoFormatted := formatImagesInfo(imagesInfoArgs)

	prompt := &survey.Select{
		Message: "select image to deploy",
		Options: imagesInfoFormatted,
	}
	selectedImage := ""
	survey.AskOne(prompt, &selectedImage)
	selectedImage = strings.Replace(selectedImage, DIVIDER, SPACE, -1)

	os.WriteFile(selectedImageTmpFilename, []byte(selectedImage), 0666)
}

func formatImagesInfo(imagesInfo []string) []string {
	var imagesInfoFormatted []string
	for i := 0; i < len(imagesInfo); i++ {
		imagesInfoFormatted = append(
			imagesInfoFormatted,
			strings.Replace(imagesInfo[i], SEPARATOR, DIVIDER, -1),
		)
	}
	return imagesInfoFormatted
}
