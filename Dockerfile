FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pack-calculator ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/pack-calculator .
COPY --from=builder /app/templates/index.html ./templates/index.html

# Create a non-root user
RUN addgroup -S norootgroup && adduser -S noroot -G norootgroup
RUN chown -R noroot:norootgroup /app
USER noroot

EXPOSE 8080

CMD ["./pack-calculator"]