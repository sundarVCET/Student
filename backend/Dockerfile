# Use the official Golang image as a builder
FROM golang:1.21 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum for dependency management
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application and name it "student-api"
RUN go build -o student-api .

# Use a minimal base image for the final container
FROM alpine:latest  

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/student-api .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./student-api"]
