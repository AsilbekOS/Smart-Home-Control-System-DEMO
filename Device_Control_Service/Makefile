generate-device_proto:
	protoc --go_out=. --go-grpc_out=. device.proto

docker-build:
	sudo docker build -t device_service .

docker-run:
	sudo docker run -p 4440:4440 device_service

generate-rpc:
	protoc \
	--go_out=./ \
	--go_opt=paths=source_relative \
	--go-grpc_out=./ \
	--go-grpc_opt=paths=source_relative \
	proto/device.proto
