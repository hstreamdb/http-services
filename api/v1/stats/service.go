package stats

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/hstreamdb/http-server/pkg/errorno"
	"net/http"
	"strconv"
)

type StatsServices interface {
	GetServerInfo() ([]string, error)
	GetStatsFromServer(string, string) (*model.TableType, error)
}

type Service struct {
	client StatsServices
}

func NewStatsServices(s StatsServices) *Service {
	return &Service{s}
}

func (s *Service) GetAppends(c *gin.Context) {
	aggAppendSum(s, c, "server stats stream appends")
}

func (s *Service) GetSends(c *gin.Context) {
	aggAppendSum(s, c, "server stats subscription sends")
}

func (s *Service) GetAppendInRecords(c *gin.Context) {
	aggAppendSum(s, c, "server stats stream append_in_record")
}

func (s *Service) GetSendOutRecords(c *gin.Context) {
	aggAppendSum(s, c, "server stats subscription send_out_records")
}

func (s *Service) GetAppendTotal(c *gin.Context) {
	aggAppendSum(s, c, "server stats stream_counter append_total")
}

func (s *Service) GetAppendFailed(c *gin.Context) {
	aggAppendSum(s, c, "server stats stream_counter append_failed")
}

func (s *Service) GetServerAppendRequestLatency(c *gin.Context) {
	aggSum(s, c, "server stats server_histogram append_request_latency")
}

func (s *Service) GetServerAppendLatency(c *gin.Context) {
	aggSum(s, c, "server stats server_histogram append_latency")
}

func adminError(err error) errorno.ErrorResponse {
	return errorno.NewErrorResponse(errorno.ADMIN_GET_STATUS_ERROR, err)
}

func getFromCluster(s *Service, cmd string) ([]model.TableResult, error) {
	info, err := s.client.GetServerInfo()
	if err != nil {
		return nil, fmt.Errorf("GetServerInfo error: %v", err)
	}

	allStatsRaw := []model.TableResult{}

	for _, addr := range info {
		err := func() error {
			respTable, err := s.client.GetStatsFromServer(addr, cmd)
			if err != nil {
				return fmt.Errorf("AdminRequest error with %v: %v", cmd, err)
			}

			stats := model.TableResult{
				Headers: respTable.Headers,
				Value:   make([]map[string]string, 0, len(respTable.Rows)),
			}

			for _, row := range respTable.Rows {
				mp := make(map[string]string, len(row))
				for idx, val := range row {
					mp[respTable.Headers[idx]] = val
				}
				stats.Value = append(stats.Value, mp)
			}

			allStatsRaw = append(allStatsRaw, stats)
			return nil
		}()
		if err != nil {
			return nil, err
		}
	}

	return allStatsRaw, nil
}

func aggAppendSum(s *Service, c *gin.Context, cmd string) {
	allStatsRaw, err := getFromCluster(s, cmd)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(err))
		return
	}

	if len(allStatsRaw) == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(fmt.Errorf("got stats of length 0")))
		return
	} else {
		headers := allStatsRaw[0].Headers

		allStats := model.TableResult{
			Headers: headers,
			Value:   []map[string]string{},
		}

		// The `allStatsRaw` contains stats tables that got from the cluster per server, which each is of the form:
		// | main_key | type_0 stats | type_1 stats | ...
		// If the main_key already contains in the container which is prepared for return the result, add to exist stats
		// for each value.  This can happen when a stream, subscription or something else if exists, which transfer to
		// be handled by another server. If main_key does not already contain in, just append the value col to the
		// table result.
		for _, stats := range allStatsRaw {
			for _, val := range stats.Value {
				mainKey := val[headers[0]]

				ok := false
				ix := -1
				for i, cur := range allStats.Value {
					if cur[mainKey] == val[mainKey] {
						ok = true
						ix = i
					}
				}

				if ok {
					for k := range allStats.Value[ix] {
						allStats.Value[ix][k] += val[k]
					}
				} else {
					allStats.Value = append(allStats.Value, val)
				}
			}
		}

		c.JSON(http.StatusOK, allStats)
	}
}

func aggSum(s *Service, c *gin.Context, cmd string) {
	allStatsRaw, err := getFromCluster(s, cmd)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(err))
		return
	}

	if len(allStatsRaw) == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(fmt.Errorf("got stats of length 0")))
		return
	} else if len(allStatsRaw[0].Value) != 1 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(fmt.Errorf("the length value should be 1")))
		return
	} else {
		headerLen := len(allStatsRaw[0].Headers)
		acc := make([]int, headerLen)
		for _, stats := range allStatsRaw {
			for _, val := range stats.Value {
				for i, k := range stats.Headers {
					val, err := strconv.ParseInt(val[k], 10, 32)
					if err != nil {
						c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(fmt.Errorf("%v", err)))
						return
					}
					acc[i] += int(val)
				}
			}
		}

		headers := allStatsRaw[0].Headers
		value := map[string]string{}
		for i, header := range headers {
			value[header] = fmt.Sprintf("%d", acc[i])
		}
		values := []map[string]string{value}
		tab := model.TableResult{Headers: headers, Value: values}

		c.JSON(http.StatusOK, tab)
		return
	}
}
