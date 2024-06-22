# Use a Go builder image
FROM golang:1.22.3

WORKDIR /app

# Copy the Go module files and download dependencies
RUN go mod init ascii-art-web

COPY . .

# Copy the rest of the source code
RUN go build -o ascii-art-web .

EXPOSE 8000

# Run the binary
CMD ["/app/ascii-art-web"]