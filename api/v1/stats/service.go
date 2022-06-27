package stats

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hstreamdb/hstreamdb-go/hstream"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/hstreamdb/http-server/pkg/errorno"
	"net/http"
	"strconv"
)

type StatsServices interface {
	GetServerInfo() ([]string, error)
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
	aggAvg(s, c, "server stats histogram append_request_latency")
}

func (s *Service) GetServerAppendLatency(c *gin.Context) {
	aggAvg(s, c, "server stats histogram append_latency")
}

func adminError(err error) errorno.ErrorResponse {
	return errorno.NewErrorResponse(errorno.ADMIN_GET_STATUS_ERROR, err)
}

func getFromCluster(s *Service, cmd string) ([]model.TableResult, error) {
	info, err := s.client.GetServerInfo()
	if err != nil {
		return nil, fmt.Errorf("GetServerInfo error: %v", err)
	}

	allStatsRaw := make([]model.TableResult, 0)

	for _, x := range info {
		err := func() error {
			client, err := hstream.NewHStreamClient(x)
			if err != nil {
				return fmt.Errorf("NewHStreamClient error with %v: %v", x, err)
			}
			defer client.Close()

			resp, err := client.AdminRequest(cmd)
			if err != nil {
				return fmt.Errorf("AdminRequest error with %v: %v", cmd, err)
			}

			var jsonObj map[string]json.RawMessage
			if err := json.Unmarshal([]byte(resp), &jsonObj); err != nil {
				return fmt.Errorf("unmarshal to RawMessage error: %v", err)
			}

			var table model.TableType
			if content, ok := jsonObj["content"]; ok {
				if err = json.Unmarshal(content, &table); err != nil {
					return fmt.Errorf("unmarshal to TableType error: %v", err)
				}
			} else {
				return fmt.Errorf("JSON object does not have the `content` field")
			}

			stats := model.TableResult{
				Headers: table.Headers,
				Value:   make([]map[string]string, 0, len(table.Rows)),
			}

			for _, row := range table.Rows {
				mp := make(map[string]string, len(row))
				for idx, val := range row {
					mp[table.Headers[idx]] = val
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

	if len(allStatsRaw) != 0 {
		headers := allStatsRaw[0].Headers

		allStats := model.TableResult{
			Headers: headers,
			Value:   make([]map[string]string, 0),
		}

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
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(fmt.Errorf("got stats of length 0")))
		return
	}
}

func aggAvg(s *Service, c *gin.Context, cmd string) {
	allStatsRaw, err := getFromCluster(s, cmd)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(err))
		return
	}

	if len(allStatsRaw) != 0 || len(allStatsRaw[0].Value) != 1 {
		headerLen := len(allStatsRaw[0].Headers)
		acc := make([]float64, headerLen)
		for _, stats := range allStatsRaw {
			for _, val := range stats.Value {
				for i, k := range stats.Headers {
					val, err := strconv.ParseFloat(val[k], 64)
					if err != nil {
						c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(fmt.Errorf("%v", err)))
						return
					}
					acc[i] += val
				}
			}
		}
		l := len(allStatsRaw)
		for i := range acc {
			acc[i] /= float64(l)
		}

		c.JSON(http.StatusOK, nil)

	} else {
		var err errorno.ErrorResponse
		if len(allStatsRaw) != 0 {
			err = adminError(fmt.Errorf("got stats of length 0"))
		} else if len(allStatsRaw[0].Value) != 1 {
			err = adminError(fmt.Errorf("the length value should be 1"))
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
}
