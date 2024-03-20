package repo

import (
	"time"
)

type Comment struct {
	ID        string    `bson:"_id,omitempty"`
	TaskID    string    `bson:"task_id"`
	Content   string    `bson:"content"`
	Author    string    `bson:"author"`
	CreatedAt time.Time `bson:"created_at"`
}

type CommentRepo interface {
	GetCommentByID(string) (*Comment, error)
	CreateComment(Comment) (*Comment, error)
}
