#!/bin/bash

# Development setup script for Library of Bears

echo "🐻 Library of Bears Development Setup"
echo "=====================================

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21+ first."
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | cut -d' ' -f3 | cut -d'o' -f2)
echo "✅ Go version: $GO_VERSION"

# Install dependencies
echo "📦 Installing dependencies..."
go mod tidy

# Build the application
echo "🔨 Building application..."
go build .
if [ $? -eq 0 ]; then
    echo "✅ Build successful!"
else
    echo "❌ Build failed!"
    exit 1
fi

# Run tests
echo "🧪 Running tests..."
go test -v
if [ $? -eq 0 ]; then
    echo "✅ All tests passed!"
else
    echo "⚠️  Some tests failed, but that's okay for development."
fi

echo "
🚀 Setup complete! 

To start the server:
  ./BearLibrary

Environment variables you might want to set:
  export DB_HOST=localhost
  export DB_USER=root
  export DB_PORT=3307
  export DB_NAME=library_of_bears
  export SECRET=your-secret-key
  export MINIO_ENDPOINT=localhost:9000
  export MINIO_ACCESS_KEY=minioadmin
  export MINIO_SECRET_KEY=minioadmin

The server will run on: http://localhost:8000
Health check: http://localhost:8000/healthcheck

📚 Check README.md for API documentation!
"