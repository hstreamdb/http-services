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

func (c *HStreamClient) DeleteStream(streamName string) error {
	return c.client.DeleteStream(streamName)
}

func (c *HStreamClient) Append(streamName string, record model.Record) (rid model.RecordId, err error) {
	producer := c.client.NewProducer(streamName)
	var r hstream.HStreamRecord
	switch record.Type {
	case "RAW":
		r = hstream.NewHStreamRawRecord(record.Key, []byte(record.Data.(string)))
	case "HRECORD":
		r = hstream.NewHStreamHRecord(record.Key, record.Data.(map[string]interface{}))
	}

	res, err := producer.Append(r).Ready()
	if err != nil {
		return rid, err
	}

	return recordIdFromHStream(res), nil
}

func recordIdFromHStream(rid *hstream.RecordId) model.RecordId {
	return model.RecordId{
		BatchId:    rid.BatchId,
		BatchIndex: rid.BatchIndex,
		ShardId:    rid.ShardId,
	}
}
