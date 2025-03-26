package handler

import (
	"miniBilling/internal/pkg/button"
	"miniBilling/internal/usecase"

	"go.mongodb.org/mongo-driver/bson"
	tele "gopkg.in/telebot.v4"
)

type Report136HanderInterface interface {
	Cdr(c tele.Context) error 
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