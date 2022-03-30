package stream

import (
	"github.com/gin-gonic/gin"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/hstreamdb/http-server/pkg/errorno"
)

type StreamServices interface {
	CreateStream(stream model.Stream) error
	ListStreams() ([]model.Stream, error)
}

type Service struct {
	client StreamServices
}

func NewStreamService(client StreamServices) *Service {
	return &Service{client}
}

// Get godoc
// @Summary Get stream
// @Description Get stream by streamName
// @Tags Stream
// @Accept  json
// @Produce  json
// @Param streamName path string true "Stream name"
// @Success 200 {object} rpcService.Stream
// @Failure 400 {object} rpcService.HTTPError
// @Failure 404 {object} rpcService.HTTPError
// @Failure 500 {object} rpcService.HTTPError
// @Router /streams/{streamName} [get]
func (s *Service) Get(c *gin.Context) {
	//name := c.Param("streamName")

}

func (s *Service) List(c *gin.Context) {
	streams, err := s.client.ListStreams()
	if err != nil {
		c.JSON(errorno.LIST_STREAMS_ERROR, gin.H{"error": err.Error()})
		return
	}
	c.JSON(errorno.SUCCESS, gin.H{"streams": streams})
}

func (s *Service) Create(c *gin.Context) {
	var stream model.Stream
	if err := c.ShouldBindJSON(&stream); err != nil {
		c.JSON(errorno.INVALID_PARAMS, gin.H{"error": err.Error()})
		return
	}
	if err := s.client.CreateStream(stream); err != nil {
		c.JSON(errorno.CREATE_STREAM_ERROR, gin.H{"error": err.Error()})
		return
	}
	c.JSON(errorno.SUCCESS, gin.H{
		"msg": "success",
	})
}

func (s *Service) Delete(c *gin.Context) {}
