package httpClient

import "time"

const (
	// defaultTimeout is the default timeout for http requests
	defaultTimeout = 10 * time.Second
)

type Config struct {
	BaseUrl   string
	KeepAlive bool
}

func DefaultConfig() *Config {
	return &Config{
		BaseUrl:   "http://localhost:8080/v1",
		KeepAlive: true,
	}
}
