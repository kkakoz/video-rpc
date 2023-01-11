package main

import (
	"context"
	"github.com/kkakoz/pkg/app"
	"github.com/kkakoz/video-rpc/internal/user/bootstrap"
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
	appName := "user-server"

	_, err = loadbalancing.NewServiceRegister(context.TODO(), etcdx.Client(), appName, viper.GetString("ip")+":"+viper.GetString("app.port"))

	app := app.NewApp(appName, servers...)

	if err = app.Start(context.TODO()); err != nil {
		log.Fatal(err)
	}

}
