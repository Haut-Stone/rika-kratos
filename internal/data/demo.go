package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
)

type demoRepo struct {
	data *Data
	log  *log.Helper
}

func NewDemoRepo(data *Data, logger log.Logger) biz.DemoRepo { // - repo 的实现
	return &demoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (v *demoRepo) Save(ctx context.Context, r *pb.CreateVideoRequest) (bool, error) {
	err := v.data.rdb.LPush(ctx, "videos", r.Name+"."+r.Format).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}
