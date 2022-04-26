package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/hstreamdb/http-server/pkg/errorno"
	"net/http"
)

type AdminServices interface {
	GetStatus() (*model.TableType, error)
	GetStats(string, []string) (*model.TableType, error)
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
		Value: make([]map[string]string, 0, len(resp.Rows)),
	}

	for _, row := range resp.Rows {
		mp := make(map[string]string, len(row))
		for idx, val := range row {
			mp[resp.Headers[idx]] = val
		}
		status.Value = append(status.Value, mp)
	}
	c.JSON(http.StatusOK, status)
}

// GetStats godoc
// @ID statsGet
// @Summary Get cluster stats
// @Param metrics query string true "Metrics"
// @Param interval query []string true "Interval collection" collectionFormat(multi)
// @Success 200 {object} model.TableResult
// @Failure 500 {object} errorno.ErrorResponse
// @Router /v1/admin/stats [get]
func (s *Service) GetStats(c *gin.Context) {
	metrics := c.Query("metrics")
	interval := c.QueryArray("interval")
	resp, err := s.client.GetStats(metrics, interval)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorno.NewErrorResponse(errorno.ADMIN_GET_STATUS_ERROR, err))
		return
	}

	stats := model.TableResult{
		Value: make([]map[string]string, 0, len(resp.Rows)),
	}

	for _, row := range resp.Rows {
		mp := make(map[string]string, len(row))
		for idx, val := range row {
			mp[resp.Headers[idx]] = val
		}
		stats.Value = append(stats.Value, mp)
	}
	c.JSON(http.StatusOK, stats)
}
