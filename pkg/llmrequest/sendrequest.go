// Package llmrequest provides functionality to send HTTP requests and handle responses.
package llmrequest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// SendRequest sends an HTTP request and handles the response.
func SendRequest(ctx context.Context, req *http.Request, headers map[string]string) (*http.Response, error) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()
		var errMessage any
		decodeErr := json.NewDecoder(res.Body).Decode(&errMessage)
		if decodeErr != nil {
			return nil, fmt.Errorf("api request failed: status %d %s %s (failed to decode error: %w)",
				res.StatusCode, res.Status, res.Request.URL, decodeErr)
		}
		return nil, fmt.Errorf("api request failed: status %d %s %s Message: %+v",
			res.StatusCode, res.Status, res.Request.URL, errMessage)
	}

	return res, nil
}
