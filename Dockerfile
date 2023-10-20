# Use the official Go image to create a build artifact.
FROM golang:1.21.3-bookworm AS builder

# Set the working directory outside $GOPATH to enable Go modules support.
WORKDIR /app

# Retrieve application dependencies.
# This allows for caching of dependencies, improving build speeds.
COPY server/go.mod .
COPY server/go.sum .
RUN go mod download

# Copy the source code as the last step so the build cache can be leveraged
# as much as possible (Go dependencies rarely change, source code often does).
COPY server/ .

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/filestore ./cmd/main.go

# Use a lightweight Alpine image for the runtime.
FROM alpine:3.14 AS runtime

WORKDIR /app

# Copy the compiled binary from the builder stage.
COPY --from=builder /app/build/filestore /app/
RUN chmod +x /app/filestore

CMD ["/app/filestore"]