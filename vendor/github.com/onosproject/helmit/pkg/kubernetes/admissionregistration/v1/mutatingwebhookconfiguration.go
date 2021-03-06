// Code generated by helmit-generate. DO NOT EDIT.

package v1

import (
	"context"
	"github.com/onosproject/helmit/pkg/kubernetes/resource"
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	"time"
)

var MutatingWebhookConfigurationKind = resource.Kind{
	Group:   "admissionregistration.k8s.io",
	Version: "v1",
	Kind:    "MutatingWebhookConfiguration",
	Scoped:  true,
}

var MutatingWebhookConfigurationResource = resource.Type{
	Kind: MutatingWebhookConfigurationKind,
	Name: "mutatingwebhookconfigurations",
}

func NewMutatingWebhookConfiguration(mutatingWebhookConfiguration *admissionregistrationv1.MutatingWebhookConfiguration, client resource.Client) *MutatingWebhookConfiguration {
	return &MutatingWebhookConfiguration{
		Resource: resource.NewResource(mutatingWebhookConfiguration.ObjectMeta, MutatingWebhookConfigurationKind, client),
		Object:   mutatingWebhookConfiguration,
	}
}

type MutatingWebhookConfiguration struct {
	*resource.Resource
	Object *admissionregistrationv1.MutatingWebhookConfiguration
}

func (r *MutatingWebhookConfiguration) Delete(ctx context.Context) error {
	client, err := kubernetes.NewForConfig(r.Config())
	if err != nil {
		return err
	}
	return client.AdmissionregistrationV1().
		RESTClient().
		Delete().
		NamespaceIfScoped(r.Namespace, MutatingWebhookConfigurationKind.Scoped).
		Resource(MutatingWebhookConfigurationResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do(ctx).
		Error()
}
