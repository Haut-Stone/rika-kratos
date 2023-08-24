package data

import (
	"context"
	pb "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type videoMysqlRepo struct {
	data *Data
	log  *log.Helper
}

func NewVideoMysqlRepo(data *Data, logger log.Logger) biz.VideoRepo { // - repo 的实现
	return &videoMysqlRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (v *videoMysqlRepo) Save(ctx context.Context, r *pb.CreateVideoRequest) (bool, error) {
	log.Info("Saving video to mysql")
	return true, nil
}
