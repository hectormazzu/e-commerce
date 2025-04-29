# Dockerfile
FROM golang:1.24-alpine

#ENV CGO_ENABLED=1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN apk add --no-cache gcc musl-dev
RUN go build -o app ./cmd/main.go
# RUN chmod +x main

EXPOSE 8081

CMD ["./app"]
