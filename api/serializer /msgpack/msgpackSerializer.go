package msgpack

import (
	"github.com/ofonimefrancis/pixels/api/models"
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
)

type User struct{}

func (u *User) Decode(input []byte) (*models.User, error) {
	user := &models.User{}

	if err := msgpack.Unmarshal(input, user); err != nil {
		return nil, errors.Wrap(err, "serializer.User.Decode")
	}

	return user, nil
}

func (u *User) Encode(input *models.User) ([]byte, error) {
	rawByte, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.User.Encode")
	}

	return rawByte, nil
}
