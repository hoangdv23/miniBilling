package initialize

import (
	"miniBilling/global"
	"miniBilling/internal/config"
	"miniBilling/internal/pkg/logger"
)

func InitLogger(logConfig config.Logger)  {
	global.Logger = logger.NewLogger(logConfig)
}