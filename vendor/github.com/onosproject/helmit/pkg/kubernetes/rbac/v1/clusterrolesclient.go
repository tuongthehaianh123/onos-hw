// Code generated by helmit-generate. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/helmit/pkg/kubernetes/resource"
)

type ClusterRolesClient interface {
	ClusterRoles() ClusterRolesReader
}

func NewClusterRolesClient(resources resource.Client, filter resource.Filter) ClusterRolesClient {
	return &clusterRolesClient{
		Client: resources,
		filter: filter,
	}
}

type clusterRolesClient struct {
	resource.Client
	filter resource.Filter
}

func (c *clusterRolesClient) ClusterRoles() ClusterRolesReader {
	return NewClusterRolesReader(c.Client, c.filter)
}
