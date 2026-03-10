# syntax=docker/dockerfile:1
FROM golang:1.24.2

# Set destination for COPY
WORKDIR /app

# Copy all go. files (i.e. go.mod)
COPY go.* ./

# Download required modules
RUN go mod download

# Copy any remaining files
COPY . .

# Build
RUN go build -o main main.go

# Expose TCP port (given to docker)
EXPOSE 9090

# Run
CMD ["./main"]