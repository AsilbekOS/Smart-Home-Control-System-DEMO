package service

import (
	"context"
	models "device-service/Models"
	"device-service/database/rabbitmq"
	getid "device-service/internal/getID"
	"device-service/proto"
	"device-service/repo"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/encoding/protojson"
)

type Server struct {
	proto.UnimplementedDeviceControlServiceServer
	DB *mongo.Collection
	LG *log.Logger
}

func NewDeviceServer(db *mongo.Collection, lg *log.Logger) *Server {
	return &Server{DB: db, LG: lg}
}

func (s *Server) AddDevice(ctx context.Context, req *proto.AddDeviceRequest) (*proto.AddDeviceResponse, error) {
	id, err := getid.GetUserIDFromToken(req.Device.Token, "Asilbek")
	if err != nil {
		s.LG.Println("AddDevice ERROR: error token:", err)
		return nil, fmt.Errorf("ERROR: error token")
	}

	user_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		s.LG.Printf("AddDevice ERROR: error converting id to ObjectID: %v", err)
		return nil, fmt.Errorf("ERROR: error converting id to ObjectID: %v", err)
	}

	device := &models.Device{
		DeviceID:     primitive.NewObjectID(),
		UserID:       user_id,
		DeviceType:   req.Device.DeviceType,
		DeviceName:   req.Device.DeviceName,
		DeviceStatus: req.Device.DeviceStatus,
		Configuration: models.Configuration{
			Brightness: req.Device.Configuration.Brightness,
			Color:      req.Device.Configuration.Color,
		},
		LastUpdated:        time.Now(),
		Location:           req.Device.Location,
		FirmwareVersion:    req.Device.FirmwareVersion,
		ConnectivityStatus: req.Device.ConnectivityStatus,
	}

	err = repo.AddDevicMongo(device)
	if err != nil {
		log.Println(err)
		s.LG.Println("AddDevice ERROR:", err)
		return nil, fmt.Errorf("ERROR:  %v", err)
	}

	err = repo.CacheUser(device)
	if err != nil {
		log.Println(err)
		s.LG.Println("AddDevice ERROR:", err)
		return nil, fmt.Errorf("ERROR:  %v", err)
	}

	response := "Successfully created device and cashed"
	log.Println(response)
	s.LG.Println(response)
	return &proto.AddDeviceResponse{Success: response}, nil
}

func (s *Server) UpdateDevice(ctx context.Context, req *proto.UpdateDeviceRequest) (*proto.UpdateDeviceResponse, error) {
	err := repo.GetUserWithID(req.Device.DeviceId)
	if err != nil {
		s.LG.Println("UpdateDevice ERROR: Bunday DeviceID topilamdi", err)
		log.Println("UpdateDevice ERROR: Bunday DeviceID topilamdi")
		return nil, fmt.Errorf("UpdateDevice ERROR: Bunday DeviceID topilamdi: %v", err)
	}

	now := time.Now()
	timeString := now.Format("02-Jan-2006 03:04:05 PM")

	update := proto.UpdateDeviceRequest{
		Device: &proto.Device{
			DeviceId:     req.Device.DeviceId,
			DeviceName:   req.Device.DeviceName,
			DeviceType:   req.Device.DeviceType,
			DeviceStatus: req.Device.DeviceStatus,
			Configuration: &proto.Configuration{
				Brightness: req.Device.Configuration.Brightness,
				Color:      req.Device.Configuration.Color,
			},
			LastUpdated:        timeString,
			Location:           req.Device.Location,
			FirmwareVersion:    req.Device.FirmwareVersion,
			ConnectivityStatus: req.Device.ConnectivityStatus,
		},
	}
	err = repo.UpdateDevicemongo(&update)
	if err != nil {
		s.LG.Println("UpdateDevice ERROR:", err)
		return &proto.UpdateDeviceResponse{Success: "Error updating device"}, fmt.Errorf("UpdateDevice ERROR: %v", err)
	}

	s.LG.Println("Successfully update device")
	log.Println("Successfully update device")
	return &proto.UpdateDeviceResponse{Success: "Successfully update device"}, nil
}

func (s *Server) DeleteDevice(ctx context.Context, req *proto.DeleteDeviceRequest) (*proto.DeleteDeviceResponse, error) {
	err := repo.DeleteDevicemongo(req)
	if err != nil {
		s.LG.Println("DeleteDevice ERROR:", err)
		log.Println("error delete device")
		return &proto.DeleteDeviceResponse{Success: "error delete device"}, err
	}

	s.LG.Println("Successfully deleted device")
	log.Println("Successfully deleted device")
	return &proto.DeleteDeviceResponse{Success: "Successfully deleted device"}, nil
}

func (s *Server) ControlDevice(ctx context.Context, req *proto.ControlRequest) (*proto.ControlResponse, error) {
	log.Println("0")
	rabbitCh, err := rabbitmq.CreateChannel()
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(err.Error(), "1")
	}

	conReq := &proto.ControlRequest{
		Command: &proto.ControlCommand{
			CommandId:   req.Command.CommandId,
			DeviceId:    req.Command.DeviceId,
			UserId:      req.Command.UserId,
			CommandType: req.Command.CommandType,
			CommandPayload: &proto.BrightnessCommandPayload{
				Brightness: req.Command.CommandPayload.Brightness,
				Color:      req.Command.CommandPayload.Color,
			},
			Timestamp:    time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"),
			DeviceStatus: req.Command.DeviceStatus,
		},
	}

	// Buyruqni RabbitMQ'ga yuborish
	messageBody, err := protojson.Marshal(conReq)
	if err != nil {
		log.Println("err-1")
		return nil, err
	}

	err = rabbitCh.Publish(
		"",
		"control_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/protobuf",
			Body:        messageBody,
		},
	)
	if err != nil {
		log.Println("err-2")
		return nil, err
	}

	log.Println("Successfully create comand")
	return &proto.ControlResponse{Success: "Command received"}, nil
}
