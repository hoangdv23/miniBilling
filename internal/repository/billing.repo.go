package repository

import (
	"log"
	"gorm.io/gorm"
)

type BillingRepo interface {
	GetCodeLogin(userCode string) (string)

}
type BillingRepository struct{
	billingDB *gorm.DB
}

func NewBillineRepository(billingDB *gorm.DB) BillingRepo  {
	return &BillingRepository{billingDB: billingDB}
}

func (r *BillingRepository) GetCodeLogin(userCode string) (string){
	var code string
	err := r.billingDB.
		Table("users").
		Where("user_code = ?", userCode).
		Pluck("two_factor_code", &code).Error
	if err != nil {
		log.Printf("❌ Lỗi khi truy vấn dữ liệu cho người dùng %s: %v", userCode, err)
		return ""
	}
	if code == "" {
		log.Printf("⚠️ Không tìm thấy mã nhân viên: %s", userCode)
		return ""
	}

	log.Printf("✅ Đã tìm thấy người dùng: %s", code)
	return code
}