package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
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

func (v *demoRepo) GetRDB() (*redis.Client, error) {
	return v.data.rdb, nil
}
