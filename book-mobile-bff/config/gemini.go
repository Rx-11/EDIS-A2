package config

import "github.com/Rx-11/EDIS-A1/ai"

var Gemini *ai.Gemini

func initGemini() {
	Gemini = ai.NewGemini(GetConfig().GeminiAPIKey)
}
