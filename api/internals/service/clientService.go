package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
    "strings"

	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/config"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
)

// Here, this service is responsible for fetching data from various APIs.
// It defines interfaces for different asset types (stocks, bonds, and cryptocurrencies)
// and implements the functions to fetch data from those APIs.

// (Lot of Dependency Injection happening here)

// For Stock:
type StockAPIClient interface {
	GetStockData(symbol string) (*models.StockData, error)
}

func NewStockClient() StockAPIClient {
    return &stockClientImpl{}
}

type stockClientImpl struct{}

func (c *stockClientImpl) GetStockData(symbol string) (*models.StockData, error) {

	fmt.Println("The Stock API Client Layer")
	fmt.Println("--------------------------------------------- \n")

    cfg := config.Get()

	// make the request to the stock API
    url := fmt.Sprintf("%s?symbols=%s", cfg.StockAPI_URL, symbol)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("request error: %v", err)
    }

    req.Header.Set("x-rapidapi-host", cfg.StockHostname)
    req.Header.Set("x-rapidapi-key", cfg.StockAPI_Key)
	req.Header.Set("Accept-Encoding", "application/json")

	// fmt.Printf("Sending the Request: %s\nRequest URL: %s\nHeaders: %v\n", symbol, req.URL.String(), req.Header)
	// fmt.Println("--------------------------------------------- \n")

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("response error: %v", err)
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("error: %s", res.Status)
    }

	bodyBytes, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response body: %v", err)
    }

    var apiResponse models.StockData
    if err := json.Unmarshal(bodyBytes, &apiResponse); err != nil {
        return nil, fmt.Errorf("decoding error: %v", err)
    }

    if len(apiResponse.Data.QuoteResponse.Result) == 0 {
        return nil, fmt.Errorf("no stock data found for symbol %s", symbol)
    }

	return &apiResponse, nil
}

// For Crypto:
type CryptoAPIClient interface {
	GetCryptoData(symbols []string) (*models.CryptoData, error)
}

func NewCryptoClient() CryptoAPIClient {
    return &cryptoClientImpl{}
}

type cryptoClientImpl struct {}

func (c *cryptoClientImpl) GetCryptoData (symbols []string) (*models.CryptoData, error) {

    fmt.Println("The Crypto API Client Layer")
	fmt.Println("--------------------------------------------- \n")

    cfg := config.Get()

    // making the request
    queryParams := map[string]string{
        "symbol": strings.Join(symbols, ","),
    }

    req, err := http.NewRequest("GET", cfg.CryptoAPI_URL, nil)
    if err != nil {
        return nil, fmt.Errorf("request error: %v", err)
    }

    // adding the queries
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

    // fmt.Printf("Sending the Request: Request URL: %s\nHeaders: %v\n", req.URL.String(), req.Header)
	// fmt.Println("--------------------------------------------- \n")

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("response error: %v ", err)
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("error: %s", res.Status)
    }

    bodyBytes, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response body: %v", err)
    }
    // fmt.Printf("The Response Body: %s\n", string(bodyBytes))
    // fmt.Println("--------------------------------------------- \n")

    var apiResponse models.CryptoData
    if err := json.Unmarshal(bodyBytes, &apiResponse); err != nil {
        return nil, fmt.Errorf("decoding error: %v", err)
    }

    return &apiResponse, nil

}
