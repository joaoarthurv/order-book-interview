FROM golang:1.19-alpine as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o users-wallet-service ./cmd/users-wallet-service

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/users-wallet-service /app/users-wallet-service

ADD /internal/config/environment ./internal/config/environment

EXPOSE 8082
EXPOSE 8080
EXPOSE 4566

CMD ["./users-wallet-service"]
