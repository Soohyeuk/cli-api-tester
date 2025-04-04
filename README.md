# cli-api-tester

A simple CLI tool written in Go for testing HTTP APIs. This project allows you to easily make GET and POST requests from the command line.

## Features

- Make GET and POST requests
- Custom headers support
- JSON body support for POST requests
- Simple and intuitive command-line interface

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/cli-api-tester.git
cd cli-api-tester
```

2. Build the project:
```bash
go build -o api ./cmd
```

3. Add to PATH (optional but recommended):
```bash
# Create a bin directory in your home folder if it doesn't exist
mkdir -p ~/bin

# Copy the executable to your bin directory
cp api ~/bin/

# Add ~/bin to your PATH
# For bash/zsh users, add this to ~/.bashrc or ~/.zshrc:
# export PATH=$PATH:~/bin
# For tcsh users, add this to ~/.tcshrc:
# set path = ($path ~/bin)
```

## Usage

Basic syntax:
```bash
api <method> <url> [headers] [body]
```

### Examples

1. Simple GET request:
```bash
api GET https://api.example.com
```

2. GET request with headers:
```bash
api GET https://api.example.com "Authorization: Bearer token,Content-Type: application/json"
```

3. POST request with JSON body:
```bash
api POST https://api.example.com "Content-Type: application/json" '{"key": "value"}'
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
api GET http://localhost:8080

# POST request
api POST http://localhost:8080/test "Content-Type: application/json" '{"message": "Hello"}'
```

## Project Structure

```
.
├── cmd/
│   └── main.go          # CLI entry point
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
