package initialize

import (
	"log"
	"miniBilling/global"
	"miniBilling/internal/config"
)

func InitConfig() *config.Config{
	var err error
	global.Config, err = config.LoadConfig()
	if err != nil {
		log.Fatalf("Lỗi tải config: %v", err)
	}
	return global.Config
}