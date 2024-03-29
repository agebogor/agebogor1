package idresolver

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
)

type fakeClient struct {
	client.Client
	nodeInspectFunc    func(string) (swarm.Node, []byte, error)
	serviceInspectFunc func(string) (swarm.Service, []byte, error)
}

func (cli *fakeClient) NodeInspectWithRaw(_ context.Context, nodeID string) (swarm.Node, []byte, error) {
	if cli.nodeInspectFunc != nil {
		return cli.nodeInspectFunc(nodeID)
	}
	return swarm.Node{}, []byte{}, nil
}

func (cli *fakeClient) ServiceInspectWithRaw(_ context.Context, serviceID string, _ types.ServiceInspectOptions) (swarm.Service, []byte, error) {
	if cli.serviceInspectFunc != nil {
		return cli.serviceInspectFunc(serviceID)
	}
	return swarm.Service{}, []byte{}, nil
}
