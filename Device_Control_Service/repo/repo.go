package repo

import (
	"context"
	models "device-service/Models"
	"device-service/database/mongodb"
	redis "device-service/database/redisdb"
	"device-service/proto"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddDevicMongo(device *models.Device) error {
	collection, err := mongodb.ConnectMDB()
	if err != nil {
		log.Println(err)
	}

	_, err = collection.InsertOne(context.Background(), bson.D{
		{Key: "device_id", Value: device.DeviceID},
		{Key: "user_id", Value: device.UserID},
		{Key: "device_type", Value: device.DeviceType},
		{Key: "device_name", Value: device.DeviceName},
		{Key: "device_status", Value: device.DeviceStatus},
		{Key: "configuration", Value: bson.D{
			{Key: "brightness", Value: device.Configuration.Brightness},
			{Key: "color", Value: device.Configuration.Color},
		}},
		{Key: "last_updated", Value: device.LastUpdated},
		{Key: "location", Value: device.Location},
		{Key: "firmware_version", Value: device.FirmwareVersion},
		{Key: "connectivity_status", Value: device.ConnectivityStatus},
	})
	if err != nil {
		log.Println(err)
	}
	return nil
}

// Foydalanuvchi ma'lumotlarini MongoDB'dan olish
func GetUserWithID(id string) (error) {
	// fmt.Println("Mongodan ma'lumotni oldi.")
	deviceid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("error changing id type: %v", err)
	}

	collection, err := mongodb.ConnectMDB()
	if err != nil {
		log.Println(err)
	}
	var user *proto.Device
	err = collection.FindOne(context.Background(), bson.D{{Key: "device_id", Value: deviceid}}).Decode(&user)
	if err != nil {
		return err
	}
	
	return nil
}

func UpdateDevicemongo(req *proto.UpdateDeviceRequest) error {
	collection, err := mongodb.ConnectMDB()
	if err != nil {
		log.Println(err)
	}

	id, err := primitive.ObjectIDFromHex(req.Device.DeviceId)
	if err != nil {
		log.Println(err)
	}

	filter := bson.D{{Key: "device_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.M{
		"device_type":   req.Device.DeviceType,
		"device_name":   req.Device.DeviceName,
		"device_status": req.Device.DeviceStatus,
		"configuration": bson.M{
			"brightness": req.Device.Configuration.Brightness,
			"color":      req.Device.Configuration.Color,
		},
		"last_updated":        req.Device.LastUpdated,
		"location":            req.Device.Location,
		"firmware_version":    req.Device.FirmwareVersion,
		"connectivity_status": req.Device.ConnectivityStatus,
	}}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func DeleteDevicemongo(req *proto.DeleteDeviceRequest) error {
	collection, err := mongodb.ConnectMDB()
	if err != nil {
		log.Println(err)
	}

	id, err := primitive.ObjectIDFromHex(req.DeviceId)
	if err != nil {
		log.Println(err)
	}

	filter := bson.D{{Key: "device_id", Value: id}}

	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func CacheUser(device *models.Device) error {
	redisClient, err := redis.ConnectRedis()
	if err != nil {
		log.Println(err)
	}

	_, err = redisClient.HSet(context.Background(), "device:"+device.DeviceName, map[string]interface{}{
		"device_type":         device.DeviceType,
		"device_name":         device.DeviceName,
		"device_status":       device.DeviceStatus,
		"configuration_b":     device.Configuration.Brightness,
		"configuration_c":     device.Configuration.Color,
		"last_updated":        device.LastUpdated,
		"location":            device.Location,
		"firmware_version":    device.FirmwareVersion,
		"connectivity_status": device.ConnectivityStatus,
	}).Result()
	if err != nil {
		log.Println(err)
	}

	return nil
}
