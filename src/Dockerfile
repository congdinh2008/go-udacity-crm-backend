# Create docker image for the application go api
# Use golang image as base image
FROM golang:1.22.3 as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Copy json file from data folder to the working directory inside the container
COPY data/customers.json data/customers.json

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app with CGO disabled to ensure compatibility with Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use alpine image as base image
FROM alpine:latest

# Add ca-certificates in case your application makes HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /app/main

# Copy the customers.json file to the correct location in the final image
COPY --from=builder /app/data/customers.json app/data/customers.json


# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["/app/main"]