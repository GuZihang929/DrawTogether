package rpc_draw

import (
	"context"
	. "draw-together/common"
	"draw-together/server/rpc/rpc_draw/proto"
	"draw-together/server/rpc/rpc_draw/server"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Start(c context.Context, rpcRegisterName, rpcTcpAddr string) error {

	// 记录 RPC 服务器的初始化信息
	SugaredLogger.Info("RPC server is initializing ", "rpcRegisterName: ", rpcRegisterName)
	// 创建 TCP 监听器

	listener, err := net.Listen(
		"tcp",
		rpcTcpAddr,
	)
	if err != nil {
		return errors.New("listen err rpcTcpAddr:" + rpcTcpAddr)
	}
	defer listener.Close()

	// 服务注册
	Register(ServerStatue{
		IP:   Config.Server.Net.Rpc,
		Name: Config.Server.Name,
		Type: Config.Server.Type,
	})

	// 初始化服务发现客户端
	s := grpc.NewServer()
	proto.RegisterGameServiceServer(s, &server.DrawServiceImpl{})
	// 启动服务器
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	return nil
}
