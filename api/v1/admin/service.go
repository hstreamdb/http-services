package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/hstreamdb/http-server/pkg/errorno"
	"github.com/hstreamdb/http-server/pkg/util"
	"net/http"
	"strconv"
	"strings"
)

type AdminServices interface {
	GetServerInfo() ([]string, error)
	GetStatus() (*model.TableType, error)
	GetStats(string, string, string, string) (*model.TableType, error)
}

type Service struct {
	client AdminServices
}

func NewAdminService(client AdminServices) *Service {
	return &Service{client}
}

func adminError(err error) errorno.ErrorResponse {
	return errorno.NewErrorResponse(errorno.ADMIN_GET_STATUS_ERROR, err)
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
	interval := c.QueryArray("interval")
	intervals := strings.Join(append([]string{""}, interval...), "-i ")

	switch category {
	case "server_histogram":
		s.aggLatencyHist(c, category, metrics, intervals)
	default:
		s.aggAppendSum(c, category, metrics, intervals)
	}
}

func (s *Service) getFromCluster(category, metrics, intervals string) (map[string]*model.TableType, error) {
	info, err := s.client.GetServerInfo()
	if err != nil {
		return nil, fmt.Errorf("GetServerInfo error: %v", err)
	}

	res := make(map[string]*model.TableType, len(info))

	for _, addr := range info {
		respTable, err := s.client.GetStats(addr, category, metrics, intervals)
		if err != nil {
			return nil, fmt.Errorf("get stats error: %s", err.Error())
		}

		res[addr] = respTable
	}
	return res, nil
}

func (s *Service) aggAppendSum(c *gin.Context, category, metrics, intervals string) {
	records, err := s.getFromCluster(category, metrics, intervals)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(err))
		return
	}
	if len(records) == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(fmt.Errorf("got stats of length 0")))
		return
	}

	dataVec := make([]*model.TableType, 0, len(records))
	servers := make([]string, 0, len(records))
	for addr, v := range records {
		dataVec = append(dataVec, v)
		servers = append(servers, addr)
	}

	res, err := sum(dataVec, servers)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(err))
		return
	}
	c.JSON(http.StatusOK, *res)
}

func (s *Service) aggLatencyHist(c *gin.Context, category, metrics, intervals string) {
	records, err := s.getFromCluster(category, metrics, intervals)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(err))
		return
	}
	if len(records) == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(fmt.Errorf("got stats of length 0")))
		return
	}

	var headers []string
	setHeader := false
	values := make([]map[string]string, 0, len(records))
	for addr, table := range records {
		if !setHeader {
			headers = make([]string, len(table.Headers)+1)
			headers[0] = "server_host"
			copy(headers[1:], table.Headers)
			setHeader = true
			util.Logger().Debug(fmt.Sprintf("headers = %s", headers))
		}

		mp := make(map[string]string, len(table.Headers)+1)
		for _, rows := range table.Rows {
			for idx := 0; idx < len(rows); idx++ {
				mp[headers[idx+1]] = rows[idx]
			}
			host := strings.Split(addr, ":")[0]
			mp[headers[0]] = host
		}
		values = append(values, mp)
		util.Logger().Debug(fmt.Sprintf("%+v", mp))
	}

	res := model.TableResult{
		Headers: headers,
		Value:   values,
	}
	c.JSON(http.StatusOK, res)
}

func sum(records []*model.TableType, servers []string) (*model.TableResult, error) {
	headers := records[0].Headers
	// statistics: {resource: {metrics1: value1, metrics2: value2}}
	statistics := map[string]map[string]int64{}
	dataSize := len(records[0].Headers) - 1
	resourceOriginMp := map[string]string{}
	for i, table := range records {
		serverUrl := strings.Split(servers[i], ":")[0]
		// e.g. table.Rows[0]
		// |stream_name|appends_1min|appends_5min|appends_10min|  <- headers
		// |    s1     |      0     |    1829    |    1829     |  <- row0
		// |    s2     |      0     |    7270    |    7270     |  <- row1
		for _, rows := range table.Rows {
			target := rows[0]
			if _, ok := statistics[target]; !ok {
				statistics[target] = make(map[string]int64, dataSize)
				resourceOriginMp[target] = serverUrl
			}

			for idx := 1; idx <= dataSize; idx++ {
				value, err := strconv.ParseInt(rows[idx], 10, 64)
				if err != nil {
					return nil, err
				}
				statistics[target][headers[idx]] += value
			}
		}
	}

	// construct result
	rows := []map[string]string{}
	for resource, metricsMp := range statistics {
		res := map[string]string{}
		res[headers[0]] = resource
		for metrics, v := range metricsMp {
			res[metrics] = strconv.FormatInt(v, 10)
		}
		res["server_host"] = resourceOriginMp[resource]
		rows = append(rows, res)
	}

	headers = append([]string{"server_host"}, headers...)
	allStats := model.TableResult{
		Headers: headers,
		Value:   rows,
	}
	util.Logger().Debug(fmt.Sprintf("res = %+v", statistics))
	return &allStats, nil
}
