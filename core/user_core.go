package core

import (
	"Moonlight_/repo"
	"log"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userCore struct {
	userRepo repo.UserRepo
}

func NewUserCore(userRepo repo.UserRepo) UserCore {
	return userCore{userRepo: userRepo}
}

func (r userCore) LoginUser(loginReq Login_req) (*string, error) {
	user, err := r.userRepo.GetUserByEmail(loginReq.Email)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.User_Id
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	t, err := token.SignedString([]byte("test")) //imple env
	if err != nil {
		log.Println("here")
		return nil, err
	}

	//return token
	return &t, nil
}

func (r userCore) NewUser(req New_user_req) (*New_user_resp, error) {
	hashedPassword, err := HashedPassword(req.Password)
	if err != nil {
		return nil, err
	}
	u := repo.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}
	newUser, err := r.userRepo.CreateUser(u)
	if err != nil {
		log.Panic(err)
		return nil, err

	}
	resp := New_user_resp{
		User_Id:  newUser.User_Id,
		Email:    newUser.Email,
		Username: newUser.Username,
		Status:   true,
	}

	return &resp, nil

}

func (r userCore) GetUser(id string) (*GetUser, error) {
	user, err := r.userRepo.GetUserById(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp := GetUser{
		Username: user.Username,
		// You might need to populate other fields of GetUser if needed
	}

	return &resp, nil
}

func (r userCore) GetUsers() (*[]GetUser, error) {
	user, err := r.userRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	custResponses := []GetUser{}
	for _, customer := range *user {
		custResponse := GetUser{
			Username: customer.Username,
		}
		custResponses = append(custResponses, custResponse)
	}

	return &custResponses, nil
}
