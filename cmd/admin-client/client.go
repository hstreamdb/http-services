package main

import (
	"github.com/hstreamdb/http-server/pkg/httpClient"
	"strings"
)

func newClient() (*httpClient.Client, error) {
	config := httpClient.DefaultConfig()
	prefix := strings.Join([]string{globalFlags.Address, globalFlags.ApiVersion}, "/")
	config.BaseUrl = prefix
	return httpClient.NewHTTPClient(config)
}
