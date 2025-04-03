package handler

import (
	"fmt"
	"miniBilling/internal/usecase"

	tele "gopkg.in/telebot.v4"
)

type BillingHandlerInterface interface {
	GetCodeLogin(c tele.Context) error 
}

type BillingHandler struct {
	billingUC usecase.Billing
	Bot    *tele.Bot
}

func NewBillingHandler(billingUC usecase.Billing, bot *tele.Bot) BillingHandlerInterface {
	return &BillingHandler{billingUC: billingUC, Bot: bot}
}

func (h *BillingHandler) GetCodeLogin(c tele.Context) error {
	user := c.Sender()
	teleId := user.ID
	userMongo,_ := h.billingUC.UserMongo(teleId)

	code := h.billingUC.GetCodeLogin(*userMongo.UserCode)
	message := fmt.Sprintf("🔑 Mã đăng nhập vào billing của bạn là:\n```\n%s\n```", code)

	return c.Send(message, tele.ModeMarkdown)
}
