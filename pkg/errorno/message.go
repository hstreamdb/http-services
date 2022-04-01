package errorno

import "github.com/hstreamdb/http-server/pkg/util"

var MsgFlags = map[ErrorCode]string{
	CREATE_STREAM_ERROR: "create stream error",
	LIST_STREAMS_ERROR:  "list stream error",
	STREAM_NOT_EXIST:    "stream not exist",
	APPEND_RECORD_ERROR: "append record error",
	DELETE_STREAM_ERROR: "delete stream error",

	CREATE_SUBSCRIPTION_ERROR: "create subscription error",
	LIST_SUBSCRIPTIONS_ERROR:  "list subscription error",
	SUBSCRIPTION_NOT_EXIST:    "subscription not exist",
	DELETE_SUBSCRIPTION_ERROR: "delete subscription error",
}

// GetMsg convert error code to error message
func GetMsg(code ErrorCode) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	util.Logger().Fatal("error code not found")
	return ""
}
