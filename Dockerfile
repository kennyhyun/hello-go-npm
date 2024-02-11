# Start from a Debian-based lightweight Go image
FROM golang:1.18-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Build the Go app
RUN go build -o helloworld .

######## Start a new stage from scratch #######
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/helloworld .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./helloworld"]
