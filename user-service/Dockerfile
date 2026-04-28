FROM golang:1.25

WORKDIR /app

COPY proto/ ./proto/

COPY user-service/ ./user-service/

WORKDIR /app/user-service

RUN go mod tidy
RUN go build -o main ./cmd/server

CMD ["./main"]