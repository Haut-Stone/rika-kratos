package data

import (
	"github.com/go-kratos/kratos/v2/log"
)

type DemoRepo struct {
	data *Data
	log  *log.Helper
}

func NewDemoRepo(data *Data, logger log.Logger) *DemoRepo { // - repo 的实现
	return &DemoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
