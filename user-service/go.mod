module user-service

go 1.25.4

require (
	github.com/golang-jwt/jwt/v5 v5.3.1
	github.com/lib/pq v1.10.9
	golang.org/x/crypto v0.18.0
	proto v0.0.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/grpc v1.62.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace proto => ../proto
