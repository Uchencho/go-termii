package gotermii

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// GetBalanceResponse is a representation of a get balance response
type GetBalanceResponse struct {
	User     string `json:"user"`
	Balance  int    `json:"balance"`
	Currency string `json:"currency"`
}

// VerifyNumberRequest is a representation of a verify phone number request
type VerifyNumberRequest struct {
	APIKey      string `json:"api_key"`
	PhoneNumber string `json:"phone_number"`
}

// VerifyNumberResponse is a representation of a verify phone number response
type VerifyNumberResponse struct {
	Number      string `json:"number"`
	Status      string `json:"status"`
	Network     string `json:"network"`
	NetworkCode string `json:"network_code"`
}

// GetBalance returns total balance and balance information from your wallet, such as currency.
// See docs https://developers.termii.com/balance for more details
func (c Client) GetBalance() (GetBalanceResponse, error) {
	rURL := fmt.Sprintf("api/get-balance?api_key=%s", c.config.APIKey)

	var Response GetBalanceResponse
	if err := c.makeRequest(http.MethodGet, rURL, nil, &Response); err != nil {
		return GetBalanceResponse{}, errors.Wrap(err, "error in making request to get balance")
	}
	return Response, nil
}

// VerifyNumber allows businesses verify phone numbers and automatically detect their status
// See docs https://developers.termii.com/search for more details
func (c Client) VerifyNumber(req VerifyNumberRequest) (VerifyNumberResponse, error) {
	rURL := "api/check/dnd"
	req.APIKey = c.config.APIKey

	var Response VerifyNumberResponse
	if err := c.makeRequest(http.MethodGet, rURL, req, &Response); err != nil {
		return VerifyNumberResponse{}, errors.Wrap(err, "error in making request to verify number")
	}
	return Response, nil
}
