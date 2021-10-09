package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption         string             `json:"caption,omitempty" bson:"caption,omitempty"`
	ImageURL        string             `json:"imageurl,omitempty" bson:"imageurl,omitempty"`
	PostedTimestamp string             `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}
