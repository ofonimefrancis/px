package json

import (
	"encoding/json"

	"github.com/ofonimefrancis/pixels/api/models"
	"github.com/pkg/errors"
)

type User struct{}

func (u *User) Decode(input []byte) (*models.User, error) {
	user := &models.User{}

	if err := json.Unmarshal(input, user); err != nil {
		return nil, errors.Wrap(err, "serializer.User.Decode")
	}

	return user, nil
}

func (u *User) Encode(input *models.User) ([]byte, error) {
	rawByte, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.User.Encode")
	}

	return rawByte, nil
}
