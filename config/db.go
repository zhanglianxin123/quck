package config

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", Config.DBUsername,
			Config.DBPassword, Config.DBHostname, Config.DBPort, Config.DBDatabase),
	)

	if err != nil {
		panic(err)
	}
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetMaxIdleConns(20)
	//mysql 保持连接时间最长为8小时
	DB.DB().SetConnMaxLifetime(time.Hour)
	//禁用表名复数
	DB.SingularTable(true)

	DB.BlockGlobalUpdate(true)

	hlog.Info("connect main succeed")
}
