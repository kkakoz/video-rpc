package main

import (
	"context"
	"github.com/kkakoz/pkg/app"
	"github.com/kkakoz/video-rpc/internal/user/bootstrap"
	userpb "github.com/kkakoz/video-rpc/pb/user"
	"log"
)

func main() {
	servers, err := bootstrap.InitApp()
	if err != nil {
		log.Fatal(err)
	}

	app := app.NewApp(userpb.AppName, servers...)

	if err = app.Start(context.TODO()); err != nil {
		log.Fatal(err)
	}

}
