package swagger

import (
	"github.com/gin-gonic/gin"
	_ "github.com/hstreamdb/http-server/docs"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func RegisterRouter(router *gin.RouterGroup) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
