package repo

import "go.mongodb.org/mongo-driver/mongo"

type commentRepo struct {
	collection *mongo.Collection
}

func NewCommentRepo(collection *mongo.Collection) CommentRepo {
	return &commentRepo{collection: collection}
}

func (r *commentRepo) GetCommentByID(id string) (*Comment, error) {
	return nil, nil
}

func (r *commentRepo) CreateComment(comment Comment) (*Comment, error) {
	return nil, nil
}
