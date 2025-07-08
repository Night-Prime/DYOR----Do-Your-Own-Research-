package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
    "strings"
    "bytes"

	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/config"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
    "github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/errors"
)

// Here, this service is responsible for fetching data from various APIs.
// It defines interfaces for different asset types (stocks, bonds, news and cryptocurrencies)
// and implements the functions to fetch data from those APIs.

// (Lot of Dependency Injection happening here)
// TODO Implement the SendRequest reusable function when refactoring

// For Stock:
type StockAPIClient interface {
    GetStockData(symbol string) ([]models.StockData, error)
}

func NewStockClient() StockAPIClient {
    return &stockClientImpl{}
}

type stockClientImpl struct{}

func (c *stockClientImpl) GetStockData(symbol string) ([]models.StockData, error) {

	fmt.Println("The Stock API Client Layer")
	fmt.Println("--------------------------------------------- \n")

    cfg := config.Get()

    queryParams := map[string]string{
        "symbols": symbol,
    }

    // construct the URL with query parameters
    url := fmt.Sprintf("%s?%s", cfg.StockAPI_URL, "symbols="+queryParams["symbols"])
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, &errors.DatabaseError{
            Message: "Request Error",
            Err: err,
        }
    }

    req.Header.Set("x-rapidapi-host", cfg.StockHostname)
    req.Header.Set("x-rapidapi-key", cfg.StockAPI_Key)
	req.Header.Set("Accept-Encoding", "application/json")


    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, &errors.DatabaseError{
            Message: "Request Error",
            Err: err,
        }
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Error: %s", res.Status)
    }

	bodyBytes, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, &errors.DatabaseError{
            Message: "Error reading response",
            Err: err,
        }
    }

    var apiResponse models.StockAPIResponse
    if err := json.Unmarshal(bodyBytes, &apiResponse); err != nil {
        return nil, &errors.DatabaseError{
            Message: "Decoding Error",
            Err: err,
        }
    }

    if len(apiResponse.Data.QuoteResponse.Result) == 0 {
        return nil, &errors.ValidationError{
            Message: "Empty Response",
        }
    }
    return apiResponse.Data.QuoteResponse.Result, nil

}

// For Crypto:
type CryptoAPIClient interface {
	GetCryptoData(symbols []string) ([]models.CryptoData, error)
}

func NewCryptoClient() CryptoAPIClient {
    return &cryptoClientImpl{}
}

type cryptoClientImpl struct {}

func (c *cryptoClientImpl) GetCryptoData (symbols []string) ([]models.CryptoData, error) {

    fmt.Println("The Crypto API Client Layer")
	fmt.Println("--------------------------------------------- \n")

    cfg := config.Get()

    // making the request
    queryParams := map[string]string{
        "symbol": strings.Join(symbols, ","),
    }

    req, err := http.NewRequest("GET", cfg.CryptoAPI_URL, nil)
    if err != nil {
        return nil, &errors.DatabaseError{
            Message: "Request Error",
            Err: err,
        }
    }

    // adding the queries (what's going on here ?)
    q := req.URL.Query()
    for key, value := range queryParams {
        if key == "symbol" {
            for _, symbol := range strings.Split(value, ",") {
                q.Add("symbols", symbol)
            }
        } else {
            q.Add(key, value)
        }
    }
    req.URL.RawQuery = q.Encode()


    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, &errors.DatabaseError{
            Message: "Response Error",
            Err: err,
        }
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Error: %s", res.Status)
    }

    bodyBytes, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, &errors.DatabaseError{
            Message: "Response Error",
            Err: err,
        }
    }

    var apiResponse models.CryptoAPIResponse
    if err := json.Unmarshal(bodyBytes, &apiResponse); err != nil {
        return nil, &errors.DatabaseError{
            Message: "Decoding Error",
            Err: err,
        }
    }

    if len(apiResponse.DataArray) == 0 {
        return nil, &errors.ValidationError{
            Message: "Request Error",
        }
    }
    return apiResponse.DataArray, nil

}

// For News:
type NewsAPIClient interface {
    GetNewsData() ([]*models.News, error)
    GetTopGainersLosers() (*models.TickerUpdates, error)
}

func NewsClient() NewsAPIClient {
    return &newsClientImpl{}
}

type newsClientImpl struct{}

func (c *newsClientImpl) GetNewsData() ([]*models.News, error) {
    
    fmt.Println("The News API Client Layer")
    fmt.Println("--------------------------------------------- \n")

    cfg := config.Get()

    // making the request
    queryParams := map[string]string{
        "function": cfg.VANTAGE_FUNCTION,
        "apikey": cfg.VANTAGE_KEY,
        "sort":"LATEST",
        "limit":"200",
    }

    req, err := http.NewRequest("GET", cfg.VANTAGE_URL, nil)
    if err != nil {
        return nil, fmt.Errorf("Request error: %v", err)
    }

    // adding the queries
    q := req.URL.Query()
    for key, value := range queryParams {
        q.Add(key, value)
    }
    req.URL.RawQuery = q.Encode()
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Error: %s", res.Status)
    }

    bodyBytes, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, fmt.Errorf("Error reading response body: %v", err)
    }

    var wrapper struct {
        Feed []models.News `json:"feed"`
    }

    if err := json.Unmarshal(bodyBytes, &wrapper); err != nil {
        return nil, fmt.Errorf("Decoding error: %v", err)
    }

    if len(wrapper.Feed) == 0 {
        return nil, fmt.Errorf("No news feed items found")
    }

    newsPtrList := make([]*models.News, len(wrapper.Feed))
    for i := range wrapper.Feed {
        newsPtrList[i] = &wrapper.Feed[i]
    }

    return newsPtrList, nil
}


func (c *newsClientImpl) GetTopGainersLosers() (*models.TickerUpdates, error) {
    fmt.Println("The News API Client Layer")
    fmt.Println("--------------------------------------------- \n")

    cfg := config.Get()

    // making the request
    queryParams := map[string]string{
        "function": cfg.VANTAGE_FUNCTION_TOP,
        "apikey": cfg.VANTAGE_KEY,
    }

    req, err := http.NewRequest("GET", cfg.VANTAGE_URL, nil)
    if err != nil {
        return nil, fmt.Errorf("Request error: %v", err)
    }

    // adding the queries
    q := req.URL.Query()
    for key, value := range queryParams {
        q.Add(key, value)
    }
    req.URL.RawQuery = q.Encode()
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Error: %s", res.Status)
    }

    bodyBytes, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, fmt.Errorf("Error reading response body: %v", err)
    }

    var apiResponse models.TickerUpdates

    if err := json.Unmarshal(bodyBytes, &apiResponse); err != nil {
        return nil, fmt.Errorf("Decoding error: %v", err)
    }

    return &apiResponse, nil
}

type AIInsightClient interface {
    GetAssetInsight(assetInfo string) (models.FinancialAnalysisResponse, error)
}

type AIInsightClientImpl struct {}

func NewAIClient() AIInsightClient {
    return &AIInsightClientImpl{}
}

func (a *AIInsightClientImpl) GetAssetInsight(assetInfo string) (models.FinancialAnalysisResponse, error) {
    fmt.Println("The AI insight Client Layer")
    fmt.Println("--------------------------------------------- \n")

    cfg := config.Get()

    prompt := fmt.Sprintf(`
    You are an expert financial analyst and investment advisor with deep expertise in:
    - Traditional financial markets (stocks, bonds, commodities)
    - Cryptocurrency and digital asset markets
    - DeFi protocols and yield farming strategies
    - NFT market dynamics and valuation
    - Blockchain metrics and on-chain analysis
    - Technical analysis and chart patterns
    - Risk assessment and portfolio management
    - Economic indicators and market trends
    - Sector analysis and industry dynamics
    - Investment strategies and recommendations.

    You are tasked with analyzing the following real-time asset data. The assets may be a mix of traditional stocks and cryptocurrencies. Perform a comprehensive, multi-layered financial analysis on each asset individually and then provide a portfolio-wide assessment.

    Injected Real-Time Data:
    %s

    Tasks:
    1. Market Sentiment: Overall market mood, confidence level, key drivers, and sentiment score
    2. Technical Analysis: Trend direction, support/resistance levels, indicators (RSI, MACD, volume, moving averages, patterns, momentum)
    3. Risk Assessment: Overall risk level, risk types (market, liquidity, volatility, regulatory, smart contract), volatility level, downside protection
    4. Actionable Insights: Investment recommendations (buy/sell/hold), price targets (short/medium/long term), time horizon, entry/exit points, probability estimates
    5. Crypto-Specific Analysis (If Type = Crypto): Tokenomics, on-chain metrics, DeFi/NFT metrics if applicable
    6. Additional Analysis: Correlation between assets, macroeconomic impact, regulatory outlook, sector dynamics, innovation trends, and competitive landscape

    Output Format:
    Always respond as an Essay Write up, Multiple Major paragraphs for each asset, (Not a markdown but plain text as you need to strip away any potential '*', newlines '\n', or double new lines '\n\n' of any kind shouldn't be there).
    

    Notes:
    - Use actual numerical values from the data provided (market cap, price changes, volume, etc.)
    - Apply contextual financial reasoning for each insight.
    - Score each metric when appropriate, even if approximate.
    `, assetInfo)

    reqBody := models.PromptRequest{
        Model: cfg.AI_MODEL,
        Messages: []models.Message{
            {
                Role:    "system",
                Content: prompt,
            },
        },
    }

    body, err := json.Marshal(reqBody)
    if err != nil {
        return models.FinancialAnalysisResponse{}, &errors.DatabaseError{
            Message: "Error marshalling request body",
            Err:     err,
        }
    }

    req, err := http.NewRequest("POST", cfg.AIEndpoint, bytes.NewBuffer(body))
    if err != nil {
        return models.FinancialAnalysisResponse{}, &errors.DatabaseError{
            Message: "Error creating request",
            Err:     err,
        }
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", cfg.AIKey))

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return models.FinancialAnalysisResponse{}, &errors.DatabaseError{
            Message: "Error making request",
            Err:     err,
        }
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return models.FinancialAnalysisResponse{}, fmt.Errorf("Error: %s", res.Status)
    }

    bodyBytes, err := io.ReadAll(res.Body)
    if err != nil {
        return models.FinancialAnalysisResponse{}, &errors.DatabaseError{
            Message: "Error reading response body",
            Err:     err,
        }
    }


    var aiResp models.AIResponse

    if err := json.Unmarshal(bodyBytes, &aiResp); err != nil {
        return models.FinancialAnalysisResponse{}, &errors.DatabaseError{
            Message: "Error unmarshaling AI response",
            Err:     err,
        }
    }

    if len(aiResp.Choices) == 0 {
        return models.FinancialAnalysisResponse{}, &errors.ValidationError{
            Message: "Error unmarshaling AI response",
        }
    }


    analysis := models.FinancialAnalysisResponse{
        Sentiment: aiResp.Choices[0].Message.Content,
    }

    return analysis, nil
}



// I think a better design pattern could be used, this feels redundant, too much complexity
// but then this is my first time writing a big project in Go.


//TODO AI-Llama Client test

// func GetAIInsightsSummary(test string) (string, error) {

// }