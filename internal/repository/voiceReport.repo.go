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
	GetCdrOutVas(telco string, vas string, year int, month int) ([]billing.CdrRecord, error)
	GetCdrInVas(telco string, vas string, year int, month int) ([]billing.CdrRecord, error)
	GetCdrOutAllVas(telco string, year int, month int) ([]billing.CdrRecord, error)
	GetCdrInAllVas(telco string, year int, month int) ([]billing.CdrRecord, error)
}

type VoiceReportRepository struct {
	db *gorm.DB
}

func NewVoiceReportRepository(db *gorm.DB) VoiceReport {
	return &VoiceReportRepository{db: db}
}

func (r *VoiceReportRepository) GetCdrOutVas(telco string, vas string, year int, month int) ([]billing.CdrRecord, error) {
	var result []billing.CdrRecord
	lastDay := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)
	daysInMonth := lastDay.Day()
	var wg sync.WaitGroup
	mu := sync.Mutex{}

	for day := 1; day <= daysInMonth; day++ {
		wg.Add(1)
		go func(day int) {
			defer wg.Done()
		tableName := fmt.Sprintf("cdr%04d%02d%02d", year, month, day)
		var cdrRecords []billing.CdrRecord

		query := global.VoiceReport.DB.
			Table(tableName).
			Select("caller", "callee", "time", "duration", "minute", "cost", "callee_gw").
			Where("(callee LIKE ? OR callee LIKE ?) AND call_type = ?", vas+"%", "84"+vas+"%", "OUT_VAS")

			if telco != "" {
				query = query.Where("callee_gw LIKE ?", "%"+telco+"%")
			}

		err := query.Find(&cdrRecords).Error
		if err != nil {
			log.Printf("Error querying table %s: %v", tableName, err)
			return
		}

		mu.Lock()
		result = append(result, cdrRecords...)
		mu.Unlock()
		}(day)
	}
	wg.Wait()
	return result, nil
}

func (r *VoiceReportRepository) GetCdrInVas(telco string, vas string, year int, month int) ([]billing.CdrRecord, error) {
	var result []billing.CdrRecord
	tableName := fmt.Sprintf("cdrdvgtgt%04d%02d", year, month)

	query := global.VoiceReport.DB.
		Table(tableName).
		Select("caller", "callee", "time", "duration", "minute", "cost", "caller_object", "caller_gw").
		Where("callee LIKE ? OR categories_code LIKE ?", vas+"%", vas)

	if telco != "" {
		query = query.Where("caller_gw LIKE ?", "%"+telco+"%")
	}

	err := query.Find(&result).Error
	if err != nil {
		log.Printf("Error querying table %s: %v", tableName, err)
		return nil, err
	}

	return result, nil
}

func (r *VoiceReportRepository) GetCdrOutAllVas(telco string, year int, month int) ([]billing.CdrRecord, error) {
	var result []billing.CdrRecord
	lastDay := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)
	daysInMonth := lastDay.Day()
	var wg sync.WaitGroup
	mu := sync.Mutex{}

	for day := 1; day <= daysInMonth; day++ {
		wg.Add(1)
		go func(day int) {
			defer wg.Done()
		tableName := fmt.Sprintf("cdr%04d%02d%02d", year, month, day)
		var cdrRecords []billing.CdrRecord

		query := global.VoiceReport.DB.
			Table(tableName).
			Select("caller", "callee", "time", "duration", "minute", "cost", "callee_gw").
			Where("(callee LIKE ? OR callee LIKE ? OR callee LIKE ? OR callee LIKE ?) AND call_type = ?",
				"1800%", "841800%", "1900%", "841900%", "OUT_VAS")

		if telco != "" {
			query = query.Where("callee_gw LIKE ?", "%"+telco+"%")
		}

		err := query.Find(&cdrRecords).Error
		if err != nil {
			log.Printf("Error querying table %s: %v", tableName, err)
			return
		}

		mu.Lock()
		result = append(result, cdrRecords...)
		mu.Unlock()
		}(day)
	}
	wg.Wait()
	return result, nil
}

func (r *VoiceReportRepository) GetCdrInAllVas(telco string, year int, month int) ([]billing.CdrRecord, error) {
	var result []billing.CdrRecord
	tableName := fmt.Sprintf("cdrdvgtgt%04d%02d", year, month)

	query := global.VoiceReport.DB.
		Table(tableName).
		Select("caller", "callee", "time", "duration", "minute", "cost", "caller_object", "caller_gw")

	if telco != "" {
		query = query.Where("caller_gw LIKE ?", "%"+telco+"%")
	}

	err := query.Find(&result).Error
	if err != nil {
		log.Printf("Error querying table %s: %v", tableName, err)
		return nil, err
	}

	return result, nil
}
