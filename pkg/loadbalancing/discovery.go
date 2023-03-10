package loadbalancing

import (
	"context"
	"github.com/kkakoz/pkg/gox"
	"github.com/kkakoz/pkg/logger"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"sync"
)

const schema = "grpclb"

//ServiceDiscovery 服务发现
type ServiceDiscovery struct {
	ctx       context.Context
	cli       *clientv3.Client //etcd client
	cc        resolver.ClientConn
	servermap map[string]resolver.Address //服务列表
	lock      sync.Mutex
}

func NewServiceDiscovery(ctx context.Context, cli *clientv3.Client) *ServiceDiscovery {
	return &ServiceDiscovery{
		ctx: ctx,
		cli: cli,
	}
}

// 为给定目标创建一个新的`resolver`，当调用`grpc.Dial()`时执行
func (s *ServiceDiscovery) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	s.cc = cc
	s.servermap = make(map[string]resolver.Address)
	prefix := "/" + target.URL.Scheme + target.URL.Path + "/"
	res, err := s.cli.Get(s.ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	for _, v := range res.Kvs {
		logger.Debug("add address: " + string(v.Value))
		s.SetServerAddress(string(v.Key), string(v.Value))
	}
	s.cc.NewAddress(s.getServices())
	//监视前缀，修改变更的server
	gox.Go(func() {
		s.watcher(prefix)
	})
	return s, nil
}

func (s *ServiceDiscovery) Scheme() string {
	return schema
}

func (s *ServiceDiscovery) SetServerAddress(key, value string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.servermap[key] = resolver.Address{Addr: value}
	s.cc.NewAddress(s.getServices())
	logger.Debug("register key:" + key + " value:" + value)
}

func (s *ServiceDiscovery) DelServerAddress(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.servermap, key)
	s.cc.NewAddress(s.getServices())
	logger.Debug("del service key:" + key)
}

func (s *ServiceDiscovery) watcher(prefix string) {
	watchChan := s.cli.Watch(s.ctx, prefix, clientv3.WithPrefix())
	for res := range watchChan {
		for _, event := range res.Events {
			switch event.Type {
			case clientv3.EventTypePut:
				s.SetServerAddress(string(event.Kv.Key), string(event.Kv.Value))
			case clientv3.EventTypeDelete:
				s.DelServerAddress(string(event.Kv.Key))
			}
		}
	}
}

func (s *ServiceDiscovery) ResolveNow(options resolver.ResolveNowOptions) {
}

func (s *ServiceDiscovery) Close() {
	logger.Info("service discover closed")
}

//GetServices 获取服务地址
func (s *ServiceDiscovery) getServices() []resolver.Address {
	addrs := make([]resolver.Address, 0, len(s.servermap))

	for _, v := range s.servermap {
		addrs = append(addrs, v)
	}
	return addrs
}
