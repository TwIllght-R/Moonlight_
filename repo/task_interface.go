package repo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Project_ID primitive.ObjectID `bson:"project_id"`
	Title      string             `bson:"title"`
	Status     string             `bson:"status"` //todo doing done
	AssignedTo string             `bson:"assigned_to"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
	DeleteAt   time.Time          `bson:"deleted_at"`
}

type TaskRepo interface {
	GetTaskByID(primitive.ObjectID) (*Task, error)
	GetTaskByProjectID(primitive.ObjectID) (*[]Task, error)
	GetAll() (*[]Task, error)
	CreateTask(Task) (*Task, error)
	UpdateTask(Task) (*Task, error)
	DeleteTask(primitive.ObjectID) (*Task, error)
}
