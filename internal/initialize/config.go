package initialize

import (
	"fmt"
	"trinity-be/config"
	"trinity-be/global"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	global.CheckErrorPanic(err, "Failed to load .env file")

	var config config.Config

	err = cleanenv.ReadEnv(&config)
	if err != nil {
		fmt.Println(err)
	}
	global.CheckErrorPanic(err, "Failed to load config")

	global.Config = config
}
