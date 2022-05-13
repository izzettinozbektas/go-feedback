package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Feedback struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Feedback   string             `json:"feedback"`
	Created_at time.Time          `json:"created_at" bson:"created_at"`
}
