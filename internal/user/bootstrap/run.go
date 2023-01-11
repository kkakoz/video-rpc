package bootstrap

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/pkg/app"
	"github.com/kkakoz/pkg/app/grpcs"
	"github.com/kkakoz/pkg/logger"
	"github.com/kkakoz/pkg/redisx"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"time"
	"video-rpc/internal/pkg/async/producer"
	"video-rpc/internal/pkg/emailx"
	"video-rpc/internal/user/handler"
	userpb "video-rpc/pb/user"
	"video-rpc/pkg/conf"
	"video-rpc/pkg/etcdx"
	"video-rpc/pkg/grpcx"
)

var cfg = pflag.StringP("config", "c", "./configs/conf.yaml", "Configuration file.")

func InitApp() ([]app.Server, error) {
	time.Local = time.FixedZone("UTC+8", 8*3600)
	pflag.Parse()
	conf.InitConfig(*cfg)
	viper.AutomaticEnv()

	logger.InitLog(viper.GetViper())
	if _, err := ormx.New(viper.GetViper()); err != nil {
		return nil, errors.WithMessage(err, "init orm failed")
	}

	ormx.DefaultErrHandler = func(err error) error {
		return errors.WithStack(err)
	}
	err := redisx.Init(viper.GetViper())
	if err != nil {
		return nil, errors.WithMessage(err, "init redis failed")
	}

	err = emailx.Init(viper.GetViper())
	if err != nil {
		return nil, err
	}

	err = producer.Init(viper.GetViper())
	if err != nil {
		return nil, err
	}

	err = etcdx.Init(viper.GetViper())
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_prometheus.UnaryServerInterceptor,
		grpc_recovery.UnaryServerInterceptor(grpcx.RecoveryInterceptor()),
	),
	))
	userpb.RegisterUserHandlerServer(s, handler.User())
	reflection.Register(s)

	servers := []app.Server{
		grpcs.NewGrpcServer(":"+viper.GetString("app.port"), s),
	}

	return servers, err

}
