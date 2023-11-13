# Builder stage
# Use an official Golang runtime as a parent image
FROM golang:1.18.10-alpine3.17 AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container
COPY . .

# Download and cache dependencies
RUN go get

# Build
RUN go build -o main main.go



# Run stage
FROM alpine:3.17
WORKDIR /app

COPY --from=builder /app/main .

# Expose port 8080 for the container
EXPOSE 8080

# Set environment variables at runtime
ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_USER=myuser
ENV DB_PASSWORD=mypassword
ENV DB_NAME=mydb
ENV SSL_MODE=disable

# Set the entry point of the container to the Go app executable
CMD ["/app/main"]
