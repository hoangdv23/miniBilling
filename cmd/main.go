package main

import (
	"miniBilling/global"
	"miniBilling/internal/handler"
	"miniBilling/internal/initialize"

)




func main() {
	cfg := 	initialize.InitConfig() // khai báo config
			initialize.InitLogger(cfg.Log) // khai báo log
	bot := 	initialize.InitTeleBot(cfg.Bot)// khai báo bot 

	handler.NewServer(bot.Bot)
	initialize.InitBilling()
	initialize.InitMongo()
	
	global.Bot.StartBot() // khởi động bot

}