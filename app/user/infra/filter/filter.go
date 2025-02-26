package filter

import (
	"context"
	"time"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/suutest/app/user/biz/dal/mysql"
	"github.com/suutest/app/user/biz/model"
)

var (
	Filter      *bloom.BloomFilter
	FilterState bool
)

func InitFilter() {
	Filter = bloom.NewWithEstimates(1000000, 0.01)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []*model.User
	var err error
	for i := 0; i < 3; i++ { // 最多重试三次
		users, err = model.GetAllUserEmail(ctx, mysql.DB) // 使用带超时的上下文
		if err == nil {
			break
		}
	}
	if err != nil {
		klog.Error("Bloom filter init failed")
		FilterState = false
		return
	}
	for _, u := range users {
		Filter.AddString(u.Email)
	}
	FilterState = true
}
