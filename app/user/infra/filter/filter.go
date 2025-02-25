package filter

import (
	"context"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/suutest/app/user/biz/dal/mysql"
	"github.com/suutest/app/user/biz/model"
)

var Filter *bloom.BloomFilter

func InitFilter() {
	Filter = bloom.NewWithEstimates(1000000, 0.01)
	users, err := model.GetAllUserEmail(context.Background(), mysql.DB)
	if err != nil {
		return
	}
	for _, u := range users {
		Filter.AddString(u.Email)
	}
}
