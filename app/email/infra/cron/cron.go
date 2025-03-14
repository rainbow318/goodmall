package cron

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

var cronIns *cron.Cron

func InitCorn() {
	cronIns = cron.New()

	cronIns.AddFunc("@every 1m", SendHotSaleEmail())
	cronIns.Start()
}

func SendHotSaleEmail() {
	fmt.Println("tick test")
}
