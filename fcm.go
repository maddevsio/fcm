package fcm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// PriorityHigh used for high notification priority
	PriorityHigh = "high"

	// PriorityNormal used for normal notification priority
	PriorityNormal = "normal"

	// HeaderRetryAfter HTTP header constant
	HeaderRetryAfter = "Retry-After"

	// ErrorKey readable error caching
	ErrorKey = "error"

	// MethodPOST indicates http post method
	MethodPOST = "POST"
)

var (
	// retryableErrors whether the error is a retryable
	retryableErrors = map[string]bool{
		"Unavailable":         true,
		"InternalServerError": true,
	}

	// fcmServerUrl for testing purposes
	fcmServerURL = "https://fcm.googleapis.com/fcm/send"
)

// FCM  stores client with api key to firebase
type FCM struct {
	APIKey             string
	AuthorizationToken string
}

// NewFCM creates a new client
func NewFCM(apiKey string) *FCM {
	return &FCM{
		APIKey:             apiKey,
		AuthorizationToken: fmt.Sprintf("key=%v", apiKey),
	}

}

// Send message to FCM
func (f *FCM) Send(message *Message) (*Response, error) {
	data, err := json.Marshal(*message)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(MethodPOST, fcmServerURL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", f.AuthorizationToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%d status code", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response := &Response{}
	if err := json.Unmarshal(body, response); err != nil {
		return nil, err
	}
	response.StatusCode = resp.StatusCode

	response.RetryAfter = resp.Header.Get(HeaderRetryAfter)

	if err := f.Failed(response); err != nil {
		return nil, err
	}
	response.Ok = true
	return response, nil

}

// Failed method incicates if the server couldn't process
// the request in time.
func (f *FCM) Failed(response *Response) error {
	for _, val := range response.Results {
		for k, v := range val {
			if k == ErrorKey && retryableErrors[v] {
				return fmt.Errorf("Failed %s", k)
			}
		}
	}
	return nil
}
