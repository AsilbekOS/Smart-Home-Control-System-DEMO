package user

import (
	proto "api_gateway/proto/user"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserApiConn(user_url string) proto.UserServiceClient {
	conn, err := grpc.NewClient(user_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed to connect user_service")
	}

	user := proto.NewUserServiceClient(conn)

	return user
}
