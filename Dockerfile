FROM golang:1.25-alpine

# Устанавливаем зависимости для сборки
RUN apk add --no-cache git curl bash

WORKDIR /app

# Ставим air для hot reload
RUN go install github.com/air-verse/air@latest

# Копируем модули отдельно, чтобы кэшировалось
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .
COPY .air.toml ./

# Запускаем air (будет слушать изменения и пересобирать проект)
CMD ["air"]
