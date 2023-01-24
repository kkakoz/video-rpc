package loadbalancing

import (
	"context"
	"github.com/kkakoz/pkg/gox"
	"github.com/kkakoz/pkg/logger"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"time"
)

type ServiceRegister struct {
	cli       *clientv3.Client
	leaseId   clientv3.LeaseID
	retry     chan struct{} // 因为网络问题断开后的重试channel
	retryTime time.Duration // 重试间隔时间
	key       string
	val       string
	cancel    context.CancelFunc

	//keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
}

func (s *ServiceRegister) Start(ctx context.Context) error {
	err := s.putKeyWithLease(ctx) // 第一次服务注册
	if err != nil {
		return err
	}
	s.RegisterAndKeepalive(ctx) // 服务注册重试
	return nil
}

func (s *ServiceRegister) Stop(ctx context.Context) error {
	logger.Info("invoke close")
	s.cancel()
	if _, err := s.cli.Revoke(ctx, s.leaseId); err != nil {
		return err
	}
	return nil
}

const leaseTime int64 = 20

// 注册到etcd
func NewServiceRegister(ctx context.Context, cli *clientv3.Client, serName, addr string) *ServiceRegister {
	ctx, cancel := context.WithCancel(ctx)
	return &ServiceRegister{
		cancel:    cancel,
		cli:       cli,
		retryTime: 20 * time.Second,
		retry:     make(chan struct{}),
		key:       "/" + schema + "/" + serName + "/" + addr,
		val:       addr,
	}
}

func (s *ServiceRegister) putKeyWithLease(ctx context.Context) error {
	logger.Info("register etcd")
	res, err := s.cli.Grant(ctx, leaseTime)
	if err != nil {
		return err
	}
	s.leaseId = res.ID
	_, err = s.cli.Put(ctx, s.key, s.val, clientv3.WithLease(s.leaseId))
	if err != nil {
		return err
	}
	resChan, err := s.cli.KeepAlive(ctx, s.leaseId)
	if err != nil {
		return err
	}
	//s.keepAliveChan = resChan
	gox.Go(func() {
		for range resChan {
		}
		// keepalive关闭
		s.retry <- struct{}{}
	})
	return nil
}

func (s *ServiceRegister) RegisterAndKeepalive(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-s.retry:
			logger.Info("retry register etcd")
			err := s.putKeyWithLease(ctx) // 重新进行服务注册
			if err != nil {
				logger.Error("注册服务失败,err:", zap.Error(err))
				time.Sleep(s.retryTime)
				s.retry <- struct{}{}
			}
			logger.Info("retry register succeeded")
		}
	}

}
