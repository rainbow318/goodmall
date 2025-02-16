package dal

import (
	"github.com/suuyh/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
