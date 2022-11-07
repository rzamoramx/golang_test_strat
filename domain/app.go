package domain

import (
	"errors"
	"golang_test_strat/adapters/api/rest"
	"golang_test_strat/domain/models"
	"golang_test_strat/domain/ports"

	"github.com/fatih/structs"
	"github.com/golang-jwt/jwt"
)

type App struct {
	userDB ports.UserDBInterface
}

func NewApp(db ports.UserDBInterface) (*App, error) {
	if db == nil {
		return nil, errors.New("db instance is required")
	}

	return &App{userDB: db}, nil
}

// Login user
func (class *App) UserLogin(req rest.RequestLogin) (rest.LoginResponse, error) {
	// retrieve user from DB
	user, err := class.userDB.RetrieveUser("user", req.User)
	if err != nil {
		return rest.LoginResponse{
			Response: rest.Response{Status: "Error",
				Message: "cannot process request"}}, err
	}

	// model logic, validate user
	err = models.ValidateUser(req.User, req.Password, user)
	if err != nil {
		return rest.LoginResponse{
			Response: rest.Response{Status: "Error",
				Message: err.Error()}}, err
	}

	// generate jwt token
	token := jwt.New(jwt.SigningMethodHS256) //SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	//claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = req.User

	tokenString, err := token.SignedString([]byte("supersecretkey"))
	if err != nil {
		return rest.LoginResponse{
			Response: rest.Response{Status: "Error",
				Message: "cannot generate token, try later"}}, err
	}

	// Ok, return token
	return rest.LoginResponse{
		Response: rest.Response{Status: "OK",
			Message: "user logged"},
		Token: tokenString}, nil
}

// Register new user
func (class *App) UserRegister(req rest.RequestRegister) (rest.Response, error) {
	// retreive from DB
	user, err := class.userDB.RetrieveUserByUniqueFields(req.Email, req.Phone, req.User)
	if err != nil {
		return rest.Response{
			Status:  "Error",
			Message: "cannot process request"}, err
	}

	// model logic, validate rules for user register
	err = models.RulesForRegisterUser(structs.Map(req), user)
	if err != nil {
		return rest.Response{
			Status:  "Error",
			Message: err.Error()}, err
	}

	// persist new user
	user.User = req.User
	user.Password = req.Password
	user.Phone = req.Phone
	user.Email = req.Email

	err = class.userDB.SaveUser(user)
	if err != nil {
		return rest.Response{
			Status:  "Error",
			Message: "cannot persist new user, try later"}, err
	}

	return rest.Response{Status: "OK", Message: "user registered"}, nil
}
