package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "helloworld/api/helloworld/v1"
	"helloworld/internal/data"
)

type DemoUsecase struct {
	data *data.DemoRepo
	log  *log.Helper
}

func NewDemoUsecase(data *data.DemoRepo, logger log.Logger) *DemoUsecase {
	return &DemoUsecase{data: data, log: log.NewHelper(logger)}
}

func (uc *DemoUsecase) Test(ctx context.Context, r *pb.TestRequest) (string, error) {
	return "调用了一个测试方法", nil
}
