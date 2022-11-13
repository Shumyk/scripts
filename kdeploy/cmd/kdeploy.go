package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	prompt "shumyk/kdeploy/cmd/prompt"

	"github.com/google/go-containerregistry/pkg/v1/google"
	"k8s.io/client-go/tools/clientcmd"
)

func KDeploy() {
	kubeConfig := resolveKubeConfig()

	go Metadata(kubeConfig)

	clientSetChannel := make(chan bool)
	go ClientSet(kubeConfig, clientSetChannel)

	imagesChannel := make(chan *google.Tags)
	go ListRepoImages(imagesChannel)

	<-clientSetChannel
	currentImage := ResolveCurrentImage()
	fmt.Fprintln(os.Stdout, "Current Image:", currentImage)

	selectedImage := prompt.PromptImageSelect(<-imagesChannel)
	fmt.Fprintln(os.Stdout, "selectedImage:", selectedImage)

	SetImage(selectedImage)
}

func resolveKubeConfig() (c clientcmd.ClientConfig) {
	k8sConfigPath := filepath.Join(clientcmd.RecommendedConfigDir, clientcmd.RecommendedFileName)
	k8sConfigBypes, _ := os.ReadFile(k8sConfigPath)
	c, _ = clientcmd.NewClientConfigFromBytes(k8sConfigBypes)
	return
}
