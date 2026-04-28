FROM golang:1.25

WORKDIR /app

COPY proto/ ./proto/
COPY api-gateway/ ./api-gateway/

WORKDIR /app/api-gateway

RUN go mod tidy
RUN go build -o main .

CMD ["./main"]