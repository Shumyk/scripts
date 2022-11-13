package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	prompt "shumyk/kdeploy/cmd/prompt"
	util "shumyk/kdeploy/cmd/util"

	"github.com/fatih/color"

	"golang.org/x/term"
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

func Metadata(config clientcmd.ClientConfig) {
	namespace, _, _ = config.Namespace()
	workloadName = namespace + "-" + microservice
	resolveWorkloadType()
	PrintEnvInfo()
}

func PrintEnvInfo() {
	width, _, _ := term.GetSize(int(os.Stdin.Fd()))
	header := color.New(color.BgHiGreen).SprintFunc()
	green := color.New(color.FgHiGreen).Add(color.Bold).SprintFunc()

	hrLine(width)
	envStr := "|   ENVIRONMENT |"
	envHeader := envStr + strings.Repeat(" ", width-len(envStr))
	fmt.Println(header(envHeader))
	hrLine(width)
	fmt.Printf("|   service  \t:  %v\t\n", green(microservice))
	fmt.Printf("|   namespace  \t:  %v\t\n", green(namespace))

	os.Exit(0)
}

func hrLine(width int) {
	fmt.Printf("%s\n", strings.Repeat("-", width))
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

func SetImage(image prompt.SelectedImage) {
	newImage := fmt.Sprintf(
		"us.gcr.io/%v%v%v@%v%v",
		REPOSITORY,
		microservice,
		util.AppendSemicolon(image.Tags[0]),
		util.DIGEST_PREFIX,
		image.Digest,
	)
	fmt.Println("newImage:", newImage)

	if isDeployment {
		deployment.Spec.Template.Spec.Containers[0].Image = newImage
		deployments.Update(context.Background(), deployment, meta.UpdateOptions{})
	} else {
		statefulSet.Spec.Template.Spec.Containers[0].Image = newImage
		statefulSets.Update(context.Background(), statefulSet, meta.UpdateOptions{})
	}
}
