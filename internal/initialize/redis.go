package initialize

import (
	"context"
	"fmt"
	"trinity-be/global"

	"github.com/go-redis/redis/v8"
)

func InitRedis() {
	configRedis := global.Config.Redis
	rd := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", configRedis.Host, configRedis.Port),
		Username: configRedis.Username,
		Password: configRedis.Password,
		DB:       configRedis.DB,
		PoolSize: configRedis.PoolSize,
	})

	var ctx = context.Background()
	_, err := rd.Ping(ctx).Result()
	global.CheckErrorPanic(err, "Failed to connect to Redis")

	global.RedisDB = rd

	global.Logger.Info("Redis connected successfully!")
}
