package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	pb "helloworld/api/helloworld/v1"
)

type DemoRepo interface {
	GetDB() (*redis.Client, error)
}

type DemoUsecase struct {
	repo DemoRepo
	log  *log.Helper
}

func NewDemoUsecase(repo DemoRepo, logger log.Logger) *DemoUsecase {
	return &DemoUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DemoUsecase) Test(ctx context.Context, r *pb.TestRequest) (res string, err error) {
	switch r.Type {
	case "test1":
		res, err = uc.TestRedis(ctx, r)
	case "test2":
		return "", nil
	default:
		return "请检查传入参数 type 设置", nil
	}
	return res, err
}

func (uc *DemoUsecase) TestRedis(ctx context.Context, r *pb.TestRequest) (string, error) {
	db, _ := uc.repo.GetDB()
	db.LPush(ctx, "test", r.Type)
	return "成功", nil
}
