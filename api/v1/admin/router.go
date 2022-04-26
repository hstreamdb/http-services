package admin

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.RouterGroup, s *Service) {
	endpoint := r.Group("/cluster")
	endpoint.GET("/status", s.GetStatus)
	// v1/admin/stats?method=append&interval=1s&interval=5s
	endpoint.GET("/stats", s.GetStats)
}
