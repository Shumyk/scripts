package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	prompt "shumyk/kdeploy/cmd/prompt"

	"github.com/google/go-containerregistry/pkg/v1/google"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clapi "k8s.io/client-go/tools/clientcmd/api"
)

func KDeploy() {
	kubeConfig := resolveKubeConfig()

	go Namespace(kubeConfig)

	clientSetChannel := make(chan bool)
	go initClientSet(kubeConfig, clientSetChannel)

	imagesChannel := make(chan *google.Tags)
	go ListRepoImages(imagesChannel)

	currentImageChannel := make(chan string)
	<-clientSetChannel
	go ResolveCurrentImage(currentImageChannel)

	workloadName = namespace + "-" + microservice
	fmt.Fprintln(os.Stdout, "Current Image:", <-currentImageChannel)

	imageOptions := prompt.ImageOptions(<-imagesChannel)
	selectedImage := prompt.PromptImageSelect(imageOptions)
	if selectedImage.IsEmpty() {
		fmt.Fprintln(os.Stdout, "heh, ctrl+C combination was gently pressed. see you")
		os.Exit(0)
	}
	fmt.Fprintln(os.Stdout, "selectedImage:", selectedImage)

	SetImage(selectedImage)
}

func resolveKubeConfig() (c clientcmd.ClientConfig) {
	k8sConfigPath := filepath.Join(clientcmd.RecommendedConfigDir, clientcmd.RecommendedFileName)
	k8sConfigBypes, _ := os.ReadFile(k8sConfigPath)
	c, _ = clientcmd.NewClientConfigFromBytes(k8sConfigBypes)
	return
}

func initClientSet(config clientcmd.ClientConfig, ch chan<- bool) {
	k8sRestConfig, _ := clientcmd.BuildConfigFromKubeconfigGetter(
		"", kubeconfigGetter(config),
	)
	clientSet, _ = kubernetes.NewForConfig(k8sRestConfig)
	ch <- true
}

func kubeconfigGetter(c clientcmd.ClientConfig) func() (*clapi.Config, error) {
	return func() (*clapi.Config, error) {
		c, err := c.RawConfig()
		return &c, err
	}
}
