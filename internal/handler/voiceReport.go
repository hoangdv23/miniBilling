package handler

import (
	"fmt"
	"miniBilling/internal/pkg/button"
	"miniBilling/internal/usecase"
	"go.mongodb.org/mongo-driver/bson"
	tele "gopkg.in/telebot.v4"
)

type Report136HanderInterface interface {
	Cdr(c tele.Context) error 
	Cdr_category_code(c tele.Context, callback string) error 
	CdrCallType(c tele.Context, callback string) error 
	CdrTelco(c tele.Context, callback string) error
	CdrMonth(c tele.Context, callback string) error
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
	fmt.Println(callback)
	 if callback == "btn_cdr|CdrVas" {
		action2 = "VAS"
	}else if callback == "btn_cdr|Cdr1900" {
		action2 = "1900"
	}else if callback == "btn_cdr|Cdr1800" {
		action2 = "1800"
	}else if callback == "btn_cdr|cdrContract" {
		action2 = "CONTRACT" 
	}else if callback == "btn_cdr|Number" {
		action2 = "NUMBER" 
		c.Send("Vui lòng chọn Call OUT hoặc Call IN", button.CDR_CallType)
	}
	h.UsersUC.UpdateUserMongo(teleId,bson.M{
		"action1": 	"CDR",
		"action2": 	action2,
		"action3": 	nil,
		"action4": 	nil,
		"action5": 	nil,
		"action6": 	nil,
	})

	return	c.Edit("Tiếp theo hãy chọn Nhà Mạng",button.GetTelcoByServices(action2))
}

func (h *Report136handler) CdrTelco(c tele.Context, callback string) error {
	var message string
	user := c.Sender()
	teleId := user.ID
	// fmt.Println(callback)
	h.UsersUC.UpdateUserMongo(teleId,bson.M{
		"action3": 	callback,
		"action4": 	nil,
		"action5": 	nil,
		"action6": 	nil,
	})
	if callback == "ALL"{
		message = "Bạn muốn lấy Call IN hay Call OUT"
	}else{
		message = fmt.Sprintf("Ok!! Tôi sẽ lấy theo nhà mạng %s. Bạn muốn lấy CTC gọi vào Digitel hay Digitel gọi ra ?",callback)
	}
	return c.Send(message,button.CDR_CallType)
}

func (h *Report136handler) CdrCallType(c tele.Context, callback string) error {
	var action4 string
	var message string
	user := c.Sender()
	teleId := user.ID
	userMongo,_ := h.UsersUC.UserMongo(user.ID)

	telco := *userMongo.Action3
	services := *userMongo.Action2
	if callback == "btn_CallIn|CdrCallIn" {
		action4 = "IN"
		message = fmt.Sprintf("Bạn đã chọn %s gọi vào %s DIGITEL. Hãy chọn tháng cần lấy",telco,services)
	}else {
		action4 = "OUT"
		message = fmt.Sprintf("Bạn đã chọn DIGITEL gọi tới %s %s. Hãy chọn tháng cần lấy",services,telco)
	}
	h.UsersUC.UpdateUserMongo(teleId,bson.M{
		"action4": 	action4,
		"action5": 	nil,
		"action6": 	nil,
	})
	return	c.Edit(message,button.BtnMonth)
}
// xử lý và xuất excel đoạn này rồi đấy
func (h *Report136handler) CdrMonth(c tele.Context, callback string) error {
	fmt.Println("✅ Đã vào CdrMonth")

	user := c.Sender()
	userMongo, _ := h.UsersUC.UserMongo(user.ID)

	if userMongo.Action2 == nil || userMongo.Action3 == nil || userMongo.Action4 == nil {
		return c.Send("❌ Thiếu thông tin Action trong hệ thống, vui lòng thao tác lại từ đầu.")
	}

	services := *userMongo.Action2 // VAS: 1800/1900
	telco := *userMongo.Action3    // Nhà mạng
	callType := *userMongo.Action4 // IN / OUT

	var (
		fileName string
		text     string
	)

	if services == "1800" || services == "1900" {
		switch callType {
		case "OUT":
			fileName = h.VoicerReport.CdrOUTVas(telco, services, callback)
			text = fmt.Sprintf("📄 Bot gửi file CTC Digitel gọi %s %s tháng %s", services, telco, callback)
		case "IN":
			fileName = h.VoicerReport.CdrINVas(telco, services, callback)
			text = fmt.Sprintf("📄 Bot gửi file CTC %s %s gọi vào Digitel tháng %s", services, telco, callback)
		default:
			return c.Send("⚠️ Kiểu gọi không hợp lệ (phải là IN hoặc OUT).")
		}
	} else {
		return c.Send("⚠️ Chỉ hỗ trợ dịch vụ 1800 hoặc 1900.")
	}

	if fileName == "" {
		return c.Send("❌ Không thể tạo file, vui lòng thử lại.")
	}
	file := &tele.Document{
		File:     tele.FromDisk("/root/mini_billing/storages/assets/" + fileName),
		FileName: fileName,
		Caption:  text,
	}
	return c.Send(file)
}
