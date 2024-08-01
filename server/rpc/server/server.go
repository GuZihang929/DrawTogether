package server

import (
	"context"
	"draw-together/server/rpc"
)

type DrawService interface {
	GetGameUser(ctx context.Context, req *rpc.GameUserRequest) (*rpc.GameUserResponse, error)
	GetList(ctx context.Context, req *rpc.GetListRequest) (*rpc.GetListResponse, error)
	RepairContent(ctx context.Context, req *rpc.RepairContentRequest) (*rpc.RepairContentResponse, error)
	GetOnlineUserList(ctx context.Context, req *rpc.GetOnlineUserListRequest) (*rpc.GetOnlineUserListResponse, error)
	GetChatList(ctx context.Context, req *rpc.GetChatListRequest) (*rpc.GetChatListResponse, error)
	AddWord(ctx context.Context, req *rpc.AddWordRequest) (*rpc.AddWordResponse, error)
}
type DrawServiceImpl struct {
	DrawService
}

func (d *DrawServiceImpl) GetGameUser(ctx context.Context, req *rpc.GameUserRequest) (*rpc.GameUserResponse, error) {
	return &rpc.GameUserResponse{}, nil
}
func (d *DrawServiceImpl) GetList(ctx context.Context, request *rpc.GetListRequest) (*rpc.GetListResponse, error) {
	return &rpc.GetListResponse{}, nil
}
func (d *DrawServiceImpl) RepairContent(ctx context.Context, req *rpc.RepairContentRequest) (*rpc.RepairContentResponse, error) {
	return &rpc.RepairContentResponse{}, nil
}

func (d *DrawServiceImpl) GetOnlineUserList(ctx context.Context, req *rpc.GetOnlineUserListRequest) (*rpc.GetOnlineUserListResponse, error) {
	return &rpc.GetOnlineUserListResponse{}, nil
}

func (d *DrawServiceImpl) GetChatList(ctx context.Context, req *rpc.GetChatListRequest) (*rpc.GetChatListResponse, error) {
	return &rpc.GetChatListResponse{}, nil
}

func (d *DrawServiceImpl) AddWord(ctx context.Context, req *rpc.AddWordRequest) (*rpc.AddWordResponse, error) {
	return &rpc.AddWordResponse{}, nil
}
