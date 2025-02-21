package dal

import (
	"github.com/suutest/app/product/biz/dal/mysql"
	"github.com/suutest/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
