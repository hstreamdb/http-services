package config

import "github.com/hstreamdb/http-server/pkg/util"

type Config struct {
	ServerUrl string
	LogLevel  util.LogLevel
}

func NewConfig(serverUrl string, logLevel string) *Config {
	var level util.LogLevel
	switch logLevel {
	case "debug":
		level = util.DEBUG
	case "info":
		level = util.INFO
	case "warn":
		level = util.WARNING
	case "error":
		level = util.ERROR
	case "fatal":
		level = util.FATAL
	case "panic":
		level = util.PANIC
	}

	return &Config{
		ServerUrl: serverUrl,
		LogLevel:  level,
	}
}
