package repointerfaces

import "teslaUsers/src/models"

type UserRepo interface {
	AddUser(user models.User) (models.User, error)
	GetUsers() ([]models.User, error)
	GetUser(mail string) (models.User, error)
	AddPaymentMethod(mail string, paymentMethod string) error
}
