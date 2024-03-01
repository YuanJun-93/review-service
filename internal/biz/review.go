package biz

import (
	"context"
	"fmt"
	v1 "review-service/api/review/v1"
	"review-service/internal/data/model"
	"review-service/pkg/snowflake"

	"github.com/go-kratos/kratos/v2/log"
)

type ReviewRepo interface {
	GetReviewByOrderID(ctx context.Context, orderID int64) ([]*model.ReviewInfo, error)
	SaveReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error)
	GetReview(ctx context.Context, reviewID int64) (*model.ReviewInfo, error)
}

type ReviewUsecase struct {
	repo ReviewRepo
	log  *log.Helper
}

func NewReviewUsecase(repo ReviewRepo, logger log.Logger) *ReviewUsecase {
	return &ReviewUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *ReviewUsecase) CreateReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	uc.log.WithContext(ctx).Debugf("[biz] CreateReview, req:%v", review)
	// 1. 数据校验，已经评价过的订单不能再评论
	reviews, err := uc.repo.GetReviewByOrderID(ctx, review.OrderID)
	if err != nil {
		return nil, v1.ErrorDbFailed("查询数据库失败")
	}
	if len(reviews) > 0 {
		fmt.Printf("订单已评价, len(reviews): %d\n", len(reviews))
		return nil, v1.ErrorOrderReviewed("订单: %d, 已评价", review.OrderID)
	}
	// 2. 生成review id
	review.ReviewID = snowflake.GenID()
	// 3. 查询订单和商品快照信息
	// 4. 组装数据入库
	return uc.repo.SaveReview(ctx, review)
}

func (uc *ReviewUsecase) GetReview(ctx context.Context, reviewID int64) (*model.ReviewInfo, error) {
	uc.log.WithContext(ctx).Debugf("[biz] GetReview, reviewID: %v", reviewID)
	return uc.repo.GetReview(ctx, reviewID)
}
