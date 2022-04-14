package main

import (
	"fmt"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/hstreamdb/http-server/pkg/util"
	"strconv"
)

func printTableResult(res *model.TableResult) {
	headers := make([]string, 0, len(res.Value[0]))
	for key := range res.Value[0] {
		headers = append(headers, key)
	}
	datas := make([][]string, 0, len(headers))
	for _, value := range res.Value {
		data := make([]string, 0, len(value))
		for _, key := range headers {
			data = append(data, value[key])
		}
		datas = append(datas, data)
	}
	printer, _ := util.RenderTable(headers, datas)
	fmt.Println(printer)
}

func printStreams(streams []model.Stream) {
	headers := []string{"StreamName", "ReplicationFactor", "BacklogDuration"}
	datas := make([][]string, 0, len(streams))
	for i := 0; i < len(streams); i++ {
		data := make([]string, 0, len(headers))
		data = append(data, streams[i].StreamName)
		data = append(data, strconv.Itoa(int(streams[i].ReplicationFactor)))
		data = append(data, strconv.Itoa(int(streams[i].BacklogDuration)))
		datas = append(datas, data)
	}
	printer, _ := util.RenderTable(headers, datas)
	fmt.Println(printer)
}

func printSubs(subs []model.Subscription) {
	headers := []string{"SubscriptionId", "StreamName", "AckTimeoutSeconds", "MaxUnackedRecords"}
	datas := make([][]string, 0, len(subs))
	for i := 0; i < len(subs); i++ {
		data := make([]string, 0, len(headers))
		data = append(data, subs[i].SubscriptionId)
		data = append(data, subs[i].StreamName)
		data = append(data, strconv.Itoa(int(subs[i].AckTimeoutSeconds)))
		data = append(data, strconv.Itoa(int(subs[i].MaxUnackedRecords)))
		datas = append(datas, data)
	}
	printer, _ := util.RenderTable(headers, datas)
	fmt.Println(printer)
}
