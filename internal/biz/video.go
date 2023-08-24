package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Video struct {
	Hello     string
	TableName string
}

type VideoRepo interface { // - repo 是个接口, 谁都可以实现
	Save(context.Context, *Video) (*Video, error)
}

type VideoUsecase struct { // - 用例对象包含 repo 对象和 log
	repo VideoRepo
	log  *log.Helper
}

func NewVideoUsecase(repo VideoRepo, logger log.Logger) *VideoUsecase {
	return &VideoUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GreeterUsecase) CreateVideo(ctx context.Context, g *Greeter) (*Video, error) {
	return nil, nil // - biz 用例内的具体执行方法
}
