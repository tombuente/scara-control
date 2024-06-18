all:
	protoc --go_out=. --go_opt=paths=/scara-proto/ --go-grpc_out=. --go-grpc_opt=paths=/scara-proto/ scara.proto
