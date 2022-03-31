package stream

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/hstreamdb/http-server/pkg/errorno"
	"net/http"
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

func (s *Service) Get(c *gin.Context) {
	//name := c.Param("streamName")

}

// List godoc
// @ID streamList
// @Summary List all streams in the cluster
// @Success 200 {object} []model.Stream
// @Failure 400 {object} errorno.ErrorResponse
// @Router /v1/streams/ [get]
func (s *Service) List(c *gin.Context) {
	streams, err := s.client.ListStreams()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorno.NewErrorResponse(errorno.LIST_STREAMS_ERROR, err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"streams": streams})
}

// Create godoc
// @ID streamCreate
// @Summary Create a stream
// @Param request body model.Stream true "Request body"
// @Success 200 {string} string "ok"
// @Failure 400 {object} errorno.ErrorResponse
// @Router /v1/streams/ [post]
func (s *Service) Create(c *gin.Context) {
	var stream model.Stream
	if err := c.ShouldBindJSON(&stream); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorno.NewErrorResponse(errorno.INVALID_PARAMETER, err))
		return
	}
	if err := s.client.CreateStream(stream); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorno.NewErrorResponse(errorno.CREATE_STREAM_ERROR, err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func (s *Service) Delete(c *gin.Context) {}
