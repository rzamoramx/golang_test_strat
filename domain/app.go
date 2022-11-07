package domain

import (
	"errors"
	"golang_test_strat/adapters/api/rest"
	"golang_test_strat/domain/ports"
)

type App struct {
	userDB ports.UserDBInterface
}

func NewApp(db ports.UserDBInterface) (*App, error) {
	if db == nil {
		return nil, errors.New("db instance is required")
	}

	return &App{}, nil
}

func (class *App) UserRegister(req rest.RequestRegister) (rest.ResponseRegister, error) {
	resp := rest.ResponseRegister{}

	// Exists?
	user, err := class.userDB.Retrieve(req.Email)
	if err != nil {
		return rest.ResponseRegister{
			Status:  "500",
			Message: "cannot process request"}, err
	}

	if user.Email != "" {
		return rest.ResponseRegister{
			Status:  "400",
			Message: "user already exists"}, err
	}

	// Persist new user
	err = class.userDB.Save(user)
	if err != nil {
		return rest.ResponseRegister{
			Status:  "500",
			Message: "cannot process request"}, err
	}

	return resp, nil
}
