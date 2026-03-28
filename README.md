## <samp> Bot Telegram </samp>

> tinh nang thong bao tin moi cua [bot discord](https://github.com/cuogne/discord-bot/blob/master/commands/fit-hcmus-news/INSTRUCTION.md) duoc tach ra mot phien ban cho telegram va duoc nang cap.

<p align="center">
<img src="https://media.licdn.com/dms/image/v2/D4E12AQE82cn9WoN0pw/article-cover_image-shrink_720_1280/article-cover_image-shrink_720_1280/0/1682192518638?e=1776297600&v=beta&t=VlhJlzqmLGJao656IEt9gf2R0HUE6h9C0KQ7tuE5ptI" alt="golangmascot" height="300" />
</p>

## <samp> 1. Advanced features </samp>

- Bot crawl tin tức của HCMUS và gửi thông báo đến user thông qua Telegram.
- New feature: Tóm tắt nội dung chính của bài viết bằng Gemini kèm link bài viết.
- Quét tin tức mới 10 phút/lần, gửi ngay tới user khi có tin mới.
- Tập trung vào 1 feature chính là thông báo tin mới (không tích hợp quá nhiều tính năng như bot discord).
- Sử dụng goroutines, channel và worker pool của Go để xử lý đồng thời, tăng hiệu suất.

<details>
	<summary> <samp> Nguồn tin tức: </samp></summary>
	<ul>
		<li>FIT@HCMUS: <a href="https://www.fit.hcmus.edu.vn/vn/Default.aspx?tabid=53">https://www.fit.hcmus.edu.vn/vn/Default.aspx?tabid=53</a></li>
		<li>Lịch thi - Phòng khảo thí: <a href="http://ktdbcl.hcmus.edu.vn/">http://ktdbcl.hcmus.edu.vn/</a></li>
		<li>Thông báo - Phòng khảo thí: <a href="http://ktdbcl.hcmus.edu.vn/">http://ktdbcl.hcmus.edu.vn/</a></li>
		<li>Thông tin dành cho sinh viên - HCMUS: <a href="https://hcmus.edu.vn/category/dao-tao/dai-hoc/thong-tin-danh-cho-sinh-vien">https://hcmus.edu.vn/category/dao-tao/dai-hoc/thong-tin-danh-cho-sinh-vien</a></li>
	</ul>
</details>

## <samp> 2. Telegram commands </samp>

- `/start`: gioi thieu va huong dan.
- `/subscribe`: dang ky nhan thong bao.
- `/unsubscribe`: huy nhan thong bao.

> Lưu ý: Bot chỉ gửi cho chat riêng, không hỗ trợ chat nhóm.

<details>
	<summary> <samp> Factos 1: </samp></summary>
	> <i>maybe toi se xoa tinh nang nay ben bot discord va chi phat trien them ben nay (neu co thoi gian) (🐳) </i>
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
	<summary> <samp> Factos 2: </samp></summary>
	> <i>t thay code bang golang suong hon js (i'm addicted to golang lol xD)</i>
</details>