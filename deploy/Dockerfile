FROM golang:1.22.5-alpine AS builder

# Копируем все файлы проекта
COPY .. /github.com/Dnlbb/auth/
WORKDIR /github.com/Dnlbb/auth/

# Скачиваем зависимости
RUN go mod download

# Собираем бинарник
RUN go build -o ./bin/auth-server cmd/main.go

# Финальный этап
FROM alpine:latest

WORKDIR /root/cmd

# Копируем бинарник auth-server из builder стадии
COPY --from=builder /github.com/Dnlbb/auth/bin/auth-server ./cmd/
COPY --from=builder /github.com/Dnlbb/auth/auth.env ./postgres/.env



CMD ["sh", "-c", "cd cmd && ./auth-server"]
