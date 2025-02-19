package dal

import (
	"github.com/suutest/app/order/biz/dal/mysql"
	"github.com/suutest/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
