# Используем официальный образ Golang
FROM golang:1.23-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY go.mod ./
COPY go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем все файлы проекта
COPY . .

# Сборка приложения
RUN go build -o /crypto-service

# Устанавливаем Swagger
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Генерация Swagger-документации
RUN swag init

# Устанавливаем рабочую директорию для Swagger
WORKDIR /app/docs

# Копируем сгенерированные файлы Swagger
COPY docs/docs.go ./docs.go
COPY docs/swagger.json ./swagger.json
COPY docs/swagger.yaml ./swagger.yaml
COPY docs/index.html ./index.html

# Устанавливаем рабочую директорию для приложения
WORKDIR /app

# Открываем порт 8080
EXPOSE 8080

# Запускаем приложение
CMD [ "/crypto-service" ]
