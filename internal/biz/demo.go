package biz

import (
	"helloworld/internal/data"
)

type DemoUsecase struct { // - 用例对象包含 repo 对象和 log
	data *data.Data
}

func NewDemoUsecase() *DemoUsecase {
	return &DemoUsecase{}
}
