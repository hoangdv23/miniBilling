package repository

import (
	"log"

	"miniBilling/internal/po/billing"
	"miniBilling/internal/po/mongo"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	mogodb "go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetAllUsers() ([]billing.Users, error)
	Check_user_billing(user_name string) string
	Check_password_billing(username string, password string) ([]billing.Users, error)
	GetUsers(teleId int64) (*mongo.Users, error)
}
type UserRepository struct {
	billingDB *gorm.DB
}

func NewUserRepository(billingDB *gorm.DB) UserRepo {
	return &UserRepository{billingDB: billingDB}
}

// Lấy danh sách users (giới hạn 3)
func (r *UserRepository) GetAllUsers() ([]billing.Users, error) {
	var users []billing.Users
	result := r.billingDB.Limit(3).Find(&users)
	return users, result.Error
}

// Kiểm tra user theo user_code
func (r *UserRepository) Check_user_billing(user_name string) string {
	var found_user_name string
	err := r.billingDB.
		Table("users").
		Where("user_code = ?", user_name).
		Pluck("user_name", &found_user_name).Error

	if err != nil {
		log.Printf("❌ Lỗi khi truy vấn dữ liệu cho người dùng %s: %v", user_name, err)
		return ""
	}

	if found_user_name == "" {
		log.Printf("⚠️ Không tìm thấy người dùng với tên: %s", user_name)
		return ""
	}

	log.Printf("✅ Đã tìm thấy người dùng: %s", found_user_name)
	return found_user_name
}

// Kiểm tra mật khẩu người dùng
func (r *UserRepository) Check_password_billing(username string, password string) ([]billing.Users, error) {
	var users []billing.Users
	result := r.billingDB.
		Where("user_code = ? AND password_show = ?", username, password).
		Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

// Lấy user từ Mongo theo TeleId
func (r *UserRepository) GetUsers(teleId int64) (*mongo.Users, error) {
	collection := mgm.Coll(&mongo.Users{})
	filter := bson.M{"tele_id": teleId}

	var user mongo.Users
	err := collection.FindOne(mgm.Ctx(), filter).Decode(&user)
	if err != nil {
		if err == mogodb.ErrNoDocuments {
			log.Println("❌ Không tìm thấy user")
			return nil, nil
		}
		log.Println("❌ Lỗi khi tìm kiếm user:", err)
		return nil, err
	}

	return &user, nil
}
