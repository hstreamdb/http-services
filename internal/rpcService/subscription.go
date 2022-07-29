package rpcService

import (
	"github.com/hstreamdb/hstreamdb-go/hstream"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/pkg/errors"
)

func (c *HStreamClient) CreateSubscription(sub model.Subscription) error {
	var offset hstream.SubscriptionOffset
	switch sub.Offset {
	case "EARLIEST":
		offset = hstream.EARLIEST
	case "LATEST":
		offset = hstream.LATEST
	default:
		return errors.New("Unknown offset")
	}
	return c.client.CreateSubscription(sub.SubscriptionId, sub.StreamName,
		hstream.WithAckTimeout(sub.AckTimeoutSeconds),
		hstream.WithMaxUnackedRecords(sub.MaxUnackedRecords),
		hstream.WithOffset(offset))
}

func (c *HStreamClient) ListSubscriptions() ([]model.Subscription, error) {
	subs, err := c.client.ListSubscriptions()
	if err != nil {
		return nil, err
	}
	var res []model.Subscription
	for _, sub := range subs {
		res = append(res, model.Subscription{
			SubscriptionId:    sub.SubscriptionId,
			StreamName:        sub.StreamName,
			AckTimeoutSeconds: sub.AckTimeoutSeconds,
			MaxUnackedRecords: sub.MaxUnackedRecords,
		})
	}
	return res, nil
}

func (c *HStreamClient) DeleteSubscription(subId string) error {
	return c.client.DeleteSubscription(subId, true)
}
