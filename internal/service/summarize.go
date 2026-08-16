package service

import (
	"context"
	"errors"
	"fmt"
	"hcmus-news-tele-bot/config"
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

func isGeminiFallbackError(err error) bool {
	var apiErr genai.APIError
	if !errors.As(err, &apiErr) {
		return false
	}

	// 429 (Too Many Requests), 503 (Service Unavailable), 500 (Internal Server Error)
	return apiErr.Code == 429 || apiErr.Code == 503 || apiErr.Code == 500
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
	geminiModels := config.Models()

	client, err := getGeminiClient(ctx)
	if err != nil {
		return model.SummaryResult{}, err
	}

	configGemini := &genai.GenerateContentConfig{
		Temperature: genai.Ptr[float32](0.4),
	}

	prompt := fmt.Sprintf(summarizePrompt, content)

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
