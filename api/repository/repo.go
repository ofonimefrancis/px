package repository

import "github.com/ofonimefrancis/pixels/api/models"

type UserRepository interface {
	Find(email string) (*models.User, error)
	Store(m models.User) error
}
