FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates

# Copy everything first
COPY . .

# Tidy and download once all source files are present
RUN go mod tidy && go mod download

# Now build
RUN go build -o chat-server .

# --- Runtime image ---
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/chat-server .

EXPOSE 8080

ENTRYPOINT ["./chat-server"]
