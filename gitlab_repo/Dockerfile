# Сборка
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY main.go .
RUN apk update
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hellokuberapp ./main.go

# Запуск
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/hellokuberapp .
EXPOSE 8081
ENTRYPOINT ["./hellokuberapp"]