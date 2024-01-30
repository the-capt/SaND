all: proto gocode

proto:
	protoc api/proto/*.proto              \
		   --go_out=pkg/proto             \
		   --proto_path=.

gocode:
	go build -o build cmd/sand/sand.go

clean:
	rm -f pkg/proto/sand/*.pb.go
	rm -f build/*
