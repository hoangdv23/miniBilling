package handler

import (
	"fmt"
	auth "miniBilling/cmd/middleware"
	"miniBilling/internal/constant"
	"miniBilling/internal/pkg/bot"
	"miniBilling/internal/pkg/button"
	"miniBilling/internal/pkg/helpers"
	"miniBilling/internal/repository"
	"miniBilling/internal/usecase"

	tele "gopkg.in/telebot.v4"
	"gopkg.in/telebot.v4/middleware"
)



func NewServer(b *tele.Bot){
	b.Use(middleware.Logger())
	b.Use(middleware.Recover())
	b.Use(bot.AutoRespond("BOT không hiểu :D"))


	userRepo := repository.NewUserRepository()
	userUC := usecase.NewUserUseCase(userRepo)

	billingRepo := repository.NewBillineRepository()
	billingUC := usecase.NewBillingUsecase(billingRepo)
	b.Use(auth.CheckUserMiddleware(userUC))

 
	userHandler := 	NewUserHandler(userUC,b)
	billingHander := NewBillingHandler(billingUC,b)
	b.Handle("/start", func(c tele.Context) error {
		return userHandler.Start(c)
	})
	b.Handle("/clear", func(c tele.Context) error {
		return userHandler.ClearAction(c)
	})
	b.Handle("/login",  func(c tele.Context) error {
		return userHandler.PreLogin(c)
	})

	
	//////////// ON TEXT ////////////////////////////
	b.Handle(tele.OnText, func(c tele.Context) error {
		message := helpers.TrimSpace(c.Text())
		user := c.Sender()
		user_info,error := userHandler.UserMongo(c,user.ID)

		if error != nil {
			return c.Send("Không thấy user")
		}

		// // Check login username
		if user_info.Action1 != nil &&  											// Yêu cầu không được nil
			*user_info.Action1 == string(constant.USER_ACTION_LOGIN) && 
			user_info.Action2 != nil && 
			*user_info.Action2 == string(constant.USER_ACTION_USERNAME) && 
			user_info.Action3 == nil{
			userHandler.Login(c,user_info,message) // check login
		}else if user_info.Action1 != nil && 
				user_info.Action2 != nil && 
				user_info.Action3 != nil && 
				*user_info.Action1 == string(constant.USER_ACTION_LOGIN) && 
				*user_info.Action2 == string(constant.USER_ACTION_USERNAME) && 
				*user_info.Action3 == string(constant.USER_ACTION_PASSWORD) {
				userHandler.Password(c,user_info,message) // check password
		}
		fmt.Println("ONTEXT UNKNOW")
		menu := button.GetMainMenu(*user_info.Role)

		return c.Send("không biết lệnh: ",message, menu)
	})

	/////////// CALL BACK (BUTTON) ////////////////////////////////////////////////////
	b.Handle(tele.OnCallback, func(ctx tele.Context) error {
		callback := helpers.TrimSpace(ctx.Callback().Data)

		if callback == "btn_login|login" {
			return userHandler.PreLogin(ctx)
		}else if callback == "btn_login|loginCode" {
			return billingHander.GetCodeLogin(ctx)
		}else if callback == "btn_intro|button_intro"{
			return ctx.Send("Xin giới thiệu với bạn, đây là bot mini Billing, phục vụ các tính năng nhanh gọn nhẹ ;D")
		}
		return ctx.Send(callback)
	})

}



