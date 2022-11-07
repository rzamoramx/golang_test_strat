package controllers

import (
	"golang_test_strat/adapters/api/rest"
	"golang_test_strat/domain"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type LoginCtrl struct {
	App *domain.App
}

func NewLoginController(app *domain.App) (*LoginCtrl, error) {
	return &LoginCtrl{App: app}, nil
}

// Handle login requests
func (class *LoginCtrl) Login(c echo.Context) error {
	var request rest.RequestLogin

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	// validate required body fields
	if err = c.Validate(request); err != nil {
		return err
	}

	// domain app handle user login
	response, err := class.App.UserLogin(request)
	log.Printf("login endpoint request: %+v", request)
	log.Printf("login endpoint response: %+v", response)

	if err != nil {
		log.Error("error: " + err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, response)
}
