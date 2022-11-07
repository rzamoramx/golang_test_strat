package models

import (
	"errors"
	"golang_test_strat/adapters/api/rest"
	"testing"

	"github.com/fatih/structs"
)

func Test_RulesForRegisterUser_ok(t *testing.T) {
	user := User{}
	req := rest.RequestRegister{Phone: " 1234567890 ", Password: "Ab3456$"}

	gotError := RulesForRegisterUser(structs.Map(req), user)

	if gotError != nil {
		t.Error("expected error must be 'password must have at least one lowercase, one uppercase, one special chars' instead: " + gotError.Error())
	}
}

func Test_RulesForRegisterUser_wrongcomplxpwd(t *testing.T) {
	user := User{}
	wantError := errors.New("password must have at least one lowercase, one uppercase, one special char")
	req := rest.RequestRegister{Phone: "1234567890", Password: "123456"}

	gotError := RulesForRegisterUser(structs.Map(req), user)

	if wantError.Error() != gotError.Error() {
		t.Error("expected error must be 'password must have at least one lowercase, one uppercase, one special chars' instead: " + gotError.Error())
	}
}

func Test_RulesForRegisterUser_wronglenpwd(t *testing.T) {
	user := User{}
	wantError := errors.New("password must be between 6 and 12 characters")
	req := rest.RequestRegister{Phone: "1234567890", Password: "1234"}

	gotError := RulesForRegisterUser(structs.Map(req), user)

	if wantError.Error() != gotError.Error() {
		t.Error("expected error must be 'password must be between 6 and 12 characters' instead: " + gotError.Error())
	}
}

func Test_RulesForRegisterUser_wrongphone(t *testing.T) {
	user := User{}
	wantError := errors.New("phone must be 10 digits")
	req := rest.RequestRegister{Phone: "1234"}

	gotError := RulesForRegisterUser(structs.Map(req), user)

	if wantError.Error() != gotError.Error() {
		t.Error("expected error must be 'phone must be 10 digits' instead: " + gotError.Error())
	}
}

func Test_RulesForRegisterUser_userexists(t *testing.T) {
	user := User{User: "1234", Password: "abcd", Email: "correo@correo.com"}
	wantError := errors.New("user already exists")
	req := rest.RequestRegister{User: "ab", Password: "1234", Email: "correo@correo.com"}

	gotError := RulesForRegisterUser(structs.Map(req), user)

	if wantError.Error() != gotError.Error() {
		t.Error("expected error must be 'user already exists' instead: " + gotError.Error())
	}
}

func Test_ValidateUser_ok(t *testing.T) {
	user := User{User: "1234", Password: "abcd"}

	gotError := ValidateUser("1234", "abcd", user)

	if gotError != nil {
		t.Error("expected error must be nil, instead: " + gotError.Error())
	}
}

func Test_ValidateUser_wrongpwd(t *testing.T) {
	user := User{User: "1234", Password: "abcd"}
	wantError := errors.New("wrong password")

	gotError := ValidateUser("1234", "qwerty", user)

	if wantError.Error() != gotError.Error() {
		t.Error("expected error must be 'wrong password', instead: " + gotError.Error())
	}
}

func Test_ValidateUser_usernotfound(t *testing.T) {
	user := User{}
	wantError := errors.New("User not found")

	gotError := ValidateUser("1234", "qwerty", user)

	if wantError.Error() != gotError.Error() {
		t.Error("expected error must be 'User not found', instead: " + gotError.Error())
	}
}
