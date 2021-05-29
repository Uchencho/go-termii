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
