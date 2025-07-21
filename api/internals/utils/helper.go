package utils

// Here we store all reusable functions : 

import(
	"fmt"
    "net/http"
    "io"
	"encoding/json"
    "strings"
    "regexp"

    	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/config"
            "github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/errors"
)

func SendRequest(url string, method string, queryParams map[string]string, queryParamsType string, body interface{}) (*http.Response, error) {
    fmt.Printf("Sending Request here: %s with method: %s\n", url, method)
    fmt.Println("--------------------------------------------- \n")

    cfg := config.Get()

    client := &http.Client{}

    req, err := http.NewRequest(method, url, nil)
    if err != nil {
        return nil,  &errors.DatabaseError{
            Message: "Request Error",
            Err: err,
        }
    }

    switch queryParamsType {
    case "GetStockData":
        q := req.URL.Query()
        for key, value := range queryParams {
            q.Add(key, value)
        }
        req.URL.RawQuery = q.Encode()

        req.Header.Set("x-rapidapi-host", cfg.StockHostname)
        req.Header.Set("x-rapidapi-key", cfg.StockAPI_Key)
        req.Header.Set("Accept-Encoding", "application/json")

    case "GetCryptoData":
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
    case "GetNewsData":
        q := req.URL.Query()
        for key, value := range queryParams {
            q.Add(key, value)
        }
        req.URL.RawQuery = q.Encode()
    default:
        // No specific queryParamsType, continue without modifying the query
    }

    if body != nil {
        jsonBody, err := json.Marshal(body)
        if err != nil {
            return nil, &errors.DatabaseError{
                Message: "Error marshalling request body",
                Err: err,
            }
        }
        req.Body = io.NopCloser(strings.NewReader(string(jsonBody)))
        req.Header.Set("Content-Type", "application/json")
    }

    resp, err := client.Do(req)
    fmt.Sprintf("Response: %v",resp)
    if err != nil {
        return nil, &errors.DatabaseError{
            Message: "Error sending request",
            Err: err,
        }
    }

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Request failed with status: %s", resp.Status)
    }

    return resp, nil

}

func StripAllFormatting(text string) string {
    text = regexp.MustCompile(`[#*_\-~`+"`]").ReplaceAllString(text, "")
    text = regexp.MustCompile(`<[^>]*>`).ReplaceAllString(text, "")
    text = regexp.MustCompile(`https?://\S+`).ReplaceAllString(text, "")
    text = regexp.MustCompile(`\[([^\]]+)\]\([^\)]+\)`).ReplaceAllString(text, `$1`)
    text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")
    text = regexp.MustCompile(`[^\w\s.,!?;:'"$€£¥¢-]`).ReplaceAllString(text, "")
    
    return strings.TrimSpace(text)
}