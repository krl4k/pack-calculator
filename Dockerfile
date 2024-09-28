FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pack-calculator ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/pack-calculator .
COPY templates/index.html ./templates/index.html

EXPOSE 8080

CMD ["./pack-calculator"]