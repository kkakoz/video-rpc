package handler

import (
	"context"
	"github.com/kkakoz/video-rpc/internal/user/logic"
	userpb "github.com/kkakoz/video-rpc/pb/user"
	"google.golang.org/protobuf/types/known/emptypb"
)

type user struct {
	userpb.UnsafeUserHandlerServer
}

func (u user) Login(ctx context.Context, req *userpb.LoginReq) (*userpb.LoginRes, error) {
	return logic.User().Login(ctx, req)
}

func (u user) Register(ctx context.Context, req *userpb.RegisterReq) (*emptypb.Empty, error) {
	return logic.User().Register(ctx, req)
}

func (u user) UserInfo(ctx context.Context, id *userpb.ID) (*userpb.UserInfoRes, error) {
	return logic.User().UserInfo(ctx, id)
}

func User() *user {
	return &user{}
}
