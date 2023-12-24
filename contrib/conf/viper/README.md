# file config
> parsing the configuration file by [viper](https://github.com/spf13/viper)

## Usage
```go
package main

import (
	"fmt"

	"github.com/xiecang/jarvis/contrib/conf/viper"
)

type Server struct {
    HTTP string `json:"http" mapstructure:"http" yaml:"http"`
}

type AllConfig struct {
	Server *Server  `json:"server" mapstructure:"server" yaml:"server"`
} 

func main() {
	c := &AllConfig{}
	config, err := viper.NewWithPath("/config/path")
	if err != nil {
		panic(err)
	}
	if err = config.Parse(&c); err != nil {
		panic(err)
	}
	fmt.Println(c.Server.HTTP)
}
```
