package main

import (
	"context"
	"github.com/whj1990/app-mine/internal/service"
	api "github.com/whj1990/app-mine/kitex_gen/api"
)

// AppMineImpl implements the last service interface defined in the IDL.
type AppMineImpl struct {
	reviewService service.ReviewService
}

// ReviewProjectList implements the AppMineImpl interface.
func (s *AppMineImpl) ReviewProjectList(ctx context.Context, req *api.ReviewProjectListParams) (resp *api.ReviewProjectListResponse, err error) {
	return s.reviewService.ReviewProjectList(ctx, req)
}

// ReviewProjectDetails implements the AppMineImpl interface.
func (s *AppMineImpl) ReviewProjectDetails(ctx context.Context, req *api.IdsInt64Params) (resp *api.ReviewProjectDetailResponse, err error) {
	return s.reviewService.ReviewProjectDetails(ctx, req)
}

// ReviewProjectSave implements the AppMineImpl interface.
func (s *AppMineImpl) ReviewProjectSave(ctx context.Context, req *api.ReviewProjectSaveParam) (resp *api.SaveResponse, err error) {
	return s.reviewService.ReviewProjectSave(ctx, req)
}

// ReviewProjectStatus implements the AppMineImpl interface.
func (s *AppMineImpl) ReviewProjectStatus(ctx context.Context, req *api.StatusParam) (resp *api.SaveResponse, err error) {
	return s.reviewService.ReviewProjectStatus(ctx, req)
}

// ReviewProjectDelete implements the AppMineImpl interface.
func (s *AppMineImpl) ReviewProjectDelete(ctx context.Context, req *api.IdParam) (resp *api.SaveResponse, err error) {
	return s.reviewService.ReviewProjectDelete(ctx, req)
}
