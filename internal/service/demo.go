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

func (s *DemoService) CreateDemo(ctx context.Context, req *pb.CreateDemoRequest) (*pb.CreateDemoReply, error) {
	return &pb.CreateDemoReply{}, nil
}
func (s *DemoService) UpdateDemo(ctx context.Context, req *pb.UpdateDemoRequest) (*pb.UpdateDemoReply, error) {
	return &pb.UpdateDemoReply{}, nil
}
func (s *DemoService) DeleteDemo(ctx context.Context, req *pb.DeleteDemoRequest) (*pb.DeleteDemoReply, error) {
	return &pb.DeleteDemoReply{}, nil
}
func (s *DemoService) GetDemo(ctx context.Context, req *pb.GetDemoRequest) (*pb.GetDemoReply, error) {
	return &pb.GetDemoReply{}, nil
}
func (s *DemoService) ListDemo(ctx context.Context, req *pb.ListDemoRequest) (*pb.ListDemoReply, error) {
	return &pb.ListDemoReply{}, nil
}
func (s *DemoService) Test(ctx context.Context, req *pb.TestRequest) (*pb.TestReply, error) {
	res, err := s.uc.Test(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return &pb.TestReply{Message: res}, nil
}
