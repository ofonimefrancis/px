package services

type UserService interface {
	GetAll() ([]m.User, error)
	GetById(id string) (m.User, error)
	Store(data *m.User) error
	Update(data map[string]interface{}, id string) (*m.User, error)
	Delete(id string) error
}
