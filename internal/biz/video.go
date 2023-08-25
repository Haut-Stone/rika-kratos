package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "helloworld/api/helloworld/v1"
)

type Video struct {
	Hello string
}

type VideoRepo interface { // - repo 是个接口, 谁都可以实现
	Save(context.Context, *pb.CreateVideoRequest) (bool, error)
}

type VideoUsecase struct { // - 用例对象包含 repo 对象和 log
	repo VideoRepo
	log  *log.Helper
}

func NewVideoUsecase(repo VideoRepo, logger log.Logger) *VideoUsecase {
	return &VideoUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *VideoUsecase) CreateVideo(ctx context.Context, r *pb.CreateVideoRequest) (bool, error) {
	success, err := uc.repo.Save(ctx, r)
	if err != nil {
		return false, err
	}
	return success, nil // - biz 用例内的具体执行方法
}
