package initialize

import (
	"miniBilling/global"
	"miniBilling/internal/config"
	"miniBilling/internal/pkg/bot"
)

func InitTeleBot(botConfig config.Bot) *bot.TeleBot {
	global.Bot = bot.NewBot(botConfig)
	if global.Bot == nil {
		panic("Khởi tạo Bot thất bại!")
	}
	return global.Bot
}