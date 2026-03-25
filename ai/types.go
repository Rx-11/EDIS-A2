package ai

type Text struct {
	Text string `json:"text,omitempty"`
}

type ChatRequest struct {
	Messages []Message
}

type ChatResponse struct {
	Response string
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
