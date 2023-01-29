package stru

import (
	"github.com/whj1990/app-mine/internal/repo"
	"github.com/whj1990/app-mine/kitex_gen/api"
)

func Convert2ReviewProjectData(data *repo.ReviewProjectData) *api.ReviewProjectData {
	return &api.ReviewProjectData{
		Id:          data.Id,
		Name:        data.Name,
		ModeCode:    data.ModeCode,
		Description: data.Description,
		Status:      data.Status,
		Deleted:     data.Deleted,
		CreatedUser: data.CreatedUser,
		UpdatedUser: data.UpdatedUser,
		CreatedDate: data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedDate: data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
