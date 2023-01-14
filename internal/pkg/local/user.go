package local

import (
	"context"
	"github.com/kkakoz/pkg/jsonx"
	userpb "github.com/kkakoz/video-rpc/pb/user"
	"google.golang.org/grpc/metadata"
)

func GetUser(ctx context.Context) *userpb.UserInfoRes {
	md, b := metadata.FromIncomingContext(ctx)
	if !b {
		return nil
	}
	strings := md.Get(userpb.UserLocalKey)
	if len(strings) == 0 {
		return nil
	}
	userBytes := strings[0]
	res, err := jsonx.Unmarshal[userpb.UserInfoRes]([]byte(userBytes))
	if err != nil {
		return nil
	}
	return res

}
