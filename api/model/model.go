package model

type Stream struct {
	StreamName        string `json:"stream_name" binding:"required"`
	ReplicationFactor uint32 `json:"replication_factor"`
	BacklogDuration   uint32 `json:"backlog_duration"`
}

type Subscription struct {
	SubscriptionId    string `json:"subscription_id" binding:"required"`
	StreamName        string `json:"stream_name" binding:"required"`
	AckTimeoutSeconds int32  `json:"ack_timeout_seconds"`
	MaxUnackedRecords int32  `json:"max_unacked_records"`
}

type Record struct {
	Key string `json:"key" binding:"required"`
	// Record Type:
	// * RAW - []byte payload
	// * HRECORD - JSON payload
	Type string      `json:"type" binding:"required" enums:"RAW,HRECORD"`
	Data interface{} `json:"data" binding:"required"`
}

type RecordId struct {
	BatchId    uint64 `json:"batch_id"`
	BatchIndex uint32 `json:"batch_index"`
	ShardId    uint64 `json:"shard_id"`
}

type TableType struct {
	Type    string `json:"type"`
	Content struct {
		Headers []string   `json:"headers"`
		Rows    [][]string `json:"rows"`
	}
}

type StatsRequestArg struct {
	Method    string   `json:"method"`
	Intervals []string `json:"intervals"`
}

type TableResult struct {
	Value []map[string]interface{} `json:"value"`
}
