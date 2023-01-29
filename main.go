package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/whj1990/app-mine/kitex_gen/api"
	"github.com/whj1990/app-mine/kitex_gen/api/appmine"
	"github.com/whj1990/go-core/launch"
)

func main() {
	logger, closer := launch.InitPremise()
	defer logger.Sync()
	defer closer.Close()

	launch.InitHttpServer()
}
func newAppMineImpl() api.AppMine {
	return &AppMineImpl{}
}

func newServer(handler api.AppMine) server.Server {
	return appmine.NewServer(handler, launch.RpcServerOptions()...)
}
