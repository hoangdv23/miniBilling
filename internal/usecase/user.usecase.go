package usecase

import (
	"log"
	"errors"
	"miniBilling/internal/po/billing"
	"miniBilling/internal/po/mongo"
	"miniBilling/internal/repository"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type Users interface {
	ListUsers() ([]billing.Users, error)
	Check_user_billing(username string) string 
	CheckTeleId(teleId int64) bool 
	UserMongo(teleId int64) (*mongo.Users, error)
	UpdateUserMongo(teleId int64, updateData bson.M) error
	Check_password_billing(username string, pasword string) ([]billing.Users, error)
	CreateUserMongo(user *mongo.Users) error
}
type UserUseCase struct {
	userRepo repository.UserRepo
}

func NewUserUseCase(userRepo repository.UserRepo) Users {
	return &UserUseCase{userRepo: userRepo}
}

func (uc *UserUseCase) ListUsers() ([]billing.Users, error) {
	return uc.userRepo.GetAllUsers()
}


func (uc *UserUseCase) Check_user_billing(username string) string {
	return uc.userRepo.Check_user_billing(username)
}

func (uc *UserUseCase) CheckTeleId(teleId int64) bool {
	var users []mongo.Users

	err := mgm.Coll(&mongo.Users{}).SimpleFind(&users, bson.M{"tele_id": teleId})
	if err != nil {
		log.Println("❌ Lỗi truy vấn MongoDB:", err)
		return false
	}
	return true
}

func (uc *UserUseCase) UserMongo(teleId int64) (*mongo.Users, error){
	user := &mongo.Users{}

	err := mgm.Coll(user).First(bson.M{"tele_id": teleId}, user)
	if err != nil {
		log.Println("❌ Không tìm thấy user:", err)
		return nil, err
	}

	return user, nil
}


func (uc *UserUseCase) UpdateUserMongo(teleId int64, updateData bson.M) error {
	collection := mgm.Coll(&mongo.Users{})
	// Tạo filter tìm user theo teleId
	filter := bson.M{"tele_id": teleId}

	// Tạo dữ liệu cập nhật (sử dụng `$set` để cập nhật chỉ các trường được chỉ định)
	update := bson.M{"$set": updateData}

	// Thực hiện cập nhật
	res, err := collection.UpdateOne(mgm.Ctx(), filter, update)
	if err != nil {
		log.Println("❌ Lỗi khi cập nhật user:", err)
		return err
	}
	// Kiểm tra xem có bản ghi nào được cập nhật không
	if res.MatchedCount == 0 {
		log.Println("⚠️ Không tìm thấy user để cập nhật.")
		return errors.New("user not found")
	}

	log.Println("✅ Cập nhật user thành công!")
	return nil
}

func (uc *UserUseCase) Check_password_billing(username string, pasword string) ([]billing.Users, error) {
	return	uc.userRepo.Check_password_billing(username, pasword)
}

func (uc *UserUseCase) CreateUserMongo(user *mongo.Users) error {
	collection := mgm.Coll(&mongo.Users{})

	// Thêm user vào MongoDB
	res, err := collection.InsertOne(mgm.Ctx(), user)
	if err != nil {
		log.Println("❌ Lỗi khi tạo user:", err)
		return err
	}

	log.Printf("✅ User được tạo với ID: %v\n", res.InsertedID)
	return nil
}
