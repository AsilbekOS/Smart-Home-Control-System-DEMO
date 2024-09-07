package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	DeviceID           primitive.ObjectID `bson:"device_id,omitempty" json:"device_id,omitempty"`
	UserID             primitive.ObjectID `bson:"user_id" json:"user_id"`
	DeviceType         string             `bson:"device_type" json:"device_type"`
	DeviceName         string             `bson:"device_name" json:"device_name"`
	DeviceStatus       string             `bson:"device_status" json:"device_status"`
	Configuration      Configuration      `bson:"configuration" json:"configuration"`
	LastUpdated        time.Time          `bson:"last_updated" json:"last_updated"`
	Location           string             `bson:"location,omitempty" json:"location,omitempty"`
	FirmwareVersion    string             `bson:"firmware_version,omitempty" json:"firmware_version,omitempty"`
	ConnectivityStatus string             `bson:"connectivity_status" json:"connectivity_status"`
}

type Configuration struct {
	Brightness int32  `bson:"brightness" json:"brightness"`
	Color      string `bson:"color" json:"color"`
}

type Command struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	DeviceID       primitive.ObjectID `bson:"device_id" json:"device_id"`
	UserID         primitive.ObjectID `bson:"user_id" json:"user_id"`
	CommandType    string             `bson:"command_type" json:"command_type"`
	CommandPayload CommandPayload     `bson:"command_payload" json:"command_payload"`
	Timestamp      time.Time          `bson:"timestamp" json:"timestamp"`
	Status         string             `bson:"status" json:"status"`
}

type CommandPayload struct {
	Brightness int `bson:"brightness" json:"brightness"`
}
