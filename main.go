package main

import (
	"context"
	"draw-together/cmd/api"
	"draw-together/cmd/rpc/rpc_draw"
	. "draw-together/common"
	"flag"
	"fmt"
)

func main() {
	defaultLogLevel := 0 // 默认info

	// 从命令行获取项目的名称，构建配置文件地址，获取配置
	configDir := ""
	flag.StringVar(&configDir, "dir", "", "conf dir")
	flag.Parse()

	configFilePath := fmt.Sprintf("./config/%s.yaml", configDir)
	loglevel := flag.Int("level", defaultLogLevel, "loglevel")
	if err := ReadConfig(configFilePath); err != nil {
		return
	}

	// 初始化日志
	loglevel8 := int8(*loglevel)
	InitSugaredLogger(loglevel8, configDir)

	// 连接数据库
	GormMysql()
	InitRedis()
	defer CloseRedis()

	// 根据命令行启动对应的服务
	// context实现优雅的协程关闭通知
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	switch Config.Server.Type {
	case ServerType_api:
		api.Start(ctx)
	case ServerType_draw:
		err := rpc_draw.Start(ctx, Config.Server.Name, Config.Server.Net.Rpc)
		if err != nil {
			SugaredLogger.Error("draw is not start: ", err)
		}
	default:
		SugaredLogger.Error("server not found")
		return
	}
}
