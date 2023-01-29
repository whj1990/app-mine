// Package service
// @Title provider
// @Description
// @Author lester 2022/10/28
package service

import (
	"github.com/google/wire"
	"github.com/whj1990/app-mine/internal/repo"
)

var ProviderSet = wire.NewSet(NewReviewService)

func NewReviewService(reviewProjectRepo repo.ReviewProjectRepo, ) ReviewService {
	return &reviewService{reviewProjectRepo}
}
