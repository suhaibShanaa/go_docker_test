# Use the official Go image
FROM golang:1.23-alpine

# Install git and Air (from the correct path)
RUN apk add --no-cache git && \
    go install github.com/air-verse/air@latest

# Set working directory
WORKDIR /app

# Copy the Go module files and install dependencies
COPY go.* ./
RUN go mod download

# Copy the source code
COPY . ./

# Expose the application port
EXPOSE 8080

# Set the default command to run Air
CMD ["air"]
