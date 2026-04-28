FROM golang:1.25

WORKDIR /app

COPY proto/ ./proto/
COPY order-service/ ./order-service/

WORKDIR /app/order-service

RUN go mod tidy
RUN go build -o main ./cmd/server

CMD ["./main"]