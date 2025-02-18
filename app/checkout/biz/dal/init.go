package dal

import (
	"github.com/suutest/app/checkout/biz/dal/mysql"
	"github.com/suutest/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
