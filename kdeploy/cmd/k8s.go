package cmd

import (
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

	updateOpts = meta.UpdateOptions{}
)

func ClientSet(config clientcmd.ClientConfig, ch chan<- bool) {
	k8sRestConfig, _ := clientcmd.BuildConfigFromKubeconfigGetter(
		"", kubeconfigGetter(config),
	)
	clientSet, _ = kubernetes.NewForConfig(k8sRestConfig)

	if isDeployment {
		deployments = clientSet.AppsV1().Deployments(namespace)
		deployment, _ = deployments.Get(ctx, workloadName, meta.GetOptions{})
	} else {
		statefulSets = clientSet.AppsV1().StatefulSets(namespace)
		statefulSet, _ = statefulSets.Get(ctx, workloadName, meta.GetOptions{})
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
		workload, _ := deployments.Get(ctx, workloadName, meta.GetOptions{})
		return workload.Spec.Template.Spec.Containers[0].Image
	} else {
		workload, _ := statefulSets.Get(ctx, workloadName, meta.GetOptions{})
		return workload.Spec.Template.Spec.Containers[0].Image
	}
}

func SetImage(image *prompt.SelectedImage) {
	var newImage = util.ComposeImagePath(Registry, Repository, microservice, image.Tag(), image.Digest)
	var updateError error
	if isDeployment {
		deployment.Spec.Template.Spec.Containers[0].Image = newImage
		_, updateError = deployments.Update(ctx, deployment, updateOpts)
	} else {
		statefulSet.Spec.Template.Spec.Containers[0].Image = newImage
		_, updateError = statefulSets.Update(ctx, statefulSet, updateOpts)
	}
	util.ErrorCheck(updateError)
	util.PrintImageInfo(util.HeaderDeployedImage, image.Tags[0], image.Digest)
}
