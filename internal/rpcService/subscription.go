package rpcService

import (
	"github.com/hstreamdb/http-server/api/model"
)

func (c *HStreamClient) CreateSubscription(sub model.Subscription) error {
	return c.client.CreateSubscription(sub.SubscriptionId, sub.StreamName, sub.AckTimeoutSeconds)
}

func (c *HStreamClient) ListSubscriptions() ([]model.Subscription, error) {
	streams, err := c.client.ListSubscriptions()
	if err != nil {
		return nil, err
	}
	var res []model.Subscription
	for ; streams.Valid(); streams.Next() {
		item := streams.Item()
		res = append(res, model.Subscription{
			SubscriptionId:    item.GetSubscriptionId(),
			StreamName:        item.GetStreamName(),
			AckTimeoutSeconds: item.GetAckTimeoutSeconds(),
		})
	}
	return res, nil
}
