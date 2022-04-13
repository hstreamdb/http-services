package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/hstreamdb/http-server/pkg/errorno"
	"net/http"
)

type AdminServices interface {
	GetStatus() (model.TableType, error)
	GetStats(string, []string) (model.TableType, error)
}

type Service struct {
	client AdminServices
}

func NewAdminService(client AdminServices) *Service {
	return &Service{client}
}

// GetStatus godoc
// @ID statusGet
// @Summary Get server status of the cluster
// @Success 200 {object} model.TableResult
// @Failure 500 {object} errorno.ErrorResponse
// @Router /v1/admin/status [get]
func (s *Service) GetStatus(c *gin.Context) {
	resp, err := s.client.GetStatus()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorno.NewErrorResponse(errorno.ADMIN_GET_STATUS_ERROR, err))
		return
	}

	status := model.TableResult{
		Value: make([]map[string]interface{}, 0, len(resp.Content.Rows)),
	}

	for _, row := range resp.Content.Rows {
		mp := make(map[string]interface{}, len(row))
		for idx, val := range row {
			mp[resp.Content.Headers[idx]] = val
		}
		status.Value = append(status.Value, mp)
	}
	c.JSON(http.StatusOK, status)
}

// // @Param request body model.StatsRequestArg true "Request body"

// GetStats godoc
// @ID statsGet
// @Summary Get cluster stats
// @Param method query string true "Method"
// @Param interval query []string true "Interval collection" collectionFormat(multi)
// @Success 200 {object} model.TableResult
// @Failure 500 {object} errorno.ErrorResponse
// @Router /v1/admin/stats [get]
func (s *Service) GetStats(c *gin.Context) {
	method := c.Query("method")
	interval := c.QueryArray("interval")
	resp, err := s.client.GetStats(method, interval)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorno.NewErrorResponse(errorno.ADMIN_GET_STATUS_ERROR, err))
		return
	}

	//var reqArg model.StatsRequestArg
	//if err := c.ShouldBindJSON(&reqArg); err != nil {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, errorno.NewErrorResponse(errorno.INVALID_PARAMETER, err))
	//	return
	//}
	//
	//resp, err := s.client.GetStats(reqArg.Method, reqArg.Intervals)

	stats := model.TableResult{
		Value: make([]map[string]interface{}, 0, len(resp.Content.Rows)),
	}

	for _, row := range resp.Content.Rows {
		mp := make(map[string]interface{}, len(row))
		for idx, val := range row {
			mp[resp.Content.Headers[idx]] = val
		}
		stats.Value = append(stats.Value, mp)
	}
	c.JSON(http.StatusOK, resp)
}
