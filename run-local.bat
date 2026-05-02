@echo off
set USER_SERVICE_ADDR=localhost:50051
set ORDER_SERVICE_ADDR=localhost:50052
set POSTGRES_HOST=localhost
set POSTGRES_PORT=5432
set POSTGRES_USER=postgres
set POSTGRES_PASSWORD=9539Abdu
set PORT=8080
start "user-service" cmd /k "cd user-service && go run ./cmd/server"
start "order-service" cmd /k "cd order-service && go run ./cmd/server"
start "api-gateway" cmd /k "cd api-gateway && go mod tidy && go run ."