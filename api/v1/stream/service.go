package stream

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/hstreamdb/http-server/pkg/errorno"
	"net/http"
	"sort"
)

type StreamServices interface {
	CreateStream(stream model.Stream) error
	ListStreams() ([]model.Stream, error)
	DeleteStream(StreamName string) error
	Append(streamName string, record model.Record) (model.RecordId, error)
}

type Service struct {
	client StreamServices
}

func NewStreamService(client StreamServices) *Service {
	return &Service{client}
}

// Get godoc
// @ID streamGet
// @Summary Get specific stream by streamName
// @Param streamName path string true "Stream name"
// @Success 200 {object} model.Stream
// @Failure 404 {object} errorno.ErrorResponse
// @Failure 500 {object} errorno.ErrorResponse
// @Router /v1/streams/{streamName} [get]
func (s *Service) Get(c *gin.Context) {
	target := c.Param("streamName")
	streams, err := s.client.ListStreams()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorno.NewErrorResponse(errorno.LIST_STREAMS_ERROR, err))
		return
	}
	sort.Slice(streams, func(i, j int) bool {
		return streams[i].StreamName < streams[j].StreamName
	})
	idx := sort.Search(len(streams), func(i int) bool {
		return streams[i].StreamName >= target
	})
	if idx < len(streams) && streams[idx].StreamName == target {
		c.JSON(http.StatusOK, streams[idx])
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound,
			errorno.NewErrorResponse(errorno.STREAM_NOT_EXIST, fmt.Errorf("stream %s not exist", target)))
	}
}

// List godoc
// @ID streamList
// @Summary List all streams in the cluster
// @Success 200 {object} []model.Stream
// @Failure 400 {object} errorno.ErrorResponse
// @Failure 500 {object} errorno.ErrorResponse
// @Router /v1/streams/ [get]
func (s *Service) List(c *gin.Context) {
	streams, err := s.client.ListStreams()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorno.NewErrorResponse(errorno.LIST_STREAMS_ERROR, err))
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
// @Failure 500 {object} errorno.ErrorResponse
// @Router /v1/streams/ [post]
func (s *Service) Create(c *gin.Context) {
	var stream model.Stream
	if err := c.ShouldBindJSON(&stream); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorno.NewErrorResponse(errorno.INVALID_PARAMETER, err))
		return
	}
	if err := s.client.CreateStream(stream); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorno.NewErrorResponse(errorno.CREATE_STREAM_ERROR, err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

// Delete godoc
// @ID streamDelete
// @Summary Delete specific stream by streamName
// @Param streamName path string true "Stream Name"
// @Success 200 {object} string "ok"
// @Failure 500 {object} errorno.ErrorResponse
// @Router /v1/streams/{streamName} [Delete]
func (s *Service) Delete(c *gin.Context) {
	streamName := c.Param("streamName")
	if err := s.client.DeleteStream(streamName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorno.NewErrorResponse(errorno.DELETE_STREAM_ERROR, err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

// Append godoc
// @ID streamAppend
// @Summary Append record to specific stream
// @Param request body model.Record true "Request body"
// @Param streamName path string true "Stream name"
// @Success 200 {object} model.RecordId
// @Failure 400 {object} errorno.ErrorResponse
// @Failure 500 {object} errorno.ErrorResponse
// @Router /v1/streams/{streamName} [post]
func (s *Service) Append(c *gin.Context) {
	var record model.Record
	streamName := c.Param("streamName")
	if err := c.ShouldBindJSON(&record); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorno.NewErrorResponse(errorno.INVALID_PARAMETER, err))
		return
	}
	rid, err := s.client.Append(streamName, record)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorno.NewErrorResponse(errorno.APPEND_RECORD_ERROR, err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"rid": rid,
	})
}
