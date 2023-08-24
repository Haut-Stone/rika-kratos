package data

import (
	"context"
	"helloworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type videoRepo struct {
	data *Data
	log  *log.Helper
}

func NewVideoRepo(data *Data, logger log.Logger) biz.VideoRepo { // - repo 的实现
	return &videoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (v *videoRepo) Save(ctx context.Context, g *biz.Video) (*biz.Video, error) {
	return nil, nil
}
