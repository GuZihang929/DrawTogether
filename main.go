package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	// 从命令行获取项目的名称，构建配置文件地址，获取配置
	configDir := ""
	flag.StringVar(&configDir, "dir", "", "conf dir")
	flag.Parse()

	v := viper.New()
	configFilePath := fmt.Sprintf("./config/%s", configDir)
	v.SetConfigFile(configFilePath)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 初始化日志

	// 根据命令行启动对应的服务

}
