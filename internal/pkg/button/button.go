package button

import (
	tele "gopkg.in/telebot.v4"
)

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
		Unique: "btn_fixed",
		Text:   "Chi tiết cước cố định",
		Data:   "CdrFixed",
	}

	BtnVAS = tele.InlineButton{
		Unique: "btn_vas",
		Text:   "Chi tiết cước GTGT",
		Data:   "CdrVas",
	}

	BtnMBS = tele.InlineButton{
		Unique: "btn_sip",
		Text:   "Chi tiết cước Mobile SIP",
		Data:   "CdrSIP",
	}

	BtnContract = tele.InlineButton{
		Unique: "btn_contract",
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
		Data:   "VIETTEL",
	}

	BtnViettel_FIXED = tele.InlineButton{
		Text:   "VIETTEL FIXED",
		Data:   "VIETTEL_FIXED",
	}

	BtnVnpt = tele.InlineButton{
		Text:   "VNPT FIXED",
		Data:   "VNPT",
	}
	BtnGPC = tele.InlineButton{
		Text:   "GPC",
		Data:   "GPC",
	}

	BtnFPT = tele.InlineButton{
		Text:   "FPT FIXED",
		Data:   "FPT",
	}

	BtnITEL = tele.InlineButton{
		Text:   "ITEL",
		Data:   "ITEL",
	}

	BtnMBF = tele.InlineButton{
		Text:   "VMS",
		Data:   "VMS",
	}

	BtnVNM = tele.InlineButton{
		Text:   "VIETNAMOBILE",
		Data:   "VNM",
	}

	BtnCMC = tele.InlineButton{
		Text:   "CMC",
		Data:   "CMC",
	}

	BtnMBC = tele.InlineButton{
		Text:   "MOBICAST",
		Data:   "MOBICAST",
	}

	BtnGTEL = tele.InlineButton{
		Text:   "GTEL FIXED",
		Data:   "MOBICAST",
	}





	

	// ================== BUTTON REPORT ================================
	BtnReport= tele.InlineButton{
		Unique: "btn_report",
		Text:   "Báo cáo sản lượng",
		Data:   "ReportQuantity",
	}
) 

