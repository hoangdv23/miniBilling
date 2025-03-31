package initialize

import (
	"fmt"
	"miniBilling/global"
	Billing "miniBilling/internal/pkg/mysql"


	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func InitVoiceReport(){
	m := global.Config.Mysql136

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var  s = fmt.Sprintf(dsn,m.Username, m.Password, m.Host, m.Port, m.VoiceReport) 
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "Init mysql Voice Report error")


	if global.VoiceReport == nil {
		fmt.Println("global.VoiceReport is nil, initializing...")
		global.Logger.Error("global.VoiceReport is nil, initializing...")
		global.VoiceReport = &Billing.VoiceReportStruct{}
	}
	global.VoiceReport.DB = db


	 setPool()
}
