package router

import (
	"api_gateway/api/handler"
	"api_gateway/grpc_client/device"
	"api_gateway/logs"
	"net/http"
)

func DeviceRouter() {
	lgs := logs.GetLogger("logs/logger.log")
	dClient := device.DeviceApiConn("deviceservice:8081")

	deviceHandler := handler.NewDeviceClient(dClient, lgs)

	http.HandleFunc("POST /devices", deviceHandler.AddDeviceApi)
	http.HandleFunc("PUT /devices/update", deviceHandler.UpdateDeviceApi)
	http.HandleFunc("DELETE /devices/delete", deviceHandler.DeleteDeviceApi)
	http.HandleFunc("POST /control", deviceHandler.SendControlCommand)
}
