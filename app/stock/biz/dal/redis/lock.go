package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

var (
	lock_ttl          = 120 * time.Second // 锁有效期
	lock_wait_timeout = 3 * time.Second   // 进程等待锁的最长时间
)

func TryLock(ctx context.Context, key string, uniqueValue string) bool {
	startTime := time.Now()
	for {
		ok, err := RedisClient.SetNX(ctx, key, uniqueValue, lock_ttl).Result()
		if err != nil {
			klog.Error("SETNX error")
			return false
		}
		if ok {
			return true
		}
		if time.Since(startTime) > lock_wait_timeout {
			fmt.Println("getting lock timeout")
			return false
		}
		// 等待一段时间后再重试，以免CPU空转
		time.Sleep(100 * time.Millisecond)
	}
}

func Unlock(ctx context.Context, key string, uniqueValue string) bool {
	luaScript := `
		if redis.call("GET",KEYS[1])==ARGV[1] then
			return redis.call("DEL",KEYS[1])
		else
			return 0
		end
	`
	result, err := RedisClient.Eval(ctx, luaScript, []string{key}, uniqueValue).Int()
	if err != nil {
		fmt.Println("delete lock fail, ", err.Error())
		return false
	}
	return result == 1
}
