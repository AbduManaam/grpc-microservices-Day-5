module order-service

go 1.22

require (
	github.com/lib/pq v1.10.9
	google.golang.org/grpc v1.62.0
	google.golang.org/protobuf v1.33.0
	proto v0.0.0
)

replace proto => ../proto
