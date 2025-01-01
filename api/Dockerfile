# Stage 1: Build the Go binary
FROM golang:1.21.3 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the project files
COPY . .

# Set the working directory to the main file directory
WORKDIR /app/cmd/chatbotapi

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /chatbotapi main.go

# Stage 2: Create a lightweight container with the binary
FROM alpine:latest

# Install certificates for HTTPS support
RUN apk --no-cache add ca-certificates

# Copy the built binary from the builder stage
COPY --from=builder /chatbotapi /chatbotapi

# Expose the port your application will run on
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["/chatbotapi"]
