// Code generated by helmit-generate. DO NOT EDIT.

package v1

import (
	"context"
	"github.com/onosproject/helmit/pkg/kubernetes/resource"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kubernetes "k8s.io/client-go/kubernetes"
	"time"
)

type SecretsReader interface {
	Get(ctx context.Context, name string) (*Secret, error)
	List(ctx context.Context) ([]*Secret, error)
}

func NewSecretsReader(client resource.Client, filter resource.Filter) SecretsReader {
	return &secretsReader{
		Client: client,
		filter: filter,
	}
}

type secretsReader struct {
	resource.Client
	filter resource.Filter
}

func (c *secretsReader) Get(ctx context.Context, name string) (*Secret, error) {
	secret := &corev1.Secret{}
	client, err := kubernetes.NewForConfig(c.Config())
	if err != nil {
		return nil, err
	}
	err = client.CoreV1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), SecretKind.Scoped).
		Resource(SecretResource.Name).
		Name(name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do(ctx).
		Into(secret)
	if err != nil {
		return nil, err
	} else {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   SecretKind.Group,
			Version: SecretKind.Version,
			Kind:    SecretKind.Kind,
		}, secret.ObjectMeta)
		if err != nil {
			return nil, err
		} else if !ok {
			return nil, errors.NewNotFound(schema.GroupResource{
				Group:    SecretKind.Group,
				Resource: SecretResource.Name,
			}, name)
		}
	}
	return NewSecret(secret, c.Client), nil
}

func (c *secretsReader) List(ctx context.Context) ([]*Secret, error) {
	list := &corev1.SecretList{}
	client, err := kubernetes.NewForConfig(c.Config())
	if err != nil {
		return nil, err
	}
	err = client.CoreV1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), SecretKind.Scoped).
		Resource(SecretResource.Name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do(ctx).
		Into(list)
	if err != nil {
		return nil, err
	}

	results := make([]*Secret, 0, len(list.Items))
	for _, secret := range list.Items {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   SecretKind.Group,
			Version: SecretKind.Version,
			Kind:    SecretKind.Kind,
		}, secret.ObjectMeta)
		if err != nil {
			return nil, err
		} else if ok {
			copy := secret
			results = append(results, NewSecret(&copy, c.Client))
		}
	}
	return results, nil
}
