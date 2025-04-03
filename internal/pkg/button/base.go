package button

import (
	tele "gopkg.in/telebot.v4"
)

// =======================================================================
func GetMainMenu(role string) *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}

	switch role {
	case "super-admin":
		menu.InlineKeyboard = [][]tele.InlineButton{
			{BtnIntro, BtnLoginCode},
			{BtnCdr, BtnReport},
		}
	case "admin": //admin
		menu.InlineKeyboard = [][]tele.InlineButton{
			{BtnIntro, BtnLoginCode},
		}
	case "Sale": //sale
		menu.InlineKeyboard = [][]tele.InlineButton{
			{BtnIntro, BtnLoginCode},
		}
	case "Auditor": // kế toán
		menu.InlineKeyboard = [][]tele.InlineButton{
			{BtnIntro, BtnLoginCode},
			{BtnCdr, BtnReport},
		}
	case "Technician": // Kỹ thuật
		menu.InlineKeyboard = [][]tele.InlineButton{
			{BtnIntro, BtnLoginCode},
			{BtnCdr, BtnReport},
		}
	default:
		menu.InlineKeyboard = [][]tele.InlineButton{
			{BtnIntro, BtnLoginCode},
		}
	}
	return menu
}


func GetTelcoByServices(services string) *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}
	// Tuỳ từng dịch vụ, chỉ cho phép lấy 1 vài telco
	switch services {
	//Nhiều quá chưa cho lấy cố định, sau này yêu cầu thì xử lý sau
	// case "FIXED":
	// 	menu.InlineKeyboard = [][]tele.InlineButton{
	// 		{BtnIntro,BtnLoginCode},
	// 	}
	case "1900":
		menu.InlineKeyboard = [][]tele.InlineButton{
			{BtnAll, BtnViettel, BtnCMC},
			{BtnVnpt, BtnCMC, BtnFPT},
			{BtnMBC, BtnGTEL, BtnHTC},
		}
	case "1800":
		menu.InlineKeyboard = [][]tele.InlineButton{
			{BtnAll, BtnViettel, BtnCMC},
			{BtnVnpt, BtnCMC, BtnFPT},
			{BtnMBC, BtnGTEL, BtnHTC},
		}
	// SIP đang rất ít nên  xuất ALL
	case "SIP":
		menu.InlineKeyboard = [][]tele.InlineButton{
			{},
		}
	default:
		menu.InlineKeyboard = [][]tele.InlineButton{
			{BtnIntro, BtnLoginCode},
			{BtnCdr, BtnReport},
		}

	}

	return menu
}


var(
	// Button Login
	Login_InlineKeys = &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{BtnLogin},
		},
	}

	// button Cdr // đang lỗi
	Cdr_InlineKeys = &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{BtnFixed, Btn1900VAS, Btn1800VAS},
			{BtnMBS, BtnContract, BtnCdrNumber},
		},
	}

	CDR_CallType = &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{BtnIN, BtnOUT},
		},
	}

	BtnMonth = &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{unique_this_month, unique_last_month, unique_last_2_month},
			// {Back_To_Main_Menu},
		},
	}

	BtnReportMenu = &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{BtnReport_3_BigCus,BtnReportRecord},
			{BtnReportTelco,BtnReportContractPostPaid},
		},
	}
)