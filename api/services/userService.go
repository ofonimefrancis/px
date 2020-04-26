package services

import "github.com/ofonimefrancis/pixels/api/models"

type UserService interface {
	Find(email string) (*models.User, error)
	Store(m models.User) error
}
