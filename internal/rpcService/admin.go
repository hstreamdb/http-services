package rpcService

import (
	"encoding/json"
	"fmt"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/pkg/errors"
)

const (
	getStatusCmd = "server status"
	getStatsCmd  = "server stats %s %s %s"
)

func (c *HStreamClient) GetServerInfo() ([]string, error) {
	return c.client.GetServerInfo()
}

func (c *HStreamClient) GetStatus() (*model.TableType, error) {
	return c.sendAdminRequest(getStatusCmd)
}

func (c *HStreamClient) GetStats(addr, category, metrics, intervals string) (*model.TableType, error) {
	return c.sendAdminRequestToServer(addr, fmt.Sprintf(getStatsCmd, category, metrics, intervals))
}

// sendAdminRequest sends an admin command to a random server in the cluster and returns a table format response
func (c *HStreamClient) sendAdminRequest(cmd string) (*model.TableType, error) {
	resp, err := c.client.AdminRequest(cmd)
	if err != nil {
		return nil, err
	}

	return c.parseAdminResponse(resp)
}

// sendAdminRequestToServer sends an admin command to a server in the cluster and returns a table format response
func (c *HStreamClient) sendAdminRequestToServer(addr, cmd string) (*model.TableType, error) {
	resp, err := c.client.AdminRequestToServer(addr, cmd)
	if err != nil {
		return nil, err
	}

	return c.parseAdminResponse(resp)
}

func (c *HStreamClient) parseAdminResponse(resp string) (*model.TableType, error) {
	var jsonObj map[string]json.RawMessage
	if err := json.Unmarshal([]byte(resp), &jsonObj); err != nil {
		return nil, err
	}

	var tableType model.TableType
	if content, ok := jsonObj["content"]; ok {
		if err := json.Unmarshal(content, &tableType); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New(fmt.Sprintf("no content fields in admin response: %+v", jsonObj))
	}

	return &tableType, nil
}
