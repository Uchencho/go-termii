package gotermii_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	termii "go-termii"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	termiiTestApiKey = "test-API"
)

func fileToStruct(filepath string, s interface{}) io.Reader {
	bb, _ := ioutil.ReadFile(filepath)
	json.Unmarshal(bb, s)
	return bytes.NewReader(bb)
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
		expectedResponse termii.FetchSenderIDResponse
	)

	termiiService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		t.Run("URL is as expected", func(t *testing.T) {
			expectedURL := fmt.Sprintf("/api/sender-id?api_key=%s", termiiTestApiKey)

			assert.Equal(t, expectedURL, req.RequestURI)
		})

		var resp termii.FetchSenderIDResponse
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
