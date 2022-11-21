package cmd

import (
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	confApps "k8s.io/client-go/applyconfigurations/apps/v1"
	core "k8s.io/client-go/applyconfigurations/core/v1"

	. "shumyk/kdeploy/cmd/model"
	. "shumyk/kdeploy/cmd/util"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

var (
	clientSet *kubernetes.Clientset

	namespace       string
	k8sResource     string
	k8sResourceName string
)

func ClientSet(config clientcmd.ClientConfig, ch chan<- bool) {
	configGetter := kubeConfigGetter(config)
	k8sRestConfig, err := clientcmd.BuildConfigFromKubeconfigGetter("", configGetter)
	ErrorCheck(err, "Building config from kube config getter failed")

	clientSet, err = kubernetes.NewForConfig(k8sRestConfig)
	ErrorCheck(err, "Creating Client Set failed")

	ch <- true
}

func kubeConfigGetter(c clientcmd.ClientConfig) clientcmd.KubeconfigGetter {
	return func() (*api.Config, error) {
		c, err := c.RawConfig()
		return &c, err
	}
}

func LoadMetadata(config clientcmd.ClientConfig) {
	var err error
	namespace, _, err = config.Namespace()
	ErrorCheck(err, "Resolving namespace failed")

	k8sResourceName = namespace + "-" + microservice
	resolveWorkloadType()

	PrintEnvironmentInfo(microservice, namespace)
}

func resolveWorkloadType() {
	// TODO: statefulsets from config
	statefulSets := map[string]any{"api-core": struct{}{}}
	if _, ok := statefulSets[microservice]; ok {
		k8sResource = "statefulsets"
	} else {
		k8sResource = "deployments"
	}
}

func GetImage() string {
	var response K8sResourceAgnosticResponse
	err := clientSet.AppsV1().RESTClient().
		Get().
		Namespace(namespace).
		Resource(k8sResource).
		Name(k8sResourceName).
		Do(ctx).
		Into(&response)
	ErrorCheck(err, "GET image failed")
	return response.Spec.Template.Spec.Containers[0].Image
}

func SetImage(image *SelectedImage) {
	newImage := ComposeImagePath(Registry, Repository, microservice, image.Tag(), image.Digest)
	imageChange := composeImagePatch(newImage)
	data, err := json.Marshal(imageChange)
	ErrorCheck(err, "Unmarshalling image change failed")

	updateError := clientSet.AppsV1().RESTClient().
		Patch(types.StrategicMergePatchType).
		Namespace(namespace).
		Resource(k8sResource).
		Name(k8sResourceName).
		Body(data).
		Do(ctx).
		Error()

	ErrorCheck(updateError, "PATCH image failed")
	PrintImageInfo(HeaderDeployedImage, image.Tags[0], image.Digest)
}

// composeImagePatch composes resource apply configuration to patch only image.
// DeploymentApplyConfiguration is used, but it's actually resource agnostic as we patch only image,
// which is located under same place among resources.
func composeImagePatch(newImage string) confApps.DeploymentApplyConfiguration {
	container := core.ContainerApplyConfiguration{Image: &newImage, Name: &microservice}
	podSpec := core.PodSpecApplyConfiguration{Containers: []core.ContainerApplyConfiguration{container}}
	templateSpec := core.PodTemplateSpecApplyConfiguration{Spec: &podSpec}
	resourceSpec := confApps.DeploymentSpecApplyConfiguration{Template: &templateSpec}
	return confApps.DeploymentApplyConfiguration{Spec: &resourceSpec}
}
