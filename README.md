# httpi

A simple CLI tool written in Go for testing HTTP APIs. This project allows you to easily make GET and POST requests from the command line.

## Features

- Make GET and POST requests
- Custom headers support
- JSON body support for POST requests
- Simple and intuitive command-line interface

## Installation

### Quick Install (Recommended)
```bash
# Download and run the installation script
curl -fsSL https://raw.githubusercontent.com/soohyeuk/cli-api-tester/main/install.sh | bash
```

Or if you prefer to clone the repository first:
```bash
git clone https://github.com/soohyeuk/cli-api-tester.git
cd cli-api-tester
./install.sh
```

The script will:
1. Install the `httpi` command
2. Add it to your PATH automatically
3. Show you how to use it

### Manual Installation

1. Clone the repository:
```bash
git clone https://github.com/soohyeuk/cli-api-tester.git
cd cli-api-tester
```

2. Install globally (choose one method):

   a. Using Go install:
   ```bash
   go install ./cmd
   ```
   This will install the binary as `httpi` in your `$GOPATH/bin`

   b. Using make (if you have make installed):
   ```bash
   make install
   ```

   c. Manual installation:
   ```bash
   go build -o $GOPATH/bin/httpi ./cmd
   ```

3. Add to PATH (if not already added):
   ```bash
   export PATH=$PATH:$(go env GOPATH)/bin
   ```
   Add this line to your ~/.bashrc, ~/.zshrc, or equivalent shell config file.

## Usage

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
