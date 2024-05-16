# Use the official Golang image as the base image
FROM golang:latest

LABEL org.opencontainers.image.source=https://github.com/fredrik1338/lore-keeper-be

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container's working directory
COPY cmd/ /app/cmd/
COPY internal/ /app/internal/
# COPY vendor/ /app/vendor/

# Copy the main.go file from the host into the container
COPY go.mod .
COPY go.sum .

# Build the Go application inside the container
RUN go build -o main ./cmd/main

ENV BACKEND_PORT="8080"
ENV BACKEND_HOST=""

# Expose the port that the Go application listens on
EXPOSE 8080

# Command to run the Go application when the container starts
CMD ["./main"]