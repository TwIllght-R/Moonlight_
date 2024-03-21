package repo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	collection *mongo.Collection
}

func NewUserRepo(collection *mongo.Collection) UserRepo {
	return &userRepo{collection: collection}
}

func (r *userRepo) GetUserById(id string) (*User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	fmt.Println(objectID)
	var user User
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetUserByEmail(email string) (*User, error) {
	var user User
	filter := bson.M{
		"email":      email,
		"is_deleted": bson.M{"$exists": false},
	}
	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepo) CreateUser(user User) (*User, error) {
	id := primitive.NewObjectID()
	user.User_Id = id.String()
	_, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (r *userRepo) GetAll() (*[]User, error) {
	var users []User
	filter := bson.M{
		"is_deleted": bson.M{"$exists": false},
	}
	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &users)
	if err != nil {
		return nil, err
	}
	return &users, nil
}
func (r *userRepo) UpdateUser(user User) (*User, error) {
	_, err := r.collection.ReplaceOne(context.Background(), bson.M{"_id": user.User_Id}, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepo) DeleteUser(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	var user User
	err = r.collection.FindOneAndDelete(context.Background(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return err
	}
	return nil
}
