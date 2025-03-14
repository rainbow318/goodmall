package dal

import (
	"github.com/suutest/app/stock/biz/dal/mysql"
	"github.com/suutest/app/stock/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
