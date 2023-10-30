# Use the official Go image for Go 1.18 as the base image
FROM golang:1.18 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application and name it differently (e.g., mycustomapp)
RUN CGO_ENABLED=0 go build -o mycustomapp

# Use a lightweight base image for the final stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary with the custom name
COPY --from=builder /app/mycustomapp .

# Set the entry point for the container
CMD ["./mycustomapp"]
