package rpcService

import (
	"github.com/hstreamdb/hstreamdb-go/hstream"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/hstreamdb/http-server/pkg/util"
	"go.uber.org/zap"
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
	for _, stream := range streams {
		res = append(res, model.Stream{
			StreamName:        stream.StreamName,
			ReplicationFactor: stream.ReplicationFactor,
			BacklogDuration:   stream.BacklogDuration,
		})
	}
	return res, nil
}

func (c *HStreamClient) DeleteStream(streamName string) error {
	return c.client.DeleteStream(streamName)
}

func (c *HStreamClient) Append(streamName string, record model.Record) (rid model.RecordId, err error) {
	producer := c.client.NewProducer(streamName)
	var r *hstream.HStreamRecord
	switch record.Type {
	case "RAW":
		r, err = hstream.NewHStreamRawRecord(record.Key, []byte(record.Data.(string)))
	case "HRECORD":
		r, err = hstream.NewHStreamHRecord(record.Key, record.Data.(map[string]interface{}))
	}
	if err != nil {
		util.Logger().Error("Error creating record: %v", zap.Error(err))
		return
	}

	res, err := producer.Append(r).Ready()
	if err != nil {
		return
	}

	return recordIdFromHStream(res), nil
}

func recordIdFromHStream(rid hstream.RecordId) model.RecordId {
	return model.RecordId{
		BatchId:    rid.BatchId,
		BatchIndex: rid.BatchIndex,
		ShardId:    rid.ShardId,
	}
}
