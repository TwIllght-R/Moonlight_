package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type projectRepo struct {
	collection *mongo.Collection
}

func NewProjectRepo(collection *mongo.Collection) ProjectRepo {
	return &projectRepo{collection: collection}
}
func (r *projectRepo) GetProjectByID(id string) (*Project, error) {
	var project Project
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *projectRepo) GetAll() (*[]Project, error) {
	var stories []Project
	filter := bson.M{
		"is_deleted": bson.M{"$exists": false},
	}
	cur, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var project Project
		err := cur.Decode(&project)
		if err != nil {
			return nil, err
		}
		stories = append(stories, project)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return &stories, nil
}

func (r *projectRepo) CreateProject(project Project) (*Project, error) {
	_, err := r.collection.InsertOne(context.Background(), project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *projectRepo) UpdateProject(project Project) (*Project, error) {
	_, err := r.collection.ReplaceOne(context.Background(), bson.M{"_id": project.ID}, project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}
func (r *projectRepo) DeleteProject(id string) (*Project, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var project Project
	err = r.collection.FindOneAndDelete(context.Background(), bson.M{"_id": objectID}).Decode(&project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}
