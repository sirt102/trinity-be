package global

import (
	"errors"
	"trinity-be/config"
	"trinity-be/pkg/logger"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config       config.Config
	Logger       *logger.LoggerZap
	PostgresQLDB *gorm.DB
	RedisDB      *redis.Client
)

func CheckErrorPanic(err error, errString string) {
	if err != nil {
		Logger.DPanic(errString, zap.Error(err))
		panic(err)
	}
}

func LogError(err error, errString string) {
	if errString == "" {
		errString = err.Error()
	}

	if err != nil {
		Logger.Error(errString, zap.Error(err))
	}
}

func NewError(errMessage string) error {
	return errors.New(errMessage)
}
