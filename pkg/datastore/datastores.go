package datastore

import model "github.com/ofonimefrancis/pixels/pkg/models"

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetBy(filter map[string]interface{}) (*model.User, error)
	Store(data *model.User) error
	Update(data map[string]interface{}, id string) (*model.User, error)
	Delete(id string) error
	Authenticate(username, password string) (bool, *model.User, error)
}
