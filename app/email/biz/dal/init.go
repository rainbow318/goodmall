package dal

import (
	"github.com/suutest/app/email/biz/dal/mysql"
	"github.com/suutest/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
