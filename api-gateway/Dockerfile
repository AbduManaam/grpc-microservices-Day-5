


FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY proto/ ./proto/
COPY api-gateway/ ./api-gateway/

WORKDIR /app/api-gateway/
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o main .

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/api-gateway/main .
CMD ["./main"]