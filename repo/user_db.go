package repo

import (
	"context"
	"errors"

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

// GetUserById gets a user by id
func (r *userRepo) GetUserById(id primitive.ObjectID) (*User, error) {
	var user User
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail gets a user by email
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

// GetUserByUsername gets a user by username
func (r *userRepo) GetUserByUsername(username string) (*User, error) {
	var user User
	filter := bson.M{
		"username":   username,
		"is_deleted": bson.M{"$exists": false},
	}
	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user
func (r *userRepo) CreateUser(user User) (*User, error) {
	id := primitive.NewObjectID()
	user.User_Id = id
	_, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

// GetAll gets all users
func (r *userRepo) GetAll() (*[]User, error) {
	var users []User
	filter := bson.M{
		"delete_at": bson.M{"$exists": false},
	}
	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return &users, nil
}

// UpdateUser updates a user
func (r *userRepo) UpdateUser(user User) (*User, error) {
	filter := bson.M{"_id": user.User_Id}
	update := bson.M{
		"$set": bson.M{
			"username":   user.Username,
			"updated_at": user.UpdatedAt,
			"email":      user.Email,
			"password":   user.Password,
			"role":       user.Role,
		},
	}
	result, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

// DeleteUser deletes a user
func (r *userRepo) DeleteUser(id primitive.ObjectID) error {
	var user User
	err := r.collection.FindOneAndDelete(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return err
	}
	return nil
}
