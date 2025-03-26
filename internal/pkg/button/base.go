package button

import (
	tele "gopkg.in/telebot.v4"
)

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
// button Cdr
var Cdr_InlineKeys = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{BtnFixed,BtnVAS},
		{BtnMBS,BtnContract},
	},
}
