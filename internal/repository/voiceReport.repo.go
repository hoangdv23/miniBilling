package repository

import (
	"fmt"
	"log"
	"miniBilling/global"
	"miniBilling/internal/po/billing"
	"sync"
	"time"

	"gorm.io/gorm"
)

type VoiceReport interface {
	GetCdrOutVas(vas string, year int, month int) ([]billing.CdrRecord, error)
	GetCdrInVas(vas string, year int, month int) ([]billing.CdrRecord, error)
}
type VoiceReportRepository struct{
	db *gorm.DB
}

func NewVoiceReportRepository(db *gorm.DB) VoiceReport {
	return &VoiceReportRepository{db: db}
}


func (r *VoiceReportRepository) GetCdrOutVas(vas string, year int, month int) ([]billing.CdrRecord, error){
	var result []billing.CdrRecord

	lastDay := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)
	daysInMonth := lastDay.Day()

	var wg sync.WaitGroup
	mu := sync.Mutex{}

	for day := 1; day <= daysInMonth; day++{
		wg.Add(1)
		go func(day int) {
			defer wg.Done()
			tableName := fmt.Sprintf("cdr%04d%02d%02d", year, month, day)
			var cdrRecords []billing.CdrRecord

			// Thực hiện truy vấn
			err := global.VoiceReport.DB.
					Table(tableName).
					Select("caller", "callee", "time", "duration","minute","cost", "callee_gw").
					Where("Callee LIKE ?", vas+"%").
					Or("Callee LIKE ?", "concat(84,"+vas+")%").
					Where("call_type LIKE ?", "OUT_VAS").
					Find(&cdrRecords).Error

			if err != nil {
				// Log lỗi nếu có nhưng không làm gián đoạn toàn bộ quá trình
				log.Printf("Error querying table %s: %v", tableName, err)
				return
			}
			// Cập nhật kết quả vào allRecords một cách thread-safe
			mu.Lock()
			result = append(result, cdrRecords...)
			mu.Unlock()
		}(day)
	}
	wg.Wait()
	return result, nil
}

func (r *VoiceReportRepository) GetCdrInVas(vas string, year int, month int) ([]billing.CdrRecord, error){
	var result []billing.CdrRecord

	tableName := fmt.Sprintf("cdrdvgtgt%04d%02d", year, month)
	
	err := global.VoiceReport.DB.
			Table(tableName).
			Select("caller", "callee", "time", "duration","minute","cost","caller_object", "caller_gw").
			Where("Callee LIKE ?", vas+"%").
			Or("categories_code LIKE ?", vas).
			Find(&result).Error

	if err != nil {
		// Log lỗi nếu có nhưng không làm gián đoạn toàn bộ quá trình
		log.Printf("Error querying table %s: %v", tableName, err)
		return nil, err
	}

	return result, nil
}