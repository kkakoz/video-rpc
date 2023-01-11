package client

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	userpb "video-rpc/pb/user"
	"video-rpc/pkg/loadbalancing"
)

func NewUserClient(ctx context.Context, etcdClient *clientv3.Client) (userpb.UserHandlerClient, error) {
	r := loadbalancing.NewServiceDiscovery(ctx, etcdClient)
	resolver.Register(r)
	const grpcServiceConfig = `{"loadBalancingPolicy":"round_robin"}`

	conn, err := grpc.Dial(
		r.Scheme()+":///"+loadbalancing.UserServName,
		grpc.WithDefaultServiceConfig(grpcServiceConfig),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	client := userpb.NewUserHandlerClient(conn)
	return client, nil
}
