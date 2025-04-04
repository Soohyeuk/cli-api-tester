package client

import (
	"bytes"
	"io"
	"net/http"
)

// CreateGetRequest creates a new GET request
// Parameters:
// - url: string - The URL to send GET request to
// - headers: map[string]string - Optional headers to include
// Returns:
// - *http.Request: The created GET request
// - error: Any error that occurred during request creation
func (c *HTTPClient) CreateGetRequest(url string, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set headers if provided
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

// CreatePostRequest creates a new POST request
// Parameters:
// - url: string - The URL to send POST request to
// - headers: map[string]string - Optional headers to include
// - body: string - The request body
// Returns:
// - *http.Request: The created POST request
// - error: Any error that occurred during request creation
func (c *HTTPClient) CreatePostRequest(url string, headers map[string]string, body string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}

	// Set headers if provided
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

// ReadResponseBody reads and returns the response body
// Parameters:
// - resp: *http.Response - The response to read body from
// Returns:
// - string: The response body as string
// - error: Any error that occurred during reading
func ReadResponseBody(resp *http.Response) (string, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return string(body), nil
}
