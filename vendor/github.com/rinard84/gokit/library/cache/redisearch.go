package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

func convertToStringKeys(input interface{}) interface{} {
	switch value := input.(type) {
	case map[interface{}]interface{}:
		newMap := make(map[string]interface{})
		for k, v := range value {
			newMap[fmt.Sprintf("%v", k)] = convertToStringKeys(v)
		}
		return newMap
	case []interface{}:
		for i, v := range value {
			value[i] = convertToStringKeys(v)
		}
	}
	return input
}

func NewRediSearch(redisCli *redis.Client, indexName string, args ...interface{}) *RediSearchCache {
	ctx := context.Background()
	newArgs := []interface{}{"FT.CREATE", indexName, "SCHEMA"}
	newArgs = append(newArgs, args...)
	_, err := redisCli.Do(ctx, newArgs...).Result()
	if err != nil {
		fmt.Println("NewRediSearchError", err)
	}
	return &RediSearchCache{redisCli: redisCli, indexName: indexName}
}

type RediSearchCache struct {
	redisCli  *redis.Client
	indexName string
}

func (c *RediSearchCache) AddNewDoc(ctx context.Context, docId string, score float32, args ...interface{}) error {
	newArgs := []interface{}{"FT.ADD", c.indexName, docId, score, "FIELDS"}
	newArgs = append(newArgs, args...)
	_, err := c.redisCli.Do(ctx, newArgs...).Result()
	if err != nil {
		return err
	}
	return nil
}

func (c *RediSearchCache) AddNewDocEx(ctx context.Context, docId string, score float32, expiration time.Duration, args ...interface{}) error {
	newArgs := []interface{}{"FT.ADD", c.indexName, docId, score, "FIELDS"}
	newArgs = append(newArgs, args...)
	_, err := c.redisCli.Do(ctx, newArgs...).Result()
	if err != nil {
		return err
	}
	_, err = c.redisCli.Expire(ctx, docId, expiration).Result()
	if err != nil {
		return err
	}
	return nil
}

func (c *RediSearchCache) SearchDoc(ctx context.Context, searchKeyWorld string, limit, offset int) ([]byte, error) {
	res, err := c.redisCli.Do(ctx, "FT.SEARCH", c.indexName, fmt.Sprintf("\"%s\"", searchKeyWorld), "LIMIT", fmt.Sprint(offset), fmt.Sprint(limit)).Result()
	if err != nil {
		return nil, err
	}
	convertedData := convertToStringKeys(res)
	jsonData, err := json.MarshalIndent(convertedData, "", "  ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func (c *RediSearchCache) DeleteDoc(ctx context.Context, docId string) (interface{}, error) {
	return c.redisCli.Do(ctx, "FT.DEL", c.indexName, docId, "DD").Result()
}
