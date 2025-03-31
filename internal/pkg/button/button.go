package button

import (
	"time"
	"fmt"
	tele "gopkg.in/telebot.v4"
)
//==============================================================
func DynamicButton(text string,data string) tele.InlineButton {
	return tele.InlineButton{
		Unique: data,
		Text:   text,
	}
}
//============= BUTTON MONTH ===================================
func GetMonthOffsetSafe(offset int) string {
	// Luôn dùng ngày 15 để tránh lỗi lùi tháng như trên
	base := time.Date(time.Now().Year(), time.Now().Month(), 15, 0, 0, 0, 0, time.Local)
	t := base.AddDate(0, offset, 0)
	return fmt.Sprintf("%02d/%d", t.Month(), t.Year())
}




var(
	BtnIntro = tele.InlineButton{
		Unique: "btn_intro",
		Text:   "Giới thiệu",
		Data:   "button_intro",
	}

	BtnLogin = tele.InlineButton{
		Unique: "btn_login",
		Text:   "🔑 Đăng nhập",
		Data:   "login",
	}

	BtnLoginCode = tele.InlineButton{
		Unique: "btn_login",
		Text:   "🔑 Lẫy mã Đăng nhập",
		Data:   "loginCode",
	}
	//=================== BUTTON CDR ===================================
	BtnCdr = tele.InlineButton{
		Unique: "btn_cdr",
		Text:   "Chi tiết cước",
		Data:   "Cdr",
	}

	BtnFixed = tele.InlineButton{
		Unique: "btn_cdr",
		Text:   "CTC cố định",
		Data:   "CdrFixed",
	}

	Btn1900VAS = tele.InlineButton{
		Unique: "btn_cdr",
		Text:   "Chi tiết cước 1900",
		Data:   "Cdr1900",
	}

	Btn1800VAS = tele.InlineButton{
		Unique: "btn_cdr",
		Text:   "Chi tiết cước 1800",
		Data:   "Cdr1800",
	}

	BtnMBS = tele.InlineButton{
		Unique: "btn_cdr",
		Text:   "CTC Mobile SIP",
		Data:   "CdrSIP",
	}
	BtnCdrNumber = tele.InlineButton{
		Unique: "btn_cdr",
		Text:   "CTC theo đầu số",
		Data:   "Number",
	}

	BtnContract = tele.InlineButton{
		Unique: "btn_cdr",
		Text:   "CTC theo Hợp đồng",
		Data:   "cdrContract",
	}
	
	BtnIN = tele.InlineButton{
		Unique: "btn_CallIn",
		Text:   "Lấy Call in",
		Data:   "CdrCallIn",
	}

	BtnOUT = tele.InlineButton{
		Unique: "btn_CallOut",
		Text:   "Lấy Call out",
		Data:   "CdrCallOUT",
	}
	
	// ================== TELCO ========================================
	BtnViettel = tele.InlineButton{
		Text:   "VIETTEL",
		Unique:   "VIETTEL",
	}

	BtnViettel_FIXED = tele.InlineButton{
		Text:   "VIETTEL FIXED",
		Unique:   "VIETTEL_FIXED",
	}

	BtnVnpt = tele.InlineButton{
		Text:   "VNPT FIXED",
		Unique:   "VNPT",
	}
	BtnGPC = tele.InlineButton{
		Text:   "GPC",
		Unique:   "GPC",
	}
	BtnHTC = tele.InlineButton{
		Text:   "HTC",
		Unique:   "HTC",
	}

	BtnFPT = tele.InlineButton{
		Text:   "FPT",
		Unique:   "FPT",
	}

	BtnITEL = tele.InlineButton{
		Text:   "ITEL",
		Unique:   "ITEL",
	}

	BtnMBF = tele.InlineButton{
		Text:   "VMS",
		Unique:   "VMS",
	}

	BtnVNM = tele.InlineButton{
		Text:   "VIETNAMOBILE",
		Unique:   "VNM",
	}

	BtnCMC = tele.InlineButton{
		Text:   "CMC",
		Unique:   "CMC",
	}

	BtnMBC = tele.InlineButton{
		Text:   "MOBICAST",
		Unique:   "MOBICAST",
	}

	BtnGTEL = tele.InlineButton{
		Text:   "GTEL",
		Unique:   "GTEL",
	}

	BtnAll = tele.InlineButton{
		Text:   "Tất cả nhà mạng",
		Unique:   "ALL",
	}
	
	unique_this_month = DynamicButton(GetMonthOffsetSafe(0),GetMonthOffsetSafe(0))
	unique_last_month = DynamicButton(GetMonthOffsetSafe(-1),GetMonthOffsetSafe(-1))
	unique_last_2_month = DynamicButton(GetMonthOffsetSafe(-2),GetMonthOffsetSafe(-2))

	// ================== BUTTON REPORT ================================
	BtnReport= tele.InlineButton{
		Unique: "btn_report",
		Text:   "Báo cáo sản lượng",
		Data:   "ReportQuantity",
	}
) 

