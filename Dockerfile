FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates

# Step 1: prepare modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Step 2: copy full source
COPY . .

# Step 3: tidy again (to resolve new imports after source code is copied)
RUN go mod tidy

# Step 4: build
RUN go build -o chat-server .

# --- Runtime image ---
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/chat-server .

EXPOSE 8080

ENTRYPOINT ["./chat-server"]
