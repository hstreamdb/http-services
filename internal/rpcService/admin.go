package rpcService

import (
	"encoding/json"
	"fmt"
	"github.com/hstreamdb/http-server/api/model"
	"strings"
)

const (
	getStatusCmd = "server status"
	getStatsCmd  = "server stats %s%s"
)

func (c *HStreamClient) GetStatus() (model.TableType, error) {
	return c.sendAdminRequest(getStatusCmd)
}

func (c *HStreamClient) GetStats(method string, intervals []string) (model.TableType, error) {
	args := strings.Join(append([]string{""}, intervals...), " -i ")
	return c.sendAdminRequest(fmt.Sprintf(getStatsCmd, method, args))
}

// sendAdminRequest sends an admin command to the server and returns a table format response
func (c *HStreamClient) sendAdminRequest(cmd string) (model.TableType, error) {
	resp, err := c.client.AdminRequest(cmd)
	if err != nil {
		return model.TableType{}, err
	}

	var tableRes model.TableType
	if err = json.Unmarshal([]byte(resp), &tableRes); err != nil {
		return model.TableType{}, err
	}
	return tableRes, nil
}
