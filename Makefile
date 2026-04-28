proto:
	protoc --go_out=. --go-grpc_out=. proto/user.proto
	protoc --go_out=. --go-grpc_out=. proto/order.proto

tidy:
	cd proto         && go mod tidy
	cd api-gateway   && go mod tidy
	cd user-service  && go mod tidy
	cd order-service && go mod tidy

run:
	docker-compose -f deployments/docker-compose.yml up --build

run-local:
	go run ./user-service/cmd/server &
	go run ./order-service/cmd/server &
	go run ./api-gateway/main.go

stop-local:
	pkill -f "user-service" || true
	pkill -f "order-service" || true
	pkill -f "api-gateway" || true	