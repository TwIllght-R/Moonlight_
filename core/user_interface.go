package core

type New_user_req struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type New_user_resp struct {
	User_Id  string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   bool
}

type Login_req struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUser struct {
	Username string `json:"username"`
}

// type Login_resp struct {
// 	Token string `json:"token"`
// }

type UserCore interface {
	NewUser(New_user_req) (*New_user_resp, error)
	GetUser(string) (*GetUser, error)
	GetUsers() (*[]GetUser, error)
	// EditUser(New_user_req) (*New_user_resp, error)
	// DelUser(int) (*New_user_resp, error)
	LoginUser(Login_req) (*string, error)
}
