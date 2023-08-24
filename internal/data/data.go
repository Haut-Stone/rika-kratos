package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"helloworld/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewVideoMysqlRepo) // - 如果需要切换数据实现,只需要在这里更换 provider 就可以了

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	// mysql database
	db, err := gorm.Open(mysql.Open(c.GetDatabase().GetSource()), &gorm.Config{}) //mysql数据库连接
	if err != nil {
		return nil, nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		DB:           int(c.Redis.Db),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	})

	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, nil, err
	}

	data := &Data{
		db:  db,
		rdb: rdb,
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		_db, err := data.db.DB()
		if err != nil {
			log.NewHelper(logger).Error("closing the mysql error", err)
		}
		_ = _db.Close()
		log.NewHelper(logger).Info("closed the mysql")
		err = data.rdb.Close()
		if err != nil {
			log.NewHelper(logger).Error("closing the redis error", err)

		}
		log.NewHelper(logger).Info("closed the redis")
	}
	return data, cleanup, nil
}
