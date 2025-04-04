// Package httpi is a simple CLI tool for testing HTTP APIs
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Soohyeuk/cli-api-tester/internal/client"
	"github.com/Soohyeuk/cli-api-tester/internal/config"
	"github.com/Soohyeuk/cli-api-tester/internal/version"
	"github.com/fatih/color"
)

// main is the entry point of the application
// It parses command line flags and routes to appropriate handlers
func main() {
	// Create color objects
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()

	if len(os.Args) < 2 {
		fmt.Printf("Hi, I'm %s, a simple CLI tool for testing HTTP APIs. Try '%s --help' or '%s -h' for more information.\n",
			cyan("httpi"), cyan("httpi"), cyan("httpi"))
		os.Exit(1)
	}

	// Check for help flag
	if os.Args[1] == "--help" || os.Args[1] == "-h" {
		fmt.Printf("Usage: %s <method> <url> [headers] [body]\n", cyan("httpi"))
		fmt.Printf("Methods Allowed: %s, %s, %s, %s\n",
			green("GET"), green("POST"), green("PUT"), green("DELETE"))
		fmt.Printf("Headers should be in the format %s\n", yellow("'Header1: value1, Header2: value2'"))
		fmt.Printf("Body should be in the format %s for GET and POST requests\n",
			yellow("'key1=value1&key2=value2'"))
		fmt.Printf("Body should be in the format %s for PUT and DELETE requests\n",
			yellow("'{\"key1\": \"value1\", \"key2\": \"value2\"}'"))
		os.Exit(0)
	}

	// Check for version flag
	if os.Args[1] == "--version" || os.Args[1] == "-v" {
		fmt.Printf("%s version %s\n", cyan("httpi"), blue(version.Version))
		os.Exit(0)
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

	fmt.Printf("Sending %s request to %s...\n", magenta(method), cyan(url))

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
		fmt.Printf("%s: Unsupported method %s\n", red("Error"), red(method))
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("%s: %s\n", red("Error"), red(err))
		os.Exit(1)
	}

	bodyStr, err := client.ReadResponseBody(resp)
	if err != nil {
		fmt.Printf("%s reading response: %s\n", red("Error"), red(err))
		os.Exit(1)
	}

	fmt.Printf("%s: %s\n", green("Response"), bodyStr)
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
