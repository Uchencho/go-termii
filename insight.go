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

// StatusRequest is a representation of a status request
type StatusRequest struct {
	APIKey      string `json:"api_key"`
	PhoneNumber string `json:"phone_number"`
	CountryCode string `json:"country_code"`
}

type RouteDetail struct {
	Number string `json:"number"`
	Ported int    `json:"ported"`
}

type CountryDetail struct {
	CountryCode       string `json:"countryCode"`
	MobileCountryCode string `json:"mobileCountryCode"`
	Iso               string `json:"iso"`
}

type OperatorDetail struct {
	OperatorCode              string `json:"operatorCode"`
	OperatorName              string `json:"operatorName"`
	MobileNumberCode          string `json:"mobileNumberCode"`
	MobileRoutingCode         string `json:"mobileRoutingCode"`
	CarrierIdentificationCode string `json:"carrierIdentificationCode"`
	LineType                  string `json:"lineType"`
}

type StatusResult struct {
	RouteDetail    RouteDetail    `json:"routeDetail"`
	CountryDetail  CountryDetail  `json:"countryDetail"`
	OperatorDetail OperatorDetail `json:"operatorDetail"`
	Status         int            `json:"status"`
}

// StatusResponse is a representation of a status response
type StatusResponse struct {
	Result []StatusResult `json:"result"`
}

// HistoryResponse is a representation of a get history response
type HistoryResponse struct {
	Sender    string      `json:"sender"`
	Receiver  string      `json:"receiver"`
	Message   string      `json:"message"`
	Amount    int         `json:"amount"`
	Reroute   int         `json:"reroute"`
	Status    string      `json:"status"`
	SmsType   string      `json:"sms_type"`
	SendBy    string      `json:"send_by"`
	MediaURL  interface{} `json:"media_url"`
	MessageID string      `json:"message_id"`
	NotifyURL interface{} `json:"notify_url"`
	NotifyID  interface{} `json:"notify_id"`
	CreatedAt string      `json:"created_at"`
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

// GetStatus allows businesses to detect if a number is fake or has ported to a new network.
// See docs https://developers.termii.com/status for more details
func (c Client) GetStatus(req StatusRequest) (StatusResponse, error) {
	rURL := "api/insight/number/query"
	req.APIKey = c.config.APIKey

	var Response StatusResponse
	if err := c.makeRequest(http.MethodGet, rURL, req, &Response); err != nil {
		return StatusResponse{}, errors.Wrap(err, "error in making request to get status")
	}
	return Response, nil
}

// GetHistory returns reports for messages sent across the sms, voice & whatsapp channels.
// See docs https://developers.termii.com/history for more details
func (c Client) GetHistory() ([]HistoryResponse, error) {
	rURL := fmt.Sprintf("api/sms/inbox?api_key=%s", c.config.APIKey)

	var Response []HistoryResponse
	if err := c.makeRequest(http.MethodGet, rURL, nil, &Response); err != nil {
		return []HistoryResponse{}, errors.Wrap(err, "error in making request to get history")
	}
	return Response, nil
}
