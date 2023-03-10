package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/whj1990/app-mine/kitex_gen/api/appmine"
	"github.com/whj1990/go-core/config"
	"github.com/whj1990/go-core/trace"
	"github.com/whj1990/go-core/util"
	"time"
)

func NewMineClient() appmine.Client {
	client, err := appmine.NewClient(
		config.GetString("app.mine.name", ""),
		util.GetClientOption(config.GetString("app.mine.port", "")),
		client.WithRPCTimeout(5*60*time.Second),
		client.WithSuite(trace.NewClientSuite()),
	)
	if err != nil {
		panic(err)
	}
	return client
}
