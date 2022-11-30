package gotermii

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// Config is a representation of config variables
type Config struct {
	APIKey   string `json:"apiKey"`
	BaseURL  string `json:"baseURL"`
	SenderID string `json:"senderId"`
}

// Client is a representation of a termii client
type Client struct {
	config Config
	client *http.Client
}

// NewClient creates a termii client using configuration variables
func NewClient(c Config, h *http.Client) Client {
	return Client{
		config: Config{
			APIKey:   c.APIKey,
			BaseURL:  c.BaseURL,
			SenderID: c.SenderID,
		},
		client: h,
	}
}

func (s *Client) makeRequest(method, rURL string, reqBody interface{}, resp interface{}) error {
	URL := fmt.Sprintf("%s/%s", s.config.BaseURL, rURL)
	var body io.Reader
	if reqBody != nil {
		bb, err := json.Marshal(reqBody)
		if err != nil {
			return errors.Wrap(err, "client - unable to marshal request struct")
		}
		body = bytes.NewReader(bb)
	}
	req, err := http.NewRequest(method, URL, body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return errors.Wrap(err, "client - unable to create request body")
	}

	res, err := s.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "client - failed to execute request")
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != 204 {
		return errors.Errorf("invalid status code received, expected 200/204, got %v", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return errors.Wrap(err, "unable to unmarshal response body")
	}
	return nil
}
