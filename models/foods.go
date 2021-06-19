package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name       string             `json:"name" validate:"required,min=2,max=100"`
	Food_image string             `json:"food_image" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
