# Сборка frontend
FROM node:latest AS frontend-build

WORKDIR /build
COPY /frontend .
COPY .env ..

RUN npm install
RUN npm run build

# Сборка backend
FROM golang:alpine AS backend-build

ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /build

COPY /backend/src .

RUN apk update --no-cache && apk add --no-cache tzdata
RUN go mod download
RUN go build -o application .

# Сборка контейнера с приложением
FROM alpine:latest

WORKDIR /app
COPY --from=frontend-build /build/dist ./dist
COPY --from=backend-build /build/application .
COPY --from=backend-build /usr/share/zoneinfo/Europe/Moscow /usr/share/zoneinfo/Europe/Moscow

ENV TZ Europe/Moscow

RUN apk update --no-cache && apk add --no-cache ca-certificates

EXPOSE 8080

CMD ["./application"]