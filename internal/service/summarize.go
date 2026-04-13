package service

import (
	"context"
	"fmt"
	"hcmus-news-tele-bot/internal/model"
	"os"
	"strings"
	"sync"

	"google.golang.org/genai"
)

var (
	geminiClient *genai.Client
	geminiOnce   sync.Once
	initErr      error
)

func getGeminiClient(ctx context.Context) (*genai.Client, error) {
	geminiOnce.Do(func() {
		geminiKey := os.Getenv("GEMINI_API_KEY")
		if geminiKey == "" {
			initErr = fmt.Errorf("GEMINI_API_KEY is not set in environment variables")
			return
		}

		client, err := genai.NewClient(ctx, &genai.ClientConfig{
			APIKey:  geminiKey,
			Backend: genai.BackendGeminiAPI,
		})

		if err != nil {
			initErr = err
			return
		}

		geminiClient = client
	})

	return geminiClient, initErr
}

func SummarizeContentWithGemini(content string) (model.SummaryResult, error) {
	ctx := context.Background()

	client, err := getGeminiClient(ctx)
	if err != nil {
		return model.SummaryResult{}, err
	}

	configGemini := &genai.GenerateContentConfig{
		Temperature: genai.Ptr[float32](0.7),
	}

	prompt := fmt.Sprintf(`
		Bạn là một biên tập viên tóm tắt tin tức chuyên nghiệp. Nhiệm vụ của bạn là:
		- Tóm tắt nội dung tin tức sau không vượt quá 3 dòng -> người dùng sẽ cảm thấy quá dài và không đọc (Bắt buộc - Key).
		- Văn phong tóm tắt phải tự nhiên, không quá máy móc, bám sát nội dung bài viết.
		- Phải đi qua đủ hết nội dung của trang web, tóm tắt lại đầy đủ -> người dùng chưa cần ấn vào link vẫn có thể nắm được sơ qua nội dung chính của bài viết.
		- Chọn những dòng quan trọng/hấp dẫn để tóm tắt -> người dùng hứng thú -> vào link đọc tiếp.
		- Không cần chào hỏi, vô thẳng nội dung chính, không cần nói thêm gì khác.
		- Nếu tóm tắt xong, nội dung có câu: Trang web này sử dụng cookie, thì không ghi đoạn này, nếu không đủ nội dung thì để rỗng.
		Nội dung bài viết như sau: %s`, content)

	resp, err := client.Models.GenerateContent(
		ctx,
		"gemini-3-flash-preview",
		genai.Text(prompt),
		configGemini,
	)

	if err != nil {
		return model.SummaryResult{}, err
	}

	summary := strings.TrimSpace(resp.Text())
	if summary == "" {
		return model.SummaryResult{
			Summary: "Không có phản hồi từ AI",
		}, nil
	}

	return model.SummaryResult{
		Summary:         summary,
		PromptToken:     int(resp.UsageMetadata.PromptTokenCount),
		CompletionToken: int(resp.UsageMetadata.CandidatesTokenCount),
	}, nil
}
