package core

import (
	"Moonlight_/repo"
	"errors"
	"log"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
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

	existingUser, err := r.userRepo.GetUserByEmail(req.Email)
	if err == nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := HashedPassword(req.Password)
	if err != nil {
		return nil, err
	}
	u := repo.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
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

func (r userCore) GetUser(id string) (*Get_user_resp, error) {
	user, err := r.userRepo.GetUserById(id)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return nil, errors.New("user not found")
		}
		log.Println(err)
		return nil, errors.New("errrrrr")
	}
	resp := Get_user_resp{
		Username: user.Username,
	}

	return &resp, nil
}

func (r userCore) GetUsers() (*[]Get_user_resp, error) {
	user, err := r.userRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	custResponses := []Get_user_resp{}
	for _, customer := range *user {
		custResponse := Get_user_resp{
			Username: customer.Username,
		}
		custResponses = append(custResponses, custResponse)
	}

	return &custResponses, nil
}

func (r userCore) EditUser(req New_user_req) (*New_user_resp, error) {
	hashedPassword, err := HashedPassword(req.Password)
	if err != nil {
		return nil, err
	}
	u := repo.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}
	newUser, err := r.userRepo.UpdateUser(u)
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

func (r userCore) DelUser(id string) (*New_user_resp, error) {
	err := r.userRepo.DeleteUser(id)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	resp := New_user_resp{
		Status: true,
	}

	return &resp, nil
}
