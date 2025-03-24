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
) 

