// Package internal provides internal utilities for the project.
package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// SendRequest sends an HTTP request and handles the response.
func SendRequest(ctx context.Context, req *http.Request, headers map[string]string) (*http.Response, error) {
	if len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			return nil, err
		}
	}

	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()
		var errMessage any
		if err := json.NewDecoder(res.Body).Decode(&errMessage); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("api request failed: status Code: %d %s %s Message: %+v",
			res.StatusCode, res.Status, res.Request.URL, errMessage)
	}

	return res, nil
}
