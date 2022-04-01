package rpcService

import (
	"github.com/hstreamdb/hstreamdb-go/hstream"
	"github.com/hstreamdb/http-server/pkg/util"
	"go.uber.org/zap"
)

type HStreamClient struct {
	client *hstream.HStreamClient
	// hstreamdb server infos
	clusterInfo string
}

func NewHStreamClient(clusterInfo string) (*HStreamClient, error) {
	client, err := hstream.NewHStreamClient(clusterInfo)
	if err != nil {
		util.Logger().Error("failed to create hstream client", zap.Error(err))
		return nil, err
	}
	return &HStreamClient{
		client:      client,
		clusterInfo: clusterInfo,
	}, nil
}

func (c *HStreamClient) Close() {
	c.client.Close()
}
