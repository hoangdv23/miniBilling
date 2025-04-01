package helpers

import (
	"fmt"
	"log"
	"github.com/xuri/excelize/v2"
	  _ "image/png"
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

	f.SetColWidth(sheetName, "A", "B", 14)
	f.SetColWidth(sheetName, "C", "C", 22)
	f.SetColWidth(sheetName, "D", "F", 9)
	f.SetColWidth(sheetName, "G", "G", 23)
	// Tạo style với màu nền
	headerStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#00bfff"}, // Màu nền
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold: true,
			Color: "#000000",
			Size: 12,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		fmt.Println("❌ Lỗi tạo style:", err)
	}
	_ = f.SetCellStyle(sheetName, "A1", "g1", headerStyle)

	style, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{"ccffff"}, Pattern: 1},
	})
	if err != nil {
		fmt.Println(err)
	}
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
		if i%2 == 1 {
			startCell := fmt.Sprintf("A%d", i+2)
			endCol := 'A' + len(row) - 1
			endCell := fmt.Sprintf("%c%d", endCol, i+2)
			_ = f.SetCellStyle(sheetName, startCell, endCell, style)
		}
	}
		lastRow := len(data) + 2 // Chừa vài dòng cách sau nội dung
		imageCell := fmt.Sprintf("A%d", lastRow)

		if err := f.AddPicture(sheetName, imageCell, "/root/mini_billing/storages/imgs/logo_digitel.png", &excelize.GraphicOptions{
			OffsetX:   	0,
            OffsetY:   	2,
			ScaleX : 	1.2,
			ScaleY :	0.79, 
		}); err != nil {
			fmt.Println("❌ Lỗi khi chèn ảnh:", err)
		}

		footerRow := lastRow + 1
		footerCell := fmt.Sprintf("A%d", footerRow)
		endMerge := fmt.Sprintf("G%d", footerRow)
		f.MergeCell(sheetName,footerCell,endMerge)

		_ = f.MergeCell(sheetName, footerCell, endMerge)
		footerText := `                                                          		CÔNG TY CỔ PHẦN HẠ TẦNG VIỄN THÔNG SỐ (DIGITEL)
                                                          					Địa chỉ giao dịch: Số OF03-19 toà OF Vinhomes West Point, Phạm Hùng, Phường Mễ Trì,
                                                          					Quận Nam Từ Liêm, Thành phố Hà Nội.
                                                          					Tel: (024-028) 8888 1111 | 1900999990 | http://digitelgroup.vn
                                                          					Email: admin@digitel.org.vn`

		_ = f.SetCellValue(sheetName, footerCell, footerText)
		_ = f.SetRowHeight(sheetName, footerRow, 60)
		f.AutoFilter(sheetName,"A1:G1",[]excelize.AutoFilterOptions{})
		f.SetActiveSheet(index)
		filepath := "/root/mini_billing/storages/assets/" + fileName + ".xlsx";

	if err := f.SaveAs(filepath); err != nil {
		log.Fatalf("Lỗi khi lưu file Excel: %v", err)
	} else {
		fmt.Println("Done, được lưu tại "+ filepath)
	}
	return filepath
}
