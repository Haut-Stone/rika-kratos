package service

import (
	"context"
	"helloworld/internal/biz"

	pb "helloworld/api/helloworld/v1"
)

type VideoService struct {
	pb.UnimplementedVideoServer
	uc *biz.VideoUsecase // - service 中注册 biz 用例
}

func NewVideoService(uc *biz.VideoUsecase) *VideoService {
	return &VideoService{uc: uc} // - 初始化 service 对象
}

func (s *VideoService) CreateVideo(ctx context.Context, req *pb.CreateVideoRequest) (*pb.CreateVideoReply, error) {
	return &pb.CreateVideoReply{}, nil
}
