@echo off
start "user-service" cmd /k "cd user-service && go run ./cmd/server"
start "order-service" cmd /k "cd order-service && go run ./cmd/server"
start "api-gateway" cmd /k "cd api-gateway && go mod tidy && go run ."