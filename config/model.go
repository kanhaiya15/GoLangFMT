package config

import "github.com/LambdaTest/mould/pkg/lumber"

// Model definition for configuration

// Config the application's configuration
type Config struct {
	Config          string
	DBConf          DBConfig
	Port            string
	SomeWeirdConfig string `json:"some-weird-config" yaml:"SomeWeirdConfig"`
	SomeAddress     string `json:"some-address" yaml:"ProxyPass"`
	LogFile         string
	LogConfig       lumber.LoggingConfig
	Env             string
	Verbose         bool
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}
