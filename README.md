## <samp> Bot Telegram </samp>

> tinh nang thong bao tin moi cua [bot discord](https://github.com/cuogne/discord-bot) duoc tach ra va nang cap, viet bang Go.

<p align="center">
<img src="https://media.licdn.com/dms/image/v2/D4E12AQE82cn9WoN0pw/article-cover_image-shrink_720_1280/article-cover_image-shrink_720_1280/0/1682192518638?e=1776297600&v=beta&t=VlhJlzqmLGJao656IEt9gf2R0HUE6h9C0KQ7tuE5ptI" alt="golangmascot" height="300" />
</p>

## <samp> 1. Advanced features </samp>

- New feature: Tóm tắt sơ lược nội dung bài viết bằng Gemini kèm link bài viết chi tiết.
- Quét mỗi 5 phút/lần (thay vì 10 phút như bên bot discord).
- Tập trung vào 1 feature chính là thông báo tin mới (không tích hợp quá nhiều tính năng như bot discord).
- Sử dụng goroutines, channel và worker pool của Go để xử lý đồng thời, tăng hiệu suất.

<details>
	<summary> <samp> Factos 1: </samp></summary>
	> <i>t thay code bang golang suong hon js (i'm addicted to golang awww)</i>
</details>

## 2. Telegram commands

- `/start`: gioi thieu va huong dan.
- `/subscribe`: dang ky nhan thong bao.
- `/unsubscribe`: huy nhan thong bao.

> Lưu ý: Bot chỉ gửi cho chat riêng, không hỗ trợ chat nhóm.

<details>
	<summary> <samp> Factos 2: </samp></summary>
	> <i>maybe toi se go tinh nang nay ben bot discord va chi phat trien them ben nay (neu co thoi gian) (🐳) </i>
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
