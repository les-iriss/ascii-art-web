FROM golang:1.22.3 AS builder

WORKDIR /app

# Copy the Go module and source code
COPY go.mod .
COPY main.go .
RUN go mod tidy

COPY pkg /app/pkg
COPY controllers /app/controllers
COPY views /app/views



# Compile the Go application to a binary named ascii-art-web
RUN go build -o ascii-art-web .

# Use a lightweight base image
FROM alpine:latest  

# Install ca-certificates for HTTPS requests (if needed)
# RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /ascii-art-web/ascii-art-web .
# COPY --from=builder /ascii-art-web/views ./views


# Run the binary
ENTRYPOINT [ "./app" ]