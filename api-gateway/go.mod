module api-gateway

go 1.22

require (
	google.golang.org/grpc v1.62.0
	google.golang.org/protobuf v1.33.0
	proto v0.0.0
)

replace proto => ../proto
