package main

import (
	"device-service/database/mongodb"
	"device-service/logs"
	"device-service/proto"
	"device-service/service"
	"log"
	"net"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lgs := logs.GetLogger("logs/logger.log")
	lis, err := net.Listen("tcp", "deviceservice:8081")
	if err != nil {
		log.Fatalf("Server listening error: %v", err)
	}
	defer func() {
		if err := lis.Close(); err != nil {
			log.Fatalf("Failed to close listener: %v", err)
		}
	}()

	db, err := mongodb.ConnectMDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	server := service.NewDeviceServer(db, lgs)
	s := grpc.NewServer()
	proto.RegisterDeviceControlServiceServer(s, server)
	reflection.Register(s)

	log.Printf("Server is listening to port: %v", os.Getenv("address"))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
