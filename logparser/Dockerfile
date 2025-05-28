FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .
RUN go mod tidy
RUN go build -o logparser main.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/logparser .

# Set environment variables (optional defaults)
ENV LOGS_DIR=/app/logs
ENV PARSED_DIR=/app/logs/parsed
ENV UDP_ADDRESS=logstash:5000

CMD ["./logparser"]
