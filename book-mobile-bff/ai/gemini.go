package ai

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

type Gemini struct {
	client *genai.Client
}

func NewGemini(apikey string) *Gemini {
	client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey:  apikey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		fmt.Println(err)
	}
	return &Gemini{client: client}
}

func (g *Gemini) Chat(req ChatRequest) (ChatResponse, error) {
	ctx := context.Background()

	var content []*genai.Content
	for _, m := range req.Messages {
		var part []*genai.Part
		part = append(part, &genai.Part{Text: m.Content})
		content = append(content, &genai.Content{
			Parts: part,
			Role:  m.Role,
		})
	}

	resp, err := g.client.Models.GenerateContent(ctx, GEMINI_MODEL, content, &genai.GenerateContentConfig{Temperature: func(f float32) *float32 { return &f }(1.2)})
	if err != nil {
		return ChatResponse{}, err
	}

	if len(resp.Candidates) == 0 {
		return ChatResponse{Response: ""}, nil
	}

	return ChatResponse{Response: resp.Candidates[0].Content.Parts[0].Text}, nil
}
