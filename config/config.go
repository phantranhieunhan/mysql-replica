package config

type config struct {
	WriteURL string `mapstructure:"WRITE_URL"`
	ReadURL  string `mapstructure:"READ_URL"`
	Database string `mapstructure:"DATABASE"`
}

var C config
