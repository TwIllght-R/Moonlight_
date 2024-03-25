package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepo struct {
	collection *mongo.Collection
}

func NewTaskRepo(collection *mongo.Collection) TaskRepo {
	return &taskRepo{collection: collection}
}

func (r *taskRepo) GetTaskByID(id primitive.ObjectID) (*Task, error) {
	var task Task
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepo) GetTaskByProjectID(id primitive.ObjectID) (*[]Task, error) {
	var tasks []Task
	filter := bson.M{
		"project_id": id,
		"delete_at":  bson.M{"$exists": false},
	}
	cur, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var task Task
		err := cur.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return &tasks, nil
}

func (r *taskRepo) GetAll() (*[]Task, error) {
	var tasks []Task
	filter := bson.M{
		"delete_at": bson.M{"$exists": false},
	}
	cur, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var task Task
		err := cur.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return &tasks, nil

}

func (r *taskRepo) CreateTask(task Task) (*Task, error) {
	_, err := r.collection.InsertOne(context.Background(), task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepo) UpdateTask(task Task) (*Task, error) {
	return nil, nil
}

func (r *taskRepo) DeleteTask(id primitive.ObjectID) (*Task, error) {
	return nil, nil
}
