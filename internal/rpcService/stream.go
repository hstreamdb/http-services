package rpcService

import (
	"github.com/hstreamdb/hstreamdb-go/hstream"
	"github.com/hstreamdb/http-server/api/model"
)

func (c *HStreamClient) CreateStream(stream model.Stream) error {
	return c.client.CreateStream(stream.StreamName,
		hstream.WithReplicationFactor(stream.ReplicationFactor),
		hstream.EnableBacklog(stream.BacklogDuration))
}

func (c *HStreamClient) ListStreams() ([]model.Stream, error) {
	streams, err := c.client.ListStreams()
	if err != nil {
		return nil, err
	}
	var res []model.Stream
	for ; streams.Valid(); streams.Next() {
		item := streams.Item()
		res = append(res, model.Stream{
			StreamName:        item.GetStreamName(),
			ReplicationFactor: item.GetReplicationFactor(),
			BacklogDuration:   item.GetBacklogDuration(),
		})
	}
	return res, nil
}
