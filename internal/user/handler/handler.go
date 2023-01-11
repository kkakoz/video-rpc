package handler

import (
	"context"
	"video-rpc/internal/user/logic"
	"video-rpc/pb/common"
	userpb "video-rpc/pb/user"
)

type user struct {
	userpb.UnsafeUserHandlerServer
}

func User() *user {
	return &user{}
}

func (u user) Login(ctx context.Context, req *userpb.LoginReq) (*userpb.LoginRes, error) {
	return logic.User().Login(ctx, req)
}

func (u user) Register(ctx context.Context, req *userpb.RegisterReq) (*common.Res, error) {
	return logic.User().Register(ctx, req)
}

func (u user) UserInfo(ctx context.Context, id *common.ID) (*userpb.UserInfoRes, error) {
	return logic.User().UserInfo(ctx, id)
}
