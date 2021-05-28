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

// FetchSenderIDResponse is a representation of a fetch senderID response
type FetchSenderIDResponse struct {
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

// FetchSenderID allows businesses retrieve the status of all registered sender ID
func (c Client) FetchSenderID() (FetchSenderIDResponse, error) {
	rURL := fmt.Sprintf("api/sender-id?api_key=%s", c.config.APIKey)

	var Response FetchSenderIDResponse
	if err := c.makeRequest(http.MethodGet, rURL, nil, &Response); err != nil {
		return FetchSenderIDResponse{}, errors.Wrap(err, "error in making request to generate token")
	}
	return Response, nil
}
