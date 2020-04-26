package serializer

import "github.com/ofonimefrancis/pixels/api/models"

type Serializer interface {
	Decode(input []byte) (*models.User, error)
	Encode(user *model.User) ([]byte, error)
}
