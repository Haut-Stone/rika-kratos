package test

import (
	"fmt"
	"github.com/spf13/cast"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"helloworld/internal/data/model"
	"testing"
)

func TestGenMySQLData(t *testing.T) {
	// mysql database
	db, _ := gorm.Open(mysql.Open("root:159753zzz@tcp(127.0.0.1:3306)/cut_table"), &gorm.Config{}) // mysql数据库连接
	for i := 0; i < 100000; i++ {
		db.Create(&model.WideCut{
			Name: "测试名称" + cast.ToString(i),
		})
	}
}

func TestGetMySQLData(t *testing.T) {
	// 获取数据看查询所有数据的时间
	db, _ := gorm.Open(mysql.Open("root:159753zzz@tcp(127.0.0.1:3306)/cut_table"), &gorm.Config{}) // mysql数据库连接
	res := db.Find(&[]model.WideCut{})
	fmt.Println(res)
}
