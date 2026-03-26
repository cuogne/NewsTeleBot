package test

import (
	"testing"

	"hcmus-news-tele-bot/internal/service"
)

func TestGetArticles(t *testing.T) {
	articles, err := service.GetArticles()
	if err != nil {
		t.Fatalf("Đã xảy ra lỗi khi lấy bài báo: %v", err)
	}

	if len(articles) == 0 {
		t.Log("Không lấy được bài báo nào. Có thể các nguồn bị lỗi kết nối hoặc chưa có data mới.")
	} else {
		t.Logf("Lấy thành công %d bài báo.", len(articles))

		for i, article := range articles {
			t.Logf("Bài báo %d: %s - %s", i+1, article.Title, article.URL)
		}

		// // Lưu toàn bộ dữ liệu vào file news.json
		// file, err := os.Create("news.json")
		// if err != nil {
		// 	t.Fatalf("Lỗi tạo file news.json: %v", err)
		// }
		// defer file.Close()

		// encoder := json.NewEncoder(file)
		// encoder.SetIndent("", "  ")
		// if err := encoder.Encode(articles); err != nil {
		// 	t.Fatalf("Lỗi ghi dữ liệu ra file json: %v", err)
		// }

		// t.Log("Đã ghi thành công bộ bài báo craw được vào test/news.json")
	}
}
