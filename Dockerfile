# Use official Go image
FROM golang:1.22-alpine AS builder

# Add CA certificates and git (for go get)
RUN apk update && apk add --no-cache ca-certificates git

# Set working directory
WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go binary
RUN go build -o chat-server ./cmd/main.go

# Final image: small and fast
FROM alpine:latest

# Add certificates
RUN apk --no-cache add ca-certificates

# Copy binary from builder
COPY --from=builder /app/chat-server /chat-server

# Expose WebSocket server port
EXPOSE 8080

# Run the server
ENTRYPOINT ["/chat-server"]
