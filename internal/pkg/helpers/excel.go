package helpers

import (
	"fmt"
	"log"
	"github.com/xuri/excelize/v2"
)

func Export_data_to_excel(fileName string, call_type string,data [][]string) (string){
	f := excelize.NewFile()
	sheetName := "sheet1"
	
	index,_ := f.NewSheet(sheetName)

	f.SetCellValue(sheetName, "A1", "Caller")
	f.SetCellValue(sheetName, "B1", "Callee")
	f.SetCellValue(sheetName, "C1", "Time")
	f.SetCellValue(sheetName, "D1", "Duration")
	f.SetCellValue(sheetName, "E1", "Minute")
	f.SetCellValue(sheetName, "F1", "Cost")
	if call_type == "OUT" {
		f.SetCellValue(sheetName, "G1", "Callee GateWay")
	}else if call_type == "IN" {
		f.SetCellValue(sheetName, "G1", "Caller GateWay")
	}
	for i, row := range data {
		for j, cell := range row {
			cellName := fmt.Sprintf("%c%d", 'A'+j, i+2) // Bắt đầu từ hàng 2
			f.SetCellValue(sheetName, cellName, cell)
		}
	}
	f.AutoFilter(sheetName,"A1:G1",[]excelize.AutoFilterOptions{})
	f.SetActiveSheet(index)
	filepath := "/root/mini_billing/storages/assets/" + fileName;

	if err := f.SaveAs(filepath); err != nil {
		log.Fatalf("Lỗi khi lưu file Excel: %v", err)
	} else {
		fmt.Println("Done, được lưu tại "+ filepath)
	}
	return filepath
}
