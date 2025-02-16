package dal

import (
	"github.com/suutest/app/user/biz/dal/mysql"
	"github.com/suutest/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
