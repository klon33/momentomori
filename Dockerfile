# Build stage
FROM golang:1.25.3-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests (if needed in future)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy templates directory
COPY --from=builder /app/templates ./templates

# Copy static files (momento.svg)
COPY --from=builder /app/momento.svg ./momento.svg

# Expose port 2129
EXPOSE 2129

# Run the application
CMD ["./main"]
