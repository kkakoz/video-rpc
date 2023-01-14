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
	"github.com/kkakoz/video-rpc/internal/pkg/async/producer"
	"github.com/kkakoz/video-rpc/internal/pkg/emailx"
	"github.com/kkakoz/video-rpc/internal/user/handler"
	userpb "github.com/kkakoz/video-rpc/pb/user"
	"github.com/kkakoz/video-rpc/pkg/conf"
	"github.com/kkakoz/video-rpc/pkg/etcdx"
	"github.com/kkakoz/video-rpc/pkg/grpcx"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"time"
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
		grpcx.ServerErrorInterceptor,
	),
	))
	userpb.RegisterUserHandlerServer(s, handler.User())
	reflection.Register(s)

	servers := []app.Server{
		grpcs.NewGrpcServer(":"+viper.GetString("app.port"), s),
	}

	return servers, err

}
