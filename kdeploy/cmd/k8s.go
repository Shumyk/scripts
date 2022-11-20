package cmd

import (
	"k8s.io/apimachinery/pkg/runtime"
	prompt "shumyk/kdeploy/cmd/model"

	util "shumyk/kdeploy/cmd/util"

	appsV1 "k8s.io/api/apps/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clapi "k8s.io/client-go/tools/clientcmd/api"
)

var (
	clientSet *kubernetes.Clientset

	isDeployment bool
	k8sResource  string

	namespace       string
	k8sResourceName string

	getOpts    = meta.GetOptions{}
	updateOpts = meta.UpdateOptions{}
)

func ClientSet(config clientcmd.ClientConfig, ch chan<- bool) {
	configGetter := kubeConfigGetter(config)
	k8sRestConfig, err := clientcmd.BuildConfigFromKubeconfigGetter("", configGetter)
	util.ErrorCheck(err)

	clientSet, err = kubernetes.NewForConfig(k8sRestConfig)
	util.ErrorCheck(err)

	ch <- true
}

func kubeConfigGetter(c clientcmd.ClientConfig) clientcmd.KubeconfigGetter {
	return func() (*clapi.Config, error) {
		c, err := c.RawConfig()
		return &c, err
	}
}

func LoadMetadata(config clientcmd.ClientConfig) {
	nm, _, err := config.Namespace()
	util.ErrorCheck(err)
	namespace = nm

	k8sResourceName = namespace + "-" + microservice
	resolveWorkloadType()

	util.PrintEnvironmentInfo(microservice, namespace)
}

func resolveWorkloadType() {
	// TODO: statefulsets from config
	statefulSets := map[string]any{"api-core": struct{}{}}
	if _, ok := statefulSets[microservice]; ok {
		k8sResource = "statefulsets"
	} else {
		k8sResource = "deployments"
	}
	//isDeployment = !ok
}

type k8sResourceAgnosticResponse struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec appsV1.DeploymentSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// TODO: maybe generated?
func (in *k8sResourceAgnosticResponse) DeepCopy() *k8sResourceAgnosticResponse {
	if in == nil {
		return nil
	}
	out := new(k8sResourceAgnosticResponse)
	in.DeepCopyInto(out)
	return out
}

func (in *k8sResourceAgnosticResponse) DeepCopyInto(out *k8sResourceAgnosticResponse) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

func (in *k8sResourceAgnosticResponse) DeepCopyObject() runtime.Object {
	if out := in.DeepCopy(); out != nil {
		return out
	}
	return nil
}

func GetImage() string {
	var response k8sResourceAgnosticResponse
	err := clientSet.AppsV1().RESTClient().
		Get().
		Namespace(namespace).
		Resource(k8sResource).
		Name(k8sResourceName).
		Do(ctx).
		Into(&response)
	util.ErrorCheck(err)
	return response.Spec.Template.Spec.Containers[0].Image
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
