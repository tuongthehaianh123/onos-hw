// Code generated by helmit-generate. DO NOT EDIT.

package v1beta1

import (
	"context"
	appsv1 "github.com/onosproject/helmit/pkg/kubernetes/apps/v1"
	corev1 "github.com/onosproject/helmit/pkg/kubernetes/core/v1"
	"github.com/onosproject/helmit/pkg/kubernetes/resource"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	"time"
)

var StatefulSetKind = resource.Kind{
	Group:   "apps",
	Version: "v1beta1",
	Kind:    "StatefulSet",
	Scoped:  true,
}

var StatefulSetResource = resource.Type{
	Kind: StatefulSetKind,
	Name: "statefulsets",
}

func NewStatefulSet(statefulSet *appsv1beta1.StatefulSet, client resource.Client) *StatefulSet {
	return &StatefulSet{
		Resource:             resource.NewResource(statefulSet.ObjectMeta, StatefulSetKind, client),
		Object:               statefulSet,
		ReplicaSetsReference: appsv1.NewReplicaSetsReference(client, resource.NewUIDFilter(statefulSet.UID)),
		PodsReference:        corev1.NewPodsReference(client, resource.NewUIDFilter(statefulSet.UID)),
	}
}

type StatefulSet struct {
	*resource.Resource
	Object *appsv1beta1.StatefulSet
	appsv1.ReplicaSetsReference
	corev1.PodsReference
}

func (r *StatefulSet) Delete(ctx context.Context) error {
	client, err := kubernetes.NewForConfig(r.Config())
	if err != nil {
		return err
	}
	return client.AppsV1beta1().
		RESTClient().
		Delete().
		NamespaceIfScoped(r.Namespace, StatefulSetKind.Scoped).
		Resource(StatefulSetResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do(ctx).
		Error()
}
