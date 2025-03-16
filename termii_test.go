package gotermii_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	termii "github.com/joshuachinemezu/go-termii"

	"github.com/stretchr/testify/assert"
)

const (
	termiiTestApiKey = "test-API"
	termiiSenderID   = "Acme"
)

func fileToStruct(filepath string, s interface{}) io.Reader {
	bb, _ := ioutil.ReadFile(filepath)
	json.Unmarshal(bb, s)
	return bytes.NewReader(bb)
}

type row struct {
	name           string
	in             interface{}
	out            interface{}
	termiiResponse string
}

func TestSendTokenSuccess(t *testing.T) {
	os.Setenv("TERMII_API_KEY", termiiTestApiKey)
	var (
		expectedTokenRequest termii.SendTokenRequest
		receivedBody         termii.SendTokenRequest
		req                  termii.SendTokenRequest
		expectedResponse     termii.SendTokenResponse
	)

	termiiService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := json.NewDecoder(req.Body).Decode(&receivedBody); err != nil {
			log.Printf("error in unmarshalling %+v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := "/api/sms/otp/send"
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		t.Run("Request is as expected", func(t *testing.T) {
			fileToStruct(filepath.Join("testdata", "send_token_request.json"), &expectedTokenRequest)
			assert.Equal(t, expectedTokenRequest, receivedBody)
		})

		var resp termii.SendTokenResponse
		fileToStruct(filepath.Join("testdata", "send_token_response.json"), &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))
	os.Setenv("TERMII_URL", termiiService.URL)
	fileToStruct(filepath.Join("testdata", "send_token_request.json"), &req)

	c := termii.NewClient()

	resp, err := c.SendToken(req)
	t.Run("No error is returned", func(t *testing.T) {
		assert.NoError(t, err)
	})

	t.Run("Response is as expected", func(t *testing.T) {
		fileToStruct(filepath.Join("testdata", "send_token_response.json"), &expectedResponse)
		assert.Equal(t, expectedResponse, resp)
	})
}

func TestVerifyTokenSuccess(t *testing.T) {
	os.Setenv("TERMII_API_KEY", termiiTestApiKey)
	var (
		expectedTokenRequest termii.VerifyTokenRequest
		receivedBody         termii.VerifyTokenRequest
		req                  termii.VerifyTokenRequest
		expectedResponse     termii.VerifyTokenResponse
	)

	termiiService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := json.NewDecoder(req.Body).Decode(&receivedBody); err != nil {
			log.Printf("error in unmarshalling %+v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := "/api/sms/otp/verify"
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		t.Run("Request is as expected", func(t *testing.T) {
			fileToStruct(filepath.Join("testdata", "verify_token_request.json"), &expectedTokenRequest)
			assert.Equal(t, expectedTokenRequest, receivedBody)
		})

		var resp termii.VerifyTokenResponse
		fileToStruct(filepath.Join("testdata", "verify_token_response.json"), &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))
	os.Setenv("TERMII_URL", termiiService.URL)
	fileToStruct(filepath.Join("testdata", "verify_token_request.json"), &req)

	c := termii.NewClient()

	resp, err := c.VerifyToken(req)
	t.Run("No error is returned", func(t *testing.T) {
		assert.NoError(t, err)
	})

	t.Run("Response is as expected", func(t *testing.T) {
		fileToStruct(filepath.Join("testdata", "verify_token_response.json"), &expectedResponse)
		assert.Equal(t, expectedResponse, resp)
	})
}

func TestGetInAppTokenSuccess(t *testing.T) {
	os.Setenv("TERMII_API_KEY", termiiTestApiKey)
	var (
		expectedTokenRequest termii.GenerateTokenRequest
		receivedBody         termii.GenerateTokenRequest
		req                  termii.GenerateTokenRequest
		expectedResponse     termii.GenerateTokenResponse
	)

	termiiService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := json.NewDecoder(req.Body).Decode(&receivedBody); err != nil {
			log.Printf("error in unmarshalling %+v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := "/api/sms/otp/generate"
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		t.Run("Request is as expected", func(t *testing.T) {
			fileToStruct(filepath.Join("testdata", "generate_token_request.json"), &expectedTokenRequest)
			assert.Equal(t, expectedTokenRequest, receivedBody)
		})

		var resp termii.GenerateTokenResponse
		fileToStruct(filepath.Join("testdata", "generate_token_response.json"), &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))
	os.Setenv("TERMII_URL", termiiService.URL)
	fileToStruct(filepath.Join("testdata", "generate_token_request.json"), &req)

	c := termii.NewClient()

	resp, err := c.GetInAppToken(req)
	t.Run("No error is returned", func(t *testing.T) {
		assert.NoError(t, err)
	})

	t.Run("Response is as expected", func(t *testing.T) {
		fileToStruct(filepath.Join("testdata", "generate_token_response.json"), &expectedResponse)
		assert.Equal(t, expectedResponse, resp)
	})
}

func TestFetchSenderIDSuccess(t *testing.T) {
	os.Setenv("TERMII_API_KEY", termiiTestApiKey)
	var (
		expectedResponse termii.FetchSenderIdResponse
	)

	termiiService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := fmt.Sprintf("/api/sender-id?api_key=%s", termiiTestApiKey)
			assert.Equal(t, http.MethodGet, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		var resp termii.FetchSenderIdResponse
		fileToStruct(filepath.Join("testdata", "fetch_sender_id_response.json"), &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))
	os.Setenv("TERMII_URL", termiiService.URL)

	c := termii.NewClient()

	resp, err := c.FetchSenderID()
	t.Run("No error is returned", func(t *testing.T) {
		assert.NoError(t, err)
	})

	t.Run("Response is as expected", func(t *testing.T) {
		fileToStruct(filepath.Join("testdata", "fetch_sender_id_response.json"), &expectedResponse)
		assert.Equal(t, expectedResponse, resp)
	})
}

func TestRegisterSenderSuccess(t *testing.T) {
	os.Setenv("TERMII_API_KEY", termiiTestApiKey)
	os.Setenv("TERMII_SENDER_ID", termiiSenderID)
	var (
		expectedTokenRequest termii.RegisterSenderIdRequest
		receivedBody         termii.RegisterSenderIdRequest
		req                  termii.RegisterSenderIdRequest
		expectedResponse     termii.RegisterSenderResponse
	)

	termiiService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := json.NewDecoder(req.Body).Decode(&receivedBody); err != nil {
			log.Printf("error in unmarshalling %+v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := "/api/sender-id/request"
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		t.Run("Request is as expected", func(t *testing.T) {
			fileToStruct(filepath.Join("testdata", "register_sender_request.json"), &expectedTokenRequest)
			assert.Equal(t, expectedTokenRequest, receivedBody)
		})

		var resp termii.RegisterSenderResponse
		fileToStruct(filepath.Join("testdata", "register_sender_response.json"), &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))
	os.Setenv("TERMII_URL", termiiService.URL)
	fileToStruct(filepath.Join("testdata", "register_sender_request.json"), &req)

	c := termii.NewClient()

	resp, err := c.RegisterSender(req)
	t.Run("No error is returned", func(t *testing.T) {
		assert.NoError(t, err)
	})

	t.Run("Response is as expected", func(t *testing.T) {
		fileToStruct(filepath.Join("testdata", "register_sender_response.json"), &expectedResponse)
		assert.Equal(t, expectedResponse, resp)
	})
}

func TestSendMessageSuccess(t *testing.T) {
	os.Setenv("TERMII_API_KEY", termiiTestApiKey)
	var (
		expectedTokenRequest termii.SendMessageRequest
		receivedBody         termii.SendMessageRequest
		req                  termii.SendMessageRequest
		expectedResponse     termii.SendMessageResponse
	)

	termiiService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := json.NewDecoder(req.Body).Decode(&receivedBody); err != nil {
			log.Printf("error in unmarshalling %+v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := "/api/sms/send"
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		t.Run("Request is as expected", func(t *testing.T) {
			fileToStruct(filepath.Join("testdata", "send_message_request.json"), &expectedTokenRequest)
			assert.Equal(t, expectedTokenRequest, receivedBody)
		})

		var resp termii.SendMessageResponse
		fileToStruct(filepath.Join("testdata", "send_message_response.json"), &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))
	os.Setenv("TERMII_URL", termiiService.URL)
	fileToStruct(filepath.Join("testdata", "send_message_request.json"), &req)

	c := termii.NewClient()

	resp, err := c.SendMessage(req)
	t.Run("No error is returned", func(t *testing.T) {
		assert.NoError(t, err)
	})

	t.Run("Response is as expected", func(t *testing.T) {
		fileToStruct(filepath.Join("testdata", "send_message_response.json"), &expectedResponse)
		assert.Equal(t, expectedResponse, resp)
	})
}

func TestSendAutoGeneratedMessageSuccess(t *testing.T) {
	os.Setenv("TERMII_API_KEY", termiiTestApiKey)
	var (
		expectedTokenRequest termii.AutoGeneratedMessageRequest
		receivedBody         termii.AutoGeneratedMessageRequest
		req                  termii.AutoGeneratedMessageRequest
		expectedResponse     termii.AutoGeneratedMessageResponse
	)

	termiiService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := json.NewDecoder(req.Body).Decode(&receivedBody); err != nil {
			log.Printf("error in unmarshalling %+v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := "/api/sms/number/send"
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		t.Run("Request is as expected", func(t *testing.T) {
			fileToStruct(filepath.Join("testdata", "send_auto_message_request.json"), &expectedTokenRequest)
			assert.Equal(t, expectedTokenRequest, receivedBody)
		})

		var resp termii.AutoGeneratedMessageResponse
		bb, _ := ioutil.ReadFile(os.Getenv("termiiResponse"))
		if err := json.Unmarshal(bb, &resp); err != nil {
			log.Printf("error in marshalling %+v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		bbb, _ := json.Marshal(resp)
		w.Write(bbb)
	}))
	os.Setenv("TERMII_URL", termiiService.URL)

	c := termii.NewClient()

	table := []row{
		{
			name:           "Balance is integer and success is expected",
			in:             filepath.Join("testdata", "send_auto_message_request.json"),
			out:            filepath.Join("testdata", "send_auto_message_response.json"),
			termiiResponse: filepath.Join("testdata", "send_auto_message_response.json"),
		},
		{
			name:           "Balance is float and success is expected",
			in:             filepath.Join("testdata", "send_auto_message_request.json"),
			out:            filepath.Join("testdata", "send_auto_message_response_float.json"),
			termiiResponse: filepath.Join("testdata", "send_auto_message_response_float.json"),
		},
	}

	for _, entry := range table {
		req = termii.AutoGeneratedMessageRequest{}
		os.Setenv("termiiResponse", entry.termiiResponse)
		fileToStruct(entry.in.(string), &req)

		resp, err := c.SendAutoGeneratedMessage(req)
		t.Run(fmt.Sprintf("%s - No error is returned", entry.name), func(t *testing.T) {
			assert.NoError(t, err)
		})

		t.Run(fmt.Sprintf("%s - Response is as expected", entry.name), func(t *testing.T) {
			fileToStruct(entry.out.(string), &expectedResponse)
			assert.Equal(t, expectedResponse, resp)
		})
	}
}

func TestSetDeviceTemplateSuccess(t *testing.T) {
	os.Setenv("TERMII_API_KEY", termiiTestApiKey)
	var (
		expectedTokenRequest termii.TemplateRequest
		receivedBody         termii.TemplateRequest
		req                  termii.TemplateRequest
		expectedResponse     []termii.TemplateResponse
	)

	termiiService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := json.NewDecoder(req.Body).Decode(&receivedBody); err != nil {
			log.Printf("error in unmarshalling %+v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := "/api/send/template"
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		t.Run("Request is as expected", func(t *testing.T) {
			fileToStruct(filepath.Join("testdata", "device_template_request.json"), &expectedTokenRequest)
			assert.Equal(t, expectedTokenRequest, receivedBody)
		})

		var resp []termii.TemplateResponse
		fileToStruct(filepath.Join("testdata", "device_template_response.json"), &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))
	os.Setenv("TERMII_URL", termiiService.URL)
	fileToStruct(filepath.Join("testdata", "device_template_request.json"), &req)

	c := termii.NewClient()

	resp, err := c.SetDeviceTemplate(req)
	t.Run("No error is returned", func(t *testing.T) {
		assert.NoError(t, err)
	})

	t.Run("Response is as expected", func(t *testing.T) {
		fileToStruct(filepath.Join("testdata", "device_template_response.json"), &expectedResponse)
		assert.Equal(t, expectedResponse, resp)
	})
}

func TestGetBalanceSuccess(t *testing.T) {
	os.Setenv("TERMII_API_KEY", termiiTestApiKey)
	var (
		expectedResponse termii.GetBalanceResponse
	)

	termiiService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := fmt.Sprintf("/api/get-balance?api_key=%s", termiiTestApiKey)
			assert.Equal(t, http.MethodGet, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		var resp termii.GetBalanceResponse
		fileToStruct(filepath.Join("testdata", "get_balance_response.json"), &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))
	os.Setenv("TERMII_URL", termiiService.URL)

	c := termii.NewClient()

	resp, err := c.GetBalance()
	t.Run("No error is returned", func(t *testing.T) {
		assert.NoError(t, err)
	})

	t.Run("Response is as expected", func(t *testing.T) {
		fileToStruct(filepath.Join("testdata", "get_balance_response.json"), &expectedResponse)
		assert.Equal(t, expectedResponse, resp)
	})
}

func TestVerifyNumberSuccess(t *testing.T) {
	os.Setenv("TERMII_API_KEY", termiiTestApiKey)
	var (
		expectedTokenRequest termii.VerifyNumberRequest
		receivedBody         termii.VerifyNumberRequest
		req                  termii.VerifyNumberRequest
		expectedResponse     termii.VerifyNumberResponse
	)

	termiiService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := json.NewDecoder(req.Body).Decode(&receivedBody); err != nil {
			log.Printf("error in unmarshalling %+v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := "/api/check/dnd"
			assert.Equal(t, http.MethodGet, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		t.Run("Request is as expected", func(t *testing.T) {
			fileToStruct(filepath.Join("testdata", "verify_number_request.json"), &expectedTokenRequest)
			assert.Equal(t, expectedTokenRequest, receivedBody)
		})

		var resp termii.VerifyNumberResponse
		fileToStruct(filepath.Join("testdata", "verify_number_response.json"), &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))
	os.Setenv("TERMII_URL", termiiService.URL)
	fileToStruct(filepath.Join("testdata", "verify_number_request.json"), &req)

	c := termii.NewClient()

	resp, err := c.VerifyNumber(req)
	t.Run("No error is returned", func(t *testing.T) {
		assert.NoError(t, err)
	})

	t.Run("Response is as expected", func(t *testing.T) {
		fileToStruct(filepath.Join("testdata", "verify_number_response.json"), &expectedResponse)
		assert.Equal(t, expectedResponse, resp)
	})
}

func TestGetStatusSuccess(t *testing.T) {
	os.Setenv("TERMII_API_KEY", termiiTestApiKey)
	var (
		expectedTokenRequest termii.StatusRequest
		receivedBody         termii.StatusRequest
		req                  termii.StatusRequest
		expectedResponse     termii.StatusResponse
	)

	termiiService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := json.NewDecoder(req.Body).Decode(&receivedBody); err != nil {
			log.Printf("error in unmarshalling %+v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := "/api/insight/number/query"
			assert.Equal(t, http.MethodGet, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		t.Run("Request is as expected", func(t *testing.T) {
			fileToStruct(filepath.Join("testdata", "get_status_request.json"), &expectedTokenRequest)
			assert.Equal(t, expectedTokenRequest, receivedBody)
		})

		var resp termii.StatusResponse
		fileToStruct(filepath.Join("testdata", "get_status_response.json"), &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))
	os.Setenv("TERMII_URL", termiiService.URL)
	fileToStruct(filepath.Join("testdata", "get_status_request.json"), &req)

	c := termii.NewClient()

	resp, err := c.GetStatus(req)
	t.Run("No error is returned", func(t *testing.T) {
		assert.NoError(t, err)
	})

	t.Run("Response is as expected", func(t *testing.T) {
		fileToStruct(filepath.Join("testdata", "get_status_response.json"), &expectedResponse)
		assert.Equal(t, expectedResponse, resp)
	})
}

func TestGetHistorySuccess(t *testing.T) {
	os.Setenv("TERMII_API_KEY", termiiTestApiKey)
	var (
		expectedResponse []termii.HistoryResponse
	)

	termiiService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		t.Run("URL and request method is as expected", func(t *testing.T) {
			expectedURL := fmt.Sprintf("/api/sms/inbox?api_key=%s", termiiTestApiKey)
			assert.Equal(t, http.MethodGet, req.Method)
			assert.Equal(t, expectedURL, req.RequestURI)
		})

		var resp []termii.HistoryResponse
		fileToStruct(filepath.Join("testdata", "get_history_response.json"), &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))
	os.Setenv("TERMII_URL", termiiService.URL)

	c := termii.NewClient()
	resp, err := c.GetHistory()
	t.Run("No error is returned", func(t *testing.T) {
		assert.NoError(t, err)
	})

	t.Run("Response is as expected", func(t *testing.T) {
		fileToStruct(filepath.Join("testdata", "get_history_response.json"), &expectedResponse)
		assert.Equal(t, expectedResponse, resp)
	})
}
