package client

import (
	"net/http"
	"time"

	"github.com/soohyeuk/cli-api-tester/internal/config"
)

// HTTPClient represents a configured HTTP client
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient creates a new HTTP client with configuration
// Parameters:
// - cfg: *config.Config - Configuration containing timeout and other settings
// Returns:
// - *HTTPClient: A configured HTTP client instance
func NewHTTPClient(cfg *config.Config) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: time.Duration(cfg.Timeout) * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:    10,
				IdleConnTimeout: time.Duration(cfg.Timeout) * time.Second,
			},
		},
	}
}

// SetHeaders sets headers for an HTTP request
// Parameters:
// - req: *http.Request - The request to set headers for
// - headers: map[string]string - Headers to set
func (c *HTTPClient) SetHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

// SendRequest sends an HTTP request and returns the response
// Parameters:
// - req: *http.Request - The request to send
// Returns:
// - *http.Response: The response from the server
// - error: Any error that occurred during the request
func (c *HTTPClient) SendRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
