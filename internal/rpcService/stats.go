package rpcService

func (c *HStreamClient) GetServerInfo() ([]string, error) {
	return c.client.GetServerInfo()
}
