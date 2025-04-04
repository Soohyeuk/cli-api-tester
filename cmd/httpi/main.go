// Package httpi is a simple CLI tool for testing HTTP APIs
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Soohyeuk/cli-api-tester/internal/client"
	"github.com/Soohyeuk/cli-api-tester/internal/config"
)

// main is the entry point of the application
// It parses command line flags and routes to appropriate handlers
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Hi, I'm httpi, a simple CLI tool for testing HTTP APIs. Try 'httpi --help' or 'httpi -h' for more information.")
		os.Exit(1)
	}

	if len(os.Args) == 2 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		fmt.Println("Usage: httpi <method> <url> [headers] [body]")
		fmt.Println("Methods Allowed: GET, POST, PUT, DELETE")
		fmt.Println("Headers should be in the format 'Header1: value1, Header2: value2'")
		fmt.Println("Body should be in the format 'key1=value1&key2=value2' for GET and POST requests")
		fmt.Println("Body should be in the format '{\"key1\": \"value1\", \"key2\": \"value2\"}' for PUT and DELETE requests")
		os.Exit(1)
	}

	method := os.Args[1]
	url := os.Args[2]
	headers := make(map[string]string)
	body := ""

	if len(os.Args) > 3 {
		headers = parseHeaders(os.Args[3])
	}

	if len(os.Args) > 4 {
		body = os.Args[4]
	}

	var cfg *config.Config = config.NewConfig()
	var httpClient *client.HTTPClient = client.NewHTTPClient(cfg)

	var resp *http.Response
	var err error

	switch method {
	case "GET":
		resp, err = handleGetRequest(httpClient, url, headers)
	case "POST":
		resp, err = handlePostRequest(httpClient, url, headers, body)
	case "PUT":
		resp, err = handlePutRequest(httpClient, url, headers, body)
	case "DELETE":
		resp, err = handleDeleteRequest(httpClient, url, headers)
	default:
		fmt.Println("Unsupported method:", method)
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bodyStr, err := client.ReadResponseBody(resp)
	if err != nil {
		fmt.Println("Error reading response:", err)
		os.Exit(1)
	}
	fmt.Println("Response:", bodyStr)
}

// parseHeaders converts a string of headers into a map
// Format: "Header1: value1, Header2: value2"
func parseHeaders(headerStr string) map[string]string {
	headers := make(map[string]string)
	if headerStr == "" {
		return headers
	}

	headerPairs := strings.Split(headerStr, ",")
	for _, pair := range headerPairs {
		parts := strings.SplitN(pair, ":", 2)
		if len(parts) == 2 {
			headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	return headers
}

// handleGetRequest processes GET requests
// Parameters:
// - client: *client.HTTPClient - The configured HTTP client
// - url: string - The URL to send GET request to
// - headers: map[string]string - Optional headers to include
// Returns:
// - *http.Response: The response from the server
// - error: Any error that occurred during the request
func handleGetRequest(client *client.HTTPClient, url string, headers map[string]string) (*http.Response, error) {
	req, err := client.CreateGetRequest(url, headers)
	if err != nil {
		return nil, err
	}
	return client.SendRequest(req)
}

// handlePostRequest processes POST requests
// Parameters:
// - client: *client.HTTPClient - The configured HTTP client
// - url: string - The URL to send POST request to
// - headers: map[string]string - Optional headers to include
// - body: string - The request body
// Returns:
// - *http.Response: The response from the server
// - error: Any error that occurred during the request
func handlePostRequest(client *client.HTTPClient, url string, headers map[string]string, body string) (*http.Response, error) {
	req, err := client.CreatePostRequest(url, headers, body)
	if err != nil {
		return nil, err
	}
	return client.SendRequest(req)
}

// handlePutRequest processes PUT requests
// Parameters:
// - client: *client.HTTPClient - The configured HTTP client
// - url: string - The URL to send PUT request to
// - headers: map[string]string - Optional headers to include
// - body: string - The request body
// Returns:
// - *http.Response: The response from the server
// - error: Any error that occurred during the request
func handlePutRequest(client *client.HTTPClient, url string, headers map[string]string, body string) (*http.Response, error) {
	req, err := client.CreatePutRequest(url, headers, body)
	if err != nil {
		return nil, err
	}
	return client.SendRequest(req)
}

// handleDeleteRequest processes DELETE requests
// Parameters:
// - client: *client.HTTPClient - The configured HTTP client
// - url: string - The URL to send DELETE request to
// - headers: map[string]string - Optional headers to include
// Returns:
// - *http.Response: The response from the server
// - error: Any error that occurred during the request
func handleDeleteRequest(client *client.HTTPClient, url string, headers map[string]string) (*http.Response, error) {
	req, err := client.CreateDeleteRequest(url, headers)
	if err != nil {
		return nil, err
	}
	return client.SendRequest(req)
}
