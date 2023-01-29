package main

import (
	api "app-mine/kitex_gen/api"
	"context"
)

// AppMineImpl implements the last service interface defined in the IDL.
type AppMineImpl struct{}

// ReviewProjectList implements the AppMineImpl interface.
func (s *AppMineImpl) ReviewProjectList(ctx context.Context, req *api.ReviewProjectListParams) (resp *api.ReviewProjectListResponse, err error) {
	// TODO: Your code here...
	return
}

// ReviewProjectDetails implements the AppMineImpl interface.
func (s *AppMineImpl) ReviewProjectDetails(ctx context.Context, req *api.IdsInt64Params) (resp *api.ReviewProjectDetailResponse, err error) {
	// TODO: Your code here...
	return
}

// ReviewProjectSave implements the AppMineImpl interface.
func (s *AppMineImpl) ReviewProjectSave(ctx context.Context, req *api.ReviewProjectSaveParam) (resp *api.SaveResponse, err error) {
	// TODO: Your code here...
	return
}

// ReviewProjectStatus implements the AppMineImpl interface.
func (s *AppMineImpl) ReviewProjectStatus(ctx context.Context, req *api.StatusParam) (resp *api.SaveResponse, err error) {
	// TODO: Your code here...
	return
}

// ReviewProjectDelete implements the AppMineImpl interface.
func (s *AppMineImpl) ReviewProjectDelete(ctx context.Context, req *api.IdParam) (resp *api.SaveResponse, err error) {
	// TODO: Your code here...
	return
}
