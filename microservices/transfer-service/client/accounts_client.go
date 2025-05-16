package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AccountsClient struct {
	BASE_URL string
}

func NewAccountsClient(baseURL string) *AccountsClient {
	return &AccountsClient{
		BASE_URL: baseURL,
	}
}

func (c *AccountsClient) PostAmount(endpoint string, accountID string, amount float64) error {
	url := fmt.Sprintf("%s/accounts/%s/%s", c.BASE_URL, accountID, endpoint)
	payload := map[string]float64{"amount": amount}
	body, _ := json.Marshal(payload)
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("operation failed: %s", resp.Status)
	}
	return nil
}
