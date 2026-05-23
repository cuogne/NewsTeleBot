package service

import (
	"context"
	"errors"
	"fmt"
	"hcmus-news-tele-bot/internal/model"
	"log"
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

var geminiModels = []string{
	"gemini-3.5-flash",
	"gemini-2.5-flash",
	"gemini-3.1-flash-lite",
}

func isGeminiFallbackError(err error) bool {
	var apiErr genai.APIError
	if !errors.As(err, &apiErr) {
		return false
	}

	return apiErr.Code == 429 || apiErr.Code == 503
}

func geminiErrorCode(err error) int {
	var apiErr genai.APIError
	if !errors.As(err, &apiErr) {
		return 0
	}

	return apiErr.Code
}

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
		Temperature: genai.Ptr[float32](0.4),
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

	var resp *genai.GenerateContentResponse

	for i, geminiModel := range geminiModels {
		resp, err = client.Models.GenerateContent(
			ctx,
			geminiModel,
			genai.Text(prompt),
			configGemini,
		)

		if err == nil {
			break
		}

		if !isGeminiFallbackError(err) || i == len(geminiModels)-1 {
			return model.SummaryResult{}, err
		}

		log.Printf("Gemini fallback: %s -> %s (%d)", geminiModel, geminiModels[i+1], geminiErrorCode(err))
	}

	if err != nil {
		return model.SummaryResult{}, err
	}

	summary := strings.TrimSpace(resp.Text())
	if summary == "" {
		return model.SummaryResult{}, nil
	}

	promptToken := 0
	completionToken := 0

	if resp.UsageMetadata != nil {
		promptToken = int(resp.UsageMetadata.PromptTokenCount)
		completionToken = int(resp.UsageMetadata.CandidatesTokenCount)
	}

	return model.SummaryResult{
		Summary:         summary,
		PromptToken:     promptToken,
		CompletionToken: completionToken,
	}, nil
}
