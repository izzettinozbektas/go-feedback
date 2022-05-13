package models

import (
	"time"
)

type Feedback struct {
	Feedback string `json:"feedback"`
	Created_at time.Time `json:"created_at" bson:"created_at"`
}
