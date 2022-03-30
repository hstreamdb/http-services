package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hstreamdb/http-server/api/v1/stream"
	"github.com/hstreamdb/http-server/api/v1/subscription"
)

type ServiceClient interface {
	stream.StreamServices
	subscription.SubServices
}

func InitRouter(client ServiceClient) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("/v1")

	streamService := stream.NewStreamService(client)
	subServices := subscription.NewSubService(client)
	stream.RegisterRouter(v1, streamService)
	subscription.RegisterRouter(v1, subServices)
	return router
}
