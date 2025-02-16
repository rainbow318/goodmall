package dal

import (
	"github.com/suutest/app/frontend/biz/dal/mysql"
	"github.com/suutest/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
