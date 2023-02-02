package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/whj1990/app-mine/kitex_gen/api/appmine"
	"github.com/whj1990/go-core/config"
	"github.com/whj1990/go-core/trace"
	"github.com/whj1990/go-core/util"
	"time"
)

func NewRecommendClient() appmine.Client {
	client, err := appmine.NewClient(
		config.GetString("app.recommend.name", ""),
		util.GetClientOption(config.GetString("app.recommend.port", "")),
		client.WithRPCTimeout(5*60*time.Second),
		client.WithSuite(trace.NewClientSuite()),
	)
	if err != nil {
		panic(err)
	}
	return client
}
