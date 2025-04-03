package usecase

import (
	"fmt"
	"miniBilling/internal/pkg/helpers"
	"miniBilling/internal/po/billing"
	"miniBilling/internal/repository"
)

var (
	cdrRecords []billing.CdrRecord
	err        error
	cdrData    [][]string
)

type VoiceReport interface{
	CdrOUTVas( telco string, services string , time string) (string,string)
	CdrINVas(telco string, services string, time string) (string,string)
	CdrSIP(time string) (string, string)
	Report3BigCustomer(time string) (string, string)
}

type VoiceReportUsecase struct {
	report repository.VoiceReport
}

func NewVoiceReportUsecase(voiceReport repository.VoiceReport) VoiceReport {
	return &VoiceReportUsecase{report: voiceReport}
}

func (uc *VoiceReportUsecase) CdrOUTVas( telco string, services string, time string) (string,string)  {
	// time := chuyển đổi từ string sang month year
		// Parse tháng/năm
		month, year, _ := helpers.ParseMonthYear(time)
	
		// Lấy dữ liệu theo dịch vụ
		if services == "1900" || services == "1800" {
			cdrRecords, err = uc.report.GetCdrOutVas(telco,services, year, month)
		} else {
			cdrRecords, err = uc.report.GetCdrOutAllVas(telco, year, month)
		}
	
		if err != nil {
			fmt.Println("❌ Lỗi khi lấy dữ liệu CDR:", err)
			return "",""
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
				fmt.Sprintf("%f", record.Cost),
				record.CallType,
				record.CalleeGw,
			}
			cdrData = append(cdrData, row)
		}
	
		fileName := fmt.Sprintf("CTC_Digitel_call_%s_%s_%d_%d.xlsx",services,telco,month,year)
		fileResult := helpers.Export_data_to_excel(fileName,"OUT",cdrData)

		return fileResult,fileName
}

func (uc *VoiceReportUsecase) CdrINVas(telco string, services string, time string) (string,string) {
	// Parse tháng/năm
	month, year, _ := helpers.ParseMonthYear(time)

	// Lấy dữ liệu theo dịch vụ
	if services == "1900" || services == "1800" {
		fmt.Println("ok cdr dịch vụ: ",services )
		cdrRecords, err = uc.report.GetCdrInVas(telco,services, year, month)
	} else {
		cdrRecords, err = uc.report.GetCdrInAllVas(telco,year, month)
	}

	if err != nil {
		fmt.Println("❌ Lỗi khi lấy dữ liệu CDR:", err)
		return "",""
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
			record.CallType,
			record.CallerGw,
		}
		cdrData = append(cdrData, row)
	}

	fileName := fmt.Sprintf("CTC_Digitel_call_%s_%s_%d_%d.xlsx",services,telco,month,year)
	fileResult := helpers.Export_data_to_excel(fileName,"IN",cdrData)

	return fileResult,fileName
}

func (uc *VoiceReportUsecase) CdrSIP(time string) (string, string){

	month, year, _ := helpers.ParseMonthYear(time)
	cdrRecords, err = uc.report.GetCdrtSIP(year, month)

	if err != nil {
		fmt.Println("❌ Lỗi khi lấy dữ liệu CDR:", err)
		return "",""
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
			record.CallType,
			record.CallerObject,
			record.CalleeObject,
			record.CallerGw,
			record.CalleeGw,
		}
		cdrData = append(cdrData, row)
	}
	fileName := fmt.Sprintf("CTC_MBS_%d_%d.xlsx",month,year)
	fileResult := helpers.Export_data_to_excel(fileName,"ALL",cdrData)

	return fileResult,fileName
}

func (uc *VoiceReportUsecase) Report3BigCustomer(time string) (string, string){
	fmt.Println("hehe2")
	month, year, _ := helpers.ParseMonthYear(time)
	report, err := uc.report.GetReport3BigCustomer(year,month)
	fmt.Println("sao ?")
	if err != nil {
		fmt.Println("❌ Lỗi khi lấy dữ liệu CDR:", err)
		return "",""
	}
	 fmt.Println(report)
	 return "nil","nil"
}