package global

import (
	"trinity-be/config"
	"trinity-be/pkg/logger"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config config.Config
	Logger *logger.LoggerZap
	PostgresQLDB *gorm.DB
	RedisDB *redis.Client
)

func CheckErrorPanic(err error, errString string) {
	if err != nil {
		Logger.DPanic(errString, zap.Error(err))
		panic(err)
	}
}
