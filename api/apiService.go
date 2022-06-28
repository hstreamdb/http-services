package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hstreamdb/http-server/api/swagger"
	"github.com/hstreamdb/http-server/api/v1/admin"
	"github.com/hstreamdb/http-server/api/v1/stats"
	"github.com/hstreamdb/http-server/api/v1/stream"
	"github.com/hstreamdb/http-server/api/v1/subscription"
)

type ServiceClient interface {
	stream.StreamServices
	subscription.SubServices
	admin.AdminServices
	stats.StatsServices
}

func InitRouter(client ServiceClient) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("/v1")
	swagger.RegisterRouter(v1)

	streamService := stream.NewStreamService(client)
	subServices := subscription.NewSubService(client)
	adminServices := admin.NewAdminService(client)
	statsServices := stats.NewStatsServices(client)
	stream.RegisterRouter(v1, streamService)
	subscription.RegisterRouter(v1, subServices)
	admin.RegisterRouter(v1, adminServices)
	stats.RegisterRouter(v1, statsServices)
	return router
}
