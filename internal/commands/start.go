package command

import (
	tele "gopkg.in/telebot.v4"
)

// Start handles the /start command
func Start(c tele.Context) error {
	msg := "Xin chào " + c.Sender().FirstName + "!\n"
	msg += "Bot này sẽ giúp bạn cập nhật thông tin từ các trang web của trường Đại học Khoa học Tự nhiên.\n"
	msg += "Các lệnh:\n"
	msg += "/start - Bắt đầu\n"
	msg += "/help - Hỗ trợ bạn cấu hình\n"
	msg += "/subscribe - Đăng ký nhận thông báo\n"
	msg += "/unsubscribe - Hủy nhận thông báo\n"

	return c.Send(msg)
}
