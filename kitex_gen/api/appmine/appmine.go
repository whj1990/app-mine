// Code generated by Kitex v0.4.3. DO NOT EDIT.

package appmine

import (
	api "app-mine/kitex_gen/api"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return appMineServiceInfo
}

var appMineServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "AppMine"
	handlerType := (*api.AppMine)(nil)
	methods := map[string]kitex.MethodInfo{
		"ReviewProjectList":    kitex.NewMethodInfo(reviewProjectListHandler, newAppMineReviewProjectListArgs, newAppMineReviewProjectListResult, false),
		"ReviewProjectDetails": kitex.NewMethodInfo(reviewProjectDetailsHandler, newAppMineReviewProjectDetailsArgs, newAppMineReviewProjectDetailsResult, false),
		"ReviewProjectSave":    kitex.NewMethodInfo(reviewProjectSaveHandler, newAppMineReviewProjectSaveArgs, newAppMineReviewProjectSaveResult, false),
		"ReviewProjectStatus":  kitex.NewMethodInfo(reviewProjectStatusHandler, newAppMineReviewProjectStatusArgs, newAppMineReviewProjectStatusResult, false),
		"ReviewProjectDelete":  kitex.NewMethodInfo(reviewProjectDeleteHandler, newAppMineReviewProjectDeleteArgs, newAppMineReviewProjectDeleteResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.3",
		Extra:           extra,
	}
	return svcInfo
}

func reviewProjectListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.AppMineReviewProjectListArgs)
	realResult := result.(*api.AppMineReviewProjectListResult)
	success, err := handler.(api.AppMine).ReviewProjectList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAppMineReviewProjectListArgs() interface{} {
	return api.NewAppMineReviewProjectListArgs()
}

func newAppMineReviewProjectListResult() interface{} {
	return api.NewAppMineReviewProjectListResult()
}

func reviewProjectDetailsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.AppMineReviewProjectDetailsArgs)
	realResult := result.(*api.AppMineReviewProjectDetailsResult)
	success, err := handler.(api.AppMine).ReviewProjectDetails(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAppMineReviewProjectDetailsArgs() interface{} {
	return api.NewAppMineReviewProjectDetailsArgs()
}

func newAppMineReviewProjectDetailsResult() interface{} {
	return api.NewAppMineReviewProjectDetailsResult()
}

func reviewProjectSaveHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.AppMineReviewProjectSaveArgs)
	realResult := result.(*api.AppMineReviewProjectSaveResult)
	success, err := handler.(api.AppMine).ReviewProjectSave(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAppMineReviewProjectSaveArgs() interface{} {
	return api.NewAppMineReviewProjectSaveArgs()
}

func newAppMineReviewProjectSaveResult() interface{} {
	return api.NewAppMineReviewProjectSaveResult()
}

func reviewProjectStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.AppMineReviewProjectStatusArgs)
	realResult := result.(*api.AppMineReviewProjectStatusResult)
	success, err := handler.(api.AppMine).ReviewProjectStatus(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAppMineReviewProjectStatusArgs() interface{} {
	return api.NewAppMineReviewProjectStatusArgs()
}

func newAppMineReviewProjectStatusResult() interface{} {
	return api.NewAppMineReviewProjectStatusResult()
}

func reviewProjectDeleteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.AppMineReviewProjectDeleteArgs)
	realResult := result.(*api.AppMineReviewProjectDeleteResult)
	success, err := handler.(api.AppMine).ReviewProjectDelete(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAppMineReviewProjectDeleteArgs() interface{} {
	return api.NewAppMineReviewProjectDeleteArgs()
}

func newAppMineReviewProjectDeleteResult() interface{} {
	return api.NewAppMineReviewProjectDeleteResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ReviewProjectList(ctx context.Context, req *api.ReviewProjectListParams) (r *api.ReviewProjectListResponse, err error) {
	var _args api.AppMineReviewProjectListArgs
	_args.Req = req
	var _result api.AppMineReviewProjectListResult
	if err = p.c.Call(ctx, "ReviewProjectList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ReviewProjectDetails(ctx context.Context, req *api.IdsInt64Params) (r *api.ReviewProjectDetailResponse, err error) {
	var _args api.AppMineReviewProjectDetailsArgs
	_args.Req = req
	var _result api.AppMineReviewProjectDetailsResult
	if err = p.c.Call(ctx, "ReviewProjectDetails", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ReviewProjectSave(ctx context.Context, req *api.ReviewProjectSaveParam) (r *api.SaveResponse, err error) {
	var _args api.AppMineReviewProjectSaveArgs
	_args.Req = req
	var _result api.AppMineReviewProjectSaveResult
	if err = p.c.Call(ctx, "ReviewProjectSave", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ReviewProjectStatus(ctx context.Context, req *api.StatusParam) (r *api.SaveResponse, err error) {
	var _args api.AppMineReviewProjectStatusArgs
	_args.Req = req
	var _result api.AppMineReviewProjectStatusResult
	if err = p.c.Call(ctx, "ReviewProjectStatus", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ReviewProjectDelete(ctx context.Context, req *api.IdParam) (r *api.SaveResponse, err error) {
	var _args api.AppMineReviewProjectDeleteArgs
	_args.Req = req
	var _result api.AppMineReviewProjectDeleteResult
	if err = p.c.Call(ctx, "ReviewProjectDelete", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
