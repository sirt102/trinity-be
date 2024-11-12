package initialize

import (
	"trinity-be/global"
	"trinity-be/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger()
}
