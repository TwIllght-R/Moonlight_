package repo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	TaskID    primitive.ObjectID `bson:"task_id"`
	Content   string             `bson:"content"`
	Author    primitive.ObjectID `bson:"author"`
	CreatedAt time.Time          `bson:"created_at"`
}
