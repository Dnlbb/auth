FROM golang:1.22.5-alpine AS builder

# Копируем все файлы проекта
COPY . /github.com/Dnlbb/auth/
WORKDIR /github.com/Dnlbb/auth/

# Скачиваем зависимости
RUN go mod download

# Собираем бинарник
RUN go build -o ./bin/auth-server cmd/main.go

# Финальный этап
FROM alpine:latest

WORKDIR /root/

# Копируем бинарник auth-server из builder стадии
COPY --from=builder /github.com/Dnlbb/auth/bin/auth-server .

# Копируем файл .env в контейнер
COPY --from=builder /github.com/Dnlbb/auth/postgres/.env /root/.env

# Указываем команду для запуска
CMD ["./auth-server"]
