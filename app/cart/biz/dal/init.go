package dal

import (
	"github.com/suutest/app/cart/biz/dal/mysql"
	"github.com/suutest/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
