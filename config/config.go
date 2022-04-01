package config

import "github.com/hstreamdb/http-server/pkg/util"

type Config struct {
	ServerUrl string
	LogLevel  util.LogLevel
}

func DefaultConfig() *Config {
	return &Config{
		ServerUrl: "localhost:6580",
		LogLevel:  util.DEBUG,
	}
}
