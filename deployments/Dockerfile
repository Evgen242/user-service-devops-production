# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем файлы зависимостей сначала для лучшего кэширования
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# Final stage  
FROM alpine:latest
RUN apk --no-cache add ca-certificates curl
RUN addgroup -S app && adduser -S app -G app
WORKDIR /root/
COPY --from=builder /app/main .
RUN chown -R app:app /root/
USER app
EXPOSE 8090
CMD ["./main"]
