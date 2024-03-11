package cache

import (
	"context"
	"errors"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/snowlyg/multi"
)

// Init 初始化缓存服务
func Init() error {
	universalOptions := &redis.UniversalOptions{
		Addrs:       strings.Split(config.CONFIG.Redis.Addr, ","),
		Password:    config.CONFIG.Redis.Password,
		PoolSize:    config.CONFIG.Redis.PoolSize,
		IdleTimeout: 300 * time.Second,
	}
	config.CACHE = redis.NewUniversalClient(universalOptions)
	err := multi.InitDriver(
		&multi.Config{
			DriverType:      config.CONFIG.System.CacheType,
			UniversalClient: config.CACHE},
	)
	if err != nil {
		return err
	}
	if multi.AuthDriver == nil {
		return errors.New("初始化认证驱动失败")
	}

	return nil
}

// SetCache 缓存数据
func SetCache(key string, value interface{}, expiration time.Duration) error {
	err := config.CACHE.Set(context.Background(), config.CONFIG.Redis.Prefix+"_"+key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// DeleteCache 删除缓存数据
func DeleteCache(key string) (int64, error) {
	return config.CACHE.Del(context.Background(), config.CONFIG.Redis.Prefix+"_"+key).Result()
}

// GetCacheString 获取字符串类型数据
func GetCacheString(key string) (string, error) {
	return config.CACHE.Get(context.Background(), config.CONFIG.Redis.Prefix+"_"+key).Result()
}

// GetCacheBytes 获取bytes类型数据
func GetCacheBytes(key string) ([]byte, error) {
	return config.CACHE.Get(context.Background(), config.CONFIG.Redis.Prefix+"_"+key).Bytes()
}

// GetCacheUint 获取uint类型数据
func GetCacheUint(key string) (uint64, error) {
	return config.CACHE.Get(context.Background(), config.CONFIG.Redis.Prefix+"_"+key).Uint64()
}
