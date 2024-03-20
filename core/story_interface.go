package core

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type New_story_req struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Priority    string    `json:"priority"`
	Labels      []string  `json:"labels"`
	AssignedTo  string    `json:"assigned_to"`
}

type Update_story_req struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Priority    string    `json:"priority"`
	Status      string    `json:"status"`
	Labels      []string  `json:"labels"`
	AssignedTo  string    `json:"assigned_to"`
}

type New_story_resp struct {
	Title string `json:"title"`
}

type Get_story_resp struct {
	ID          primitive.ObjectID `json:"_id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	DueDate     time.Time          `json:"due_date"`
	Priority    string             `json:"priority"`
	Status      string             `json:"status"`
	Labels      []string           `json:"labels"`
	AssignedTo  string             `json:"assigned_to"`
}

type StoryCore interface {
	NewStory(New_story_req) (*New_story_resp, error)
	GetStory(string) (*Get_story_resp, error)
	ListStory() (*[]Get_story_resp, error)
	UpdateStory(string, Update_story_req) (*New_story_resp, error)
	DeleteStory(string) (*Get_story_resp, error)
}
