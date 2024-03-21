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

type Get_user_resp struct {
	User_Id  string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// type Login_resp struct {
// 	Token string `json:"token"`
// }

type UserCore interface {
	NewUser(New_user_req) (*New_user_resp, error)
	GetUser(string) (*Get_user_resp, error)
	GetUsers() (*[]Get_user_resp, error)
	EditUser(New_user_req) (*New_user_resp, error)
	DelUser(string) (*New_user_resp, error)
	LoginUser(Login_req) (*string, error)
}
