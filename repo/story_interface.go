package repo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Story struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	DueDate     time.Time          `bson:"due_date"`
	Priority    string             `bson:"priority"`
	Status      string             `bson:"status"`
	Labels      []string           `bson:"labels"`
	AssignedTo  string             `bson:"assigned_to"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	IsDeleted   bool               `bson:"is_deleted"`
}

type StoryRepo interface {
	GetStoryByID(string) (*Story, error)
	CreateStory(Story) (*Story, error)
}
