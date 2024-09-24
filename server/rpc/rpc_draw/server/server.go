package server

import (
	"context"
	"draw-together/common"
	"draw-together/model"
	"draw-together/server/rpc/rpc_draw/proto"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
)

type DrawServiceImpl struct {
	proto.UnimplementedGameServiceServer
}

func (d *DrawServiceImpl) GetGameUser(ctx context.Context, req *proto.GameUserRequest) (*proto.GameUserResponse, error) {
	registerData := model.GameUser{}
	userId, err := model.CreateGameUser(&registerData)
	if err != nil {
		common.SugaredLogger.Error("user", fmt.Sprintf("CreateGameUser error: %v", err))
		return nil, err
	}

	userInfo, err := model.GetGameUserById(userId)
	if err != nil {
		common.SugaredLogger.Error("user", fmt.Sprintf("GetUserByUsername error: %v", err))

		return nil, err
	}
	// 创建token
	token := jwt.New(jwt.SigningMethodHS256)
	// 创建一个token的声明
	claims := token.Claims.(jwt.MapClaims)
	// 可以添加自定义的用户ID或其他信息
	claims["user_id"] = strconv.FormatInt(userInfo.Id, 10)
	claims["created_at"] = userInfo.CreatedAt
	t, err := token.SignedString([]byte("tebie6.com"))
	if err != nil {
		common.SugaredLogger.Error("user", fmt.Sprintf("create token error: %v", err))
		return nil, err
	}

	userInfo.AccessToken = t
	err = model.UpdateGameUser(userInfo)
	if err != nil {
		common.SugaredLogger.Error("user", fmt.Sprintf("SaveUserLoginTime error: %v", err))
		return nil, err
	}

	return &proto.GameUserResponse{AccessToken: t}, nil
}
func (d *DrawServiceImpl) GetList(ctx context.Context, request *proto.GetListRequest) (*proto.GetListResponse, error) {

	return &proto.GetListResponse{}, nil
}
func (d *DrawServiceImpl) RepairContent(ctx context.Context, req *proto.RepairContentRequest) (*proto.RepairContentResponse, error) {
	return &proto.RepairContentResponse{}, nil
}

func (d *DrawServiceImpl) GetOnlineUserList(ctx context.Context, req *proto.GetOnlineUserListRequest) (*proto.GetOnlineUserListResponse, error) {
	return &proto.GetOnlineUserListResponse{}, nil
}

func (d *DrawServiceImpl) GetChatList(ctx context.Context, req *proto.GetChatListRequest) (*proto.GetChatListResponse, error) {
	return &proto.GetChatListResponse{}, nil
}

func (d *DrawServiceImpl) AddWord(ctx context.Context, req *proto.AddWordRequest) (*proto.AddWordResponse, error) {
	return &proto.AddWordResponse{}, nil
}
