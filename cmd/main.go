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
	
	initialize.InitBilling()
	initialize.InitVoiceReport()
	initialize.InitMongo()
	handler.NewServer(bot.Bot)
	
	
	global.Bot.StartBot() // khởi động bot

}