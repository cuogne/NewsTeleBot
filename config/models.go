package config

import (
	"log"
	"os"
	"strings"
	"sync"
)

// default models to use if GEMINI_MODELS env is not set (fallback order)
var defaultModels = []string{
	"gemini-3.5-flash-lite",
	"gemini-3.7-flash",
	"gemini-3.6-flash",
	"gemini-3.5-flash",
}

var (
	models     []string
	modelsOnce sync.Once
)

// Models returns gemini models from GEMINI_MODELS env (comma-separated),
// falling back to default list if unset.
func Models() []string {
	modelsOnce.Do(func() {
		raw := strings.TrimSpace(os.Getenv("GEMINI_MODELS"))
		if raw == "" {
			models = defaultModels
			log.Printf("GEMINI_MODELS not set, using default models: %v", models)
			return
		}

		for m := range strings.SplitSeq(raw, ",") {
			if m = strings.TrimSpace(m); m != "" {
				models = append(models, m)
			}
		}

		if len(models) == 0 {
			models = defaultModels
		}
	})

	return models
}
