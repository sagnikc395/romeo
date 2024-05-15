package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID         primitive.ObjectID `bson:"_id"`
	Text       string
	Title      string
	Created_at time.Time
	Updated_at time.Time
	Note_id    
}
