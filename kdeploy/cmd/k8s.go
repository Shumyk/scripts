package cmd

import (
	"context"
	"fmt"
	prompt "shumyk/kdeploy/cmd/model"

	util "shumyk/kdeploy/cmd/util"

	apps "k8s.io/api/apps/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/tools/clientcmd"
	clapi "k8s.io/client-go/tools/clientcmd/api"
)

var (
	clientSet *kubernetes.Clientset

	isDeployment bool
	deployments  v1.DeploymentInterface
	deployment   *apps.Deployment
	statefulSets v1.StatefulSetInterface
	statefulSet  *apps.StatefulSet

	namespace    string
	workloadName string
)

func ClientSet(config clientcmd.ClientConfig, ch chan<- bool) {
	k8sRestConfig, _ := clientcmd.BuildConfigFromKubeconfigGetter(
		"", kubeconfigGetter(config),
	)
	clientSet, _ = kubernetes.NewForConfig(k8sRestConfig)

	if isDeployment {
		deployments = clientSet.AppsV1().Deployments(namespace)
		deployment, _ = deployments.Get(context.Background(), workloadName, meta.GetOptions{})
	} else {
		statefulSets = clientSet.AppsV1().StatefulSets(namespace)
		statefulSet, _ = statefulSets.Get(context.Background(), workloadName, meta.GetOptions{})
	}

	ch <- true
}

func kubeconfigGetter(c clientcmd.ClientConfig) func() (*clapi.Config, error) {
	return func() (*clapi.Config, error) {
		c, err := c.RawConfig()
		return &c, err
	}
}

func LoadMetadata(config clientcmd.ClientConfig) {
	namespace, _, _ = config.Namespace()
	workloadName = namespace + "-" + microservice
	resolveWorkloadType()
	util.PrintEnvironmentInfo(microservice, namespace)
}

func resolveWorkloadType() {
	// TODO: statefulsets from config
	statefulSets := map[string]any{"api-core": struct{}{}}
	_, ok := statefulSets[microservice]
	isDeployment = !ok
}

func ResolveCurrentImage() string {
	if isDeployment {
		workload, _ := deployments.Get(context.Background(), workloadName, meta.GetOptions{})
		return workload.Spec.Template.Spec.Containers[0].Image
	} else {
		workload, _ := statefulSets.Get(context.Background(), workloadName, meta.GetOptions{})
		return workload.Spec.Template.Spec.Containers[0].Image
	}
}

func SetImage(image *prompt.SelectedImage) {
	newImage := fmt.Sprintf(
		"us.gcr.io/%v%v%v@%v%v",
		Repository,
		microservice,
		util.AppendSemicolon(image.Tags[0]),
		util.DigestPrefix,
		image.Digest,
	)

	// TODO: uncomment
	if isDeployment {
		// deployment.Spec.Template.Spec.Containers[0].Image = newImage
		// deployments.Update(context.Background(), deployment, meta.UpdateOptions{})
	} else {
		// statefulSet.Spec.Template.Spec.Containers[0].Image = newImage
		// statefulSets.Update(context.Background(), statefulSet, meta.UpdateOptions{})
	}
	fmt.Println("New Image:", newImage)
	util.PrintDeployedImageInfo(image.Tags[0], image.Digest)
}
