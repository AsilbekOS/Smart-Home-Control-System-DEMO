package device

import (
	proto "api_gateway/proto/device"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DeviceApiConn(device_url string) proto.DeviceControlServiceClient {
	conn, err := grpc.NewClient(device_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed to connect device_service")
	}

	device := proto.NewDeviceControlServiceClient(conn)

	return device
}
