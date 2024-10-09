# Use a base image with Go
FROM golang:1.22-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the Go module files and the source code
COPY go.mod .
COPY . .

# Download Go dependencies
RUN go mod download

# Build the Go application
RUN go build -o model-app .

# Run the binary when the container starts
CMD ["/app/model-app"]
