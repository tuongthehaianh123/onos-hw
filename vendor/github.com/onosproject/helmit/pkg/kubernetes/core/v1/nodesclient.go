// Code generated by helmit-generate. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/helmit/pkg/kubernetes/resource"
)

type NodesClient interface {
	Nodes() NodesReader
}

func NewNodesClient(resources resource.Client, filter resource.Filter) NodesClient {
	return &nodesClient{
		Client: resources,
		filter: filter,
	}
}

type nodesClient struct {
	resource.Client
	filter resource.Filter
}

func (c *nodesClient) Nodes() NodesReader {
	return NewNodesReader(c.Client, c.filter)
}
