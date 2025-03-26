package handler

import (
	"miniBilling/internal/pkg/button"
	"miniBilling/internal/usecase"

	"go.mongodb.org/mongo-driver/bson"
	tele "gopkg.in/telebot.v4"
)

type Report136HanderInterface interface {
	Cdr(c tele.Context) error 
	Cdr_category_code(c tele.Context, callback string) error 
	CdrCallType (c tele.Context, callback string) error 
}

type Report136handler struct {
	VoicerReport 	usecase.VoiceReport
	UsersUC 		usecase.Users
	Bot 			*tele.Bot
}

func NewVoiceReportHandler(voiceReport usecase.VoiceReport, userMongo usecase.Users,bot *tele.Bot) Report136HanderInterface {
	return &Report136handler{VoicerReport: voiceReport, UsersUC: userMongo, Bot: bot}
}

func (h *Report136handler) Cdr(c tele.Context) error {
	user := c.Sender()
	teleId := user.ID

	h.UsersUC.UpdateUserMongo(teleId,bson.M{
		"action1": 	"CDR",
		"action2": 	nil,
		"action3": 	nil,
		"action4": 	nil,
		"action5": 	nil,
		"action6": 	nil,
	})
	return c.Send("Bạn muốn lấy CDR dịch vụ nào?", button.Cdr_InlineKeys)

}

func (h *Report136handler) Cdr_category_code(c tele.Context, callback string) error {
	var action2 string
	user := c.Sender()
	teleId := user.ID
	if callback == "btn_cdr|CdrFixed" {
		action2 = "FIXED"
	}else if callback == "btn_cdr|CdrVas" {
		action2 = "VAS"
	}else if callback == "btn_cdr|CdrSIP" {
		action2 = "SIP"
	}else if callback == "btn_cdr|cdrContract" {
		action2 = "CONTRACT"
	}else{
		return	c.Send("Vui lòng /start để chọn lại menu")
	}
	h.UsersUC.UpdateUserMongo(teleId,bson.M{
		"action1": 	"CDR",
		"action2": 	action2,
		"action3": 	nil,
		"action4": 	nil,
		"action5": 	nil,
		"action6": 	nil,
	})

	return	c.Edit("Tiếp theo hãy chọn call IN hay call OUT",button.CDR_CallType)
}

func (h *Report136handler) CdrCallType (c tele.Context, callback string) error {
	var action3 string
	var message string
	user := c.Sender()
	teleId := user.ID

	if callback == "btn_CallIn|CdrCallIn" {
		action3 = "IN"
		message = "Tiếp theo hãy chọn nhà mạng gọi vào Digitel"
	}else {
		action3 = "OUT"
		message = "Tiếp theo hãy chọn nhà mạng mà Digitel gọi tới"

	}
	h.UsersUC.UpdateUserMongo(teleId,bson.M{
		"action3": 	action3,
		"action4": 	nil,
		"action5": 	nil,
		"action6": 	nil,
	})
	// đoạn này xử lý từng nhà mạng theo rule đi, vì kp nhà mạng nào cũng cho lấy
	return	c.Edit(message)

}

