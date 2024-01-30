all: proto gocode

proto:
	protoc api/proto/*.proto              \
		   --go_out=pkg/proto             \
		   --proto_path=.

gocode:
	go build cmd/sand/sand.go
