package stats

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.RouterGroup, s *Service) {
	endpoint := r.Group("/stats")

	endpoint.GET("/append/bytes", s.GetAppends)
	endpoint.GET("/send/bytes", s.GetSends)

	endpoint.GET("/append/records", s.GetAppendInRecords)
	endpoint.GET("/send/records", s.GetSendOutRecords)

	endpoint.GET("/append/success", s.GetAppendTotal)
	endpoint.GET("/append/failed", s.GetAppendFailed)

	endpoint.GET("/histogram/server_append_request_latency", s.GetServerAppendRequestLatency)
	endpoint.GET("/histogram/server_append_latency", s.GetServerAppendLatency)
}
