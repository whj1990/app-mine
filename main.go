package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/whj1990/app-mine/internal/service"
	"github.com/whj1990/app-mine/kitex_gen/api"
	"github.com/whj1990/app-mine/kitex_gen/api/appmine"
	"github.com/whj1990/go-core/launch"
)

func main() {
	logger, closer := launch.InitPremise()
	defer logger.Sync()
	defer closer.Close()
	server, err := initServer()
	if err != nil {
		panic(err)
	}
	go launch.RunServer(server)
	launch.InitHttpServer()
}
func newAppMineImpl(reviewService service.ReviewService) api.AppMine {
	return &AppMineImpl{reviewService}
}

func newServer(handler api.AppMine) server.Server {
	return appmine.NewServer(handler, launch.RpcServerOptions()...)
}
