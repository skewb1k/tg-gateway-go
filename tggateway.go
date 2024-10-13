package tggateway

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	token      string
	httpClient *http.Client
}

func NewClient(token string) Client {
	return Client{
		token:      token,
		httpClient: http.DefaultClient,
	}
}

func (c Client) makeAPIRequest(ctx context.Context, endpoint string, body any, result any) error {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "https://gatewayapi.telegram.org/"+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		return err
	}

	return nil
}
