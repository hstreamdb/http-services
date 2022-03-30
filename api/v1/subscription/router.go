package subscription

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.RouterGroup, s *Service) {
	endpoint := r.Group("/subscription")
	endpoint.GET("/", s.List)
	endpoint.GET("/:subId", s.Get)
	endpoint.POST("/", s.Create)
	endpoint.DELETE("/:subId", s.Delete)
}
