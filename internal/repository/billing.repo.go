package repository

import (
	"miniBilling/global"
	"log"
)

type BillingRepository struct{}

func NewBillineRepository() *BillingRepository  {
	return &BillingRepository{}
}

func (r *BillingRepository) GetCodeLogin(userCode string) (string){
	var code string
	err := global.Billing.DB.
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