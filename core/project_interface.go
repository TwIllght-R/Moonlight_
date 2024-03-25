package core

import (
	"time"
)

type New_project_req struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Priority    string    `json:"priority"`
	Labels      []string  `json:"labels"`
	//AssignedTo  string       `json:"assigned_to"`
	Attachments []Attachment `json:"attachments"`
	// Tasks       []Task       `json:"tasks"`
}

type Update_project_req struct {
	ID          string    `json:"project_id"` //for update
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Priority    string    `json:"priority"`
	Status      string    `json:"status"`
	Labels      []string  `json:"labels"`
	//AssignedTo  string       `json:"assigned_to"`
	Attachments []Attachment `json:"attachments"`
	//Tasks       []Task       `json:"tasks"`
}

type New_project_resp struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Priority    string    `json:"priority"`
	Labels      []string  `json:"labels"`
	//AssignedTo  string    `json:"assigned_to"`
	//Tasks []Task `json:"tasks"`
}

type Get_project_resp struct {
	ID          string    `json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Priority    string    `json:"priority"`
	Status      string    `json:"status"`
	Labels      []string  `json:"labels"`
	//AssignedTo  string             `json:"assigned_to"`
	Attachments []Attachment `json:"attachments"`
	//	Tasks       []Task       `json:"tasks"`
}

// type Task struct {
// 	Title      string `json:"title"`
// 	Status     string `json:"status"`
// 	AssignedTo string `json:"assigned_to"`
// }

type Attachment struct {
	Filename string `json:"filename"`
	URL      string `json:"url"`
	Type     string `json:"type"`
	Size     int64  `json:"size"`
}
type ProjectCore interface {
	NewProject(New_project_req) (*New_project_resp, error)
	GetProject(string) (*Get_project_resp, error)
	ListProject() (*[]Get_project_resp, error)
	UpdateProject(Update_project_req) (*New_project_resp, error)
	DeleteProject(string) (*Get_project_resp, error)
}
