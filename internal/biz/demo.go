package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "helloworld/api/helloworld/v1"
)

type DemoRepo interface {
	Save(context.Context, *pb.CreateVideoRequest) (bool, error)
}

type DemoUsecase struct {
	repo DemoRepo
	log  *log.Helper
}

func NewDemoUsecase(repo DemoRepo, logger log.Logger) *DemoUsecase {
	return &DemoUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DemoUsecase) Test(ctx context.Context, r *pb.TestRequest) (string, error) {
	return "调用了一个测试方法", nil
}
