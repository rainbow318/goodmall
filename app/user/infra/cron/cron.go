package cron

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/suutest/app/user/biz/dal/mysql"
	"github.com/suutest/app/user/biz/model"
)

var cronIns *cron.Cron

func InitCorn() {
	cronIns = cron.New()

	cronIns.AddFunc("@every 1m", HardDeleteUser())
	cronIns.Start()
}

func HardDeleteUser() {
	fmt.Println("HardDeleteUser")
	mysql.DB.Unscoped().Where("deleted_at IS NOT NULL").Delete(&model.User{})
}
