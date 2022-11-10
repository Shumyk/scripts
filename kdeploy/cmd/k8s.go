package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/google/go-containerregistry/pkg/authn"
	gcr "github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/google"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

const REPOSITORY = ""

var (
	k8sClientConfig clientcmd.ClientConfig
	k8sRestConfig   *rest.Config
	clientSet       *kubernetes.Clientset

	namespace    string
	workloadName string
	currentImage string

	statefulSets = map[string]any{"api-core": struct{}{}}
)

func init() {
	k8sConfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	k8sConfigBypes, _ := os.ReadFile(k8sConfigPath)
	k8sClientConfig, _ = clientcmd.NewClientConfigFromBytes(k8sConfigBypes)
	k8sRestConfig, _ = clientcmd.BuildConfigFromFlags("", k8sConfigPath)
}

func ResolveResources() {
	Namespace()
	fmt.Fprintln(os.Stdout, "Namespace:", namespace)

	workloadName = namespace + "-" + microservice
	workloadResource := resolveWorkloadResource()
	fmt.Fprintln(os.Stdout, "Workload resource:", workloadResource)

	resolveImage()
	fmt.Fprintln(os.Stdout, "Current Image:", currentImage)

	tags := getImages()
	imageOptions := sorted(toImageOptions(tags))
	selectedImage := PromptImageSelect(imageOptions)
	fmt.Fprintln(os.Stdout, "selectedImage:", selectedImage)

	setImage(selectedImage)
}

func resolveWorkloadResource() string {
	if microservice == "api-core" {
		return "statefulset"
	}
	return "deployment"
}

func Namespace() {
	namespace, _, _ = k8sClientConfig.Namespace()
}

func resolveImage() string {
	clientSet, _ = kubernetes.NewForConfig(k8sRestConfig)
	if _, ok := statefulSets[microservice]; ok {
		workload, _ := clientSet.AppsV1().StatefulSets(namespace).Get(context.Background(), workloadName, v1.GetOptions{})
		currentImage = workload.Spec.Template.Spec.Containers[0].Image
		return currentImage
	} else {
		workload, _ := clientSet.AppsV1().Deployments(namespace).Get(context.Background(), workloadName, v1.GetOptions{})
		currentImage = workload.Spec.Template.Spec.Containers[0].Image
		return currentImage
	}
}

func getImages() (tags *google.Tags) {
	google.NewGcloudAuthenticator()
	repo, _ := gcr.NewRepository(REPOSITORY+microservice, gcr.WithDefaultRegistry("us.gcr.io"))
	tags, _ = google.List(repo, google.WithAuthFromKeychain(authn.DefaultKeychain))
	return
}

func toImageOptions(tags *google.Tags) (options []ImageOption) {
	for digest, manifest := range tags.Manifests {
		options = append(options, ImageOption{manifest.Created, manifest.Tags, digest})
	}
	return
}

func sorted(options []ImageOption) []ImageOption {
	sort.SliceStable(options, sortByCreated(options))
	return options
}

func sortByCreated(options []ImageOption) func(i, j int) bool {
	return func(i, j int) bool {
		return options[i].Created.After(options[j].Created)
	}
}

func setImage(image SelectedImage) {
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
