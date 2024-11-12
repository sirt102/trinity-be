package initialize

import (
	"fmt"
	"time"
	"trinity-be/global"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresQL() {
	config := global.Config.PostgresQL

	connectionString := fmt.Sprintf(
    "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    config.Host, config.Port, config.User, config.Password, config.DBName,
)

  db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
        SkipDefaultTransaction: false,
    })
	global.CheckErrorPanic(err, "PostgresQL connection failed")

	global.Logger.Info("PostgresQL connected successfully!")
	global.PostgresQLDB = db

	SetPool()
}

func SetPool() {
	config := global.Config.PostgresQL

	sqlDB, err := global.PostgresQLDB.DB()
	global.CheckErrorPanic(err, "PostgresQL went wrong!")

	sqlDB.SetConnMaxIdleTime(time.Duration(config.MaxIdleConns))
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime))
}
