generate-device_proto:
	protoc --go_out=. --go-grpc_out=. device.proto

docker-build:
	sudo docker build -t api_gateway .

docker-run:
	sudo docker run -p 4440:4440 api_gateway

generate-device:
	protoc \
	--go_out=./ \
	--go_opt=paths=source_relative \
	--go-grpc_out=./ \
	--go-grpc_opt=paths=source_relative \
	proto/device/device.proto

generate-user:
	protoc \
	--go_out=./ \
	--go_opt=paths=source_relative \
	--go-grpc_out=./ \
	--go-grpc_opt=paths=source_relative \
	proto/user/user.proto

swag-gen:
	~/go/bin/swag init -g ./api/cmd/main.go -o api/docs force 1