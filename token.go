package gotermii

import (
	"net/http"

	"github.com/pkg/errors"
)

// SendTokenRequest is a representation of a send token request
type SendTokenRequest struct {
	APIKey         string `json:"api_key"`
	MessageType    string `json:"message_type"`
	To             string `json:"to"`
	From           string `json:"from"`
	Channel        string `json:"channel"`
	PinAttempts    int    `json:"pin_attempts"`
	PinTimeToLive  int    `json:"pin_time_to_live"`
	PinLength      int    `json:"pin_length"`
	PinPlaceholder string `json:"pin_placeholder"`
	MessageText    string `json:"message_text"`
	PinType        string `json:"pin_type"`
}

// SendTokenResponse is a representation of a send token response
type SendTokenResponse struct {
	PinID     string `json:"pinId"`
	To        string `json:"to"`
	SmsStatus string `json:"smsStatus"`
}

// SendTokenFunc provides the functionality of sending a token request
type SendTokenFunc func(req SendTokenRequest) (SendTokenResponse, error)

// SendToken sends a token request
func SendToken(c Client) SendTokenFunc {
	return func(req SendTokenRequest) (SendTokenResponse, error) {
		req.APIKey = c.config.APIKey
		rURL := "api/sms/otp/send"

		var tokenResponse SendTokenResponse
		if err := c.makeRequest(http.MethodPost, rURL, req, &tokenResponse); err != nil {
			return SendTokenResponse{}, errors.Wrap(err, "error in making request to send otp token")
		}
		return tokenResponse, nil
	}
}
