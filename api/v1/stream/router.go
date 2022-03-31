package stream

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup, s *Service) {
	endpoint := r.Group("/streams")
	endpoint.GET("/", s.List)
	endpoint.GET("/:streamName", s.Get)
	endpoint.POST("/", s.Create)
	endpoint.POST("/:streamName", s.Append)
	endpoint.DELETE("/:streamName", s.Delete)
}
