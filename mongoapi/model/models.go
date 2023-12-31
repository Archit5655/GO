package model

import (
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie     string             `json:"movie,omitempty"`
	Iswatched bool               `json:"is_watched,omitempty"`
}
