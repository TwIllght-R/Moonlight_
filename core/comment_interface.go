package core

import (
	"time"
)

type Comment_resp struct {
	ID        string    `json:"id"`
	TaskID    string    `json:"task_id"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}
type New_comment_req struct {
	StoryID string `json:"task_id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
type New_comment_resp struct {
	CommentID string `json:"comment_id"`
}

type CommentCore interface {
	NewComment(New_comment_req) (*Comment_resp, error)
	GetComment(string) (*Comment_resp, error)
}
