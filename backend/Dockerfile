# Stage 1: Build the Go application
FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY server.go .
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o server server.go

# Stage 2: Create a minimal runtime container
FROM scratch
WORKDIR /app
COPY server.env .
COPY --from=builder /app/server .
CMD ["/app/server"]
