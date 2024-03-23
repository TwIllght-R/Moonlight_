package repo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	User_Id   primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	DelatedAt time.Time          `bson:"deleted_at"`
	Username  string             `bson:"username"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Role      string             `bson:"role"`
}
type Role struct {
	Role_Id    string    `bson:"_id"`
	CreatedAt  time.Time `bson:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at"`
	DelatedAt  time.Time `bson:"deleted_at"`
	RoleName   string    `bson:"role_name"`
	RoleAccess []string  `bson:"role_access"`
}

type UserRepo interface {
	GetAll() (*[]User, error)
	GetUserById(primitive.ObjectID) (*User, error)
	GetUserByEmail(string) (*User, error)
	GetUserByUsername(string) (*User, error)
	CreateUser(User) (*User, error)
	UpdateUser(User) (*User, error)
	DeleteUser(primitive.ObjectID) error
}
