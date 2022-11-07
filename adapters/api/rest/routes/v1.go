package routes

import (
	"errors"

	"github.com/labstack/echo/v4"

	"golang_test_strat/adapters/api/rest/controllers"
	"golang_test_strat/domain"
)

var AppDep *domain.App

func Boostrap(e *echo.Echo) error {
	loginController, err := controllers.NewLoginController(AppDep)
	if err != nil {
		return errors.New("cannot get login controller instance: " + err.Error())
	}

	registerController, err := controllers.NewRegisterController(AppDep)
	if err != nil {
		return errors.New("cannot get register controller instance: " + err.Error())
	}

	v3 := e.Group("/v1")
	v3.POST("/login", loginController.Login)
	v3.POST("/register", registerController.Register)

	return nil
}
