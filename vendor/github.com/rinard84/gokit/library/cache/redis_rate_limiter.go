package cache

import (
	"context"

	"github.com/go-redis/redis_rate/v10"
)

func newRateLimiter(limiter *redis_rate.Limiter, cache *Cache) *RateLimitCache {

	return &RateLimitCache{limiter: limiter, cache: cache}
}

type RateLimitCache struct {
	limiter *redis_rate.Limiter
	cache   *Cache
}

func (p *RateLimitCache) Reset(cxt context.Context, scope, subject string) error {
	return p.limiter.Reset(cxt, p.cache.refineCacheKey("rate:"+scope+"-"+subject))
}

func (p *RateLimitCache) RateLimit(cxt context.Context, scope, subject string, limit redis_rate.Limit) (*redis_rate.Result, error) {
	return p.limiter.Allow(cxt, p.cache.refineCacheKey("rate:"+scope+"-"+subject), limit)
}

func (p *RateLimitCache) GetRateLimit(cxt context.Context, scope, subject string, limit redis_rate.Limit) (*redis_rate.Result, error) {
	return p.limiter.AllowN(cxt, p.cache.refineCacheKey("rate:"+scope+"-"+subject), limit, 0)
}
