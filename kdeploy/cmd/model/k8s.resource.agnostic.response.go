package model

import (
	v12 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type K8sResourceAgnosticResponse struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec v12.DeploymentSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// TODO: maybe generated?
func (in *K8sResourceAgnosticResponse) DeepCopy() *K8sResourceAgnosticResponse {
	if in == nil {
		return nil
	}
	out := new(K8sResourceAgnosticResponse)
	in.DeepCopyInto(out)
	return out
}

func (in *K8sResourceAgnosticResponse) DeepCopyInto(out *K8sResourceAgnosticResponse) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

func (in *K8sResourceAgnosticResponse) DeepCopyObject() runtime.Object {
	if out := in.DeepCopy(); out != nil {
		return out
	}
	return nil
}
