// Code generated by helmit-generate. DO NOT EDIT.

package v1beta1

import (
	"context"
	"github.com/onosproject/helmit/pkg/kubernetes/resource"
	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	"time"
)

var IngressKind = resource.Kind{
	Group:   "networking",
	Version: "v1beta1",
	Kind:    "Ingress",
	Scoped:  true,
}

var IngressResource = resource.Type{
	Kind: IngressKind,
	Name: "ingresses",
}

func NewIngress(ingress *networkingv1beta1.Ingress, client resource.Client) *Ingress {
	return &Ingress{
		Resource: resource.NewResource(ingress.ObjectMeta, IngressKind, client),
		Object:   ingress,
	}
}

type Ingress struct {
	*resource.Resource
	Object *networkingv1beta1.Ingress
}

func (r *Ingress) Delete(ctx context.Context) error {
	client, err := kubernetes.NewForConfig(r.Config())
	if err != nil {
		return err
	}
	return client.NetworkingV1beta1().
		RESTClient().
		Delete().
		NamespaceIfScoped(r.Namespace, IngressKind.Scoped).
		Resource(IngressResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do(ctx).
		Error()
}
