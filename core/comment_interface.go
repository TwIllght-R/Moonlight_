package core

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment_resp struct {
	Text      string             `json:"text"`
	UserID    primitive.ObjectID `json:"user_id"`
	CreatedAt time.Time          `json:"created_at"`
}
