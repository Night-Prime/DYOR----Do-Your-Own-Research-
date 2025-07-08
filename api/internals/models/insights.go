package models

type PromptRequest struct {
	Model string `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

type AIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
			// Reasoning string `json:"reasoning,omitempty"`
		} `json:"message"`
	} `json:"choices"`
}

type FinancialAnalysisResponse struct {
	Sentiment string `json:"sentiment"`
	// Reasoning string `json:"reasoning"`
}

type AssetInsightsRequest struct {
	AssetInfo string `json:"asset_info"`
}