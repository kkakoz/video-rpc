package client

import (
	"context"
	userpb "github.com/kkakoz/video-rpc/pb/user"
	"github.com/kkakoz/video-rpc/pkg/loadbalancing"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func NewUserClient(ctx context.Context, etcdClient *clientv3.Client) (userpb.UserHandlerClient, error) {
	r := loadbalancing.NewServiceDiscovery(ctx, etcdClient)
	resolver.Register(r)
	const grpcServiceConfig = `{"loadBalancingPolicy":"round_robin"}`

	conn, err := grpc.Dial(
		r.Scheme()+":///"+userpb.AppName,
		grpc.WithDefaultServiceConfig(grpcServiceConfig),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	client := userpb.NewUserHandlerClient(conn)
	return client, nil
}
