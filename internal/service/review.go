package service

import (
	"context"
	"github.com/whj1990/app-mine/internal/repo"
	"github.com/whj1990/app-mine/internal/service/stru"
	"github.com/whj1990/app-mine/kitex_gen/api"
)

type ReviewService interface {
	ReviewProjectList(ctx context.Context, req *api.ReviewProjectListParams) (resp *api.ReviewProjectListResponse, err error)
	ReviewProjectDetails(ctx context.Context, req *api.IdsInt64Params) (resp *api.ReviewProjectDetailResponse, err error)
	ReviewProjectSave(ctx context.Context, req *api.ReviewProjectSaveParam) (resp *api.SaveResponse, err error)
	ReviewProjectStatus(ctx context.Context, req *api.StatusParam) (resp *api.SaveResponse, err error)
	ReviewProjectDelete(ctx context.Context, req *api.IdParam) (resp *api.SaveResponse, err error)
}

type reviewService struct {
	reviewProjectRepo repo.ReviewProjectRepo
}

func (s *reviewService) ReviewProjectList(ctx context.Context, req *api.ReviewProjectListParams) (*api.ReviewProjectListResponse, error) {
	list, count, err := s.reviewProjectRepo.List(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := &api.ReviewProjectListResponse{
		Data: make([]*api.ReviewProjectData, 0),
	}
	for _, data := range *list {
		resp.Data = append(resp.Data, stru.Convert2ReviewProjectData(&data))
	}
	resp.PageNum = req.PageNum
	resp.PageSize = req.PageSize
	resp.Total = count
	return resp, err
}

func (s *reviewService) ReviewProjectDetails(ctx context.Context, req *api.IdsInt64Params) (resp *api.ReviewProjectDetailResponse, err error) {
	list, err := s.reviewProjectRepo.Detail(ctx, req)
	if err != nil {
		return nil, err
	}
	resp = &api.ReviewProjectDetailResponse{
		Data: make([]*api.ReviewProjectData, 0),
	}
	for _, data := range *list {
		resp.Data = append(resp.Data, stru.Convert2ReviewProjectData(&data))
	}
	return resp, err
}

func (s *reviewService) ReviewProjectSave(ctx context.Context, req *api.ReviewProjectSaveParam) (resp *api.SaveResponse, err error) {
	resp = &api.SaveResponse{}
	if req.Action == "add" {
		resp.RowsAffected, err = s.reviewProjectRepo.Add(ctx, req.Data)
	} else if req.Action == "update" {
		resp.RowsAffected, err = s.reviewProjectRepo.Update(ctx, req.Data)
	}
	return resp, err
}

func (s *reviewService) ReviewProjectStatus(ctx context.Context, req *api.StatusParam) (resp *api.SaveResponse, err error) {
	resp = &api.SaveResponse{}
	resp.RowsAffected, err = s.reviewProjectRepo.UpdateStatus(ctx, req)
	return resp, err
}

func (s *reviewService) ReviewProjectDelete(ctx context.Context, req *api.IdParam) (resp *api.SaveResponse, err error) {
	resp = &api.SaveResponse{}
	resp.RowsAffected, err = s.reviewProjectRepo.Delete(ctx, req)
	return resp, err
}
