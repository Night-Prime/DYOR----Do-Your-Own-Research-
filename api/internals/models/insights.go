package models

type AIRequest struct {
    Test string `json:"test"`
}


type FinancialAnalysisResponse struct {
	Success           bool               `json:"success"`
	FinancialAnalysis FinancialAnalysis  `json:"financial_analysis"`
	RawAnalysis       string             `json:"raw_analysis"`
	Metadata          Metadata           `json:"metadata"`
	InputData         string             `json:"input_data"`
}

type FinancialAnalysis struct {
	MarketSentiment   MarketSentiment          `json:"market_sentiment"`
	TechnicalAnalysis TechnicalAnalysis        `json:"technical_analysis"`
	RiskAssessment    RiskAssessment           `json:"risk_assessment"`
	ActionableInsights ActionableInsights      `json:"actionable_insights"`
	RawTextAnalysis   string                   `json:"raw_text_analysis"`
}

type MarketSentiment struct {
	OverallSentiment string   `json:"overall_sentiment"`
	ConfidenceLevel  string   `json:"confidence_level"`
	KeyDrivers       []string `json:"key_drivers"`
}

type TechnicalAnalysis struct {
    Trend            string           `json:"trend"`
    SupportLevels    []float64        `json:"support_levels"`    // Changed from map to slice
    ResistanceLevels []float64        `json:"resistance_levels"` // Changed from map to slice
    KeyIndicators    map[string]interface{} `json:"key_indicators"`
    Patterns         []string         `json:"patterns"`
}
type RiskAssessment struct {
	RiskLevel           string            `json:"risk_level"`
	KeyRisks            []string          `json:"key_risks"`
	RiskFactors         map[string]string `json:"risk_factors"`
	VolatilityAssessment string           `json:"volatility_assessment"`
}

type ActionableInsights struct {
	Recommendations        []string          `json:"recommendations"`
	PriceTargets           map[string]float64 `json:"price_targets"`
	TimeHorizon            string            `json:"time_horizon"`
	ProbabilityAssessments map[string]string `json:"probability_assessments"`
}

type Metadata struct {
	Model                  string `json:"model"`
	AnalysisType           string `json:"analysis_type"`
	Temperature            float64 `json:"temperature"`
	TokensUsed             int     `json:"tokens_used"`
	PromptTokens           int     `json:"prompt_tokens"`
	CompletionTokens       int     `json:"completion_tokens"`
	FinishReason           string  `json:"finish_reason"`
	Timestamp              string  `json:"timestamp"`
	IncludeRiskAnalysis    bool    `json:"include_risk_analysis"`
	IncludeTechnicalAnalysis bool  `json:"include_technical_analysis"`
}
