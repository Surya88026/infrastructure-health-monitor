
# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./

COPY . .

RUN go build -o healthchecker .

# Stage 2: Run the application
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/healthchecker .

ENTRYPOINT ["./healthchecker"]

