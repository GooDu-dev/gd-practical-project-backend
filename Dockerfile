# Start from the official Golang image with version 1.20
FROM golang:1.20 as builder

# Set the working directory in the container
WORKDIR /app

# Copy go.mod and go.sum files to cache dependencies
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Start a new stage to build a minimal Docker image with the compiled binary
FROM alpine:latest

# Install necessary packages
RUN apk --no-cache add ca-certificates

# Set working directory in the final image
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
