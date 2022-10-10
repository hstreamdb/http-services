package rpcService

import "github.com/hstreamdb/http-server/api/model"

//func (c *HStreamClient) GetServerInfo() ([]string, error) {
//	return c.client.GetServerInfo()
//}

func (c *HStreamClient) GetStatsFromServer(addr, cmd string) (*model.TableType, error) {
	return c.sendAdminRequestToServer(addr, cmd)
}
