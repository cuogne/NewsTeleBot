<h2 align="center"><samp>🐳 HCMUS News Telegram Bot 🐳</samp></h2>

<p align="center">
  <img src="./assets/fly_gopher.png" alt="golangmascot" width="400" />
</p>

<!-- <p align="center">
  <img src="./assets/kite_gopher_yellow.png" alt="golangmascot" width="200" />
</p> -->

<p align="center"><samp> Bot Telegram crawl tin tức của HCMUS </samp></p>

<p align="center">
<a href="https://t.me/hcmus_tintuc_bot"><img src="https://img.shields.io/badge/Telegram-Start%20Bot-26A5E4?style=for-the-badge&logo=telegram&logoColor=white" alt="telegram" /></a>
<img src="https://img.shields.io/badge/Go-1.26.1-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="go" />
<img src="https://img.shields.io/badge/AI-Gemini-4285F4?style=for-the-badge&logo=google&logoColor=white" alt="gemini" />
</p>

## <samp> why </samp>

> tinh nang thong bao tin moi cua [bot discord](https://github.com/cuogne/discord-bot/blob/master/commands/fit-hcmus-news/INSTRUCTION.md) duoc tach ra mot phien ban cho telegram va nang cap hon.

## <samp> 0. How to use </samp>

<details>
<summary><samp> click here (for users) </samp></summary>

<br />

> Cài đặt app Telegram trên [PC/Laptop](https://desktop.telegram.org/) | [Android](https://play.google.com/store/apps/details?id=org.telegram.messenger&hl=vi) | [iOS](https://apps.apple.com/vn/app/telegram-messenger/id686449807?l=vi)

> Start bot: [https://t.me/hcmus_tintuc_bot](https://t.me/hcmus_tintuc_bot)

</details>

## <samp> 1. Advanced features </samp>

- Bot crawl tin tức của HCMUS và gửi thông báo đến user thông qua Telegram.
- New feature: Tóm tắt nội dung chính của bài viết bằng Gemini, giúp nắm nhanh thông tin trước khi đọc tiếp.
- Quét tin tức mới 10 phút/lần, gửi ngay tới user khi có tin mới.
- Tập trung vào 1 feature chính là thông báo tin mới (không tích hợp quá nhiều tính năng như bot discord).
- Sử dụng goroutines, channel và worker pool của Go để xử lý đồng thời, tăng hiệu suất.

> Lưu ý: Bot chỉ hỗ trợ private chat (chat riêng/cá nhân), không hỗ trợ group chat.

<!-- <details>
<summary><samp>Nguồn tin tức: HCMUS</samp></summary>

<br />

| Nguồn | URL |
| ---- | --- |
| Thông tin dành cho sinh viên | https://hcmus.edu.vn/category/dao-tao/dai-hoc/thong-tin-danh-cho-sinh-vien |
| Lịch thi/Thông báo phòng khảo thí | http://ktdbcl.hcmus.edu.vn/ |
| Khoa CNTT - FIT@HCMUS | https://www.fit.hcmus.edu.vn/vn/Default.aspx?tabid=53 |

> Đây là các nguồn chính thống được bot crawl, có thể sẽ được update thêm.

</details> -->

## <samp> 2. News Resources </samp>

| Category | URL |
| ---- | --- |
| Thông tin dành cho sinh viên | https://hcmus.edu.vn/category/dao-tao/dai-hoc/thong-tin-danh-cho-sinh-vien |
| Lịch thi - Phòng khảo thí | http://ktdbcl.hcmus.edu.vn/ |
| Thông báo - Phòng khảo thí | http://ktdbcl.hcmus.edu.vn/ |
| Khoa CNTT - FIT@HCMUS | https://www.fit.hcmus.edu.vn/vn/Default.aspx?tabid=53 |

<details>
<summary><samp>ảnh meme của Gopher: </samp></summary>
  
<br />
  
> [https://github.com/nlepage/gophers](https://github.com/nlepage/gophers)

</details>

## <samp> 3. Telegram commands </samp>

<table>
  <tr>
    <td valign="top" width="70%">
      <p><b>Command</b></p>
      <p>
        <code>/start</code> - Giới thiệu và hướng dẫn<br />
        <code>/subscribe</code> - Đăng ký nhận thông báo<br />
        <code>/unsubscribe</code> - Hủy nhận thông báo
      </p>
      <img src="./assets/command.png" alt="command-preview" width="700" />
    </td>
    <td align="right" valign="top" width="30%">
      <img src="./assets/kite_gopher_yellow.png" alt="golangmascot" width="170" />
    </td>
  </tr>
</table>

<!-- | Command |
| ------------------------------------------------------------ |
|<img src="./assets/command.png" width="400" /> | -->


### Demo

<table align="center">
  <thead>
    <tr>
      <th>PC/Laptop</th>
      <th>Mobile</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td align="center"><img src="./assets/demo_laptop.png" width="530" /></td>
      <td align="center"><img src="./assets/demo_mobile.png" height="440" /></td>
    </tr>
  </tbody>
</table>

<table align="center">
  <thead>
    <tr>
      <th>Thông báo</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td align="center"><img src="./assets/demo_notice.png" width="755" /></td>
    </tr>
  </tbody>
</table>


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

B1. Đảm bảo đã cài Go 1.26.1 thông qua: [https://go.dev/doc/install](https://go.dev/doc/install)

```zsh
brew install go # for homebrew (macOS)

go version # check version
```

B2. Clone repo này về và di chuyển vào thư mục:
```zsh
git clone https://github.com/cuogne/NewsTeleBot.git

cd NewsTeleBot
```

B3. chạy file [setup.sh](setup.sh) đã được thiết lập sẳn trong repo để setup:
```bash
bash setup.sh
```

B4. Sau khi chạy xong, bạn sẽ có file `.env` trong thư mục với các biến môi trường cần thiết. 

```env
TELEGRAM_BOT_TOKEN=your_telegram_bot_token
SUPABASE_URL=your_supabase_url
GEMINI_API_KEY=your_gemini_api_key
```

Thay các token trong file `.env` bằng token của bạn, cách lấy như sau:

- **Telegram Bot Token**: Cài đặt Telegram (link có ở trên), tạo bot trên Telegram bằng cách nhắn tin với [BotFather](https://t.me/BotFather), gõ `/newbot` và làm theo hướng dẫn.

- **Supabase URL**: Login và tạo project trên [Supabase](https://supabase.com/), dán script tạo database trong [db/database.sql](db/database.sql) vào và chạy nó, sau đó lấy URL trong `Connect > Session pooler`.

- **Gemini API Key**: tạo tài khoản trên [Google AI Studio](https://aistudio.google.com/), chọn `Get API Key` và lấy key.

B5. Run bot:
```zsh
go run ./cmd/bot # chạy trực tiếp

make run # chạy bằng Makefile, lệnh như chạy trực tiếp

make dev # chạy ở chế độ dev, có hot reload (sử dụng air)

go build -o ./bin/bot ./cmd/bot # build thành binary
```

Bonus: Run thông qua docker:

```zsh
docker build -t hcmus-news-tele-bot .
docker run -d --env-file .env --name my-tele-bot hcmus-news-tele-bot
```

</details>

<p align="center">
  <img src="./assets/heart_gopher.png" alt="golangmascot" width="150" />
</p>