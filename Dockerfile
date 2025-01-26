# Сборка frontend
FROM node:latest AS frontend-build

WORKDIR /build
COPY /frontend .
COPY .env ..

RUN npm install
RUN npm run build

# Сборка backend
FROM golang:latest AS backend-build

WORKDIR /build

COPY /backend/src .

RUN go mod download
RUN go build -o application .

# Сборка контейнера с приложением
FROM ubuntu:latest

WORKDIR /app
COPY --from=frontend-build /build/dist ./dist
COPY --from=backend-build /build/application .

RUN apt-get update
RUN apt-get  install --reinstall -y ca-certificates

EXPOSE 8080

CMD ["./application"]