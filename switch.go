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

// Media is a representation of a media request
type Media struct {
	URL     string `json:"url"`
	Caption string `json:"caption"`
}

// AutoGeneratedMessageRequest is a representation of a send message request with an auto generated "source" number
type AutoGeneratedMessageRequest struct {
	To     string `json:"to"`
	Sms    string `json:"sms"`
	APIKey string `json:"api_key"`
}

// AutoGeneratedMessageResponse is a representation of a send message response from an auto generated "source" number
type AutoGeneratedMessageResponse struct {
	Code      string `json:"code"`
	MessageID string `json:"message_id"`
	Message   string `json:"message"`
	Balance   float32    `json:"balance"`  //the balance returned is a float
	User      string `json:"user"`
}

// SendMesageRequest is a representation of a send message request
type SendMessageRequest struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Sms     string `json:"sms"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	APIKey  string `json:"api_key"`
	Media   *Media  `json:"media,omitempty"`  //so the it can have a nil value on requests in which it ought to be absent
}

// SendMessageResponse is a representation of a send message response
type SendMessageResponse struct {
	MessageID string `json:"message_id"`
	Message   string `json:"message"`
	Balance   float32    `json:"balance"`
	User      string `json:"user"`
}

// TemplateData is a representation of a template data
type TemplateData struct {
	ProductName string `json:"product_name"`
	Otp         int    `json:"otp"`
	ExpiryTime  string `json:"expiry_time"`
}

// TemplateRequest is a representation of a template request
type TemplateRequest struct {
	PhoneNumber string       `json:"phone_number"`
	DeviceID    string       `json:"device_id"`
	TemplateID  string       `json:"template_id"`
	APIKey      string       `json:"api_key"`
	Data        TemplateData `json:"data"`
}

// TemplateResponse is a representation of a set template response
type TemplateResponse struct {
	Code      string `json:"code"`
	MessageID string `json:"message_id"`
	Message   string `json:"message"`
	Balance   string `json:"balance"`
	User      string `json:"user"`
}

// FetchSenderID allows businesses retrieve the status of all registered sender ID
// See docs https://developers.termii.com/sender-id#fetch-sender-id for more details
func (c Client) FetchSenderID() (FetchSenderIdResponse, error) {
	rURL := fmt.Sprintf("api/sender-id?api_key=%s", c.config.APIKey)

	var Response FetchSenderIdResponse
	if err := c.makeRequest(http.MethodGet, rURL, nil, &Response); err != nil {
		return FetchSenderIdResponse{}, errors.Wrap(err, "error in making request to fetch sender id")
	}
	return Response, nil
}

// RegisterSender allows businesses register a sender.
// See docs https://developers.termii.com/sender-id#request-sender-id for more details
func (c Client) RegisterSender(req RegisterSenderIdRequest) (RegisterSenderResponse, error) {
	rURL := "api/sender-id/request"
	req.APIKey = c.config.APIKey
	req.SenderID = c.config.SenderID

	var Response RegisterSenderResponse
	if err := c.makeRequest(http.MethodPost, rURL, req, &Response); err != nil {
		return RegisterSenderResponse{}, errors.Wrap(err, "error in making request to register sender")
	}
	return Response, nil
}

// SendSMS allows a business to send sms. See docs https://developers.termii.com/messaging for more details
func (c Client) SendMessage(req SendMessageRequest) (SendMessageResponse, error) {
	rURL := "api/sms/send"
	req.APIKey = c.config.APIKey

	var Response SendMessageResponse
	if err := c.makeRequest(http.MethodPost, rURL, req, &Response); err != nil {
		return SendMessageResponse{}, errors.Wrap(err, "error in making request to send message")
	}
	return Response, nil
}

// SendAutoGeneratedMessage allows businesses send messages to customers using auto-generated messaging numbers.
// See docs https://developers.termii.com/number for more details
func (c Client) SendAutoGeneratedMessage(req AutoGeneratedMessageRequest) (AutoGeneratedMessageResponse, error) {
	rURL := "api/sms/number/send"
	req.APIKey = c.config.APIKey

	var Response AutoGeneratedMessageResponse
	if err := c.makeRequest(http.MethodPost, rURL, req, &Response); err != nil {
		return AutoGeneratedMessageResponse{}, errors.Wrap(err, "error in making request to send message from an auto generated number")
	}
	return Response, nil
}

// Templates is a feature used to set a template for the one-time-passwords (pins) sent to their customers via whatsapp or sms.
// See docs https://developers.termii.com/templates for more details
func (c Client) SetDeviceTemplate(req TemplateRequest) ([]TemplateResponse, error) {
	rURL := "api/send/template"
	req.APIKey = c.config.APIKey

	var Response []TemplateResponse
	if err := c.makeRequest(http.MethodPost, rURL, req, &Response); err != nil {
		return []TemplateResponse{}, errors.Wrap(err, "error in making request to set device template")
	}
	return Response, nil
}
