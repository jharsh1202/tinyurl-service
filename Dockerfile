# Use the official Golang image to build the Go application
FROM golang:1.18-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download and cache the Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o /tiny-url-service

# Use a minimal image to run the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built application from the build stage
COPY --from=build /tiny-url-service .

# Expose the port on which the service will run
EXPOSE 8080

# Command to run the application
CMD ["./tiny-url-service"]
