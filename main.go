package main

import (
	"context"
	"log"
	"net/http"

	"golang_test_strat/adapters/api/rest/routes"
	"golang_test_strat/adapters/database"
	"golang_test_strat/domain"
	"golang_test_strat/domain/ports"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

var (
	serverPort string = "8080"
	ctx        context.Context
)

type CustomValidator struct {
	validator *validator.Validate
}

func main() {
	ctx = context.Background()

	// Http Server
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	wireUp(e)

	// Bootstrap routes
	err := routes.Boostrap(e)
	if err != nil {
		log.Fatal(err.Error())
	}

	// ping endpoint
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})

	e.Logger.Fatal(e.Start(":" + serverPort))
}

func wireUp(e *echo.Echo) {
	dummyDB := database.NewDummyDD()
	db, ok := dummyDB.(ports.UserDBInterface)
	if !ok {
		log.Fatal("dummyDB not implement UserDBInterface")
	}

	app, err := domain.NewApp(db)
	if err != nil {
		log.Fatal("Cannot instantiate app: " + err.Error())
	}

	routes.AppDep = app
	/*e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("App", app)
			return next(c)
		}
	})*/
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
