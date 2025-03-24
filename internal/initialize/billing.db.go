package initialize

import (
	"fmt"
	"time"
	"miniBilling/global"
	Billing "miniBilling/internal/pkg/mysql"


	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func checkErrorPanic(err error, errString string){
	if	err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitBilling(){
	m := global.Config.Mysql

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var  s = fmt.Sprintf(dsn,m.Username, m.Password, m.Host, m.Port, m.Billing) 
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "Init mysql Billing error")


	if global.Billing == nil {
		fmt.Println("global.Billing is nil, initializing...")
		global.Logger.Error("global.Billing is nil, initializing...")
		global.Billing = &Billing.BillingStruct{}
	}
	global.Billing.DB = db
	// global.Logger.Info("Init mysql Billing success")


	//set Pool
	 setPool()
}
func setPool(){
	
	if global.Billing == nil || global.Billing.DB == nil {
		// global.Logger.Error("Database connection is not initialized")
		return
	}

	sqlDB, err := global.Billing.DB.DB()
	if err != nil {
		// global.Logger.Error("MySQL Billing error", zap.Error(err))
		return
	}
	m := global.Config.Mysql
	sqlDB.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}