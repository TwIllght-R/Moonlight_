package repo

type User struct {
	User_Id  string `bson:"_id"`
	Username string `bson:"username"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type UserRepo interface {
	GetAll() (*[]User, error)
	GetUserById(string) (*User, error)
	GetUserByEmail(string) (*User, error)
	CreateUser(User) (*User, error)
	//UpdateUser(User) (*User, error)
	//DeleteUser(User) (*User, error)
}
