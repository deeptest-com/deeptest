package serverConsts

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	VIPER      *viper.Viper          // viper
	CACHE      redis.UniversalClient // 缓存
	PermRoutes []map[string]string   // 权限路由
)
