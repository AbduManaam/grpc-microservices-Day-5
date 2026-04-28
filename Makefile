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