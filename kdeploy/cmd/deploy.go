package cmd

import (
	"os"
	"path/filepath"
	prompt "shumyk/kdeploy/cmd/prompt"
	printer "shumyk/kdeploy/cmd/util"

	"github.com/google/go-containerregistry/pkg/v1/google"
	"k8s.io/client-go/tools/clientcmd"
)

type ImageSelecter func(chan bool) prompt.SelectedImage

func DeployNew() {
	deployTemplate(func(clientSetChannel chan bool) prompt.SelectedImage {
		imagesChannel := make(chan *google.Tags)
		go ListRepoImages(imagesChannel)

		<-clientSetChannel
		currentImage := ResolveCurrentImage()
		curTag, curDigest := printer.PrintImageInfo(currentImage)

		selectedImage := prompt.PromptImageSelect(<-imagesChannel)
		go SavePreviouslyDeployed(curTag, curDigest)
		return selectedImage
	})
}

func DeployPrevious(p []prompt.PrevImage) {
	deployTemplate(func(clientSet chan bool) prompt.SelectedImage {
		s := prompt.PromptPrevImageSelect(p)
		<-clientSet
		return s
	})
}

func deployTemplate(selecter ImageSelecter) {
	kubeConfig := resolveKubeConfig()
	go LoadMetadata(kubeConfig)

	clientSetChannel := make(chan bool)
	go ClientSet(kubeConfig, clientSetChannel)

	selectedImage := selecter(clientSetChannel)
	SetImage(&selectedImage)
}

func resolveKubeConfig() (c clientcmd.ClientConfig) {
	k8sConfigPath := filepath.Join(clientcmd.RecommendedConfigDir, clientcmd.RecommendedFileName)
	k8sConfigBypes, _ := os.ReadFile(k8sConfigPath)
	c, _ = clientcmd.NewClientConfigFromBytes(k8sConfigBypes)
	return
}
