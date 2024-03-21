package repo

import "time"

type User struct {
	User_Id   string    `bson:"_id"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	DelatedAt time.Time `bson:"deleted_at"`
	IsDeleted bool      `bson:"is_deleted"`
	Username  string    `bson:"username"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	Role      string    `bson:"role"`
}

type UserRepo interface {
	GetAll() (*[]User, error)
	GetUserById(string) (*User, error)
	GetUserByEmail(string) (*User, error)
	CreateUser(User) (*User, error)
	UpdateUser(User) (*User, error)
	DeleteUser(string) error
}
