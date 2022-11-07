package controllers

import (
	"golang_test_strat/domain"

	"github.com/labstack/echo/v4"
)

type LoginCtrl struct {
	App *domain.App
}

func NewLoginController(app *domain.App) (*LoginCtrl, error) {
	return &LoginCtrl{App: app}, nil
}

func (class *LoginCtrl) Login(c echo.Context) error {
	return nil
}
