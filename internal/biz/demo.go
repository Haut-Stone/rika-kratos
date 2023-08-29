package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	pb "helloworld/api/helloworld/v1"
)

type DemoRepo interface {
	GetRDB() (*redis.Client, error)
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
		res, err = uc.TestRedisHash(ctx, r)
	case "test2":
		res, err = uc.TestRedisSet(ctx, r)
	default:
		return "请检查传入参数 type 设置", nil
	}
	return res, err
}

func (uc *DemoUsecase) TestRedisHash(ctx context.Context, r *pb.TestRequest) (res string, err error) {
	rdb, _ := uc.repo.GetRDB()

	// - 联通测试
	rdb.Ping(ctx)

	// - 可以直接放入单层 map, 但不可以直接放入一个有多于一个元素的结构体, 多层的结构里面每个都需要实现出一个编码器才可以
	err = rdb.HSet(ctx, "hashMapData:1", map[string]interface{}{
		"age":  "hitori",
		"name": "boqi",
	}).Err()

	// - 因为 hash 存储单个对象, 但是对象一般都有很多,所以一般用文件夹的 key 标识来构成数组
	res = rdb.HGet(ctx, "hashMapData:1", "key1").Val()
	keys := rdb.HKeys(ctx, "hashMapData:1").Val()
	fmt.Println(keys)
	data := rdb.HScan(ctx, "hashMapData:1", 0, "*1", 10)
	fmt.Println("get data: ===> ", data)

	// - list
	rdb.LPush(ctx, "test", r.Type)
	return res, err
}

func (uc *DemoUsecase) TestRedisSet(ctx context.Context, r *pb.TestRequest) (res string, err error) {
	rdb, _ := uc.repo.GetRDB()

	// - set
	rdb.SAdd(ctx, "test", r.Type)
	return res, err
}
