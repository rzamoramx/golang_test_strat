package controllers

import (
	"golang_test_strat/adapters/api/rest"
	"net/http"

	"golang_test_strat/domain"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type RegisterCtrl struct {
	App *domain.App
}

func NewRegisterController(app *domain.App) (*RegisterCtrl, error) {
	return &RegisterCtrl{App: app}, nil
}

// handle register user requests
func (class *RegisterCtrl) Register(c echo.Context) error {
	var request rest.RequestRegister

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	if err = c.Validate(request); err != nil {
		return err
	}

	// domain app handle user register
	response, err := class.App.UserRegister(request)
	log.Printf("register endpoint request: %+v", request)
	log.Printf("register endpoint response: %+v", response)

	if err != nil {
		log.Error("error: " + err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, response)
}
