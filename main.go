package main

import (
	api "app-mine/kitex_gen/api/appmine"
	"github.com/whj1990/app-mine/kitex_gen/api/appmine"
	"github.com/whj1990/go-core/launch"
	"log"
)

func main() {
	logger, closer := launch.InitPremise()
	defer logger.Sync()
	defer closer.Close()
	svr := api.NewServer(new(AppMineImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
	launch.InitHttpServer()
}
func newRecommendImpl(service service.RecommendService, tagService service.TagService, questionService service.QuestionService, streamService service.StreamService, reviewService service.ReviewService, aliService service.AliService) api.Recommend {
	return &RecommendImpl{service, tagService, questionService, streamService, reviewService, aliService}
}

func newServer(handler api.Recommend) server.Server {
	return appmine.NewServer(handler, launch.RpcServerOptions()...)
}
