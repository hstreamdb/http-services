package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/hstreamdb/http-server/pkg/errorno"
	"net/http"
)

type AdminServices interface {
	GetStatus() (*model.TableType, error)
	GetStats(string, string, []string) (*model.TableType, error)
	GetStatsFromAddr(string, string, string, []string) (*model.TableType, error)
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
// @Router /v1/cluster/status [get]
func (s *Service) GetStatus(c *gin.Context) {
	resp, err := s.client.GetStatus()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorno.NewErrorResponse(errorno.ADMIN_GET_STATUS_ERROR, err))
		return
	}

	status := model.TableResult{
		Headers: resp.Headers,
		Value:   make([]map[string]string, 0, len(resp.Rows)),
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
// @Param category query string true "Category"
// @Param metrics query string true "Metrics"
// @Param interval query []string false "Interval collection" collectionFormat(multi)
// @Success 200 {object} model.TableResult
// @Failure 500 {object} errorno.ErrorResponse
// @Router /v1/cluster/stats [get]
func (s *Service) GetStats(c *gin.Context) {
	category := c.Query("category")
	metrics := c.Query("metrics")
	intervals := c.QueryArray("interval")
	addr := c.Query("addr")

	var resp *model.TableType
	var err error
	if addr != "" {
		resp, err = s.client.GetStatsFromAddr(addr, category, metrics, intervals)

	} else {
		resp, err = s.client.GetStats(category, metrics, intervals)
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorno.NewErrorResponse(errorno.ADMIN_GET_STATUS_ERROR, err))
		return
	}

	stats := model.TableResult{
		Headers: resp.Headers,
		Value:   make([]map[string]string, 0, len(resp.Rows)),
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
