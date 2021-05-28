package gotermii

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// FetchSenderIdData is a representation of a fetch senderId data nested object
type FetchSenderIdData struct {
	SenderID  string      `json:"sender_id"`
	Status    string      `json:"status"`
	Company   string      `json:"company"`
	Usecase   interface{} `json:"usecase"`
	Country   interface{} `json:"country"`
	CreatedAt string      `json:"created_at"`
}

// FetchSenderIdResponse is a representation of a fetch senderID response
type FetchSenderIdResponse struct {
	CurrentPage  int                 `json:"current_page"`
	Data         []FetchSenderIdData `json:"data"`
	FirstPageURL string              `json:"first_page_url"`
	From         int                 `json:"from"`
	LastPage     int                 `json:"last_page"`
	LastPageURL  string              `json:"last_page_url"`
	NextPageURL  string              `json:"next_page_url"`
	Path         string              `json:"path"`
	PerPage      int                 `json:"per_page"`
	PrevPageURL  interface{}         `json:"prev_page_url"`
	To           int                 `json:"to"`
	Total        int                 `json:"total"`
}

// RegisterSenderIdRequest is a representation of a register sender reuest
type RegisterSenderIdRequest struct {
	APIKey   string `json:"api_key"`
	SenderID string `json:"sender_id"`
	Usecase  string `json:"usecase"`
	Company  string `json:"company"`
}

// RegisterSenderResponse is a repreentation of a register sender response
type RegisterSenderResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// FetchSenderID allows businesses retrieve the status of all registered sender ID
func (c Client) FetchSenderID() (FetchSenderIdResponse, error) {
	rURL := fmt.Sprintf("api/sender-id?api_key=%s", c.config.APIKey)

	var Response FetchSenderIdResponse
	if err := c.makeRequest(http.MethodGet, rURL, nil, &Response); err != nil {
		return FetchSenderIdResponse{}, errors.Wrap(err, "error in making request to fetch sender id")
	}
	return Response, nil
}

// RegisterSender allows businesses register a sender
func (c Client) RegisterSender(req RegisterSenderIdRequest) (RegisterSenderResponse, error) {
	rURL := "api/sender-id/request"
	req.APIKey = c.config.APIKey

	var Response RegisterSenderResponse
	if err := c.makeRequest(http.MethodPost, rURL, req, &Response); err != nil {
		return RegisterSenderResponse{}, errors.Wrap(err, "error in making request to register sender")
	}
	return Response, nil
}
