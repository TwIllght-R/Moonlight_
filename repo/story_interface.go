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
	//Priority    Priority
	Status string `bson:"status"` //todo doing done
	//Progress    Progress           `bson:"progress"`
	Labels      []string     `bson:"labels"`      //tag
	AssignedTo  string       `bson:"assigned_to"` //[]
	CreatedAt   time.Time    `bson:"created_at"`
	UpdatedAt   time.Time    `bson:"updated_at"`
	IsDeleted   bool         `bson:"is_deleted"`
	Comments    []Comment    `bson:"comments"`
	Attachments []Attachment `bson:"attachments"`
}

type StoryRepo interface {
	GetStoryByID(string) (*Story, error)
	GetAll() (*[]Story, error)
	CreateStory(Story) (*Story, error)
	UpdateStory(Story) (*Story, error)
	DeleteStory(string) (*Story, error)
}
type Progress struct {
	Completed  int `bson:"completed"`
	AlmostDone int `bson:"almost_done"`
	HalfWay    int `bson:"in_progress"`
	Started    int `bson:"started"`
	NotStarted int `bson:"not_started"`
}

type TimeEntry struct {
	ID        string    `bson:"_id,omitempty"`
	TaskID    string    `bson:"task_id"`
	StartTime time.Time `bson:"start_time"`
	EndTime   time.Time `bson:"end_time"`
	Duration  int       `bson:"duration"` // Duration in seconds
}

type Notification struct {
	ID      string    `bson:"_id,omitempty"`
	TaskID  string    `bson:"task_id"`
	Content string    `bson:"content"`
	Created time.Time `bson:"created_at"`
}

type Task struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `bson:"title"`
	//Description string             `bson:"description"`
	Status     string    `bson:"status"` //todo doing done
	AssignedTo string    `bson:"assigned_to"`
	CreatedAt  time.Time `bson:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at"`
	IsDeleted  bool      `bson:"is_deleted"`
}

type Attachment struct {
	Filename string `bson:"filename"`
	URL      string `bson:"url"`
	Type     string `bson:"type"`
	Size     int64  `bson:"size"`
}

// func (s *Story) GetProgress() Progress {
// 	progress := Progress{}

// 	// Calculate the progress based on the status of the story
// 	switch s.Status {
// 	case "to":
// 		progress.NotStarted++
// 	case "doing":
// 		progress.InProgress++
// 	case "done":
// 		progress.Completed++
// 	}

// 	return progress
// }

type Priority struct {
	Low      int `bson:"low"`
	Medium   int `bson:"medium"`
	High     int `bson:"high"`
	Critical int `bson:"critical"`
}
