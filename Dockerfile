# Use a Go builder image
FROM golang:alpine AS builder
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.* ./
RUN go mod download

# Copy the rest of the source code
COPY . ./
RUN go build -o ascii-art-web .
# Compile the Go application to a binary named ascii-art-web
FROM alpine
WORKDIR /app


# RUN ./ascii-art-web 
COPY --from=builder /app/ascii-art-web  /app/ascii-art-web 


# Run the binary
CMD ["/app/ascii-art-web"]