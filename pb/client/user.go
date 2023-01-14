package client

import (
	"context"
	userpb "github.com/kkakoz/video-rpc/pb/user"
	"github.com/kkakoz/video-rpc/pkg/loadbalancing"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func NewUserClient(ctx context.Context, etcdClient *clientv3.Client, interceptors ...grpc.UnaryClientInterceptor) (userpb.UserHandlerClient, error) {
	r := loadbalancing.NewServiceDiscovery(ctx, etcdClient)
	resolver.Register(r)
	const grpcServiceConfig = `{"loadBalancingPolicy":"round_robin"}`

	opts := []grpc.DialOption{grpc.WithDefaultServiceConfig(grpcServiceConfig), grpc.WithInsecure()}

	for _, interceptor := range interceptors {
		opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	}

	conn, err := grpc.Dial(
		r.Scheme()+":///"+userpb.AppName,
		opts...,
	)
	if err != nil {
		return nil, err
	}
	client := userpb.NewUserHandlerClient(conn)
	return client, nil
}
