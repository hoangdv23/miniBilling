package bot

import (
	"gopkg.in/telebot.v4"
	
)

// CustomAutoRespond gửi tin nhắn khi không có handler nào xử lý
func AutoRespond(message string) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			// Gọi middleware tiếp theo
			err := next(c)

			// Nếu err == nil thì tin nhắn đã được xử lý => không cần auto respond
			if err == nil {
				return nil
			}

			// Nếu err khác nil, bot sẽ tự động phản hồi
			return c.Send(message)
		}
	}
}