package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type storyRepo struct {
	collection *mongo.Collection
}

func NewStoryRepo(collection *mongo.Collection) StoryRepo {
	return &storyRepo{collection: collection}
}

func (r *storyRepo) GetStoryByID(id primitive.ObjectID) (*Story, error) {
	var story Story
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&story)
	if err != nil {
		return nil, err
	}
	//fmt.Println(user)
	return &story, nil
}

func (r *storyRepo) CreateStory(story Story) (*Story, error) {
	_, err := r.collection.InsertOne(context.Background(), story)
	if err != nil {
		return nil, err
	}
	return &story, nil
}
