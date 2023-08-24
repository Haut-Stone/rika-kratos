package data

import (
	"context"
	pb "helloworld/api/helloworld/v1"
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

func (v *videoRepo) Save(ctx context.Context, r *pb.CreateVideoRequest) (bool, error) {
	err := v.data.rdb.LPush(ctx, "videos", r.Name+"."+r.Format).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}
