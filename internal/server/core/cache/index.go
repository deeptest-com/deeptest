package cache

import (
	"context"
	"errors"
	serverConfig "github.com/aaronchen2k/deeptest/internal/server/config"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/snowlyg/multi"
)

// Init 初始化缓存服务
func Init() error {
	universalOptions := &redis.UniversalOptions{
		Addrs:       strings.Split(serverConfig.CONFIG.Redis.Addr, ","),
		Password:    serverConfig.CONFIG.Redis.Password,
		PoolSize:    serverConfig.CONFIG.Redis.PoolSize,
		IdleTimeout: 300 * time.Second,
	}
	serverConfig.CACHE = redis.NewUniversalClient(universalOptions)
	err := multi.InitDriver(
		&multi.Config{
			DriverType:      serverConfig.CONFIG.System.CacheType,
			UniversalClient: serverConfig.CACHE},
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
	err := serverConfig.CACHE.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// DeleteCache 删除缓存数据
func DeleteCache(key string) (int64, error) {
	return serverConfig.CACHE.Del(context.Background(), key).Result()
}

// GetCacheString 获取字符串类型数据
func GetCacheString(key string) (string, error) {
	return serverConfig.CACHE.Get(context.Background(), key).Result()
}

// GetCacheBytes 获取bytes类型数据
func GetCacheBytes(key string) ([]byte, error) {
	return serverConfig.CACHE.Get(context.Background(), key).Bytes()
}

// GetCacheUint 获取uint类型数据
func GetCacheUint(key string) (uint64, error) {
	return serverConfig.CACHE.Get(context.Background(), key).Uint64()
}
