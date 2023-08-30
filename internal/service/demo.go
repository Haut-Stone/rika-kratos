package service

import (
	"context"
	"fmt"
	"helloworld/internal/biz"

	pb "helloworld/api/helloworld/v1"
)

type DemoService struct {
	pb.UnimplementedDemoServer
	uc *biz.DemoUsecase
}

func NewDemoService(uc *biz.DemoUsecase) *DemoService {
	return &DemoService{uc: uc}
}

func (s *DemoService) CreateDemo(_ context.Context, _ *pb.CreateDemoRequest) (*pb.CreateDemoReply, error) {
	return &pb.CreateDemoReply{}, nil
}
func (s *DemoService) UpdateDemo(_ context.Context, _ *pb.UpdateDemoRequest) (*pb.UpdateDemoReply, error) {
	return &pb.UpdateDemoReply{}, nil
}
func (s *DemoService) DeleteDemo(_ context.Context, _ *pb.DeleteDemoRequest) (*pb.DeleteDemoReply, error) {
	return &pb.DeleteDemoReply{}, nil
}
func (s *DemoService) GetDemo(_ context.Context, _ *pb.GetDemoRequest) (*pb.GetDemoReply, error) {
	return &pb.GetDemoReply{}, nil
}
func (s *DemoService) ListDemo(_ context.Context, _ *pb.ListDemoRequest) (*pb.ListDemoReply, error) {
	return &pb.ListDemoReply{}, nil
}
func (s *DemoService) Test(ctx context.Context, req *pb.TestRequest) (*pb.TestReply, error) {
	res, err := s.uc.Test(ctx, req)
	if err != nil {
		return &pb.TestReply{Message: res, ErrorMsg: err.Error()}, nil
	}
	fmt.Println(res)
	return &pb.TestReply{Message: res, ErrorMsg: ""}, nil
}
