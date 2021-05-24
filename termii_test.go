package gotermii_test

import (
	"bytes"
	"encoding/json"
	gotermii "go-termii"
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

func fileToStruct(filepath string, s interface{}) io.Reader {
	bb, _ := ioutil.ReadFile(filepath)
	json.Unmarshal(bb, s)
	return bytes.NewReader(bb)
}

func TestSendTokenSuccess(t *testing.T) {
	os.Setenv("TERMII_API_KEY", "test-API")
	var (
		expectedTokenRequest gotermii.SendTokenRequest
		receivedBody         gotermii.SendTokenRequest
		req                  gotermii.SendTokenRequest
		expectedResponse     gotermii.SendTokenResponse
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

		var resp gotermii.SendTokenResponse
		fileToStruct(filepath.Join("testdata", "send_token_response.json"), &resp)

		w.WriteHeader(http.StatusOK)
		bb, _ := json.Marshal(resp)
		w.Write(bb)
	}))
	os.Setenv("TERMII_URL", termiiService.URL)
	fileToStruct(filepath.Join("testdata", "send_token_request.json"), &req)

	c := gotermii.NewClient()

	sendToken := gotermii.SendToken(c)
	resp, err := sendToken(req)
	t.Run("No error is returned", func(t *testing.T) {
		assert.NoError(t, err)
	})

	t.Run("Response is as expected", func(t *testing.T) {
		fileToStruct(filepath.Join("testdata", "send_token_response.json"), &expectedResponse)
		assert.Equal(t, expectedResponse, resp)
	})
}
