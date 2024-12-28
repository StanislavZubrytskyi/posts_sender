# Етап збірки
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копіюємо весь код проекту
COPY . .

# Завантажуємо залежності та збираємо застосунок
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o main .

# Фінальний етап
FROM alpine:3.19

WORKDIR /app

# Копіюємо бінарний файл
COPY --from=builder /app/main .

# Встановлюємо необхідні права
RUN chmod +x /app/main

# Відкриваємо порт для веб-сервера
EXPOSE 8080

# Запускаємо застосунок
CMD ["/app/main"]
