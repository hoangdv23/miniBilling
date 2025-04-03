package repository

import (
	"fmt"
	"log"
	"miniBilling/global"
	"miniBilling/internal/po/billing"
	"sync"
	"time"
	"strconv"

	"gorm.io/gorm"
)

type VoiceReport interface {
	GetCdrOutVas(telco string, vas string, year int, month int) ([]billing.CdrRecord, error)
	GetCdrInVas(telco string, vas string, year int, month int) ([]billing.CdrRecord, error)
	GetCdrOutAllVas(telco string, year int, month int) ([]billing.CdrRecord, error)
	GetCdrInAllVas(telco string, year int, month int) ([]billing.CdrRecord, error)
	GetCdrtSIP( year int, month int) ([]billing.CdrRecord, error)
	GetReport3BigCustomer(year int, month int) ([]map[string]interface{}, error)
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
			Select("caller", "callee", "time", "duration", "minute", "cost","call_type", "callee_gw").
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
		Select("caller", "callee", "time", "duration", "minute", "cost","call_type", "caller_object", "caller_gw").
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
		Select("caller", "callee", "time", "duration", "minute", "cost","call_type", "caller_object","callee_object", "caller_gw","caller_gw")

	err := query.Find(&result).Error
	if err != nil {
		log.Printf("Error querying table %s: %v", tableName, err)
		return nil, err
	}

	return result, nil
}

func (r *VoiceReportRepository) GetCdrtSIP( year int, month int) ([]billing.CdrRecord, error) {
	var result []billing.CdrRecord
	tableName := fmt.Sprintf("cdrdsip%04d%02d", year, month)

	query := global.VoiceReport.DB.
		Table(tableName).
		Select("caller", "callee", "time", "duration", "minute", "cost","call_type", "caller_object","callee_object", "caller_gw","callee_gw")

	err := query.Find(&result).Error
	if err != nil {
		log.Printf("Error querying table %s: %v", tableName, err)
		return nil, err
	}

	return result, nil
}

func (r *VoiceReportRepository) GetReport3BigCustomer(year int, month int) ([]map[string]interface{}, error) {
	lastDay := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)
	daysInMonth := lastDay.Day()

	var wg sync.WaitGroup
	mu := sync.Mutex{}

	// Kết quả tổng theo customer_code
	total := make(map[string]map[string]interface{})

	for day := 1; day <= daysInMonth; day++ {
		wg.Add(1)
		go func(day int) {
			defer wg.Done()

			tableName := fmt.Sprintf("cdr%04d%02d%02d", year, month, day)

			rows, err := global.VoiceReport.DB.Raw(`
				SELECT 
					customer_code,
					SUM(CASE WHEN caller_gw LIKE '%vietnamobile%' THEN 1 ELSE 0 END) AS vietnamobile_count,
					SUM(CASE WHEN caller_gw LIKE '%vietnamobile%' THEN duration ELSE 0 END) AS vietnamobile_duration,
					SUM(CASE WHEN caller_gw LIKE '%vietnamobile%' THEN minute ELSE 0 END) AS vietnamobile_minute,

					SUM(CASE WHEN caller_gw LIKE '%vnpt%' THEN 1 ELSE 0 END) AS vnpt_count,
					SUM(CASE WHEN caller_gw LIKE '%vnpt%' THEN duration ELSE 0 END) AS vnpt_duration,
					SUM(CASE WHEN caller_gw LIKE '%vnpt%' THEN minute ELSE 0 END) AS vnpt_minute,

					SUM(CASE WHEN caller_gw LIKE '%viettel%' THEN 1 ELSE 0 END) AS viettel_count,
					SUM(CASE WHEN caller_gw LIKE '%viettel%' THEN duration ELSE 0 END) AS viettel_duration,
					SUM(CASE WHEN caller_gw LIKE '%viettel%' THEN minute ELSE 0 END) AS viettel_minute,

					SUM(CASE WHEN caller_gw LIKE '%mobifone%' THEN 1 ELSE 0 END) AS mobifone_count,
					SUM(CASE WHEN caller_gw LIKE '%mobifone%' THEN duration ELSE 0 END) AS mobifone_duration,
					SUM(CASE WHEN caller_gw LIKE '%mobifone%' THEN minute ELSE 0 END) AS mobifone_minute
				FROM `+"`"+tableName+"`"+`
				WHERE customer_code IN ('DG00080', 'DG00095', 'DG00155') 
					AND call_type = 'IN_MOBILE'
					AND provider IN (888,555)
				GROUP BY customer_code
			`).Rows()

			if err != nil {
				log.Printf("❌ Error querying table %s: %v", tableName, err)
				return
			}
			defer rows.Close()

			columns, _ := rows.Columns()
			for rows.Next() {
				cols := make([]interface{}, len(columns))
				colPtrs := make([]interface{}, len(columns))
				for i := range cols {
					colPtrs[i] = &cols[i]
				}

				_ = rows.Scan(colPtrs...)

				row := make(map[string]interface{})
				for i, colName := range columns {
					val := cols[i]
					if b, ok := val.([]byte); ok {
						row[colName] = string(b)
					} else {
						row[colName] = val
					}
				}

				customerCode := row["customer_code"].(string)

				mu.Lock()
				if _, ok := total[customerCode]; !ok {
					total[customerCode] = make(map[string]interface{})
					total[customerCode]["customer_code"] = customerCode
				}

				// Cộng dồn các field
				for key, val := range row {
					if key == "customer_code" {
						continue
					}
					// Nếu chưa có key này thì khởi tạo bằng 0
					if _, ok := total[customerCode][key]; !ok {
						total[customerCode][key] = 0.0
					}
					// Convert về float64 để cộng dồn
					current := total[customerCode][key].(float64)
					value := toFloat(val)
					total[customerCode][key] = current + value
				}
				mu.Unlock()
			}
		}(day)
	}

	wg.Wait()

	// Chuyển map tổng thành slice kết quả
	var result []map[string]interface{}
	for _, v := range total {
		result = append(result, v)
	}
	return result, nil
}

// Chuyển mọi loại value thành float64 để cộng dồn
func toFloat(val interface{}) float64 {
	switch v := val.(type) {
	case int64:
		return float64(v)
	case int:
		return float64(v)
	case float64:
		return v
	case []byte:
		f, _ := strconv.ParseFloat(string(v), 64)
		return f
	default:
		return 0
	}
}
