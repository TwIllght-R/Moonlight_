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
	err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	//fmt.Println(user)
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
	//auto id increment
}

func (r *userRepo) GetAll() (*[]User, error) {
	var users []User
	cursor, err := r.collection.Find(context.Background(), bson.M{})
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

// type user_RepoDB struct {
// 	db *sqlx.DB
// }

// func New_user_RepoDB(db *sqlx.DB) UserRepo {
// 	return user_RepoDB{db: db}
// }
// func (r user_RepoDB) GetAll() ([]User, error) {
// 	customers := []User{}
// 	q := "select * from customers"
// 	err := r.db.Select(&customers, q)
// 	if err != nil {
// 		return nil, err

// 	}
// 	return customers, nil
// }

// func (r user_RepoDB) CreateUser(user User) (*User, error) {
// 	// query := "insert into users (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
// 	// result, err := r.db.Exec(
// 	// 	query,
// 	// user.Email
// 	// )

// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// id, err := result.LastInsertId()
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// acc.AccountID = int(id)

// 	return &user, nil
// }
