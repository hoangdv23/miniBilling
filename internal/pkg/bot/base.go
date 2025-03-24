package bot

import (
	"fmt"
	"log"
	"time"

	"miniBilling/internal/config"

	tele "gopkg.in/telebot.v4"
)

// Định nghĩa struct TeleBot
type TeleBot struct {
	*tele.Bot
}

// NewBot khởi tạo bot Telegram
func NewBot(botConfig config.Bot) *TeleBot {
	bot, err := tele.NewBot(tele.Settings{
		Token:  botConfig.Token_bot,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	return &TeleBot{bot} // Trả về con trỏ TeleBot
}

// StartBot bắt đầu bot
func (t *TeleBot) StartBot() {
	fmt.Println("Bot bắt đầu chạy...")
	t.Bot.Start() // Gọi bot.Start() từ TeleBot struct
}
