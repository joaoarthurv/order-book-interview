FROM golang:1.19-alpine as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o orders-management-service ./cmd/orders-management-service

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/orders-management-service /app/orders-management-service

ADD /internal/config/environment ./internal/config/environment

EXPOSE 8080
EXPOSE 8082
EXPOSE 5432

CMD ["./orders-management-service"]
