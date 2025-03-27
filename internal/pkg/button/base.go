package button

import (
	
	tele "gopkg.in/telebot.v4"
)


//=======================================================================
func GetMainMenu(role string) *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}

	switch role {
	case "super-admin":
		menu.InlineKeyboard = [][]tele.InlineButton{
			{BtnIntro,BtnLoginCode},
			{BtnCdr,BtnReport},
		}
	case "admin": //admin
		menu.InlineKeyboard = [][]tele.InlineButton{
			{BtnIntro,BtnLoginCode}, 
		}
	case "Sale": //sale
		menu.InlineKeyboard = [][]tele.InlineButton{
			{BtnIntro,BtnLoginCode}, 
		}
	case "Accountant": // kế toán
		menu.InlineKeyboard = [][]tele.InlineButton{
			{BtnIntro,BtnLoginCode}, 
		}
	case "Technician": // Kỹ thuật
		menu.InlineKeyboard = [][]tele.InlineButton{
			{BtnIntro,BtnLoginCode}, 
		}
	default:
		menu.InlineKeyboard = [][]tele.InlineButton{
			{BtnIntro,BtnLoginCode}, 
		}
	}
	return menu
}


// Button Login
var Login_InlineKeys = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{BtnLogin},
	},
}
// button Cdr // đang lỗi
var Cdr_InlineKeys = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{BtnFixed, Btn1900VAS, Btn1800VAS},
		{BtnMBS, BtnContract, BtnCdrNumber},
	},
}

var CDR_CallType = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{BtnIN, BtnOUT},
	},
}

func GetTelcoByServices(services string) *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}
	// Tuỳ từng dịch vụ, chỉ cho phép lấy 1 vài telco
	switch services{
		//Nhiều quá chưa cho lấy cố định, sau này yêu cầu thì xử lý sau
		// case "FIXED": 
		// 	menu.InlineKeyboard = [][]tele.InlineButton{
		// 		{BtnIntro,BtnLoginCode},
		// 	}
		case "VAS":
			menu.InlineKeyboard = [][]tele.InlineButton{
				{BtnAll,BtnViettel,BtnCMC},
				{BtnVnpt,BtnCMC,BtnFPT},
				{BtnMBC,BtnGTEL,BtnHTC},
			}
		// SIP đang rất ít nên  xuất ALL
		// case "SIP":
		// 	menu.InlineKeyboard = [][]tele.InlineButton{
		// 		{BtnIntro,BtnLoginCode},
		// 		{BtnCdr,BtnReport},
		// 	}
		default:
			menu.InlineKeyboard = [][]tele.InlineButton{
				{BtnIntro,BtnLoginCode},
				{BtnCdr,BtnReport},
			}
		
	}
	
	return menu
}

var BtnMonth = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{unique_this_month, unique_last_month, unique_last_2_month},
		// {Back_To_Main_Menu},
	},
}