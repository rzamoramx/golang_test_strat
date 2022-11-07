package domain

import (
	"golang_test_strat/adapters/api/rest"
	"golang_test_strat/adapters/database"
	"golang_test_strat/domain/ports"
	"log"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

var app *App

func init() {
	var err error

	dummyDB := database.NewSqlite()
	db, ok := dummyDB.(ports.UserDBInterface)
	if !ok {
		log.Fatal("dummyDB not implement UserDBInterface")
	}

	app, err = NewApp(db)
	if err != nil {
		panic("cannot instantiate db: " + err.Error())
	}
}

func Test_UserLogin_ok(t *testing.T) {
	user := "fulanito"
	token, err := helperToken(user)
	if err != nil {
		t.Error(err)
	}

	reqReg := rest.RequestRegister{User: user, Email: "fulanito@correo.com", Phone: "1234567890", Password: "Ab3456$"}
	reqLog := rest.RequestLogin{User: "fulanito", Password: "Ab3456$"}
	wantResponse := rest.LoginResponse{
		Response: rest.Response{Status: "OK",
			Message: "user logged"},
		Token: token}

	app.UserRegister(reqReg)

	gotResponse, err := app.UserLogin(reqLog)
	if err != nil {
		t.Error(err)
	}

	assert.Equalf(t, wantResponse, gotResponse, "different response: %+v", gotResponse)
}

func Test_UserRegister_ok(t *testing.T) {
	req := rest.RequestRegister{User: "fulanito", Email: "fulanito@correo.com", Phone: "1234567890", Password: "Ab3456$"}
	wantResponse := rest.Response{Status: "OK", Message: "user registered"}

	gotResponse, err := app.UserRegister(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equalf(t, wantResponse, gotResponse, "different response: %+v", gotResponse)
}

func helperToken(user string) (string, error) {
	// Generate jwt token
	token := jwt.New(jwt.SigningMethodHS256) //SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	//claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = user

	tokenString, err := token.SignedString([]byte("supersecretkey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
