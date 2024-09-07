package main

import (
	"log"
	"net"
	"os"
	mongodb "user_service/database/MongoDB"
	"user_service/logs"
	"user_service/proto"
	"user_service/service"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lgs := logs.GetLogger("logs/logger.log")
	lis, err := net.Listen("tcp", "userservice:8081")
	if err != nil {
		log.Println("Server listening error:", err)
	}

	defer lis.Close()

	db, err := mongodb.ConnectMDB()
	if err != nil {
		log.Println(err)
	}
	server := service.NewUserServer(db, lgs)
	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, server)
	reflection.Register(s)

	log.Printf("Server is listening to port: %v", os.Getenv("address"))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error to Server - %v", err)
	}
}
