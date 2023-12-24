package viper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type AllConf struct {
	Base struct {
		WebPort string `json:"web_port" yaml:"web_port"`
	} `json:"base" yaml:"base"`
	Logger struct {
		Level string
	}
	Database []struct {
		Connection   string
		MaxIdleConns int
	}
}

func TestNewWithPath(t *testing.T) {
	var c AllConf

	config, err := NewWithPath("./testdata/config.yaml")
	assert.NoError(t, err)
	assert.NotNil(t, config)

	err = config.Parse(&c)
	assert.NoError(t, err)

	base := config.Get("base")
	assert.NotNil(t, base)

	assert.Equal(t, c.Base.WebPort, "8080")
	assert.Equal(t, c.Logger.Level, "debug")
}
