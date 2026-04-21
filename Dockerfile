# Use an official Golang runtime as a parent image
FROM golang:1.22.5-alpine AS builder

# Set the working directory to /app
WORKDIR /app

# Set the build argument for the app version number
ARG APP_VERSION=0.1.0

# Copy go.mod and go.sum first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go app
RUN go build -ldflags "-X main.version=$APP_VERSION" -o makeline-service .

# Run the app on alpine
FROM alpine:latest AS runner

ARG APP_VERSION=0.1.0

# Install ca-certificates for HTTPS calls to Product Service
RUN apk --no-cache add ca-certificates

# Copy the build output from the builder container
COPY --from=builder /app/makeline-service .

# Expose port for the container (changed to 3004 for Best Buy)
EXPOSE 3004

# Set environment variables for Best Buy service
ENV APP_VERSION=$APP_VERSION
ENV PORT=3004
ENV ORDER_DB_NAME=bestbuy
ENV ORDER_DB_COLLECTION_NAME=orders
ENV PRODUCT_SERVICE_URL=http://product-service:3002

# Run the Go app when the container starts
CMD ["./makeline-service"]