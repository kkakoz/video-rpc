package main

import (
	"context"
	"github.com/kkakoz/pkg/app"
	"github.com/kkakoz/video-rpc/internal/user/bootstrap"
	userpb "github.com/kkakoz/video-rpc/pb/user"
	"github.com/kkakoz/video-rpc/pkg/etcdx"
	"github.com/kkakoz/video-rpc/pkg/loadbalancing"
	"github.com/spf13/viper"
	"log"
)

func main() {
	servers, err := bootstrap.InitApp()
	if err != nil {
		log.Fatal(err)
	}

	_, err = loadbalancing.NewServiceRegister(context.TODO(), etcdx.Client(), userpb.AppName, viper.GetString("KUBERNETES_SERVICE_HOST")+":"+viper.GetString("app.port"))

	app := app.NewApp(userpb.AppName, servers...)

	if err = app.Start(context.TODO()); err != nil {
		log.Fatal(err)
	}

}
