package middleware

import (
	"fmt"
	"log"

	"miniBilling/internal/constant"
	"miniBilling/internal/pkg/button"
	"miniBilling/internal/usecase"

	tele "gopkg.in/telebot.v4"
)

func CheckUserMiddleware(uc usecase.Users) tele.MiddlewareFunc {
	// Danh sách route không cần kiểm tra quyền
	excludedRoutes := map[string]bool{
		"/start":   true,
		"/clear":   true,
		"/login":   true,
		"\fbtn_login|login": true,
	}

	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			user := c.Sender()
			if user == nil {
				return c.Send("❌ Không thể lấy thông tin user!")
			}

			// Xác định route từ Text hoặc CallbackQuery
			var route string
			if c.Callback() != nil {
				route = c.Callback().Data // Nếu là button, lấy nội dung callback
			} else {
				route = c.Text() // Nếu là tin nhắn, lấy nội dung text
			}

			if excludedRoutes[route] {
				return next(c)
			}

			// Gọi CheckUserConnectBilling để kiểm tra user
			userInfo, err := uc.UserMongo(user.ID)
			if err != nil {
				log.Println("❌ Lỗi kiểm tra user:", err)
				return c.Send("❌ Lỗi khi kiểm tra tài khoản. Vui lòng thử lại sau!")
			}

			if userInfo.Action1 != nil && *userInfo.Action1 == string(constant.USER_ACTION_LOGIN) {
				return next(c)
			}

			// Kiểm tra nếu user chưa đăng nhập vào hệ thống billing
			if userInfo.UserCode == nil || userInfo.UserName == nil {
				fmt.Println("lỗi chưa liên kết tài khoản billing")
				return c.Send("⚠️ Bạn chưa liên kết tài khoản với hệ thống billing! Hãy ấn nút dưới đây để đăng nhập tài khoản", &button.Login_InlineKeys)
			}

			return next(c)
		}
	}
}


