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

var (
	k8sClientConfig clientcmd.ClientConfig
	k8sRestConfig   *rest.Config

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
	workloadName = namespace + "-" + microservice
	workloadResource := resolveWorkloadResource()
	resolveImage(namespace)
	getImages()

	fmt.Fprintln(os.Stdout, "Namespace:", namespace)
	fmt.Fprintln(os.Stdout, "Workload resource:", workloadResource)
	fmt.Fprintln(os.Stdout, "Current Image:", currentImage)
	// fmt.Fprintln(os.Stdout, "tags: ", tags)
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

func resolveImage(namespace string) string {
	clientSet, _ := kubernetes.NewForConfig(k8sRestConfig)
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

func getImages() {
	fmt.Println()

	google.NewGcloudAuthenticator()
	repo, _ := gcr.NewRepository("your-repo-"+microservice, gcr.WithDefaultRegistry("us.gcr.io"))
	tags, _ := google.List(repo, google.WithAuthFromKeychain(authn.DefaultKeychain))

	options := make([]ImageOption, 0)
	for k, v := range tags.Manifests {
		if len(k) > 0 {
			options = append(options, ImageOption{v.Created, v.Tags, k})
		}
	}

	sort.SliceStable(options, func(i, j int) bool {
		return options[i].Created.After(options[j].Created)
	})

	selectedImage := PromptImageSelect(options)
	fmt.Fprintln(os.Stdout, "selectedImage:", selectedImage)

	fmt.Println()
}
