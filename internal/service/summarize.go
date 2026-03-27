package service

import (
	"context"
	"fmt"
	"hcmus-news-tele-bot/internal/model"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func SummarizeContentWithGemini(content string, link string) (model.SummaryResult, error) {
	ctx := context.Background()
	geminiKey := os.Getenv("GEMINI_API_KEY")
	if geminiKey == "" {
		return model.SummaryResult{}, fmt.Errorf("GEMINI_API_KEY is not set in environment variables")
	}

	client, err := genai.NewClient(ctx,
		option.WithAPIKey(geminiKey),
	)
	if err != nil {
		return model.SummaryResult{}, err
	}
	defer client.Close()

	geminiModel := client.GenerativeModel("gemini-2.5-flash")
	geminiModel.SetTemperature(0.7)

	prompt := fmt.Sprintf(`
		Bạn là một biên tập viên tóm tắt tin tức chuyên nghiệp. Nhiệm vụ của bạn là:
		- Tóm tắt nội dung tin tức sau không vượt quá 3 dòng -> người dùng sẽ cảm thấy quá dài và không đọc (Bắt buộc - Key).
		- Phải đi qua đủ hết nội dung của trang web, tóm tắt lại đầy đủ -> người dùng chưa cần ấn vào link vẫn có thể nắm được sơ qua nội dung chính của bài viết.
		- Chọn những dòng quan trọng/hấp dẫn để tóm tắt -> người dùng hứng thú -> vào link đọc tiếp.
		- Không cần chào hỏi, vô thẳng nội dung chính, không cần nói thêm gì khác.
		- Nếu tóm tắt xong, nội dung có câu: Trang web này sử dụng cookie, thì không ghi đoạn này, nếu không đủ nội dung thì để rỗng.
		- Ở cuối tóm tắt, xuống dòng 1 cái, ghi câu: Chi tiết xem tại: %s.
		Nội dung bài viết như sau: %s`, link, content)

	resp, err := geminiModel.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return model.SummaryResult{}, err
	}
	if len(resp.Candidates) == 0 {
		return model.SummaryResult{Summary: "Không có phản hồi từ AI"}, nil
	}

	var summary string
	for _, part := range resp.Candidates[0].Content.Parts {
		summary += fmt.Sprintf("%v", part)
	}

	return model.SummaryResult{
		Summary:         summary,
		PromptToken:     int(resp.UsageMetadata.PromptTokenCount),
		CompletionToken: int(resp.UsageMetadata.CandidatesTokenCount),
	}, nil
}
