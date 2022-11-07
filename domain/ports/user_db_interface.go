package ports

import (
	"golang_test_strat/domain/models"
)

type UserDBInterface interface {
	RetrieveUserByUniqueFields(email string, phone string, user string) (models.User, error)
	RetrieveUser(fieldName string, value string) (models.User, error)
	SaveUser(user models.User) error
}
