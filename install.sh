#!/bin/bash

# Get the GOPATH
GOPATH=$(go env GOPATH)

# Install the binary
echo "Installing httpi..."
go install ./cmd

# Check if the installation was successful
if [ $? -eq 0 ]; then
    echo "✅ httpi installed successfully at $GOPATH/bin/httpi"
else
    echo "❌ Installation failed"
    exit 1
fi

# Check if GOPATH/bin is already in PATH
if [[ ":$PATH:" != *":$GOPATH/bin:"* ]]; then
    # Detect shell config file
    if [ -n "$ZSH_VERSION" ]; then
        SHELL_RC="$HOME/.zshrc"
    elif [ -n "$BASH_VERSION" ]; then
        SHELL_RC="$HOME/.bashrc"
    else
        SHELL_RC="$HOME/.profile"
    fi

    # Add GOPATH/bin to PATH if not already present
    if ! grep -q "export PATH=\$PATH:\$GOPATH/bin" "$SHELL_RC"; then
        echo "Adding GOPATH/bin to PATH in $SHELL_RC"
        echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> "$SHELL_RC"
        echo "✅ Added to PATH. Please restart your terminal or run: source $SHELL_RC"
    else
        echo "✅ GOPATH/bin is already in your PATH"
    fi
else
    echo "✅ GOPATH/bin is already in your PATH"
fi

echo "
🎉 Installation complete! You can now use 'httpi' from anywhere.

Try it out:
  httpi GET http://localhost:8080
" 