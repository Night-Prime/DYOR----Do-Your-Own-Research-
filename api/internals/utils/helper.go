package utils

// Here we store all reusable functions : 

import(
	"fmt"
    "net/http"
    "io"
	"encoding/json"
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