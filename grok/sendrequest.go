package grok

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) sendRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.APIKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.client.Do(req)
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
