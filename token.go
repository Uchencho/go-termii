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

// VerifyTokenRequest is a representation of a verify token request
type VerifyTokenRequest struct {
	APIKey string `json:"api_key"`
	PinID  string `json:"pin_id"`
	Pin    string `json:"pin"`
}

// VerifyTokenResponse is a representation of a verify token response
type VerifyTokenResponse struct {
	PinID    string `json:"pinId"`
	Verified string `json:"verified"`
	Msisdn   string `json:"msisdn"`
}

// SendTokenResponse is a representation of a send token response
type SendTokenResponse struct {
	PinID     string `json:"pinId"`
	To        string `json:"to"`
	SmsStatus string `json:"smsStatus"`
}

// SendToken sends a token request
func (c Client) SendToken(req SendTokenRequest) (SendTokenResponse, error) {
	req.APIKey = c.config.APIKey
	rURL := "api/sms/otp/send"

	var tokenResponse SendTokenResponse
	if err := c.makeRequest(http.MethodPost, rURL, req, &tokenResponse); err != nil {
		return SendTokenResponse{}, errors.Wrap(err, "error in making request to send otp token")
	}
	return tokenResponse, nil
}

// VerifyToken sends a request to verify token
func (c Client) VerifyToken(req VerifyTokenRequest) (VerifyTokenResponse, error) {
	req.APIKey = c.config.APIKey
	rURL := "api/sms/otp/verify"

	var tokenResponse VerifyTokenResponse
	if err := c.makeRequest(http.MethodPost, rURL, req, &tokenResponse); err != nil {
		return VerifyTokenResponse{}, errors.Wrap(err, "error in making request to verify otp token")
	}
	return tokenResponse, nil
}
