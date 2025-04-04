# httpi

A simple CLI tool written in Go for testing HTTP APIs. This project allows you to easily make GET and POST requests from the command line.

## Features

- Make GET, POST, PUT, and DELETE requests
- Custom headers support
- JSON body support for POST and PUT requests
- Simple and intuitive command-line interface

## Installation

### Method 1: Direct Install (Recommended)
```bash
# Install the binary
go install github.com/Soohyeuk/cli-api-tester/cmd/httpi@latest

# Add to PATH (if not already added)
# For bash/zsh:
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc  # or ~/.zshrc
source ~/.bashrc  # or source ~/.zshrc

# For tcsh:
echo 'setenv PATH ${PATH}:$(go env GOPATH)/bin' >> ~/.tcshrc
source ~/.tcshrc
```

### Method 2: Manual Build
```bash
# Clone the repository
git clone https://github.com/Soohyeuk/cli-api-tester.git
cd cli-api-tester

# Build and install the binary
go build -o $(go env GOPATH)/bin/httpi ./cmd/httpi

# Add to PATH (same as Method 1)
```## Usage

Basic syntax:
```bash
httpi <method> <url> [headers] [body]
```

### Examples

1. Simple GET request:
```bash
httpi GET https://api.example.com
```

2. GET request with headers:
```bash
httpi GET https://api.example.com "Authorization: Bearer token,Content-Type: application/json"
```

3. POST request with JSON body:
```bash
httpi POST https://api.example.com "Content-Type: application/json" '{"key": "value"}'
```

4. PUT request with JSON body:
```bash
httpi PUT https://api.example.com "Content-Type: application/json" '{"user": {"name": "John", "age": 30}}'
```

5. DELETE request:
```bash
httpi DELETE https://api.example.com
```

6. DELETE request with custom headers:
```bash
httpi DELETE https://api.example.com "Authorization: Bearer token"
```

### Testing with Local Server

The project includes a simple Flask server for testing. To use it:

1. Start the Flask server:
```bash
python3 test/main.py
```

2. Test with the CLI:
```bash
# GET request
httpi GET http://localhost:8080

# POST request
httpi POST http://localhost:8080/test "Content-Type: application/json" '{"message": "Hello"}'

# PUT request
httpi PUT http://localhost:8080/test "Content-Type: application/json" '{"message": "Updated"}'

# DELETE request
httpi DELETE http://localhost:8080/test
```

## Project Structure

```
.
├── cmd/
│   └── httpi
│       └── main.go          # CLI entry point
├── internal/
│   └── client/
│       ├── client.go    # HTTP client implementation
│       ├── config.go    # Configuration management
│       └── request.go   # Request handling
└── test/
    └── main.py         # Test server
```

## Contributing

Feel free to submit issues and enhancement requests! 



