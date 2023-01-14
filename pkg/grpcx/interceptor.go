package grpcx

import (
	"context"
	"errors"
	"fmt"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/kkakoz/video-rpc/pkg/errs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func RecoveryInterceptor() grpc_recovery.Option {
	return grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
		err, ok := p.(error)
		if ok {
			return errors.New(err.Error())
		}
		return errors.New(fmt.Sprintf("%+v", p))
	})
}

func ServerErrorInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if _, ok := status.FromError(err); ok {
		return resp, err
	}
	return resp, errs.New(500, fmt.Sprintf("%+v", err), "服务器发生错误")
}
