generate-user_proto:
	protoc --go_out=. --go-grpc_out=. proto/user.proto

docker-build:
	sudo docker build -t user_service .

docker-run:
	sudo docker run -p 4440:4440 user_service

generate-rpc:
	protoc \
	--go_out=./ \
	--go_opt=paths=source_relative \
	--go-grpc_out=./ \
	--go-grpc_opt=paths=source_relative \
	proto/user.proto

