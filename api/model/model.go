package model

type Stream struct {
	StreamName        string `json:"stream_name" binding:"required"`
	ReplicationFactor uint32 `json:"replication_factor"`
	BacklogDuration   uint32 `json:"backlog_duration"`
}

type Subscription struct {
	SubscriptionId    string `json:"subscription_id" binding:"required"`
	StreamName        string `json:"streamName" binding:"required"`
	AckTimeoutSeconds int32  `json:"ack_timeout_seconds"`
}
