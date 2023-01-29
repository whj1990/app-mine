//go:build wireinject
// +build wireinject

package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/google/wire"
	"github.com/whj1990/app-mine/internal/repo"
	"github.com/whj1990/app-mine/internal/service"
)

func initServer() (server.Server, error) {
	panic(wire.Build(service.ProviderSet, repo.ProviderSet, newAppMineImpl, newServer))
}
