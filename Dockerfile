# Use the official Golang image as the base image
FROM golang:1.23-alpine AS builder

# Set the working directory in the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Install the necessary dependencies
RUN go mod download 

# Copy the project files into the container
COPY . .

# Build the server binary
RUN go build -o bin/server ./cmd/service/service.go

# Use a smaller base image for the final image
FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/bin/server .

ENV JWT_SECRET='jwt-secret'

# Expose the port that your gRPC server listens on
EXPOSE 9001

# Run the gRPC server when the container starts
ENTRYPOINT ["./server"]