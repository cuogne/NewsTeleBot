## <samp> Bot Telegram </samp>

> tinh nang thong bao tin moi cua [bot discord](https://github.com/cuogne/discord-bot/blob/master/commands/fit-hcmus-news/INSTRUCTION.md) duoc tach ra mot phien ban cho telegram va duoc nang cap.

<p align="center">
<img src="./assets/goaway.png" alt="golangmascot" height="250" />
</p>

## <samp> 0. Cài đặt </samp>

<details>
<summary><samp> click here (for users) </samp></summary>

<br />

> Cài đặt app Telegram trên [PC/Laptop](https://desktop.telegram.org/) | [Android](https://play.google.com/store/apps/details?id=org.telegram.messenger&hl=vi) | [iOS](https://apps.apple.com/vn/app/telegram-messenger/id686449807?l=vi)

> Start bot: [https://t.me/hcmus_tintuc_bot](https://t.me/hcmus_tintuc_bot)

</details>

## <samp> 1. Advanced features </samp>

- Bot crawl tin tức của HCMUS và gửi thông báo đến user thông qua Telegram.
- New feature: Tóm tắt nội dung chính của bài viết bằng Gemini kèm link bài viết.
- Quét tin tức mới 10 phút/lần, gửi ngay tới user khi có tin mới.
- Tập trung vào 1 feature chính là thông báo tin mới (không tích hợp quá nhiều tính năng như bot discord).
- Sử dụng goroutines, channel và worker pool của Go để xử lý đồng thời, tăng hiệu suất.

<details>
<summary><samp>Nguồn tin tức: HCMUS</samp></summary>

<br />

| Nguồn | URL |
| ---- | --- |
| Thông tin dành cho sinh viên | https://hcmus.edu.vn/category/dao-tao/dai-hoc/thong-tin-danh-cho-sinh-vien |
| Lịch thi/Thông báo phòng khảo thí | http://ktdbcl.hcmus.edu.vn/ |
| Khoa CNTT - FIT@HCMUS | https://www.fit.hcmus.edu.vn/vn/Default.aspx?tabid=53 |

> Đây là các nguồn chính thống được bot crawl, có thể sẽ được update thêm.

</details>

## <samp> 2. Telegram commands </samp>

- `/start`: gioi thieu va huong dan.
- `/subscribe`: dang ky nhan thong bao.
- `/unsubscribe`: huy nhan thong bao.

| Command |
| ------------------------------------------------------------ |
|<img src="./assets/command.png" width="860" /> |

### Demo

| PC/Laptop| Mobile |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| <img src="./assets/demo_laptop.png" width="600" /> | <img src="./assets/demo_mobile.png" height="500" /> |

| Thông báo |
| ------------------------------------------------------------ |
|<img src="./assets/demo_notice.png" width="860" /> |

> Lưu ý: Bot chỉ gửi cho chat riêng, không hỗ trợ chat nhóm.

<details>
<summary><samp>Factos 1:</samp></summary>

> <i>maybe toi se xoa tinh nang nay ben bot discord va chi phat trien them ben nay (neu co thoi gian) (🐳)</i>

</details>


## <samp> 3. Tech stack </samp>

- **Language**: Go 1.26.1
- **Database**: PostgreSQL (Supabase)
- **Crawler**: Colly (HTML), gofeed (RSS)
- **Content extractor**: go-readability
- **AI Summarization**: Gemini
- **Telegram Bot**: Telebot v4
- **Scheduler**: robfig/cron
- **Concurrency**: Goroutines, Channels, sync.WaitGroup, Worker Pool
- **Deployment**: Docker + CI/CD pipeline (GitHub Actions)

<details>
<summary><samp>Factos 2:</samp></summary>

> <i>t thay code bang golang suong hon js (i'm addicted to golang lol xD)</i>

</details>

## <samp> 4. Run (for dev) </samp>

<details>
<summary><samp> click here (if u want to run it locally)</samp></summary>

<br />

B1. Đảm bảo đã cài go 1.26.1: [Go](https://go.dev/doc/install)

```zsh
brew install go@1.26 # for homebrew (macOS)

go --version # check version
```

B2. Clone repo:
```zsh
git clone https://github.com/cuogne/NewsTeleBot.git
```

cd vào thư mục
```zsh
cd NewsTeleBot
```

B3. Cài dependencies:
```go
go mod tidy
```

hoặc:
```zsh
make deps
```


B4. tạo file `.env` dựa theo file [.env.example](.env.example) và điền thông tin
```zsh
TELEGRAM_BOT_TOKEN=your_telegram_bot_token
SUPABASE_URL=your_supabase_url
GEMINI_API_KEY=your_gemini_api_key
```

> Telegram Bot Token: Tạo bot trên Telegram bằng cách nhắn tin với [BotFather](https://t.me/BotFather), gõ `/newbot` và làm theo hướng dẫn.

> Supabase URL: Tạo project trên [Supabase](https://supabase.com/), chạy script tạo database trong [db/database.sql](db/database.sql) và lấy URL trong `Connect > Session pooler`.

> Gemini API Key: tạo tài khoản trên [Google AI Studio](https://aistudio.google.com/), chọn `Get API Key` và lấy key.

B5. Chạy bot:
```zsh
go run ./cmd/bot # chạy trực tiếp

go build -o ./bin/bot ./cmd/bot # build thành executable

make run # chạy bằng Makefile, lệnh như chạy trực tiếp

make dev # chạy ở chế độ dev, có hot reload (sử dụng air)
```
</details>