package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

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

	fmt.Fprintln(os.Stdout, "Namespace:", namespace)
	fmt.Fprintln(os.Stdout, "Workload resource:", workloadResource)
	fmt.Fprintln(os.Stdout, "Current Image:", currentImage)
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
