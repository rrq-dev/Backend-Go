package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ParkingSession struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    VehicleID string             `bson:"vehicle_id"`
    EntryTime int64              `bson:"entry_time"`
    ExitTime  int64              `bson:"exit_time,omitempty"`
    IsActive  bool               `bson:"is_active"`
}
