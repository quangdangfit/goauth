package cache

import (
	"context"
	"time"

	rl "github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
)

func newRedisLock(redis *redis.Client, cache *Cache) *RedisLockClient {
	return &RedisLockClient{locker: rl.New(redis), cache: cache}
}

type RedisLockClient struct {
	locker *rl.Client
	cache  *Cache
}

func (p *RedisLockClient) Obtain(cxt context.Context, scope, subject string, duration time.Duration) (*rl.Lock, error) {
	return p.locker.Obtain(cxt, p.cache.refineCacheKey("lock:"+scope+"-"+subject), duration, nil)
}
