package ports

import (
	"golang_test_strat/domain/models"
)

type UserDBInterface interface {
	Retrieve(email string) (models.User, error)
	Save(user models.User) error
}
