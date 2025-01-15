# Use the official Golang image as the base image
FROM golang:1.18-alpine AS builder

# Set the working directory in the container
WORKDIR /cmd/service

# Copy the project files into the container
COPY . .

# Install the necessary dependencies
RUN go mod download 

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Use a smaller base image for the final image
FROM alpine:latest

# Set the working directory in the container
WORKDIR /

# Copy the binary from the builder stage
COPY --from=builder /cmd/service/service.go .

# Expose the port that your gRPC server listens on
EXPOSE 9001

# Run the gRPC server when the container starts
CMD ["./cmd/service/service.go"]