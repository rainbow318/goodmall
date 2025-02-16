package dal

import (
	"github.com/suuyh/demo/demo_thrift/biz/dal/mysql"
	"github.com/suuyh/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
