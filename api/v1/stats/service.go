package stats

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hstreamdb/hstreamdb-go/hstream"
	"github.com/hstreamdb/http-server/api/model"
	"net/http"
	"strconv"
)

type Service struct {
	client *hstream.HStreamClient
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
	aggAvg(s, c, "server stats histogram server_append_latency")
}

func (s *Service) GetServerAppendLatency(c *gin.Context) {
	aggAvg(s, c, "server stats histogram server_append_latency")
}

// FIXME
func getFromCluster(s *Service, c *gin.Context, cmd string) []model.TableResult {
	info, err := s.client.GetServerInfo()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Errorf("%v", err))
		return nil
	}

	allStatsRaw := make([]model.TableResult, 0)

	for _, x := range info {
		func() {
			client, err := hstream.NewHStreamClient(x)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Errorf("%v", err))
				return
			}
			defer client.Close()

			resp, err := client.AdminRequest(cmd)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Errorf("%v", err))
				return
			}

			var jsonObj map[string]json.RawMessage
			if err := json.Unmarshal([]byte(resp), &jsonObj); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Errorf("%v", err))
				return
			}

			var table model.TableType
			if content, ok := jsonObj["content"]; ok {
				if err = json.Unmarshal(content, &table); err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Errorf("%v", err))
					return
				}
			} else {
				// FIXME
				c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Errorf("%v", err))
				return
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
		}()
	}

	return allStatsRaw
}

func aggAppendSum(s *Service, c *gin.Context, cmd string) {

	allStatsRaw := getFromCluster(s, c, cmd)
	if allStatsRaw == nil {
		// FIXME
		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Errorf("%v", nil))
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
				// FIXME
				allStats.Value = append(allStats.Value, val)
			}
		}

		c.JSON(http.StatusOK, allStats)
	} else {
		// FIXME
		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Errorf("%v", nil))
		return
	}
}

func aggAvg(s *Service, c *gin.Context, cmd string) {
	allStatsRaw := getFromCluster(s, c, cmd)
	if allStatsRaw == nil {
		// FIXME
		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Errorf("%v", nil))
		return
	}

	if len(allStatsRaw) != 0 || len(allStatsRaw[0].Value) != 1 {
		headerLen := len(allStatsRaw[0].Headers)
		acc := make([]float64, headerLen, 0.0)
		for _, stats := range allStatsRaw {
			for _, val := range stats.Value {
				for i, k := range stats.Headers {
					val, err := strconv.ParseFloat(val[k], 64)
					if err != nil {
						c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Errorf("%v", err))
						return
					}
					acc[i] += val
				}
			}
		}

	} else {
		// FIXME
		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Errorf("%v", nil))
		return
	}
}
