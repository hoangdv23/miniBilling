package usecase

import (
	"fmt"
	"miniBilling/internal/pkg/helpers"
	"miniBilling/internal/po/billing"
	"miniBilling/internal/repository"
)



type VoiceReport interface{
	CdrOUTVas( telco string, services string , time string) (string, string)
	CdrINVas(telco string, services string, time string) (string, string)
}

type VoiceReportUsecase struct {
	report repository.VoiceReport
}

func NewVoiceReportUsecase(voiceReport repository.VoiceReport) VoiceReport {
	return &VoiceReportUsecase{report: voiceReport}
}

func (uc *VoiceReportUsecase) CdrOUTVas( telco string, services string, time string) (string, string)  {
	// time := chuyển đổi từ string sang month year
		// Parse tháng/năm
		month, year, _ := helpers.ParseMonthYear(time)

		var (
			cdrRecords []billing.CdrRecord
			err        error
			cdrData    [][]string
		)
	
		// Lấy dữ liệu theo dịch vụ
		if services == "1900" || services == "1800" {
			cdrRecords, err = uc.report.GetCdrOutVas(telco,services, year, month)
		} else {
			cdrRecords, err = uc.report.GetCdrOutAllVas(telco, year, month)
		}
	
		if err != nil {
			fmt.Println("❌ Lỗi khi lấy dữ liệu CDR:", err)
			return "", "error"
		}
	
		// Xử lý dữ liệu
		for _, record := range cdrRecords {
			timeFormatted := record.Time.Format("2006-01-02 15:04:05")
			row := []string{
				record.Caller,
				record.Callee,
				timeFormatted,
				fmt.Sprintf("%d", record.Duration),
				fmt.Sprintf("%d", record.Minute),
				fmt.Sprintf("%.2f", record.Cost),
				record.CallerGw,
			}
			cdrData = append(cdrData, row)
		}
	
		// Debug kết quả
		fmt.Println("✅ Dữ liệu CDR:", cdrData)
	
		return "ok OUT Vas", ""
}

func (uc *VoiceReportUsecase) CdrINVas(telco string, services string, time string) (string, string) {
	// Parse tháng/năm
	month, year, _ := helpers.ParseMonthYear(time)
	var (
		cdrRecords []billing.CdrRecord
		err        error
		cdrData    [][]string
	)

	// Lấy dữ liệu theo dịch vụ
	if services == "1900" || services == "1800" {
		fmt.Println("ok cdr dịch vụ: ",services )
		cdrRecords, err = uc.report.GetCdrInVas(telco,services, year, month)
	} else {
		cdrRecords, err = uc.report.GetCdrInAllVas(telco,year, month)
	}

	if err != nil {
		fmt.Println("❌ Lỗi khi lấy dữ liệu CDR:", err)
		return "", "error"
	}

	// Xử lý dữ liệu
	for _, record := range cdrRecords {
		timeFormatted := record.Time.Format("2006-01-02 15:04:05")
		row := []string{
			record.Caller,
			record.Callee,
			timeFormatted,
			fmt.Sprintf("%d", record.Duration),
			fmt.Sprintf("%d", record.Minute),
			fmt.Sprintf("%.2f", record.Cost),
			record.CallerGw,
		}
		cdrData = append(cdrData, row)
	}

	// Debug kết quả
	fmt.Println("✅ Dữ liệu CDR:", cdrData)

	return "ok IN Vas", ""
}
