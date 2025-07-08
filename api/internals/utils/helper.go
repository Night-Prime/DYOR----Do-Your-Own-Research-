package utils

// Here we store all reusable functions : 

import(
	"fmt"
    "net/http"
    "io"
	"encoding/json"
    "strings"
    "regexp"
)

func SendRequest(client *http.Client, req *http.Request, v interface{}) error {
    res, err := client.Do(req)
    if err != nil {
        return err
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return fmt.Errorf("unexpected status code: %d", res.StatusCode)
    }

    bodyBytes, err := io.ReadAll(res.Body)
    if err != nil {
        return err
    }

    return json.Unmarshal(bodyBytes, v)
}

func StripAllFormatting(text string) string {
	text = regexp.MustCompile(`[#*_\-~`+"`]").ReplaceAllString(text, "")
	text = regexp.MustCompile(`<[^>]*>`).ReplaceAllString(text, "")
	text = regexp.MustCompile(`https?://\S+`).ReplaceAllString(text, "")
	text = regexp.MustCompile(`\[([^\]]+)\]\([^\)]+\)`).ReplaceAllString(text, `$1`)
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")
	text = regexp.MustCompile(`[^\w\s.,!?;:'"-]`).ReplaceAllString(text, "")
	
	return strings.TrimSpace(text)
}