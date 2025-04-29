# Dockerfile
FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN apk add --no-cache gcc musl-dev
RUN go build -o app ./cmd/main.go

EXPOSE 8082

CMD ["./app"]


