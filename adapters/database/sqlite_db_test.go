package database

import (
	"golang_test_strat/domain/models"
	"golang_test_strat/domain/ports"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var db ports.UserDBInterface

func init() {
	dbI, ok := NewSqlite().(ports.UserDBInterface)
	if !ok {
		log.Fatal("dummyDB not implement UserDBInterface")
	}
	db = dbI
}

func Test_RetrieveUserByEmailOrPhone_notFound(t *testing.T) {
	user := models.User{User: "1234", Email: "correo@correo.com", Password: "Ab3456$", Phone: "1234567890"}

	db.SaveUser(user)

	actualUser, err := db.RetrieveUserByUniqueFields("correo@email.com", "123", "12346")
	if err != nil {
		t.Error(err)
	}

	if actualUser.User != "" {
		t.Errorf("actualUser is populated but should be empty: %+v", actualUser)
	}
}

func Test_RetrieveUserByUniqueFields_ok(t *testing.T) {
	user := models.User{User: "1234", Email: "correo@correo.com", Password: "Ab3456$", Phone: "1234567890"}

	db.SaveUser(user)

	expectedUser := models.User{User: "1234", Email: "correo@correo.com", Password: "Ab3456$", Phone: "1234567890"}
	actualUser, err := db.RetrieveUserByUniqueFields("correo@correo.com", "1234567890", "1234")
	if err != nil {
		t.Error(err)
	}

	assert.Equalf(t, expectedUser, actualUser, "Results are different: %+v")
}

func Test_RetrieveUser_okByPhone(t *testing.T) {
	user := models.User{User: "1234", Email: "correo@correo.com", Password: "Ab3456$", Phone: "1234567890"}

	db.SaveUser(user)

	expectedUser := models.User{User: "1234", Email: "correo@correo.com", Password: "Ab3456$", Phone: "1234567890"}
	actualUser, err := db.RetrieveUser("Phone", "1234567890")
	if err != nil {
		t.Error(err)
	}

	assert.Equalf(t, expectedUser, actualUser, "Results are different: %+v")
}

func Test_SaveUser_ok(t *testing.T) {
	user := models.User{User: "1234", Email: "correo@correo.com", Password: "Ab3456$"}

	gotError := db.SaveUser(user)
	if gotError != nil {
		t.Error("Error must be nil: " + gotError.Error())
	}
}
