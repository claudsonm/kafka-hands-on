# Builder image steps
FROM golang:1.16-alpine as builder

RUN apk update && apk add --no-cache \
    bash \
    ca-certificates \
    git \
    gcc \
    g++ \
    libc-dev \
    librdkafka-dev \
    pkgconf && \
    update-ca-certificates

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags musl -o go-kafka-consumer

# Execution image steps
FROM alpine:latest as runner
WORKDIR /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/go-kafka-consumer .
ENTRYPOINT /app/go-kafka-consumer