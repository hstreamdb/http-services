package rpcService

import (
	"encoding/json"
	"fmt"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/pkg/errors"
	"strings"
)

const (
	getStatusCmd = "server status"
	getStatsCmd  = "server stats %s %s%s"
)

func (c *HStreamClient) GetStatus() (*model.TableType, error) {
	return c.sendAdminRequest(getStatusCmd)
}

func (c *HStreamClient) GetStats(category string, metrics string, intervals []string) (*model.TableType, error) {
	args := strings.Join(append([]string{""}, intervals...), " -i ")
	return c.sendAdminRequest(fmt.Sprintf(getStatsCmd, category, metrics, args))
}

// sendAdminRequest sends an admin command to the server and returns a table format response
func (c *HStreamClient) sendAdminRequest(cmd string) (*model.TableType, error) {
	resp, err := c.client.AdminRequest(cmd)
	if err != nil {
		return nil, err
	}

	var jsonObj map[string]json.RawMessage
	if err = json.Unmarshal([]byte(resp), &jsonObj); err != nil {
		return nil, err
	}

	var tableType model.TableType
	if content, ok := jsonObj["content"]; ok {
		if err = json.Unmarshal(content, &tableType); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New(fmt.Sprintf("no content fields in admin response: %+v", jsonObj))
	}

	return &tableType, nil
}
