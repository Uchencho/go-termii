package gotermii

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

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

// ConfigFromEnvVars provides the default config from env vars for termii
func ConfigFromEnvVars() Config {
	return Config{
		APIKey:   os.Getenv("TERMII_API_KEY"),
		BaseURL:  os.Getenv("TERMII_URL"),
		SenderID: os.Getenv("TERMII_SENDER_ID"),
	}
}

// NewClient creates a termii client using config variables
func NewClient() Client {
	return Client{config: ConfigFromEnvVars(), client: &http.Client{Timeout: 30 * time.Second}}
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

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return errors.Wrap(err, "unable to unmarshal request body")
	}
	return nil
}
