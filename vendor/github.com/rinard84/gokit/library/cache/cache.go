package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	redis_rate "github.com/go-redis/redis_rate/v10"
	redis "github.com/redis/go-redis/v9"
)

type Cache struct {
	redis       *redis.Client
	log         logr.Logger
	RateLimiter *RateLimitCache
	RedisLock   *RedisLockClient
	RediSearch  *RediSearchCache
	keyPattern  string
}

type Option = func(*Cache)
type ProcessFunc func(keys []string) error

func NewCache(log logr.Logger, redis *redis.Client, options ...Option) *Cache {
	redis.Set(context.Background(), "NewCache", "value", 0).Err()

	cl := &Cache{
		log:        log,
		redis:      redis,
		keyPattern: "%s",
	}
	cl.RateLimiter = newRateLimiter(redis_rate.NewLimiter(redis), cl)
	cl.RedisLock = newRedisLock(redis, cl)
	for _, o := range options {
		o(cl)
	}
	return cl
}

func WithKeyPattern(keyPattern string) Option {
	return func(c *Cache) { c.keyPattern = keyPattern }
}

func (r *Cache) Ping(ctx context.Context) error {
	return r.redis.Ping(ctx).Err()
}

func (r *Cache) Close() error {
	//return r.redis.Close()
	return nil
}

func (r *Cache) Client() *redis.Client {
	return r.redis
}

func (r *Cache) refineCacheKey(key string) string {
	return fmt.Sprintf(r.keyPattern, key)
}

func (r *Cache) refineMultipleCacheKey(keys []string) []string {
	res := make([]string, len(keys))
	for idx, item := range keys {
		res[idx] = fmt.Sprintf(r.keyPattern, item)
	}
	return res
}

func (r *Cache) refineMultipleMSetCacheKey(pairs map[string]interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	for idx, item := range pairs {
		res[fmt.Sprintf(r.keyPattern, idx)] = item
	}
	return res
}

func (r *Cache) Set(ctx context.Context, key string, value interface{}, expireSecond time.Duration) error {
	jsonValue, err := json.Marshal(value)

	if err != nil {
		return err
	}
	return r.redis.Set(ctx, r.refineCacheKey(key), jsonValue, expireSecond).Err()
}
func (r *Cache) Get(ctx context.Context, key string, output interface{}) error {
	value, err := r.redis.Get(ctx, r.refineCacheKey(key)).Result()
	if err != nil {
		return err
	}
	if err = json.Unmarshal([]byte(value), output); err != nil {
		return err
	}
	return nil
}
func (r *Cache) MGet(ctx context.Context, keys []string) ([]interface{}, error) {
	multipleKey := r.refineMultipleCacheKey(keys)
	values, err := r.redis.MGet(ctx, multipleKey...).Result()
	if err != nil {
		return nil, err
	}
	return values, nil
}
func (r *Cache) MSet(ctx context.Context, pairs map[string]interface{}) error {
	pairs = r.refineMultipleMSetCacheKey(pairs)
	return r.redis.MSet(ctx, pairs).Err()
}

func (r *Cache) MSetWithExpire(ctx context.Context, pairs map[string]interface{}, expire time.Duration) error {
	pairs = r.refineMultipleMSetCacheKey(pairs)
	err := r.redis.MSet(ctx, pairs).Err()
	if err != nil {
		return err
	}

	for key := range pairs {
		err := r.redis.Expire(ctx, key, expire).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Cache) SetString(ctx context.Context, key string, value string, expireSecond time.Duration) error {
	return r.redis.Set(ctx, r.refineCacheKey(key), value, expireSecond).Err()
}

func (r *Cache) GetString(ctx context.Context, key string) (string, error) {
	return r.redis.Get(ctx, r.refineCacheKey(key)).Result()
}

func (r *Cache) Exists(ctx context.Context, key string) bool {
	if vl, err := r.redis.Exists(ctx, r.refineCacheKey(key)).Result(); err == nil && vl == 1 {
		return true
	}
	return false
}
func (r *Cache) Del(ctx context.Context, key string) error {
	cacheKey := r.refineCacheKey(key)
	return r.redis.Del(ctx, cacheKey).Err()
}
func (r *Cache) DelMultiple(ctx context.Context, keys []string) error {
	cacheKeys := make([]string, 0, len(keys))
	for _, key := range keys {
		cacheKeys = append(cacheKeys, r.refineCacheKey(key))
	}
	return r.redis.Del(ctx, cacheKeys...).Err()
}

func (r *Cache) DelByPattern(ctx context.Context, pattern string) error {
	pattern = r.refineCacheKey(pattern)
	keys, _ := r.redis.Keys(ctx, pattern).Result()
	for _, key := range keys {
		err := r.redis.Del(ctx, key).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Cache) Incr(ctx context.Context, key string) error {
	cacheKey := r.refineCacheKey(key)
	return r.redis.Incr(ctx, cacheKey).Err()
}

func (r *Cache) IncrResVal(ctx context.Context, key string, expireSecond time.Duration) (int64, error) {
	cacheKey := r.refineCacheKey(key)
	idInc, err := r.redis.Incr(ctx, cacheKey).Result()
	if err == nil {
		r.redis.Expire(ctx, cacheKey, expireSecond)
	}
	return idInc, err
}

func (r *Cache) Decr(ctx context.Context, key string) error {
	cacheKey := r.refineCacheKey(key)
	return r.redis.Decr(ctx, cacheKey).Err()
}

func (r *Cache) Keys(ctx context.Context, pattern string) ([]string, error) {
	pattern = r.refineCacheKey(pattern)
	return r.redis.Keys(ctx, pattern).Result()
}

func (r *Cache) IsTTL(ctx context.Context, key string) bool {
	cacheKey := r.refineCacheKey(key)
	// return r.redis.Decr(ctx, cacheKey).Err()
	ttl, err := r.redis.TTL(ctx, cacheKey).Result()
	if err != nil {
		return true
	}
	if ttl < 1 {
		return true
	}
	return false
}

// use for multiple keys, we limit the number of key, if key > 1000 key, I limit 1000
func (r *Cache) Scan(ctx context.Context, pattern string, limit int) ([]string, error) {
	pattern = r.refineCacheKey(pattern)
	if limit > 1000 {
		limit = 1000
	}
	var cursor uint64
	var keys []string

	for {
		var nkeys []string
		var err error

		nkeys, cursor, err = r.redis.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			return nil, err
		}

		keys = append(keys, nkeys...)
		if len(keys) >= limit || cursor == 0 {
			break
		}
	}

	if len(keys) > limit {
		keys = keys[:limit]
	}

	return keys, nil
}

// handle key batch to avoid over memory of service
func (r *Cache) ScanAndProcessKeys(ctx context.Context, pattern string, numOfKey int64, processFunc ProcessFunc) error {
	pattern = r.refineCacheKey(pattern)
	var cursor uint64
	for {
		keys, nextCursor, err := r.redis.Scan(ctx, cursor, pattern, numOfKey).Result()
		if err != nil {
			return err
		}

		if err := processFunc(keys); err != nil {
			return err
		}

		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}
	return nil
}

func (r *Cache) HScanAndProcessKeys(ctx context.Context, hgetallKeys, patter string, numOfKey int64, processFunc ProcessFunc) error {
	hgetallKeys = r.refineCacheKey(hgetallKeys)
	var cursor uint64
	for {
		keys, nextCursor, err := r.redis.HScan(ctx, hgetallKeys, cursor, patter, numOfKey).Result()
		if err != nil {
			return err
		}

		if err := processFunc(keys); err != nil {
			return err
		}

		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}
	return nil
}

func (r *Cache) HGet(ctx context.Context, key, subkey string, output interface{}) error {
	cacheKey := r.refineCacheKey(key)
	value, err := r.redis.HGet(ctx, cacheKey, subkey).Result()
	if err != nil {
		return err
	}
	if err = json.Unmarshal([]byte(value), output); err != nil {
		return err
	}
	return nil
}

func (r *Cache) HSet(ctx context.Context, key, subkey string, value interface{}) error {
	cacheKey := r.refineCacheKey(key)
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.redis.HSet(ctx, cacheKey, subkey, jsonValue).Err()
}

func (r *Cache) HDel(ctx context.Context, key, subkey string) error {
	cacheKey := r.refineCacheKey(key)
	return r.redis.HDel(ctx, cacheKey, subkey).Err()
}

func (r *Cache) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	cacheKey := r.refineCacheKey(key)
	return r.redis.HGetAll(ctx, cacheKey).Result()
}

func (r *Cache) Expire(ctx context.Context, key string, expire time.Duration) error {
	cacheKey := r.refineCacheKey(key)
	return r.redis.Expire(ctx, cacheKey, expire).Err()
}
