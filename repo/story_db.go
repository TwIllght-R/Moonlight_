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

func (r *storyRepo) GetStoryByID(id string) (*Story, error) {
	var story Story
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&story)
	if err != nil {
		return nil, err
	}
	return &story, nil
}

func (r *storyRepo) GetAll() (*[]Story, error) {
	var stories []Story
	filter := bson.M{
		"is_deleted": bson.M{"$exists": false},
	}
	cur, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var story Story
		err := cur.Decode(&story)
		if err != nil {
			return nil, err
		}
		stories = append(stories, story)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return &stories, nil
}

func (r *storyRepo) CreateStory(story Story) (*Story, error) {
	_, err := r.collection.InsertOne(context.Background(), story)
	if err != nil {
		return nil, err
	}
	return &story, nil
}

func (r *storyRepo) UpdateStory(story Story) (*Story, error) {
	_, err := r.collection.ReplaceOne(context.Background(), bson.M{"_id": story.ID}, story)
	if err != nil {
		return nil, err
	}
	return &story, nil
}
func (r *storyRepo) DeleteStory(id string) (*Story, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var story Story
	err = r.collection.FindOneAndDelete(context.Background(), bson.M{"_id": objectID}).Decode(&story)
	if err != nil {
		return nil, err
	}
	return &story, nil
}
