package repository

import (
	"log"
	"miniBilling/global"
	"miniBilling/internal/po/billing"
	"miniBilling/internal/po/mongo"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	mogodb "go.mongodb.org/mongo-driver/mongo"
)

// UserRepository xử lý database cho User
type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Lấy danh sách users
func (r *UserRepository) GetAllUsers() ([]billing.Users, error) {
	var users []billing.Users
	result := global.Billing.DB.Limit(3).Find(&users)
	return users, result.Error
}

func (r *UserRepository) Check_user_billing(user_name string) string {
	var found_user_name string

	err := global.Billing.DB.
		Table("users").
		Where("user_code = ?", user_name).
		Pluck("user_name", &found_user_name).Error

	// Kiểm tra lỗi kết nối hoặc lỗi truy vấn
	if err != nil {
		log.Printf("❌ Lỗi khi truy vấn dữ liệu cho người dùng %s: %v", user_name, err)
		return ""
	}

	// Nếu không có dữ liệu, `found_user_name` sẽ là chuỗi rỗng
	if found_user_name == "" {
		log.Printf("⚠️ Không tìm thấy người dùng với tên: %s", user_name)
		return ""
	}

	log.Printf("✅ Đã tìm thấy người dùng: %s", found_user_name)
	return found_user_name
}


// Nếu có tk thì nhập mk vào
func (r *UserRepository) Check_password_billing(username string, password string) ([]billing.Users, error) {
	var users []billing.Users
	result := global.Billing.DB.Where("user_code = ? AND password_show = ?", username, password).Find(&users)

	// Kiểm tra lỗi trong truy vấn
	if result.Error != nil {
		return nil, result.Error
	}

	// Nếu không tìm thấy user nào, trả về slice rỗng thay vì nil
	if len(users) == 0 {
		return []billing.Users{}, nil
	}

	return users, nil
}




func (r *UserRepository) GetUsers(teleId int64) (*mongo.Users, error) {
	collection := mgm.Coll(&mongo.Users{})

	// Tạo filter để kiểm tra user theo teleId
	filter := bson.M{"tele_id": teleId}

	// Tìm user thỏa mãn điều kiện
	var user mongo.Users
	err := collection.FindOne(mgm.Ctx(), filter).Decode(&user)
	if err != nil {
		if err == mogodb.ErrNoDocuments { // Sửa lại lỗi sai tên package
			log.Println("❌ Không tìm thấy user")
			return nil, nil // Trả về nil thay vì false
		}
		log.Println("❌ Lỗi khi tìm kiếm user:", err)
		return nil, err // Trả về error nếu có lỗi
	}

	return &user, nil // Trả về con trỏ đến user
}
