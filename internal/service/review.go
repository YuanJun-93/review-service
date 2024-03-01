package service

import (
	"context"
	"fmt"

	pb "review-service/api/review/v1"
	"review-service/internal/biz"
	"review-service/internal/data/model"
)

type ReviewService struct {
	pb.UnimplementedReviewServer
	uc *biz.ReviewUsecase
}

func NewReviewService(uc *biz.ReviewUsecase) *ReviewService {
	return &ReviewService{uc: uc}
}

func (s *ReviewService) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.CreateReviewReply, error) {
	fmt.Printf("[service] CreateReview, req: %#v\n", req)
	// 调用biz层
	var anonymous int32
	if req.Anonymous {
		anonymous = 1
	}

	reviewInfo, err := s.uc.CreateReview(ctx, &model.ReviewInfo{
		UserID:       req.UserID,
		OrderID:      req.OrderID,
		Score:        req.Score,
		ServiceScore: req.Score,
		ExpressScore: req.ExpressScore,
		Content:      req.Content,
		PicInfo:      req.PicInfo,
		VideoInfo:    req.VideoInfo,
		Anonymous:    anonymous,
		Status:       0,
	})
	if err != nil {
		return nil, err
	}

	// 拼装返回结果
	return &pb.CreateReviewReply{ReviewID: reviewInfo.ReviewID}, nil
}
func (s *ReviewService) GetReview(ctx context.Context, req *pb.GetReviewRequest) (*pb.GetReviewReply, error) {
	fmt.Printf("[service] GetReview, req: %#v\n", req)
	reviewInfo, err := s.uc.GetReview(ctx, req.ReviewID)
	if err != nil {
		return nil, err
	}
	return &pb.GetReviewReply{Data: &pb.ReviewInfo{
		ReviewID:     reviewInfo.ReviewID,
		UserID:       reviewInfo.UserID,
		OrderID:      reviewInfo.OrderID,
		Score:        reviewInfo.Score,
		ServiceScore: reviewInfo.ServiceScore,
		ExpressScore: reviewInfo.ExpressScore,
		Content:      reviewInfo.Content,
		PicInfo:      reviewInfo.PicInfo,
		VideoInfo:    reviewInfo.VideoInfo,
		Status:       reviewInfo.Status,
	}}, nil
}
func (s *ReviewService) AuditReview(ctx context.Context, req *pb.AuditReviewRequest) (*pb.AuditReviewReply, error) {
	fmt.Printf("[service] AuditReview, req: %#v\n", req)
	return &pb.AuditReviewReply{}, nil
}
func (s *ReviewService) ReplyReview(ctx context.Context, req *pb.ReplyReviewRequest) (*pb.ReplyReviewReply, error) {
	fmt.Printf("[service] ReplyReview, req: %#v\n", req)
	
	return &pb.ReplyReviewReply{}, nil
}
func (s *ReviewService) AuditAppeal(ctx context.Context, req *pb.AuditAppealRequest) (*pb.AuditAppealReply, error) {
	return &pb.AuditAppealReply{}, nil
}
func (s *ReviewService) ListReviewByUserID(ctx context.Context, req *pb.ListReviewByUserIDRequest) (*pb.ListReviewByUserIDReply, error) {
	return &pb.ListReviewByUserIDReply{}, nil
}
