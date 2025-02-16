package dal

import (
	"github.com/suutest/app/product/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
