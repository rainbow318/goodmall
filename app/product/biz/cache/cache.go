package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var LocalCache *cache.Cache

func InitLocalCache() {
	LocalCache = cache.New(5*time.Minute, 10*time.Minute)
}
