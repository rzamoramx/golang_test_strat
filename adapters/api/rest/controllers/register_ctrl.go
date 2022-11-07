package controllers

import (
	"fmt"
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

func (class *RegisterCtrl) Register(c echo.Context) error {
	var request rest.RequestRegister

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	if err = c.Validate(request); err != nil {
		return err
	}

	response, err := class.App.UserRegister(request)
	fmt.Printf("request: %+v\n", request)

	if err != nil {
		log.Error("error: " + err.Error())
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, response)
}
