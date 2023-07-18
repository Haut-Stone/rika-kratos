package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"helloworld/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

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

	data := &Data{
		db: db,
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		_db, err := data.db.DB()
		if err != nil {
			log.NewHelper(logger).Error("closing the mysql error", err)
		}
		_ = _db.Close()
		log.NewHelper(logger).Info("closing the mysql")
	}
	return data, cleanup, nil
}
