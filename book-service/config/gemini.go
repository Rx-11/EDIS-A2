package config

import "github.com/Rx-11/EDIS-A2/book-service/ai"

var Gemini *ai.Gemini

func initGemini() {
	Gemini = ai.NewGemini(GetConfig().GeminiAPIKey)
}
