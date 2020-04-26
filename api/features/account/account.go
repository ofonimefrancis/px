package account

import (
	"github.com/ofonimefrancis/pixels/api/models"
	"github.com/ofonimefrancis/pixels/api/repository"
)

type userService struct {
	datastore repository.UserRepository
}

func NewUserService(ds repository.UserRepository) repository.UserRepository {
	return &userService{
		ds,
	}
}

func (u *userService) Find(email string) (*models.User, error) {
	return u.datastore.Find(email)
}

func (u *userService) Store(m models.User) error {
	return u.datastore.Store(m)
}
