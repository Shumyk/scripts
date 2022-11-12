package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	prompt "shumyk/kdeploy/cmd/prompt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/google/go-containerregistry/pkg/authn"
	gcr "github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/google"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

const REPOSITORY = ""

var (
	clientSet *kubernetes.Clientset

	namespace    string
	workloadName string

	statefulSets = map[string]any{"api-core": struct{}{}}
)

func init() {
	k8sConfigPath := filepath.Join(clientcmd.RecommendedConfigDir, clientcmd.RecommendedFileName)
	k8sConfigBypes, _ := os.ReadFile(k8sConfigPath)
	k8sClientConfig, _ := clientcmd.NewClientConfigFromBytes(k8sConfigBypes)

	go Namespace(k8sClientConfig)
	initClientSet(k8sClientConfig)
}

func initClientSet(config clientcmd.ClientConfig) {
	k8sRestConfig, _ := clientcmd.BuildConfigFromKubeconfigGetter(
		"",
		func() (*clientcmdapi.Config, error) {
			c, err := config.RawConfig()
			return &c, err
		},
	)
	clientSet, _ = kubernetes.NewForConfig(k8sRestConfig)
}

func ResolveResources() {
	imagesChannel := make(chan *google.Tags)
	go getImages(imagesChannel)

	currentImageChannel := make(chan string)
	go resolveCurrentImage(currentImageChannel)

	workloadName = namespace + "-" + microservice
	workloadResource := resolveWorkloadResource()
	fmt.Fprintln(os.Stdout, "Workload resource:", workloadResource)

	fmt.Fprintln(os.Stdout, "Current Image:", <-currentImageChannel)

	imageOptions := prompt.ImageOptions(<-imagesChannel)
	selectedImage := prompt.PromptImageSelect(imageOptions)
	if selectedImage.IsEmpty() {
		fmt.Fprintln(os.Stdout, "heh, ctrl+C combination was gently pressed. see you")
		os.Exit(0)
	}
	fmt.Fprintln(os.Stdout, "selectedImage:", selectedImage)

	setImage(selectedImage)
}

func resolveWorkloadResource() string {
	if microservice == "api-core" {
		return "statefulset"
	}
	return "deployment"
}

func Namespace(config clientcmd.ClientConfig) {
	namespace, _, _ = config.Namespace()
	fmt.Fprintln(os.Stdout, "Namespace:", namespace)
}

func resolveCurrentImage(ch chan<- string) {
	if _, ok := statefulSets[microservice]; ok {
		workload, _ := clientSet.AppsV1().StatefulSets(namespace).Get(context.Background(), workloadName, v1.GetOptions{})
		ch <- workload.Spec.Template.Spec.Containers[0].Image
	} else {
		workload, _ := clientSet.AppsV1().Deployments(namespace).Get(context.Background(), workloadName, v1.GetOptions{})
		ch <- workload.Spec.Template.Spec.Containers[0].Image
	}
}

func getImages(ch chan<- *google.Tags) {
	google.NewGcloudAuthenticator()
	repo, _ := gcr.NewRepository(REPOSITORY+microservice, gcr.WithDefaultRegistry("us.gcr.io"))
	tags, _ := google.List(repo, google.WithAuthFromKeychain(authn.DefaultKeychain))
	ch <- tags
}

func setImage(image prompt.SelectedImage) {
	deployments := clientSet.AppsV1().Deployments(namespace)
	deployment, _ := deployments.Get(context.Background(), workloadName, v1.GetOptions{})
	newImage := fmt.Sprintf(
		"us.gcr.io/%v%v%v@sha256:%v",
		REPOSITORY,
		microservice,
		optionallyAppendSemicolon(image.Tags[0]),
		image.Digest,
	)
	fmt.Fprintln(os.Stdout, "newImage:", newImage)
	deployment.Spec.Template.Spec.Containers[0].Image = newImage
	deployments.Update(context.Background(), deployment, v1.UpdateOptions{})
}

func optionallyAppendSemicolon(tag string) string {
	if len(tag) > 0 {
		return fmt.Sprintf(":%v", tag)
	}
	return ""
}
