package initialize

import (
	"log"
	"fmt"
	"miniBilling/global"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo() {
	m := global.Config.Mongo_db
	
	clientOptions := options.Client().
		ApplyURI(m.Url_mongo).
		SetMaxPoolSize(100). // Tối đa 100 kết nối trong pool
		SetMinPoolSize(10)   // Giữ sẵn 10 kết nối


	err := mgm.SetDefaultConfig(nil, m.DB_mongo, clientOptions)
	if err != nil {
		log.Fatal("❌ Lỗi kết nối MongoDB:", err)
	}
	fmt.Println("✅ Kết nối MongoDB thành công!")
}