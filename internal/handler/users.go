package handler

import (
	"fmt"
	"miniBilling/internal/constant"
	"miniBilling/internal/pkg/button"
	"miniBilling/internal/po/mongo"
	"miniBilling/internal/usecase"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	tele "gopkg.in/telebot.v4"
)

type UserHanderInterface interface{
	Start(c tele.Context) error 
	ClearAction(c tele.Context) error
	UserMongo(c tele.Context, teleId int64) (*mongo.Users, error)
	PreLogin(c tele.Context) error
	Login(c tele.Context, user_info *mongo.Users, usercode string) error
	Password(c tele.Context, user_info *mongo.Users, password string) error 
}
type UserHandler struct {
	userUC usecase.Users
	Bot    *tele.Bot
}

func NewUserHandler(userUC usecase.Users, bot *tele.Bot) UserHanderInterface {
	return &UserHandler{userUC: userUC, Bot: bot}
}

func (h *UserHandler) Start(c tele.Context) error {
	user := c.Sender()
	teleId := user.ID
	teleUsername := fmt.Sprintf("%s %s", user.LastName, user.FirstName)
	tele_user_code := user.Username

	// Kiểm tra user có tồn tại trong MongoDB chưa
	userMongo,_ := h.userUC.UserMongo(user.ID)
	
	if userMongo == nil {
		newUser := &mongo.Users{
			TeleId:       &teleId,
			TeleName:     &teleUsername,
			TeleUsername: &tele_user_code,
			Status:       constant.USER_STATUS_INACTIVED.Pointer(),
		}
		if err := h.userUC.CreateUserMongo(newUser); err != nil {
			return c.Send("❌ Lỗi khi lưu thông tin user!")
		}
		intro := "Chào mừng bạn đến với mini Billing. Để sử dụng dịch vụ, vui lòng đăng nhập vào ứng dụng"
		return c.Send(intro, button.Login_InlineKeys)
	}else {
		menu := button.GetMainMenu(*userMongo.Role)
		message := fmt.Sprintf("Chào mừng %s quay trở lại. Mình có thể giúp gì cho bạn ?",*userMongo.UserName)
		return c.Send(message,menu)
	}

	// Gửi tin nhắn chào mừng và nút đăng nhập
	
}

func (h *UserHandler) ClearAction(c tele.Context) error {
	user := c.Sender()
	teleId := user.ID

	userMongo,_ := h.userUC.UserMongo(user.ID)

	h.userUC.UpdateUserMongo(teleId,bson.M{
		"action1": 	nil,
		"action2": 	nil,
		"action3": 	nil,
		"action4": 	nil,
		"action5": 	nil,
		"action6": 	nil,
	})
	menu := button.GetMainMenu(*userMongo.Role)
	return c.Send("Clear ok", menu)
}
func (h *UserHandler) UserMongo(c tele.Context, teleId int64) (*mongo.Users, error) {
	return h.userUC.UserMongo(teleId)
}

func (h *UserHandler) PreLogin(c tele.Context) error {
	user := c.Sender()
	teleId := user.ID
	userMongo,_ := h.userUC.UserMongo(teleId)

	if(userMongo.Email == nil || userMongo.Password == nil || userMongo.UserName ==  nil){
		h.userUC.UpdateUserMongo(teleId,bson.M{
			"status":    	constant.USER_STATUS_ACTIVED,
			"action1": 		constant.USER_ACTION_LOGIN,
			"action2": 		constant.USER_ACTION_USERNAME,
		})
		return c.Send("Mời bạn nhập tài khoản Billing")
	}
	h.userUC.UpdateUserMongo(teleId,bson.M{
		"action1": 		nil,
		"action2": 		nil,
		"action3": 		nil,
		"action4": 		nil,
		"action5": 		nil,
	})
	message := fmt.Sprintf("Bạn đã đăng nhập với tài khoản %s (%s). Vui lòng chọn thao tác dưới đây.", *userMongo.UserName, *userMongo.UserCode)
	return c.Send(message)
}

func (h *UserHandler) Login(c tele.Context, user_info *mongo.Users, usercode string) error {
	user_name := h.userUC.Check_user_billing(usercode)
	teleid := *user_info.TeleId
	tele_user := *user_info.TeleUsername

	if user_name != "" {
		h.userUC.UpdateUserMongo(teleid,bson.M{
			"username": 	user_name,
			"user_code": 	usercode,
			"tele_user": 	tele_user,
			"status":    	constant.USER_STATUS_ACTIVED,
			"action3": 		constant.USER_ACTION_PASSWORD,
		})
		message := fmt.Sprintf("Xin chào %s, Vui lòng nhập mật khẩu của bạn", user_name)
		return c.Send(message)
	}
	return c.Send("Xin chào,khong tháy user rùi")
}

func (h *UserHandler) Password(c tele.Context, user_info *mongo.Users, password string) error {
	user_code := user_info.UserCode
	tele_id := user_info.TeleId

	users, err := h.userUC.Check_password_billing(*user_code, password)
	if err != nil {
		return err 
	}
	// Kiểm tra nếu có user trong billing
	if len(users) > 0 {
		user_name_billing := users[0].User_name 
		
		password := users[0].Password_show
		email := users[0].Email
		role := users[0].Role
		company := users[0].Company_name

		re := regexp.MustCompile(`\["(.*?)"\]`)
		match := re.FindStringSubmatch(role)
		if len(match) > 1 {
			role = match[1] 
		}
		err := h.userUC.UpdateUserMongo(*tele_id, bson.M{
			"username"	: user_name_billing,
			"password" 	: password,
			"role" 		: role,
			"company" 	: company,
			"email" 	: email,
			"action1"	: nil,
			"action2"	: nil,
			"action3"	: nil,
		})
		if err != nil {
			return err
		}
		return c.Send("Bạn đã đăng nhập thành công", button.GetMainMenu(role))
	}

	return c.Send("k vào điều kiện")
}





