package usecase

import (
	"log"
	"miniBilling/internal/po/mongo"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"miniBilling/internal/repository"
)
type BillingUseCase struct {
	billingRepo *repository.BillingRepository
}

func NewBillingUsecase(billingRepo *repository.BillingRepository) *BillingUseCase {
	return &BillingUseCase{billingRepo: billingRepo}
}

func (uc *BillingUseCase) GetCodeLogin(userCode string) string {
	return	uc.billingRepo.GetCodeLogin(userCode)
}

func (uc *BillingUseCase) UserMongo(teleId int64) (*mongo.Users, error){
	user := &mongo.Users{}
	err := mgm.Coll(user).First(bson.M{"tele_id": teleId}, user)
	if err != nil {
		log.Println("❌ Không tìm thấy user trong billing:", err)
		return nil, err
	}

	return user, nil
}
