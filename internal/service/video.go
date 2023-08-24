package service

import (
	"context"

	pb "helloworld/api/helloworld/v1"
)

type VideoService struct {
	pb.UnimplementedVideoServer
}

func NewVideoService() *VideoService {
	return &VideoService{}
}

func (s *VideoService) CreateVideo(ctx context.Context, req *pb.CreateVideoRequest) (*pb.CreateVideoReply, error) {
	return &pb.CreateVideoReply{}, nil
}
