package serviceLogic

import (
	ds "github.com/ofonimefrancis/pixels/pkg/datastore"
  m "github.com/ofonimefrancis/pixels/pkg/models"
  "github.com/pkg/errors"
  services "github.com/ofonimefrancis/pixels/pkg/services/contracts"
  commons "github.com/ofonimefrancis/pixels/pkg/commons/error"
  "gopkg.in/dealancer/validate.v2"
)

type userService struct {
	userDatastore ds.UserRepository
}

func NewUserService(userDatastore ds.UserRepository) services.UserService {
	return &userService{
		userDatastore,
	}
}


func (u *userService) GetAll() ([]m.User, error) {
	users := []m.User 
	if res, err := u.userDatastore.GetAll(); err != nil {
		return res, err 
	}
	return res, nil 
}

func (u *userService) GetById(id string) (*m.User, error) {
	res, e := u.userRepo.GetBy(map[string]interface{}{"id": id})
	if e != nil {
		return res, e
	}

	return res, nil
}


func (u *userService) Store(data *m.User) error {
	if e := validate.Validate(data); e != nil {
		return errors.Wrap(commons.ErrUserInvalid, "service.User.Store")
	}
	if data.ID == "" {
		data.ID = shortid.MustGenerate()
	}
	if isFound, _, _ := u.GetByUsername(data.Username); isFound {
		return errors.Wrap(commons.ErrUserNameDuplicate, "service.User.Store")
	}
  
  //TODO: Change this password hashing strategy...
  data.Password = datastore.EncryptPassword(data.Password)
	return u.userDatastore.Store(data)

}

func (u *userService) Update(data map[string]interface{}, id string) (*m.User, error) {
	user := new(m.User)
	var e error
	if data["id"].(string) == "" {
		return user, errs.Wrap(helper.ErrUserInvalid, "service.User.Update")
	}
	if data["password"].(string) != "" {
		data["password"] = datastore.EncryptPassword(data["password"].(string))
	}
	user, e = u.userRepo.Update(data, id)
	if e != nil {
		return user, errs.Wrap(e, "service.User.Update")
	}
	return user, nil

}

func (u *userService) Delete(id string) error {
	if id == "" {
		return errs.Wrap(commons.ErrUserInvalid, "service.User.Delete")
	}
	if e := u.userDatastore.Delete(id); e != nil {
		return e
	}
	return nil

}

func (u *userService) GetByUsername(email string) (bool, *m.User, error) {
	res, e := u.userDatastore.GetBy(map[string]interface{}{"email": email})
	if e != nil {
		return false, res, e
	}

	return true, res, nil
}