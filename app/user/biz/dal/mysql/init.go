package mysql

import (
	"fmt"
	"os"

	"github.com/suutest/app/user/biz/model"
	"github.com/suutest/app/user/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
	) // 实现在启动时通过环境变量读取配置
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
