package goslog

import (
	_ "github.com/jessevdk/go-flags"
)

type LogConfig struct {
	Verbose bool      `short:"v" long:"verbose" env:"VERBOSE" description:"Short hand for log level debug. Will override any setting for level"`
	Level   LogLevel  `short:"l" long:"level" env:"LEVEL" default:"warn" choice:"debug" choice:"info" choice:"warn" choice:"error" description:"Log level"`
	Format  LogFormat `short:"f" long:"format" env:"FORMAT" default:"text" choice:"text" choice:"json" choice:"pretty" description:"Log format"`
}
