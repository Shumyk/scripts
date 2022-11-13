package cmd

import (
	"context"
	"fmt"
	"os"

	prompt "shumyk/kdeploy/cmd/prompt"
	util "shumyk/kdeploy/cmd/util"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	clientSet *kubernetes.Clientset

	namespace    string
	workloadName string

	statefulSets = map[string]any{"api-core": struct{}{}}
)

func Namespace(config clientcmd.ClientConfig) {
	namespace, _, _ = config.Namespace()
	fmt.Fprintln(os.Stdout, "Namespace:", namespace)
}

func ResolveCurrentImage(ch chan<- string) {
	if _, ok := statefulSets[microservice]; ok {
		workload, _ := clientSet.AppsV1().StatefulSets(namespace).Get(context.Background(), workloadName, v1.GetOptions{})
		ch <- workload.Spec.Template.Spec.Containers[0].Image
	} else {
		workload, _ := clientSet.AppsV1().Deployments(namespace).Get(context.Background(), workloadName, v1.GetOptions{})
		ch <- workload.Spec.Template.Spec.Containers[0].Image
	}
}

func SetImage(image prompt.SelectedImage) {
	newImage := fmt.Sprintf(
		"us.gcr.io/%v%v%v@%v%v",
		REPOSITORY,
		microservice,
		util.AppendSemicolon(image.Tags[0]),
		util.DIGEST_PREFIX,
		image.Digest,
	)
	fmt.Fprintln(os.Stdout, "newImage:", newImage)

	if _, ok := statefulSets[microservice]; ok {
		statefulsets := clientSet.AppsV1().StatefulSets(namespace)
		statefulset, _ := statefulsets.Get(context.Background(), workloadName, v1.GetOptions{})
		statefulset.Spec.Template.Spec.Containers[0].Image = newImage
		statefulsets.Update(context.Background(), statefulset, v1.UpdateOptions{})
	} else {
		deployments := clientSet.AppsV1().Deployments(namespace)
		deployment, _ := deployments.Get(context.Background(), workloadName, v1.GetOptions{})
		deployment.Spec.Template.Spec.Containers[0].Image = newImage
		deployments.Update(context.Background(), deployment, v1.UpdateOptions{})
	}
}
