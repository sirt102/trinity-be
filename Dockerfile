# Step 1: Build the Go binary
FROM golang:1.23.2-alpine3.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the application code
COPY . .

# Build the Go binary from the correct path
RUN go build -o main ./cmd/server/main.go

# Step 2: Create a small image for running the Go binary
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder image
COPY --from=builder /app/main .

# Copy .env file into the container
COPY .env .env

# Download wait-for-it.sh and ensure it's executable
RUN apk --no-cache add curl bash && \
    curl -o wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh && \
    chmod +x wait-for-it.sh

# Expose port for the application
EXPOSE 8080

# Use wait-for-it.sh to ensure PostgresQL is ready before starting the Go app
CMD ["./wait-for-it.sh", "db:5432", "--", "./main"]
