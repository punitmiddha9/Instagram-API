package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id",omitempty`
	Caption   string             `json:"Caption" bson:"Caption"`
	ImageUrl  string             `json:"ImageUrl" bson:"ImageURL"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
	UserID    string             `json:"UserID" bson:"UserID"`
}
