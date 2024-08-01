package server

import (
	"context"
	"draw-together/server/rpc/rpc_draw"
	"fmt"
)

//type DrawService interface {
//	GetGameUser(ctx context.Context, req *rpc_draw.GameUserRequest) (*rpc_draw.GameUserResponse, error)
//	GetList(ctx context.Context, req *rpc_draw.GetListRequest) (*rpc_draw.GetListResponse, error)
//	RepairContent(ctx context.Context, req *rpc_draw.RepairContentRequest) (*rpc_draw.RepairContentResponse, error)
//	GetOnlineUserList(ctx context.Context, req *rpc_draw.GetOnlineUserListRequest) (*rpc_draw.GetOnlineUserListResponse, error)
//	GetChatList(ctx context.Context, req *rpc_draw.GetChatListRequest) (*rpc_draw.GetChatListResponse, error)
//	AddWord(ctx context.Context, req *rpc_draw.AddWordRequest) (*rpc_draw.AddWordResponse, error)
//}

type DrawServiceImpl struct {
	rpc_draw.UnimplementedGameServiceServer
}

func (d *DrawServiceImpl) GetGameUser(ctx context.Context, req *rpc_draw.GameUserRequest) (*rpc_draw.GameUserResponse, error) {
	fmt.Println("1234")
	return &rpc_draw.GameUserResponse{}, nil
}
func (d *DrawServiceImpl) GetList(ctx context.Context, request *rpc_draw.GetListRequest) (*rpc_draw.GetListResponse, error) {
	return &rpc_draw.GetListResponse{}, nil
}
func (d *DrawServiceImpl) RepairContent(ctx context.Context, req *rpc_draw.RepairContentRequest) (*rpc_draw.RepairContentResponse, error) {
	return &rpc_draw.RepairContentResponse{}, nil
}

func (d *DrawServiceImpl) GetOnlineUserList(ctx context.Context, req *rpc_draw.GetOnlineUserListRequest) (*rpc_draw.GetOnlineUserListResponse, error) {
	return &rpc_draw.GetOnlineUserListResponse{}, nil
}

func (d *DrawServiceImpl) GetChatList(ctx context.Context, req *rpc_draw.GetChatListRequest) (*rpc_draw.GetChatListResponse, error) {
	return &rpc_draw.GetChatListResponse{}, nil
}

func (d *DrawServiceImpl) AddWord(ctx context.Context, req *rpc_draw.AddWordRequest) (*rpc_draw.AddWordResponse, error) {
	return &rpc_draw.AddWordResponse{}, nil
}
