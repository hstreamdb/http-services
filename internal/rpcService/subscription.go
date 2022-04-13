package rpcService

import (
	"github.com/hstreamdb/http-server/api/model"
)

func (c *HStreamClient) CreateSubscription(sub model.Subscription) error {
	return c.client.CreateSubscription(sub.SubscriptionId, sub.StreamName, sub.AckTimeoutSeconds)
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
	return c.client.DeleteSubscription(subId)
}
