FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Set Go proxy to use direct fallback
ENV GOPROXY=https://proxy.golang.org,direct

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download


# Install swag
RUN go install github.com/swaggo/swag/cmd/swag@latest


# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN swag init --parseInternal --parseDependency

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /root/


# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8082 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]