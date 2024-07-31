package common

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"strings"
)

// 配置
type config struct {
	Server *configServer
}

var Config = &config{
	Server: &configServer{},
}

func ReadConfig(p string) error {
	v := viper.New()

	v.SetConfigFile(p)

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = v.Unmarshal(Config.Server)
	if err != nil {
		panic(fmt.Errorf("read config file to struct err: %s \n", err))
	}

	name := path.Base(p)
	Config.Server.Type = strings.Split(name, "_")[0]
	Config.Server.Id = strings.Split(name, "_")[1]
	Config.Server.Name = name
	return err
}
