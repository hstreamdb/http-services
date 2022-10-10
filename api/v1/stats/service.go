package stats

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/hstreamdb/http-server/api/model"
//	"github.com/hstreamdb/http-server/pkg/errorno"
//	"github.com/hstreamdb/http-server/pkg/util"
//	"net/http"
//	"strconv"
//)
//
//type StatsServices interface {
//	GetServerInfo() ([]string, error)
//	GetStatsFromServer(string, string) (*model.TableType, error)
//}
//
//type Service struct {
//	client StatsServices
//}
//
//func NewStatsServices(s StatsServices) *Service {
//	return &Service{s}
//}
//
//// GetAppends godoc
//// @ID streamAppendsStats
//// @Summary Get rate of bytes successfully written to the stream.
//// @Success 200 {object} model.TableResult
//// @Failure 500 {object} errorno.ErrorResponse
//// @Router /v1/stats/append/bytes [get]
//func (s *Service) GetAppends(c *gin.Context) {
//	s.aggAppendSum(c, "server stats stream appends")
//}
//
//// GetSends godoc
//// @ID subscriptionSendsStats
//// @Summary Get rate of bytes sent by the server per subscription
//// @Success 200 {object} model.TableResult
//// @Failure 500 {object} errorno.ErrorResponse
//// @Router /v1/stats/sends/bytes [get]
//func (s *Service) GetSends(c *gin.Context) {
//	s.aggAppendSum(c, "server stats subscription sends")
//}
//
//// GetAppendInRecords godoc
//// @ID appendInRecordStats
//// @Summary Get rate of records successfully written to the stream
//// @Success 200 {object} model.TableResult
//// @Failure 500 {object} errorno.ErrorResponse
//// @Router /v1/stats/append/records [get]
//func (s *Service) GetAppendInRecords(c *gin.Context) {
//	s.aggAppendSum(c, "server stats stream append_in_record")
//}
//
//// GetSendOutRecords godoc
//// @ID sendOutRecords
//// @Summary Get rate of records successfully sent by the server per subscription
//// @Success 200 {object} model.TableResult
//// @Failure 500 {object} errorno.ErrorResponse
//// @Router /v1/stats/send/records [get]
//func (s *Service) GetSendOutRecords(c *gin.Context) {
//	s.aggAppendSum(c, "server stats subscription send_out_records")
//}
//
//// GetAppendTotal godoc
//// @ID totalAppendStats
//// @Summary Get total number of success append requests of a stream
//// @Success 200 {object} model.TableResult
//// @Failure 500 {object} errorno.ErrorResponse
//// @Router /v1/stats/append/success [get]
//func (s *Service) GetAppendTotal(c *gin.Context) {
//	s.aggAppendSum(c, "server stats stream_counter append_total")
//}
//
//// GetAppendFailed godoc
//// @ID failedAppendStats
//// @Summary Get total number of failed append request of a stream
//// @Success 200 {object} model.TableResult
//// @Failure 500 {object} errorno.ErrorResponse
//// @Router /v1/stats/append/failed [get]
//func (s *Service) GetAppendFailed(c *gin.Context) {
//	s.aggAppendSum(c, "server stats stream_counter append_failed")
//}
//
//// GetAppendQPS godoc
//// @ID appendQPS
//// @Summary Get qps for each stream
//// @Success 200 {object} model.TableResult
//// @Failure 500 {object} errorno.ErrorResponse
//// @Router /v1/stats/append/qps [get]
//func (s *Service) GetAppendQPS(c *gin.Context) {
//	s.aggAppendSum(c, "server stats stream append_in_requests")
//}
//
//// GetServerAppendRequestLatency godoc
//// @ID appendRequestLatency
//// @Summary Get stream append request latency stats
//// @Success 200 {object} model.TableResult
//// @Failure 500 {object} errorno.ErrorResponse
//// @Router /v1/stats/histogram/server_append_request_latency [get]
//func (s *Service) GetServerAppendRequestLatency(c *gin.Context) {
//	s.aggLatencyHist(c, "server stats server_histogram append_request_latency")
//}
//
//// GetServerAppendLatency godoc
//// @ID appendLatency
//// @Summary Get stream append latency stats
//// @Success 200 {object} model.TableResult
//// @Failure 500 {object} errorno.ErrorResponse
//// @Router /v1/stats/histogram/server_append_latency [get]
//func (s *Service) GetServerAppendLatency(c *gin.Context) {
//	s.aggLatencyHist(c, "server stats server_histogram append_latency")
//}
//
//func adminError(err error) errorno.ErrorResponse {
//	return errorno.NewErrorResponse(errorno.ADMIN_GET_STATUS_ERROR, err)
//}
//
//func getFromCluster(s *Service, cmd string) (map[string]*model.TableType, error) {
//	info, err := s.client.GetServerInfo()
//	if err != nil {
//		return nil, fmt.Errorf("GetServerInfo error: %v", err)
//	}
//
//	res := make(map[string]*model.TableType, len(info))
//
//	for _, addr := range info {
//		respTable, err := s.client.GetStatsFromServer(addr, cmd)
//		if err != nil {
//			return nil, fmt.Errorf("AdminRequest error with %v: %v", cmd, err)
//		}
//
//		res[addr] = respTable
//	}
//	return res, nil
//}
//
//func (s *Service) aggAppendSum(c *gin.Context, cmd string) {
//	records, err := getFromCluster(s, cmd)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(err))
//		return
//	}
//	if len(records) == 0 {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(fmt.Errorf("got stats of length 0")))
//		return
//	}
//
//	dataVec := make([]*model.TableType, 0, len(records))
//	for _, v := range records {
//		dataVec = append(dataVec, v)
//	}
//
//	res, err := sum(dataVec)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(err))
//		return
//	}
//	c.JSON(http.StatusOK, *res)
//}
//
//func (s *Service) aggLatencyHist(c *gin.Context, cmd string) {
//	records, err := getFromCluster(s, cmd)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(err))
//		return
//	}
//	if len(records) == 0 {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, adminError(fmt.Errorf("got stats of length 0")))
//		return
//	}
//
//	var headers []string
//	setHeader := false
//	values := make([]map[string]string, 0, len(records))
//	for addr, table := range records {
//		if !setHeader {
//			headers = make([]string, len(table.Headers)+1)
//			headers[0] = "server-address"
//			copy(headers[1:], table.Headers)
//			setHeader = true
//			util.Logger().Debug(fmt.Sprintf("headers = %s", headers))
//		}
//
//		mp := make(map[string]string, len(table.Headers)+1)
//		for _, rows := range table.Rows {
//			for idx := 0; idx < len(rows); idx++ {
//				mp[headers[idx+1]] = rows[idx]
//			}
//			mp[headers[0]] = addr
//		}
//		values = append(values, mp)
//		util.Logger().Debug(fmt.Sprintf("%+v", mp))
//	}
//
//	res := model.TableResult{
//		Headers: headers,
//		Value:   values,
//	}
//	c.JSON(http.StatusOK, res)
//}
//
//func sum(records []*model.TableType) (*model.TableResult, error) {
//	headers := records[0].Headers
//	// statistics: {resource: {metrics1: value1, metrics2: value2}}
//	statistics := map[string]map[string]int64{}
//	dataSize := len(records[0].Headers) - 1
//	for _, table := range records {
//		// e.g. table.Rows[0]
//		// |stream_name|appends_1min|appends_5min|appends_10min|  <- headers
//		// |    s1     |      0     |    1829    |    1829     |  <- row0
//		// |    s2     |      0     |    7270    |    7270     |  <- row1
//		for _, rows := range table.Rows {
//			target := rows[0]
//			if _, ok := statistics[target]; !ok {
//				statistics[target] = make(map[string]int64, dataSize)
//			}
//
//			for idx := 1; idx <= dataSize; idx++ {
//				value, err := strconv.ParseInt(rows[idx], 10, 64)
//				if err != nil {
//					return nil, err
//				}
//				statistics[target][headers[idx]] += value
//			}
//			util.Logger().Debug(fmt.Sprintf("[%s] = %s, val: %+v", headers[0], target, rows))
//		}
//	}
//
//	// construct result
//	rows := []map[string]string{}
//	for resource, metricsMp := range statistics {
//		res := map[string]string{}
//		res[headers[0]] = resource
//		for metrics, v := range metricsMp {
//			res[metrics] = strconv.FormatInt(v, 10)
//		}
//		rows = append(rows, res)
//	}
//
//	allStats := model.TableResult{
//		Headers: headers,
//		Value:   rows,
//	}
//	util.Logger().Debug(fmt.Sprintf("res = %+v", statistics))
//	return &allStats, nil
//}
