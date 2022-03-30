package errorno

var MsgFlags = map[uint16]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "invalid params",

	CREATE_STREAM_ERROR: "create stream error",
	LIST_STREAMS_ERROR:  "list stream error",

	CREATE_SUBSCRIPTION_ERROR: "create subscription error",
	LIST_SUBSCRIPTIONS_ERROR:  "list subscription error",
}

// GetMsg convert error code to error message
func GetMsg(code uint16) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
