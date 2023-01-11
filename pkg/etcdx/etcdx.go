package etcdx

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"time"
)

type EtcdOptions struct {
	Endpoints []string `json:"endpoints"`
}

var client *clientv3.Client

func Client() *clientv3.Client {
	return client
}

func Init(viper *viper.Viper) (err error) {
	viper.SetDefault("etcd.endpoints", []string{"127.0.0.1:2379"})
	options := &EtcdOptions{}
	err = viper.UnmarshalKey("etcd", options)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal失败")
	}
	config := clientv3.Config{
		Endpoints:   options.Endpoints,
		DialTimeout: 5 * time.Second,
		DialOptions: []grpc.DialOption{
			grpc.WithBlock(),
		},
	}
	client, err = clientv3.New(config)
	if err != nil {
		return errors.Wrap(err, "etcd连接失败")
	}
	return nil
}
