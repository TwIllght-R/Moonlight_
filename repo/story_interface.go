package repo

import "go.mongodb.org/mongo-driver/bson/primitive"

// type StoryContent struct {
// 	ID            primitive.ObjectID `bson:"_id,omitempty"`
// 	StoryID       primitive.ObjectID `bson:"story_id"`
// 	SectionNumber int                `bson:"section_number"`
// 	ContentType   string             `bson:"content_type"`
// 	Content       map[int]string     `bson:"content"`
// }

//type story_type struct{}
type Story struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Title   string             `bson:"title"`
	Content map[int]struct {
		Text string `bson:"text"`
		Type string `bson:"type"`
	} `bson:"content"`
	UserID primitive.ObjectID `bson:"user_id"`
}

type StoryRepo interface {
	GetStoryByID(primitive.ObjectID) (*Story, error)
	CreateStory(Story) (*Story, error)
}

// type StoryContentRepo interface {
// 	GetStoryContentByID(primitive.ObjectID) (StoryContent, error)
// }
